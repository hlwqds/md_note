package main
import (
	"fmt"
	"strconv"
)

func main() {
	x := 12333
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	//FormatInt可以将int型按传递的进制数转换为对应的字符串
	fmt.Println(strconv.FormatInt(int64(x), 13))
}