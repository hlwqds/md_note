package main

import (
	"html/template"
	"log"
	"os"

	"github.com/my/repo/go程序设计语言/practice/4/github"

)

const template2 = `
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
	<th>#</th>
	<th>State<th>
	<th>User</th>
	<th>Title</th>
</tr>
{{range .Items}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>`
const htmlFileName = "tmp.html"

var templateExample = template.Must(template.New("htmlTemp").Parse(template2))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	ptr, err := os.Create(htmlFileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer ptr.Close()

	if err := templateExample.Execute(ptr, result); err != nil {
		log.Fatal(err)
	}
}
