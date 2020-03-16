package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const (
	rootPath              = "."
	ModuleSkyDrive        = "/skyDrive"
	SubModuleSkyDriveFile = "/file"
)

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
	var name string
	var urlT url.URL
	dirs, err := f.Readdir(-1)
	if err != nil {
		http.Error(w, "Error reading directory", http.StatusInternalServerError)
		return
	}
	upath := r.URL.Path
	sort.Slice(dirs, func(i, j int) bool { return dirs[i].ModTime().Before(dirs[j].ModTime()) })

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<pre>\n")
	if upath != ModuleSkyDrive+"/" {
		name = upath + ".."
		urlT = url.URL{Path: name}
		fmt.Fprintf(w, "<a href=\"%s\">%s</a>\n", urlT.String(), "..")
	}

	for _, d := range dirs {
		var sizeS string
		if d.IsDir() {
			sizeS = "-"
			name = upath + d.Name() + "/"
		} else {
			sizeS = fmt.Sprintf("%dKB", (d.Size()+(1<<10+1))/(1<<10))
			rPath := strings.TrimLeft(upath, ModuleSkyDrive)
			name = ModuleSkyDrive + SubModuleSkyDriveFile + "/" + rPath + d.Name() + "/"
		}
		// name may contain '?' or '#', which must be escaped to remain
		// part of the URL path, and not indicate the start of a query
		// string or fragment.
		urlT = url.URL{Path: name}
		fmt.Fprintf(w, "<a href=\"%s\">%s</a>", urlT.String(), html.EscapeString(d.Name()))

		fmt.Fprintf(w, "\t%2v-%3v-%4v\t%20v\n", d.ModTime().Day(), d.ModTime().Month(), d.ModTime().Year(), sizeS)
	}
	fmt.Fprintf(w, "</pre>\n")
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
	}
	msg, code := "404 page not found", http.StatusNotFound
	http.Error(w, msg, code)
	return
}

func fileHttpServeFunc(w http.ResponseWriter, r *http.Request) {
	uPath := path.Clean(r.URL.Path)
	pathT, fileName := filepath.Split(uPath)
	dirPath := strings.TrimLeft(pathT, ModuleSkyDrive+SubModuleSkyDriveFile)

	http.ServeFile(w, r, rootPath+"/"+dirPath+fileName)

	return
}

func main() {
	//文件夹处理
	http.HandleFunc(ModuleSkyDrive+"/", dirFormatServe)

	http.HandleFunc(ModuleSkyDrive+SubModuleSkyDriveFile+"/", fileHttpServeFunc)
	log.Fatal(http.ListenAndServe("localhost:80", nil))
}
