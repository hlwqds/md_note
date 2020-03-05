package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

)

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Get(url)
		if err == nil {
			return nil
		}
		log.Printf("Server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
	log.SetPrefix("wait:")
	log.SetFlags(0)
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	err := WaitForServer(os.Args[1])
	if err != nil {
		log.Fatal("Site is down:%v\n", err)
	}
	return
}
