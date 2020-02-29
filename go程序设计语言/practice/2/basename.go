package main

import "fmt"

//base移除路径部分和.后缀
func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func main() {
	str := "/mod/hlwqds/test.txt"
	str = basename(str)
	fmt.Println(str)
}
