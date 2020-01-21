package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string "学生姓名"
	Age  int    `a:"111"b:"3333"`
}

func main() {
	s := Student{}
	rt := reflect.TypeOf(s)
	fieldName, ok := rt.FieldByName("Name")

	//取tag数据
	if ok {
		fmt.Println(fieldName.Tag)
	}
	fieldAge, ok2 := rt.FieldByName("Age")
	//可以像JSON一样，取tag里的数据，多个tag之间没有逗号，tag不需要引号
	if ok2 {
		fmt.Println(fieldAge.Tag.Get("a"))
		fmt.Println(fieldAge.Tag.Get("b"))
	}
}
