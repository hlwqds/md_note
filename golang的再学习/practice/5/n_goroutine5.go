package main

import (
	"sync"
	"time"
	"fmt"
	"github.com/go-redis/redis/v7"
)

type taskStruct struct{
	fd *redis.Client
	result chan int
}

func (t *taskStruct)do(){
	t.fd.Set("key", "value", 0).Err()

	t.result <- 1
}

func InitTask(taskChan chan taskStruct, r chan int, client *redis.Client)(){

	for i := 0; i < 10000000; i++{
		task := taskStruct{
			fd: client,
			result: r,
		}
		taskChan <- task
	}

	close(taskChan)
}

func ProcessingTask(task taskStruct, wait *sync.WaitGroup){
	task.do()
	wait.Done()
}

func distriubteTask(taskChan chan taskStruct, wait *sync.WaitGroup, result chan int)(){
	for v := range taskChan{
		wait.Add(1)
		//fmt.Println(v)
		go ProcessingTask(v, wait)
	}

	wait.Wait()
	close(result)
}

func ProcessResult(resultChan chan int) (int){
	sum := 0
	for v := range resultChan{
		sum += v
	}

	return sum
}

func main(){
	client := redis.NewClient(&redis.Options{
		Addr:     "148.70.52.135:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	taskChan := make(chan taskStruct, 10)
	resultChan := make(chan int, 10)
	wait := &sync.WaitGroup{}
	time1 := time.Now().Unix()
	go InitTask(taskChan, resultChan, client)

	go distriubteTask(taskChan, wait, resultChan)

	sum := ProcessResult(resultChan)
	time2 := time.Now().Unix()
	fmt.Println(time2 - time1)
	fmt.Println(sum)
}