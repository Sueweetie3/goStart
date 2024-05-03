package main

import (
	"fmt"
	"os"
	"time"
)

func main(){
	fmt.Println("Commencing countdown.")
	abort := make(chan struct{})
	go func(){
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}();
	go func(){
		tick := time.Tick(1 * time.Second)
		for countdown := 10; countdown >0; countdown--{
			fmt.Println(countdown)
			<-tick
		}
	}()
	select{
		case <-abort:
			fmt.Println("Launch aborted!")
			return;
		case <-time.After(10 * time.Second):
	}
	fmt.Println("Launch!")
}