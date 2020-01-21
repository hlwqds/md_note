package main

import (
	"net/http"
	"html/template"
)

func index(w http.ResponseWriter, r *http.Request){
	files := []string{
		"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html",
	}

	templates := template.Must(template.ParseFiles(files...))
	if threads, err := data.Threads(); err == nil{
		templates.ExecuteTemplate(w, "layout", threads)
	}
}

func main(){
	//创建一个多路复用器
	mux := http.NewServeMux()

	//创建一个能够为指定目录中的静态文件服务的处理器，目录为/public
	files := http.FileServer(http.Dir(config.Static))
	//当有static开头的url到来时，会去除url中的static字段，然后到public目录中寻找静态文件
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//当有针对根URL的请求到达时，多路复用器将请求重定向到名为index的处理器函数
	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/sigup", sigup)
	mux.HandleFunc("/sigup_account", sigup_account)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux
	}

	server.ListenAndServe()
}