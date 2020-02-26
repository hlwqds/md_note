package main

import (
	"strings"
	"io"
	"os"
	"io/ioutil"
	"net/http"
	"fmt"
)

func main(){
	for _, url := range os.Args[1:]{
		if !strings.HasPrefix(url, "http://"){
			url = "http://" + url
		}
		
		resp, err := http.Get(url)
		if err != nil{
			fmt.Printf("fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil{
			fmt.Printf("ReadAll: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)

		resp, err = http.Get(url)
		if err != nil{
			fmt.Printf("fetch: %v\n", err)
			os.Exit(1)
		}
		/*
		for{
			n, err := io.CopyN(os.Stdout, resp.Body, 4096 * 4)
			if err != nil && n < 4096 * 4{
				fmt.Printf("Copy: %d %v\n", n, err)
				break
			}
		}
		*/
		for{
			n, err := io.Copy(os.Stdout, resp.Body)
			if err != nil || n == 0{
				fmt.Printf("Copy: %d %v\n", n, err)
				break
			}
		}
		resp.Body.Close()
		fmt.Printf("%s\n", resp.Status)
	}
}