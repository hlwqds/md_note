//生成器
//最简单的带缓冲的生成器
package main

import (
	"fmt"
	"math/rand"
)

func GenerateIntA() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			select {
			case ch <- rand.Int():
			}
		}
	}()

	return ch
}

func main() {
	ch := GenerateIntA()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
