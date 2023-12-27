package main

import "fmt"

func main() {
	a1 := []string{1: "one", 2: "two"}
	a2 := [3]string{"", "one", "two"}
	a3 := [...]string{"", "one", "two"}
	fmt.Println(a1, a2, a3, a2 == a3) // can't compare a1 == a2

	var m0 map[string]int
	m1 := make(map[string]int)
	m2 := map[string]int{}
	fmt.Println(m0, m1, m2, m0 == nil, m1 == nil, m2 == nil, len(m0))

	type Point struct {
		X, Y int
	}

	p := Point{1, 2}
	p1 := Point{X: 1}
	fmt.Println(p, p1)
}
