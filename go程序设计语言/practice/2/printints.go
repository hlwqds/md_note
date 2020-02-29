package main

import (
	"bytes"
	"fmt"
)

//比起string来说[]byte更为高效，因为在赋值过程中[]byte不会在生成新的底层数组

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteByte(',')
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	buf.WriteRune('话')
	buf.WriteRune(']')
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3}))
}
