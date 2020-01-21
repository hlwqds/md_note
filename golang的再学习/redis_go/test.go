package main

import (
	"github.com/go-redis/redis/v7"
	"fmt"
	"time"
)

func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "148.70.52.135:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	done := make(chan struct{}, 10)

	time1 := time.Now().Unix()
	
	for i := 0; i < 10000; i++{
		go func(client *redis.Client){
			client.Set("key", "value", 0).Err()

			done <- struct{}{}
		}(client)
	}

	for i := 0; i < 10000; i++{
		<- done
	}
	time2 := time.Now().Unix()
	fmt.Println(time2 - time1)

	// Output: PONG <nil>
}

func main(){
	ExampleNewClient()
}