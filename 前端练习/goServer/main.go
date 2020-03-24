package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"sort"
	"strings"
	"text/template"
	"time"
)

const (
	rootPath       = "."
	ModuleSkyDrive = "/skyDrive"
)

type FileType int

const (
	File FileType = iota
	Dir
	Root
)

type FileInfo struct {
	Name string
	Path string
	Time string
	Size string
	Type FileType
	Even bool
}

type condResult int

const (
	condNone condResult = iota
	condTrue
	condFalse
)

func toHTTPError(err error) (msg string, httpStatus int) {
	if os.IsNotExist(err) {
		return "404 page not found", http.StatusNotFound
	}
	if os.IsPermission(err) {
		return "403 Forbidden", http.StatusForbidden
	}
	// Default:
	return "500 Internal Server Error", http.StatusInternalServerError
}
func localRedirect(w http.ResponseWriter, r *http.Request, newPath string) {
	if q := r.URL.RawQuery; q != "" {
		newPath += "?" + q
	}
	w.Header().Set("Location", newPath)
	w.WriteHeader(http.StatusMovedPermanently)
}

var unixEpochTime = time.Unix(0, 0)

func isZeroTime(t time.Time) bool {
	return t.IsZero() || t.Equal(unixEpochTime)
}
func checkIfModifiedSince(r *http.Request, modtime time.Time) condResult {
	if r.Method != "GET" && r.Method != "HEAD" {
		return condNone
	}
	ims := r.Header.Get("If-Modified-Since")
	if ims == "" || isZeroTime(modtime) {
		return condNone
	}
	t, err := http.ParseTime(ims)
	if err != nil {
		return condNone
	}
	// The Last-Modified header truncates sub-second precision so
	// the modtime needs to be truncated too.
	modtime = modtime.Truncate(time.Second)
	if modtime.Before(t) || modtime.Equal(t) {
		return condFalse
	}
	return condTrue
}
func setLastModified(w http.ResponseWriter, modtime time.Time) {
	if !isZeroTime(modtime) {
		w.Header().Set("Last-Modified", modtime.UTC().Format(http.TimeFormat))
	}
}
func writeNotModified(w http.ResponseWriter) {
	// RFC 7232 section 4.1:
	// a sender SHOULD NOT generate representation metadata other than the
	// above listed fields unless said metadata exists for the purpose of
	// guiding cache updates (e.g., Last-Modified might be useful if the
	// response does not have an ETag field).
	h := w.Header()
	delete(h, "Content-Type")
	delete(h, "Content-Length")
	if h.Get("Etag") != "" {
		delete(h, "Last-Modified")
	}
	w.WriteHeader(http.StatusNotModified)
}

func dirList(w http.ResponseWriter, r *http.Request, f os.File) {
	var name, path string
	var urlT url.URL
	dirs, err := f.Readdir(-1)
	if err != nil {
		http.Error(w, "Error reading directory", http.StatusInternalServerError)
		return
	}
	upath := r.URL.Path
	sort.Slice(dirs, func(i, j int) bool { return dirs[i].ModTime().Before(dirs[j].ModTime()) })

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	list := make([]FileInfo, 0)
	var element FileInfo

	t, err := template.ParseFiles("app.html")
	if err != nil {
		fmt.Println(err)
	}

	if upath != ModuleSkyDrive+"/" {
		name = upath + ".."
		urlT = url.URL{Path: name}
		element.Name = ".."
		element.Path = urlT.String()
		element.Size = "-"
		element.Time = ""
		element.Type = Root
		list = append(list, element)
	}

	for i, d := range dirs {
		if i%2 == 0 {
			element.Even = true
		} else {
			element.Even = false
		}
		if d.IsDir() {
			element.Size = "-"
			element.Type = Dir
			element.Name = d.Name() + "/"
		} else {
			element.Size = fmt.Sprintf("%dKB", (d.Size()+(1<<10+1))/(1<<10))
			element.Type = File
			element.Name = d.Name()
		}
		// name may contain '?' or '#', which must be escaped to remain
		// part of the URL path, and not indicate the start of a query
		// string or fragment.
		path = upath + d.Name() + "/"
		urlT = url.URL{Path: path}
		element.Path = fmt.Sprintf("%s", urlT.String())
		element.Time = fmt.Sprintf("%2v-%3v-%4v", d.ModTime().Day(), d.ModTime().Month(), d.ModTime().Year())
		list = append(list, element)
	}
	t.Execute(w, list)
}

func dirFormatServe(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}

	//根目录
	name := rootPath + "/" + strings.TrimLeft(upath, ModuleSkyDrive)

	file, err := os.Open(path.Clean(name))
	//是否有此路径
	if err != nil {
		msg, code := toHTTPError(err)
		http.Error(w, msg, code)
		return
	}
	defer file.Close()

	//是否是文件夹
	fileInfo, err := file.Stat()
	if err != nil {
		msg, code := toHTTPError(err)
		http.Error(w, msg, code)
		return
	}
	//如果是文件夹，我们需要填写模板，并且将html文件发送过去
	if fileInfo.IsDir() {
		url := r.URL.Path
		// redirect if the directory name doesn't end in a slash
		if url == "" || url[len(url)-1] != '/' {
			localRedirect(w, r, path.Base(url)+"/")
			return
		}

		if checkIfModifiedSince(r, fileInfo.ModTime()) == condFalse {
			writeNotModified(w)
			return
		}
		setLastModified(w, fileInfo.ModTime())

		w.Header().Set("Last-Modified", fileInfo.ModTime().UTC().Format(http.TimeFormat))
		dirList(w, r, *file)
		return
	} else {
		// serveContent will check modification time
		http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)
	}
	msg, code := "404 page not found", http.StatusNotFound
	http.Error(w, msg, code)
	return
}

func main() {
	http.HandleFunc(ModuleSkyDrive+"/", dirFormatServe)

	log.Fatal(http.ListenAndServe("localhost:80", nil))
}
