package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func Distance(p, q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func main() {
	q := Point{}
	p := Point{1, 2}
	fmt.Println(Distance(q, p))
	fmt.Println(q.Distance(p))
	fmt.Println(p.Distance(q))

	preim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(preim.Distance())
}

type Path []Point

func (p Path) Distance() float64 {
	sum := 0.0
	for i, _ := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}
