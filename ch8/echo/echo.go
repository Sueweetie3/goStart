package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func main(){
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration){
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn){
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan(){
		echo(c, input.Text(), 1*time.Second)
	}
}
