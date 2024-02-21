package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	var name = "Bob"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}
