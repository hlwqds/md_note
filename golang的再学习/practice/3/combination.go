package main

type X struct{
	a int
}

type Y struct{
	X
	b int
}

type Z struct{
	Y
	c int
}

func main() (){
	x := X{a:1}
	y := Y{
		X: x,
		b: 2,
	}
	z := Z{
		Y: y,
		c: 3,
	}
	println(z.a, z.Y.a, z.Y.X.a)

	z = Z{}
	z.a = 2
	println(z.a, z.Y.a, z.Y.X.a)
}