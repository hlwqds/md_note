package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

)

func main() {
	ptr, err := os.Open("EOF.go")
	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewReader(ptr)
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("read failed: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%c", r)
	}
}
