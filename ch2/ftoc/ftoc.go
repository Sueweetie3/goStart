package main

import "fmt"

func main() {
	var f1, f2 = 32.0, 123.8
	fmt.Printf("%g F = %g C \n", f1, ftoc(f1))
	fmt.Printf("%g F = %g C \n", f2, ftoc(f2))
}

func ftoc(f float64) float64 {
	return (f - 32) * 5 / 9
}
