package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", period)
	//duration实现String()和Seconds(),事实上是转换成了durationValue类型，而
	//durationValue实现了Set方法
	fmt.Println(period.Seconds())
	time.Sleep(*period)
	fmt.Println()
}
