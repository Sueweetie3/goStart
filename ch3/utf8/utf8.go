package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	fmt.Println(len(s))                    // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"
	n := 0
	for range s {
		n++
	}
	fmt.Println(n) // "9"

	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

}
