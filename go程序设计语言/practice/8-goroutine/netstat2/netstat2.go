package main

import (
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
func main() {
	/*
		ip := []byte{172, 0, 0, 1}
		localhost := net.TCPAddr{IP: ip, Port: 8000}
		conn, err := net.DialTCP("tcp", nil, &localhost)
		if err != nil {
			log.Fatal(err)
		}
	*/
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	tcpConn := conn.(*net.TCPConn)
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, tcpConn)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(tcpConn, os.Stdin)
	tcpConn.CloseWrite()
	<-done
}
