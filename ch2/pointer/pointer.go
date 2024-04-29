package main

import "fmt"

func main() {
	x := 1
	p := &x
	fmt.Println(*p)
	*p++
	fmt.Println(x)

	var z, y int
	fmt.Println(&z == &z, &z == &y, &z == nil)
}
