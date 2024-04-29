package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{ X, Y float64 }
type Path []Point

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func Test() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	dist := p.Distance(q.Point)
	p.ScaleBy(2)
	fmt.Println(dist, p)
}

func main() {
	p := Point{3, 4}
	q := Point{1, 2}
	fmt.Println(p.Distance(q))

	(&p).ScaleBy(2)
	fmt.Println(p)
	p.ScaleBy(2)
	fmt.Println(p)

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
}
