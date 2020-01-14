package main
import (
	"fmt"
)

type T struct{
	a int
}

func (t *T) Get() (int){
	return t.a
}

func (t *T) Set(v int) (){
	t.a = v
}

func (t *T) Print(){
	fmt.Printf("%p, %v, %d \n", t, t, t.a)
}

func main()(){
	var t = &T{}
	f := t.Set
	f(2)
	t.Print()

	f(3)
	t.Print()
}