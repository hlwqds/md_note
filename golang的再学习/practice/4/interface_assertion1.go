package main

import (
	"fmt"
)

type Inner interface {
	Ping()
	Pong()
}

type Outer interface {
	Inner
	Print()
}

type St struct {
	Name string
}

func (s St)Ping(){
	println("Ping")
}

func (s St)Pong(){
	println("Pong")
}

func main() (){
	var i interface{} = St{Name: "huanglin"}
	
	if o, ok := i.(Inner); ok{
		o.Ping()
		o.Pong()
	}

	if o, ok := i.(Outer); ok{
		o.Ping()
		o.Pong()
		o.Print()
	}

	if o, ok := i.(St); ok{
		o.Ping()
		o.Pong()
		fmt.Printf("%s\n", o.Name)
	}
}