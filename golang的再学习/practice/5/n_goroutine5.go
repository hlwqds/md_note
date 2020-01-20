package main

import (
	"sync"
	"time"
	"fmt"
)

type taskStruct struct{
	begin int
	end int
	result chan int
}

func (t *taskStruct)do(){
	sum := 0
	//fmt.Println(t)
	for i := t.begin; i < t.end; i++{
		sum += i
	}
	t.result <- sum
}

func InitTask(taskChan chan taskStruct, r chan int, p int)(){
	qu := p / 10
	mod := p % 10
	high := qu *10

	for j := 0; j < qu; j++{
		b := 10 * j + 1
		e := 10 * (j + 1)
		task := taskStruct{
			begin: b,
			end: e,
			result: r,
		}
		taskChan <- task
	}

	if mod != 0{
		task := taskStruct{
			begin: high + 1,
			end: p,
			result: r,
		}
		taskChan <- task
	}
	time.Sleep(5 * time.Second)
	fmt.Println("close to break loop")
	close(taskChan)
}

func ProcessingTask(task taskStruct, wait *sync.WaitGroup){
	task.do()
	wait.Done()
}

func distriubteTask(taskChan chan taskStruct, wait *sync.WaitGroup, result chan int)(){
	fmt.Printf("%T\n", taskChan)
	for v := range taskChan{
		wait.Add(1)
		//fmt.Println(v)
		go ProcessingTask(v, wait)
	}

	fmt.Println("通道关闭我才会退出循环")
	wait.Wait()
	close(result)
}

func ProcessResult(resultChan chan int) (int){
	sum := 0
	for v := range resultChan{
		fmt.Println(v)
		sum += v
	}

	return sum
}

func main(){
	taskChan := make(chan taskStruct, 10)
	resultChan := make(chan int, 10)
	wait := &sync.WaitGroup{}

	go InitTask(taskChan, resultChan, 74)

	go distriubteTask(taskChan, wait, resultChan)

	sum := ProcessResult(resultChan)
	println(sum)
}