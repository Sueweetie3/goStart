package main

import "fmt"

func rev(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	emptySlice := a[:0]
	// emptySlice := a[:][6:]
	var nilSlice []int
	fmt.Println(emptySlice, nilSlice, emptySlice == nil, nilSlice == nil)
	rev(a[:])
	fmt.Println(a)

}
