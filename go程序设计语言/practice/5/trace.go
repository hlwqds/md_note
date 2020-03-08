package main

import (
	"log"
	"time"
)

func wait() {
	defer trace("wait")()

	time.Sleep(time.Second * 10)
	return
}

func trace(message string) func() {
	start := time.Now()
	log.Printf("enter %s", message)
	return func() { log.Printf("exit %s (%s)", message, time.Since(start)) }
}

func main() {
	wait()
}
