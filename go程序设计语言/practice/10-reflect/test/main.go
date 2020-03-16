package main

import (
	"fmt"
	"unsafe"

	format "github.com/my/repo/go程序设计语言/practice/10-reflect/formatReflect"
)

func main() {
	testMap := make(map[string]string)
	testMap["huang"] = "lin"
	fmt.Println(format.Any(testMap))
	fmt.Println(unsafe.Sizeof(int32(1)))
}
