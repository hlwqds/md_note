package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress message")
var done = make(chan struct{})

//限制目录并发数
var sema = make(chan struct{}, 20)

func main() {
	flag.Parse()
	var n sync.WaitGroup
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	for _, root := range roots {
		go duSingleDir(root, &n)
		n.Add(1)
	}
	n.Wait()
}

func duSingleDir(dir string, w *sync.WaitGroup) {
	defer w.Done()
	ch := make(chan int64)
	var n sync.WaitGroup
	go WalkDir(dir, &n, ch)
	n.Add(1)
	go func() {
		n.Wait()
		close(ch)
	}()

	var tick <-chan time.Time
	if *verbose {
		ticker := time.NewTicker(500 * time.Millisecond)
		tick = ticker.C
	}

	var nfiles, nbytes int64

loop:
	for {
		select {
		case <-done:
			//耗尽已有的数据以使所有的工作函数都能够正常关闭，
			//不会阻塞在发送消息到ch
			for range ch {

			}
			fmt.Println("Cancel")
			return
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		case size, ok := <-ch:
			if !ok {
				break loop
				//标签话的语句可以跳出select和for两层循环
				//当ch关闭时会返回false
			}
			nfiles++
			nbytes += size
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("There are %v files and %v Mb at all.\n", nfiles, float64(nbytes)/1e6)
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func dirents(dir string) []os.FileInfo {
	//获取令牌
	select {
	case sema <- struct{}{}:
	case <-done:
		//取消通知
		return nil
	}

	defer func() {
		//释放令牌
		<-sema
	}()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func WalkDir(dir string, w *sync.WaitGroup, ch chan<- int64) {
	defer w.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			w.Add(1)
			WalkDir(subDir, w, ch)
		} else {
			ch <- entry.Size()
		}
	}
}
