### 第一章

整型和字符转之间的强制转换

```go
b,error := strconv.Atoi(a)
d := strconv.Itoa(c)   //数字变成字符串
//强制转换函数已包括在strconv包中
```



p6

```go
找出两种程序的性能差距
```



从map中获取的值是副本

```go
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
```

