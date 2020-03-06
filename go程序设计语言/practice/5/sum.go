package main

import (
	"fmt"
	"strings"
)

func sum(vals ...int) int {
	total := 0
	for _, v := range vals {
		total += v
	}

	return total
}

func sum2(vals []int) int {
	total := 0
	for _, v := range vals {
		total += v
	}

	return total
}

func join(sep string, a ...string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return a[0]
	}
	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(a[0])
	for _, s := range a[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}
	return b.String()
}

func main() {
	list := []int{}
	fmt.Println(sum())
	fmt.Println(sum(list...))
	fmt.Printf("%T\n", sum)
	fmt.Printf("%T\n", sum2)
	//变长函数和变量为切片的函数类型是不同的

	fmt.Println(join(" ", "huang", "lin"))
}
