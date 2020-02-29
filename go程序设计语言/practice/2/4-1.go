package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("huanglin"))

	fmt.Println(byteCount(&c1))
}

func byteCount(s *[sha256.Size]byte) int {
	byteMap := map[byte]int{}
	for _, v := range s {
		byteMap[v]++
	}
	fmt.Println(byteMap)
	return len(byteMap)
}
