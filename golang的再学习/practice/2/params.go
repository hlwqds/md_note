package main

import (
	"fmt"
)

func chvalue(a int) (int){
	a = a + 1
	return a
}

func chpointer(a *int) (){
	*a = *a + 1
	return
}

func main() (){
	var a int = 10
	chvalue(a)
	fmt.Println(a)

	chpointer(&a)
	fmt.Println(a)
}