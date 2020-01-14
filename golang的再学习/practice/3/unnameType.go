package main

import (
	"fmt"
)

type Person struct{
	name string
	age int
}

func main() (){
	a := struct{
		name string
		age int
	}{"huanglin1", 18}

	fmt.Printf("%T\n", a)
	fmt.Printf("%v\n", a)

	b := Person{"huanglin2", 24}
	fmt.Printf("%T\n", b)
	fmt.Printf("%v\n", b)
}