package main

import (
	"io"
	"net"
	"time"
)

func main(){
	listner, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
		
	}

	for {
		conn, err := listner.Accept()
		if err != nil {
			panic(err)
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn){
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}