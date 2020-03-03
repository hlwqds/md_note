package main

import "fmt"

const testLen = 10

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse2(s *[testLen]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func appendInt(s []int, e ...int) []int {
	var z []int
	zlen := len(s) + len(e)
	if zlen <= cap(s) {
		z = s[:zlen]
	} else {
		//为了兼容len(s)为0的情况
		zcap := zlen
		if zcap < 2*len(s) {
			zcap = 2 * len(s)
		}
		z = make([]int, zlen, zcap)
		copy(z, s)
	}

	copy(z[len(s):], e)
	return z
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)

	//访问以len为标准，扩展或者创建切片时以cap为标准(因为底层的数组是不能修改的，
	//如果想要扩展，则需要新建一个更大的底层数组，再将原有的数据拷贝到其中)
	slice1 := make([]int, 3, 5)
	slice2 := slice1[1:4]
	slice2[2] = 2
	fmt.Println(slice1)
	slice1 = appendInt(slice1, 8, 7, 5)
	slice1 = appendInt(slice1, 8, 7, 5)

	fmt.Println(slice1)
	fmt.Println(cap(slice1))

	array := [testLen]int{2, 3, 4, 3, 5, 6, 3, 5}
	reverse2(&array)
	fmt.Println(array)

	fmt.Println(array[:0])
}
