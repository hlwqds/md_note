package main

import (
	"fmt"
)

func recusionMulti(list []int, index int) []int {
	length := len(list)
	if length <= index {
		return list
	} else {
		tmp := list[index]
		for i := 0; i < length; i++ {
			list[i] *= tmp
		}
		return recusionMulti(list, index+1)
	}
}

func recusionMulti2(list []int, task []int) []int {

	if len(task) <= 0 {
		return list
	} else {
		length := len(list)
		for i := 0; i < length; i++ {
			list[i] *= task[0]
		}
		return recusionMulti2(list, task[1:])
	}
}

func main() {
	list := []int{2, 3, 7, 8, 10}
	list = recusionMulti(list, 0)
	fmt.Println(list)

	list2 := []int{2, 3, 7, 8, 10}
	tmp := []int{}
	for i := 0; i < len(list2); i++ {
		tmp = append(tmp, list2[i])
	}
	list2 = recusionMulti2(list2, tmp)
	fmt.Println(list2)
}
