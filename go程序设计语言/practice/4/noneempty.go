package main

import (
	"fmt"
	"unicode"
)

func noneempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func removeDupStringNearby(strings []string) []string {
	i := 0
	var prevEle string = ""
	for _, s := range strings {
		if s != prevEle {
			strings[i] = s
			prevEle = s
			i++
		}
	}

	return strings[:i]
}

func replaceSpace(s []byte) []byte {
	i := 0

	var prevEle byte = 0

	for _, b := range s {
		//unicode空白字符的utf-8实现只需要一个字节就行，可以遍历byte找出
		if unicode.IsSpace(rune(b)) {
			if b != prevEle {
				s[i] = ' '
				i++
			}
		} else {
			s[i] = b
			i++
		}
		prevEle = b
	}
	return s[:i]

}

func remove(strings []string, i int) []string {
	copy(strings[i:], strings[i+1:])
	return strings[:len(strings)-1]
}

func main() {
	strings := []string{"dss", "", "sdad"}
	fmt.Println(strings)
	fmt.Println(len(strings))
	strings = noneempty(strings)
	fmt.Println(strings)
	fmt.Println(len(strings))
	strings = remove(strings, len(strings)-1)
	fmt.Println(strings)
	fmt.Println(len(strings))

	strings2 := []string{"aa", "bb", "bb", "bb", "aa", "aa"}
	strings2 = removeDupStringNearby(strings2)
	fmt.Println(strings2)

	strings3 := "  dawd   awd\n\n\t\tda"
	fmt.Println(strings3)
	strings3 = string(replaceSpace([]byte(strings3)))
	fmt.Println(strings3)
}
