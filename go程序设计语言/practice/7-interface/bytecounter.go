package main

import (
	"bufio"
	"fmt"
)

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello,%s", name)
	fmt.Println(c)

	wordLen := 0
	index := 0
	testString := "hello, every one, nice to\n"
	wordLen, buf, err := bufio.ScanWords([]byte(testString), true)
	index += wordLen
	fmt.Println(wordLen, err)
	fmt.Printf("%s\n", buf)
	wordLen, buf, err = bufio.ScanWords([]byte(testString)[index:], true)
	index += wordLen
	fmt.Println(wordLen, err)
	fmt.Printf("%s\n", buf)
	wordLen, buf, err = bufio.ScanWords([]byte(testString)[index:], true)
	fmt.Println(wordLen, err)
	fmt.Printf("%s\n", buf)
	index += wordLen
	wordLen, buf, err = bufio.ScanWords([]byte(testString)[index:], true)
	fmt.Println(wordLen, err)
	fmt.Printf("%s\n", buf)
	index += wordLen
	wordLen, buf, err = bufio.ScanWords([]byte(testString)[index:], true)
	fmt.Println(wordLen, err)
	fmt.Printf("%s\n", buf)
	index += wordLen
	wordLen, buf, err = bufio.ScanWords([]byte(testString)[index:], true)
	fmt.Println(wordLen, err)
	fmt.Printf("%s\n", buf)
	index += wordLen
}
