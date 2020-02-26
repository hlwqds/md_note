package main

import (
	"fmt"
)

type test struct{
	str string
	number int
	list []string
}

func main(){
	var v test = test{
		str: "sda",
		number: 11,
		list: []string{"ada", "sa"},
	}

	fmt.Println(v)

	var a []string
	fmt.Println(a == nil)
}