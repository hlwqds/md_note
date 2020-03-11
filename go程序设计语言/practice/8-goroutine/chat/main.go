package main

import (
	"log"
	"net"
)

func main(){
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil{
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil{
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}

type client chan<- string
var (
	entering = make(chan client)
	leaving = make(chan client)
	message = make(chan string)
)
func broadcaster(){
	clients := make(map[client]bool)
	for {
		select{
		case msg:= <-message:
			//把所有接受的消息广播给所有的客户
			//发送消息通道
			for cli := range clients{
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

