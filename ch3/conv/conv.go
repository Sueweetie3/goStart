package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 100
	y := fmt.Sprintf("%d", x)
	z := strconv.Itoa(x)
	fmt.Println(y, z)
	fmt.Println(strconv.FormatInt(int64(x), 2))
	fmt.Printf("%b\n", x)

	j, _ := strconv.Atoi("100")
	fmt.Println(j)

	i, _ := strconv.ParseInt("100", 10, 32)
	fmt.Println(i)
}
