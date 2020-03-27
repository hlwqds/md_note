package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

//HashString Gen HashValue for string
func HashString(str []byte, hashSize uint) uint {
	var seed uint = 131
	var hash uint = 0
	length := len(str)

	for i := 0; i < length; i++ {
		hash = hash*seed + uint(str[i])
	}

	return (hash & (hashSize - 1))
}

func hashValue(a, b uint, bits uint) uint {
	var hash uint = (((b & 0xF0F0F0F0) >> 4) | ((b & 0x0F0F0F0F) << 4))
	hash = hash ^ (hash >> 16)
	hash = hash ^ a ^ ((a >> 4) | (a << 4))
	hash = (hash ^ (hash >> 8)) & ((1 << bits) - 1)
	return hash
}

//HashStringAB return hashvalue for AB
func HashStringAB(a, b string) uint {
	hash1 := HashString([]byte(a), 1<<16)
	hash2 := HashString([]byte(b), 1<<16)
	return hashValue(hash1, hash2, 16)
}

func main() {
	testMap := make(map[uint]uint)
	for i := 0; i < 1000; i++ {
		A := strconv.Itoa(rand.Int())
		B := strconv.Itoa(rand.Int())
		testMap[HashStringAB(A, B)]++
	}

	fmt.Println(testMap)
}
