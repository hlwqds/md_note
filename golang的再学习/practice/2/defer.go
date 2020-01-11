package main

import(
	"os"
)

func f() int{
	a := 0
	defer func(i int) (){
		println("defer i=", i) 
	}(a)
	
	a++
	return a
}

func main() (){
	//先进后出
	defer func() (){
		println("first")		
	}()

	defer func() (){
		println("second")
	}()

	println(f())


	println("func body")

	os.Exit(1)

	return

	//在return后企图注册defer，失败
	defer func() (){
		println("register after return")
	}()
}