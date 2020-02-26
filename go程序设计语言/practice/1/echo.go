//package main
package main

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main(){
	var s, sep string
	for i := 1; i < len(os.Args); i++{
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	k, kep := "", ""

	for _, arg := range os.Args[1:]{
		k += kep + arg
		kep = " "
	}

	fmt.Println(k)

	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])

	fmt.Println(os.Args[0])

	k, kep = "", "\n"
	for i, arg := range os.Args[1:]{
		k += strconv.Itoa(i + 1) + ". " + arg + kep

	}
	fmt.Printf(k)
}