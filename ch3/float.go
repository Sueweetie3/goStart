package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)
	fmt.Println(cmplx.Sqrt(-1))
}
