package main

import (
	"flag"
	"fmt"

	"github.com/my/repo/go程序设计语言/practice/7-interface/celflag"
)

var temp = celflag.CelsiusFlag("temp", 20.0, "the Temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
