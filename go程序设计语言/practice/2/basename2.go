package main

import (
	"fmt"
	"strings"
)

func baseName2(s string) string {
	slash := strings.LastIndex(s, "/")
	//如果没有找到/，返回-1
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s
}

func main() {
	s := "/dad/dawd/dd.hh"
	s = baseName2(s)
	fmt.Println(s)
}
