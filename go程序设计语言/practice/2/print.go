package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	o := 0666
	fmt.Printf("%d, %[1]o, %#[1]o\n", o)

	x := int(0xdeadbeef)
	fmt.Printf("%d, %[1]x, %#[1]X\n", x)

	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]c %[1]q\n", newline)

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d, e^x = %e\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)
	fmt.Println(math.Inf(-1), math.Inf(1))
	fmt.Println(math.NaN())

	//NaN的比较结果总不成立
	nan := math.NaN()
	fmt.Println(nan == nan, nan > nan, nan < nan)

	var m complex128 = complex(1, 2)
	var n complex128 = complex(3, 4)

	fmt.Println(m * n)
	fmt.Println(real(m * n))
	fmt.Println(imag(m * n))

	fmt.Println(1i * 1i)

	str := "\u4e16"

	fmt.Println("\u4e16")
	fmt.Println(len(str))
	str = `awddawd\n`
	fmt.Println(str)

	fptr, _ := os.Create("test.txt")
	fmt.Fprintln(fptr, "\u4e16")
	fptr.Close()
}
