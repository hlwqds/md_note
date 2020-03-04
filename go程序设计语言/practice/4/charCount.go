package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utfLen [utf8.UTFMax + 1]int
	invalid := 0
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount err:%v\n", err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		counts[r]++
		utfLen[n]++
	}

	fmt.Printf("rune\tcount\t\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\t\n", c, n)
	}

	fmt.Printf("\nlen\tcount\n")
	for i, n := range utfLen {
		fmt.Printf("%d\t%d\t\n", i, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
