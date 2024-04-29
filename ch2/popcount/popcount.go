package main

import (
	"fmt"
	"os"
	"strconv"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(1&i)
	}
}

func popcount(x uint64) int {
	var count int
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

func main() {
	for _, arg := range os.Args[1:] {
		argnumber, _ := strconv.ParseInt(arg, 10, 64)
		fmt.Printf("%d has %d bit 1\n", argnumber, popcount(uint64(argnumber)))
	}
}
