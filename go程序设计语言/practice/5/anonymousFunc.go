package main

import "fmt"

func square() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := square()
	p := square()
	fmt.Println(f())
	fmt.Println(p())
	fmt.Println(f())
	fmt.Println(p())
	fmt.Println(f())
	fmt.Println(p())
	fmt.Println(f())
	fmt.Println(p())
}
