//多个goroutine增强型生成器，这个实例生成器无限膨胀了

package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func GenerateIntA() chan int {
	println("我们不会重复调用")
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
			println(ch)
		}
	}()

	return ch
}

func GenerateIntB() chan int {
	println("我们不会重复调用")
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
			println(ch)
		}
	}()

	return ch
}

func GenerateInt() chan int {
	ch := make(chan int, 20)
	go func() {
		for {
			select {
			case ch <- <-GenerateIntA():
			case ch <- <-GenerateIntB():
			}
		}
	}()

	return ch
}

func main() {
	ch := GenerateInt()
	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("NumGoroutine=", runtime.NumGoroutine())
}
