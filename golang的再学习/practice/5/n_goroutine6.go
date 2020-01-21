package main

import (
	"fmt"
)

const (
	NUMBER = 10
)

type task struct{
	begin int
	end int
	result chan int
}

func (t *task)do(){
	sum := 0
	for i := t.begin; i <= t.end; i++{
		sum += i
	}

	t.result <- sum
}

func InitTask(taskChan chan task, r chan int, p int)(){
	qu := p / 10
	mod := p % 10
	high := qu *10

	for j := 0; j < qu; j++{
		b := 10 * j + 1
		e := 10 * (j + 1)
		task := task{
			begin: b,
			end: e,
			result: r,
		}
		taskChan <- task
	}

	if mod != 0{
		task := task{
			begin: high + 1,
			end: p,
			result: r,
		}
		taskChan <- task
	}

	close(taskChan)
}

func ProcessTask(taskChan chan task, done chan struct{}){
	for t := range taskChan{
		t.do()
	}

	done <- struct{}{}
}

func DistirbuteTask(taskChan chan task, workers int, done chan struct{}){
	for i := 0; i < workers; i++{
		go ProcessTask(taskChan, done)
	}
}

func CloseResult(done chan struct{}, resultChan chan int, workers int){
	for i := 0; i < workers; i++{
		<- done
	}
	close(done)
	close(resultChan)
}

func ProcessResult(resultChan chan int)(int){
	sum := 0
	for r := range resultChan{
		sum += r
	}

	return sum
}

func main(){
	workers := NUMBER

	taskChan := make(chan task, 10)
	resultChan := make(chan int, 10)
	done := make(chan struct{}, 10)

	go InitTask(taskChan, resultChan, 100)
	go DistirbuteTask(taskChan, workers, done)
	go CloseResult(done, resultChan, workers)
	sum := ProcessResult(resultChan)
	fmt.Println(sum)
}