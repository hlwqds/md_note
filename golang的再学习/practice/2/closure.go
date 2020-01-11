package main
var (
	b = 0
)

func fa(a int) (func(i int) (int)){
	return func(i int) (int){
		println(&a, a)
		a = a + i
		return a
	}
}

func da() (func(i int) (int)){
	return func(i int) (int){
		println(&b, b)
		b = b + i
		return b
	}
}

func main() (){
	f := fa(1)
	//g引用的外部的闭包环境包括本次函数调用的形参a的值1
	g := fa(1)
	//g引用的外部的闭包环境包括本次函数调用的形参a的值1
	//此时f,g引用的闭包环境中的a的值并不是同一个， 而是两次函数调用产生的副本

	println(f(1))
	println(f(1))

	println(g(1))
	println(g(1))

	println(&b, b)
	h := da()
	i := da()
	println(h(1))
	println(h(1))

	println(i(1))
	println(i(1))
}