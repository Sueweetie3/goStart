package main

import "fmt"

// slice是数组一部分的别名，不是数组一部分的引用
func processSlice(s []int, v int) []int {
	s = append(s, v)
	return s
}

func main() {
	var s0 []int = nil
	var s1 = []int{}
	var s2 = processSlice(s1, 5)
	var s3 = make([]int, 0, 2)
	var s4 = processSlice(s3, 5)
	fmt.Println(s0, s0 == nil, s1, s1 == nil)
	fmt.Println(s1, s2)
	fmt.Println(s3, s4)
}
