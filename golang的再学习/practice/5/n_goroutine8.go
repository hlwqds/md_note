package main

import (
	"fmt"
	"time"
	"context"
)

type otherContext struct{
	context.Context
}

func work(ctx context.Context, name string){
	for{
		select{
		case <- ctx.Done():
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			fmt.Printf("%s is running \n", name)
			time.Sleep(1 * time.Second)
		}
	}
}

func workWithValue(ctx context.Context, name string){
	for{
		select{
		case <- ctx.Done():
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			value := ctx.Value("key").(string)
			fmt.Printf("%s is running value=%s\n", name, value)
			time.Sleep(1 * time.Second)
		}
	}
}

func main(){

	//使用context.Background()构建一个WithCancel类型的上下文
	ctxa, cancel := context.WithCancel(context.Background())
	//work模拟运行并检测前端的推出通知
	go work(ctxa, "work1")

	//使用WithDeadline包装前面的上下文对象
	tm := time.Now().Add(3 * time.Second)
	//ctxb为ctxa的子节点
	ctxb, cancel2 := context.WithDeadline(ctxa, tm)

	go work(ctxb, "work2")

	oc := otherContext{ctxb}
	//ctxc为oc即ctxb的子节点
	ctxc := context.WithValue(oc, "key", "andes, pass from main ")

	go workWithValue(ctxc, "work3")

	//故意等待10s，让work2、work3超时退出
	cancel2()
	time.Sleep(10 * time.Second)

	//显示调用ctxa的cancel
	cancel()
	time.Sleep(5 * time.Second)
}