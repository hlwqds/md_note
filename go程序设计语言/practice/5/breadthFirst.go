package main

import (
	"fmt"
	"log"
)

func breadthFirst(f func(string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) != 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if seen[item] != true {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}

	return list
}

func main() {
	breadthFirst(crawl, []string{"https://baidu.com"})
}
