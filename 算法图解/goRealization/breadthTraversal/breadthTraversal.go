package main

import (
	"fmt"
	"strings"
)

type personInfo struct {
	friends []string
	visited bool
}

func main() {
	chart := make(map[string]personInfo)
	chart["you"] = personInfo{
		friends: []string{"alice", "bob", "claire"},
	}
	chart["bob"] = personInfo{
		friends: []string{"anuj", "peggy"},
	}
	chart["aclice"] = personInfo{
		friends: []string{"peggy"},
	}
	chart["claire"] = personInfo{
		friends: []string{"thom", "jonny"},
	}
	chart["anuj"] = personInfo{
		friends: []string{},
	}
	chart["peggy"] = personInfo{
		friends: []string{},
	}
	chart["thom"] = personInfo{
		friends: []string{},
	}
	chart["jonny"] = personInfo{
		friends: []string{},
	}

	taskCh := make(chan []string, 1)
	resultCh := make(chan []string, 1)
	person := chart["you"].friends
	taskCh <- person
	answerCh := make(chan string)
	producer := make(chan struct{})
	consumer := make(chan struct{})

	go func() {
		defer close(producer)
		for friends := range resultCh {
			if len(friends) != 0 {
				taskCh <- friends
			} else {
				close(taskCh)
				close(answerCh)
			}
		}
	}()

	go func() {
		defer close(resultCh)
		defer close(consumer)
		for people := range taskCh {
			toBeVisisted := []string{}
			for _, person := range people {
				personinfo := chart[person]
				for _, friend := range personinfo.friends {
					if strings.HasSuffix(friend, "adawd") {
						answerCh <- friend
						return
					}
					if !personinfo.visited {
						toBeVisisted = append(toBeVisisted, personinfo.friends...)
						personinfo.visited = true
					}
				}
			}
			fmt.Println(toBeVisisted)
			resultCh <- toBeVisisted
		}
	}()

	fmt.Println(<-answerCh)

	<-consumer
	<-producer

	return
}
