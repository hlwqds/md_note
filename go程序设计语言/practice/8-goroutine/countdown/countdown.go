package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown.")
	/*
		tick := time.Tick(1 * time.Second) //这个tick将在程序的整个生命周期内运行，不正确使用会造成goroutine泄漏
	*/
	ticker := time.NewTicker(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		select {
		case <-ticker.C:
			fmt.Println(countdown)
		case <-abort:
			fmt.Println("Abort!")
			ticker.Stop()
			return
			/*
				default:
					fmt.Println("Not block!")
			*/
		}
	}
	launch()
}

func launch() {
	fmt.Println("Launch!")
}
