package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, i := range os.Args[1:] {
		s += sep + i
		sep = " "
	}
	fmt.Println(s)
}
