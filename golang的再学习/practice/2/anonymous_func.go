package main

import (
	"fmt"
)

var sum = func(a, b int) int{
	return a + b
}

type Aa func(int, int) (int)

func doinput(f Aa, a, b int) (int){
	return f(a, b)
}

//匿名函数作为返回值
func wrap(op string) (Aa){
	switch op{
	case "add":
		return func(a, b int) (int){
			return a + b
		}
	case "sub":
		return func(a, b int) (int){
			return a - b
		}
	default:
		return nil
	}
}

func main() (){
	//直接调用匿名函数
	defer func() (){
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()

	sum(1, 2)

	//匿名函数作为实参
	doinput(func(x, y int) (int){
		return x + y
	}, 3, 4)

	opFunc := wrap("add")
	result := opFunc(2, 3)

	fmt.Println(result)
}