package main

import "fmt"

func processSlice(s []int, v int) []int {
	s = append(s, v)
	return s
}

func main() {
	var s1 = []int{}
	var s2 = processSlice(s1, 5)
	fmt.Println(s1, s2)
}
