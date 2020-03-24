package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type FileType int

const (
	File FileType = iota
	Dir
)

type TestTemp struct {
	Name string
	Path string
	Time string
	Size int
	Type FileType
}

func view(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println(r.Header)
	fmt.Println(r.URL)
	t, _ := template.ParseFiles("app.html")
	//t,_ := template.ParseGlob("*.tpl")

	list := [3]TestTemp{
		{Path: "/skyDrive/file/main/", Name: "huanglin", Time: "2019", Size: 0, Type: 0},
		{Path: "/skyDrive/file/main.go/", Name: "huangmain", Time: "2020", Size: 1, Type: 1},
		{Path: "/skyDrive/cc/", Name: "huanghao", Time: "2021", Size: 2, Type: 1},
	}

	t.Execute(w, list)
}

func main() {
	http.HandleFunc("/", view)
	fmt.Println(http.ListenAndServe("localhost:8080", nil))
	/*
	   1.声明一个Template对象并解析模板文本
	   func New(name string) *Template
	   func (t *Template) Parse(text string) (*Template, error)

	   2.从html文件解析模板
	   func ParseFiles(filenames ...string) (*Template, error)

	   3.模板生成器的包装
	   template.Must(*template.Template, error )会在Parse返回err不为nil时，调用panic。
	   func Must(t *Template, err error) *Template

	   t := template.Must(template.New("name").Parse("html"))
	*/

}
