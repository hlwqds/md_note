package main

import "fmt"

var list [100]int
var times int = 0

func init() {
	length := len(list)
	for i := 0; i < length; i++ {
		list[i] = i
	}
}

func binarySearch(num int, left, right int) int {
	b := (left + right) / 2
	if num == list[b] {
		return b
	} else if num > list[b] {
		left = b + 1
	} else {
		right = b - 1
	}

	return binarySearch(num, left, right)
}

func main() {
	var a int = 88
	index := binarySearch(a, 0, 99)
	fmt.Println(index)
}
