package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func GenerateIntA(done chan struct{}) (chan int){
	ch := make(chan int)
	go func(){
		Lable:
		for{
			select{
			case ch<- rand.Int():
				println("choose ch")
			case <-done:
				println("choose done")
				break Lable
			}
		}

		close(ch)
	}()

	return ch
}

func main(){
	done := make(chan struct{})
	ch := GenerateIntA(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	
	//因为操作系统的原因，select可能会同一时间检测到，done和ch，并选取其中一个
	close(done)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("NumGoroutine=", runtime.NumGoroutine())
}