package main

import "fmt"

func changeFun(f func(int) int) func(int) int {
	f = func(x int) int {
		return x
	}
	return f
}

func main() {
	fp := func(x int) int {
		return x + 1
	}
	fp2 := changeFun(fp)
	fmt.Println(fp(4))
	fmt.Println(fp2(4))
}
