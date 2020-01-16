package main

import (
	"time"
	"runtime"
)

func main() (){
	go func(){
		sum := 0
		for i := 0; i < 10000; i++{
			sum += i
		}

		println(sum)
		time.Sleep(1 * time.Second)

	}()

	println("NumGoroutine=", runtime.NumGoroutine())

	time.Sleep(5 * time.Second)
}