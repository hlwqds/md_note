package main

import "fmt"

func sum(num []int) int {
	if len(num) == 0 {
		return 0
	} else {
		return num[0] + sum(num[1:])
	}
}

func max(num []int, cmp int) int {
	if len(num) == 0 {
		return cmp
	} else {
		if num[0] > cmp {
			cmp = num[0]
		}
		return max(num[1:], cmp)
	}
}

func main() {
	num := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(num))
	fmt.Println(max(num[1:], num[0]))
}
