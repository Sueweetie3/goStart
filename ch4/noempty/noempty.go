package main

import "fmt"

func nonempty(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	input := []string{"one", "", "three"}
	input = nonempty(input) // re-assign slice, the key of slice is **don't care much about origin array**
	fmt.Println(input)
}
