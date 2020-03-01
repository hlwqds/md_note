package main

import "fmt"

func main() {
	number := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	slice1 := number[5:7]

	//创建切片时endIndex不能超过底层数据的容量
	//以数组创建新的切片
	//slice的容量决定于底层数组的容量，cap = len - startIndex
	fmt.Println(slice1)
	fmt.Println(cap(slice1))

	//在切片的基础上创建新的切片
	//新的slice的容量决定于元slice的容量，cap = len - startIndex
	slice2 := slice1[1:5]
	fmt.Println(slice2)
	fmt.Println(cap(slice2))
	slice2[0] = 16
	fmt.Println(number)
	fmt.Println(slice1)
	fmt.Println(slice2)
}
