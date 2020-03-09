package main

import (
	"fmt"
	"log"

	"github.com/my/repo/go程序设计语言/practice/8-goroutine/links"

)

func cral(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
}
