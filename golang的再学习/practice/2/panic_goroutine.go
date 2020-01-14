package main

import (
	"time"
	"fmt"
)

func da() (){
	defer func() (){
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()

	panic("panic da")
	for i := 0; i < 10; i++{
		fmt.Println(i)
	}
}

func db() (){
	for i := 0; i < 10; i++{
		fmt.Println(i)
	}
}

func do() (){
	defer func() (){
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()

	go da()
	go db()
	time.Sleep(3 * time.Second)
}

func main() (){
	do()
}