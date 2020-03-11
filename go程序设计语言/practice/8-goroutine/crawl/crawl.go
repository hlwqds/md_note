package main

import (
	"fmt"
	"log"
	"os"

	"github.com/my/repo/go程序设计语言/practice/8-goroutine/links"

)

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int
	n++
	go func() { worklist <- os.Args[1:] }()
	seen := make(map[string]bool)
	for n > 0 {
		for list := range worklist {
			n--
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					n++
					go func(link string) {
						worklist <- crawl(link)
					}(link)
				}
			}
			if n == 0 {
				close(worklist)
			}
		}
	}
}
