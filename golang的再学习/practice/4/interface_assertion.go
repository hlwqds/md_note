package main
import (
	"fmt"
)

type Inter interface {
	Ping()
	Pong()
}

type Anter interface {
	Inter
	String()
}

type St struct {
	Name string
}

func (s St) Ping(){
	println("ping")
}

func (s St) Pong(){
	println("pong")
}

func main() (){
	st := &St{Name: "huanglin"}
	var i interface{} = st
	o := i.(Inter)
	o.Ping()
	o.Pong()

//	p := i.(Anter)
//	p.String()

	s := i.(*St)
	s.Ping()
	s.Pong()
	fmt.Printf("%s\n", s.Name)
}