package links

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"

)

//Extract crawl a html web
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("get resp: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("get resp return: %d", resp.StatusCode)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parse resp: %v", err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, prev func(n *html.Node), post func(n *html.Node)) {
	if prev != nil {
		prev(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, prev, post)
	}

	if post != nil {
		post(n)
	}
}
