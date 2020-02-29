package main

import (
	"bytes"
	"fmt"
	"strings"
)

func commaInt(s string) string {
	var r bytes.Buffer
	var index int

	l := len(s)
	//如果想要尽可能地减少内存分配，我们就要提前分配好空间并且依次将子串追加到空间中
	//让待分割的字符串长度对3取模，得出第一个子串的位数
	mod := l % 3
	if mod > 0 && l > 3 {
		//将第一个子串写入，这是一个特殊情况，单独讨论，如果第一个串不是唯一的串，则我们追加它，并且后面跟分割符
		r.Write([]byte(s[:mod] + ","))
		index = mod
	}
	for index+3 < l {
		//循环追加子串，直到最后一个子串
		r.Write([]byte(s[mod:mod+3] + ","))
		index += 3
	}

	//最后一个子串，将其追加到尾部
	r.Write([]byte(s[index:l]))
	return r.String()
}

func commaFormatDigit(s string) string {
	var r bytes.Buffer
	l := len(s)
	prefixIndex := 0 //是否有符号前缀

	if l >= 1 {
		//如果长度大于1，我们尝试去除其符号，然后做进一步的处理
		if s[0] == '+' || s[0] == '-' {
			r.Write([]byte(s[:1]))
			prefixIndex = 1
		}
	} else {
		return s
	}

	if dpIndex := strings.IndexRune(s, '.'); dpIndex > 0 {
		r.Write([]byte(commaInt(s[prefixIndex:dpIndex])))
		r.WriteByte('.')
		r.Write([]byte(commaInt(s[dpIndex+1:])))
	} else {
		r.Write([]byte(commaInt(s[prefixIndex:])))
	}

	return r.String()
}

func main() {
	fmt.Println(commaFormatDigit("-3123124124"))
	fmt.Println(1 % 4)
}
