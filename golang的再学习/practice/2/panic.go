package main

import (
	"fmt"
)

func except() {
    recover()
}

func test(){
	//这个会捕获失败
	defer recover()

	//这个会捕获失败
	defer fmt.Println(recover())

	//这个嵌套两层也会捕获失败
	defer func(){
		func(){
			println("defer inneer")
			recover()
		}()
	}()

	//如下的场景会被捕捉成功
	defer func() {
		println("defer inner")
		recover()
	}()

	defer except()
	
	panic("test panic")
	println("hello")
}

func main() (){
	defer println("skip")
	test()
	println("Be skipped")
}