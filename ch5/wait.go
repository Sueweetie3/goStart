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
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("sever not responding %s; retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to responde after %s", url, timeout)
}

func main() {
	url := os.Args[1]
	if err := WaitForServer(url); err != nil {
		log.Printf("Site is down: %v", err)
		os.Exit(1)
	}
}
