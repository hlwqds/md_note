package main

import (
	"fmt"
	"runtime"
	"math/rand"
)

func GenerateIntA(done chan struct{}) (chan int){
	ch := make(chan int, 10)
	go func(){
		Lable:
			for{
				select{
				case ch <- rand.Int():
				case <- done:

					break Lable
				}
			}
		close(ch)
	}()

	return ch
}

func GenerateIntB(done chan struct{}) (chan int){
	ch := make(chan int, 10)
	go func(){
		Lable:
			for{
				select{
				case ch <- rand.Int():
				case <- done:
					break Lable
				}
			}
		close(ch)
	}()

	return ch
}

func GenerateInt(done chan struct{}) (chan int){
	ch := make(chan int, 20)
	send := make(chan struct{})
	A := GenerateIntA(send)
	B := GenerateIntB(send)

	go func(){
		Lable:
			for{
				select{
				case ch <- <-A:
				case ch <- <-B:
				case <-done:
					//通知所有的生产者
					close(send)
					break Lable
				}
			}
		
		close(ch)
	}()

	return ch
}

func main(){
	done := make(chan struct{})
	ch := GenerateInt(done)

	for i := 0; i < 100; i++{
		fmt.Println(<-ch)
	}

	close(done)
	fmt.Println(<-ch)
	fmt.Println("NumGoroutine=", runtime.NumGoroutine())

}