package main

import (
	"log"
	"os"
	"text/template"
	"time"

	"github.com/my/repo/go程序设计语言/practice/4/github"

)

const template1 = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:	{{.User.Login}}
Title: 	{{.Title | printf "%s.64s"}}
Age:	{{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(template1))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
