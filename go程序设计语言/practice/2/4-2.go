package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

const (
	SHA256 int = iota
	SHA384
	SHA512
)

func main() {
	shatypePtr := flag.Int("ttt", SHA256, "type 0-sha256 1-sha384 2-sha512")
	flag.Parse()
	args := flag.Args()
	fmt.Println(args)
	//返回未带flag的其他参数，即向转换的字符串
	if len(args) < 1 {
		fmt.Fprintln(os.Stdout, "please input:")
		return
	}

	switch *shatypePtr {
	case SHA256:
		fmt.Fprintln(os.Stdout, sha256.Sum256([]byte(args[0])))
	case SHA384:
		fmt.Fprintln(os.Stdout, sha512.Sum384([]byte(args[0])))
	case SHA512:
		fmt.Fprintln(os.Stdout, sha512.Sum512([]byte(args[0])))
	default:
		fmt.Fprintf(os.Stdout, "do not support type %d\n", *shatypePtr)
	}
}
