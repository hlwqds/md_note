package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func finlinks2(url string) []string {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.Status != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parse %s as html: %v", url, err)
	}
}.

func main() {
	for _, url := range os.Args[1:] {
		links, err := findlinks2(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}

		for _, link := range links {
			fmt.Println(link)
		}
	}
}
