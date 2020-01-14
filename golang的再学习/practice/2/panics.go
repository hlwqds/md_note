package main

import (
	"fmt"
)

func main() (){
	defer func() (){
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()

	defer func() (){
		if err := recover(); err != nil{
			fmt.Println(err)
		}
		panic("first defer panic")
	}()

	defer func() (){
		panic("second defer panic")
	}()

	println("func body")
}