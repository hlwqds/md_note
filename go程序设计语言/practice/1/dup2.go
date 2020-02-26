package main
import (
	"bufio"
	"fmt"
	"os"
)

type lineInfo struct{
	len int
	fileMap map[string]int
}

func main(){
	counts := make(map[string]lineInfo)
	files := os.Args[1:]
	if len(files) == 0{
		countlines(os.Stdin, "stdin", counts)
	}else{
		for _, arg := range files{
			f, err := os.Open(arg)
			if err != nil{
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countlines(f, arg, counts)
			f.Close()
		}
	}

	fmt.Printf("%v\n", counts)

	for line, lineInfo := range counts{
		if lineInfo.len > 1{
			fmt.Printf("%d\t%s\t%v\n", lineInfo.len, line, lineInfo.fileMap)
		}
	}
}

func countlines(f *os.File, name string, counts map[string]lineInfo){
	inputs := bufio.NewScanner(f)
	for inputs.Scan(){
		lineinfo, ok := counts[inputs.Text()]
		if ok{
			//lineinfo.len++ 返回的是结构副本
			lineinfo.fileMap[name]++
			lineinfo.len++
		}else{
			filemap := make(map[string]int)
			filemap[name] = 1
			lineinfo = lineInfo{
				len: 1,
				fileMap: filemap,
			}
		}
		counts[inputs.Text()] = lineinfo
	}
}