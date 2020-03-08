package main

import (
	"bufio"
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type WordsScanner struct {
	wordsNum int
}

func (w *WordsScanner) Write(p []byte) (int, error) {
	index := 0
	w.wordsNum = 0
	for {
		wordLen, _, _ := bufio.ScanWords([]byte(p)[index:], true)
		if wordLen == 0 {
			break
		}
		index += wordLen
		w.wordsNum++
	}

	return w.wordsNum, nil
}

func main() {
	w := WordsScanner{
		wordsNum: 9,
	}

	fmt.Println(w)
	fmt.Fprintf(&w, "hello, every one")
	fmt.Println(w)

	reader := strings.NewReader("<h1>haha</h1>")
	node, err := html.Parse(reader)
	fmt.Println(err)
	fmt.Println(node)
}
