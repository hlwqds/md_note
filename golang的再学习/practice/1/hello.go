package main

import "fmt"

type Person struct{
	Name string
	Age int
}

type Student struct{
	*Person
	Number int
}


func main(){
	p := &Person{
		Name: "huanglin",
		Age: 24,
	}

	s := Student{
		Person: p,
		Number: 110,
	}
	fmt.Printf("%s\n",s.Name)

	L1:
		fmt.Println("L1")
		return


	if true{
		fmt.Println("haha")
		goto L1
	}

}
