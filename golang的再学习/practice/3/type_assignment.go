package main

import (
	"fmt"
)

type Map map[string] string

func (m Map) Print() {
	for _, v := range m{
		fmt.Println(v)
	}
}

type iMap Map

func (m iMap) Print() {
	for _, v := range m{
		fmt.Println(v)
	}
}

type slice []int
func (s slice) Print() {
	for _, v := range s{
		fmt.Println(v)
	}
}

func main() (){
	mp := make(map[string]string, 10)
	mp["hi"] = "huanglin"

	var mb Map = mp
	var ma iMap = mp

	ma.Print()
	mb.Print()

	var i interface{
		Print()
	} = ma
	i.Print()
	//它居然也能打印，数据是哪里来的
    //im与ma虽然拥有相同的底层类型，但是二者中没有一个是字面量类型，不能直接赋值，可以进行强制类型转换
    //var im iMap = ma
	var im iMap = (iMap) (ma)
	im.Print()
	s1 := []int{1,2,3}
	var s2 slice
	s2 = s1
	s2.Print()

	s := "hello,世界"
	var a []byte
	a = ([]byte) (s)
	var b string
	b = (string) (a)
	var c []rune
	c = ([]rune) (s)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}