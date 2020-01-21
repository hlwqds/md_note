package main

import (
	"fmt"
	"time"
)

type query struct{
	sql chan string
	result chan string
}

func execQuery(q query)(){
	go func(){
		sql := <-q.sql

		q.result <- "result from " + sql
		close(q.result)
	}()
}

func main(){
	q := query{
		sql: make(chan string, 1),
		result: make(chan string ,1),
	}

	execQuery(q)
	q.sql <- "select *from tables"
	close(q.sql)
	
	//do something
	time.Sleep(1 * time.Second)

	fmt.Println(<-q.result)
}