package main

import "fmt"

func quickSort(list []int) {
	length := len(list)
	if length <= 1 {
		return
	} else {
		last := 1
		base := length / 2
		list[0], list[base] = list[base], list[0]
		for i := 1; i < length; i++ {
			if list[i] < list[last] {
				list[i], list[last] = list[last], list[i]
				last++
				list[i], list[last] = list[last], list[i]
			}
			list[0], list[last-1] = list[last-1], list[0]
		}

		quickSort(list[0 : last-1])
		quickSort(list[last:])
	}
}

func main() {
	list := []int{4, 2, 6, 8, 1, 2, 8}
	quickSort(list)
	fmt.Println(list)
}
