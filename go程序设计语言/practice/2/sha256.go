package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T", c1, c2, c1 == c2, c1)

	c := [32]byte{1, 2, 3, 4, 5}
	fmt.Println(c)
	zero(&c)
	fmt.Println(c)
}

func zero(ptr *[32]byte) {
	*ptr = [32]byte{}
}
