package main

import (
	"fmt"
	"log"
	"os"

	"github.com/my/repo/go程序设计语言/practice/4/github"

)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("%#v\n", item)
	}
}
