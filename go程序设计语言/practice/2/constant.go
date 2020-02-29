package main

import (
	"fmt"
	"math"
)

type Weekday int

const (
	Sunday Weekday = 1 << iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	const pi = 3.1415926
	const test int = 10

	const (
		a = 1
		b
		c = 2
		d
	)
	math.Pi
	fmt.Println(Saturday)
}
