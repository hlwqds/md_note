package main

type Printer interface {
	Print() ()
}

type S struct{}

func (s S) Print() (){
	println("print")
}

func main() (){
	var i Printer
	//i.Print()
	//为初始化的接口调用会产生panic

	i = S{}
	i.Print()
}