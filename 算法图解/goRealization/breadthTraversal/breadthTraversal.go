package main

import (
	"sync"
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

	var wg sync.WaitGroup
	task := make(chan []string)
	result := make(chan []string)

	go func() {
		var friends []string
		person := chart["you"].friends
		for {
			task <- person
			wg.Add(1)
			friends := <-result
			for _, friend := range friends {

			}
		}
	}()
}
