package main
import (
	"fmt"
)

func add(a, b int) (int){
	return a + b
}

func sub(a, b int) (int){
	return a - b
}

type op func(int, int) (int)

func do(f op,a int,b int) (int){
	return f(a, b)
}

func main() (){
	fmt.Printf("%T\n", add)
	fmt.Println(do(add, 1, 2))
	fmt.Println(do(sub, 1, 2))
}