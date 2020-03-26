package main

import (
	"fmt"
	"math/rand"
)

func quickSort(list []int) {
	length := len(list)
	if length <= 1 {
		return
	} else {
		last := 0
		base := length / 2

		list[0], list[base] = list[base], list[0]
		for i := 1; i < length; i++ {

			if list[i] <= list[0] {
				//当list[i] <= list[0]时，我们可以确定这次找到了一个比基准值还小的数
				//将位置预留给他
				last++
				if i > last {
					//如果需要交换，即这次找的数不在预留的位置上，交互位置
					list[i], list[last] = list[last], list[i]
				}
			}
		}

		//将基准值和最后一个比他小的值交换
		list[0], list[last] = list[last], list[0]

		quickSort(list[0:last])
		quickSort(list[last+1:])
	}
}

func main() {
	list := []int{}
	for i := 0; i < 100000; i++ {
		list = append(list, rand.Int())
	}
	quickSort(list)
	fmt.Println(list)

	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			fmt.Println("error")
		}
	}
}
