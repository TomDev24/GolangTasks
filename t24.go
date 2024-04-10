package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) CalcDistance(p2 *Point) float64 {
	dist := math.Pow(p2.x - p.x, 2) + math.Pow(p2.y - p.y, 2)
	return math.Sqrt(dist)
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func main(){
	p1 := NewPoint(-5.0, 10.0)
	p2 := NewPoint(20.0, 20.0)

	fmt.Println(p1.CalcDistance(p2))
}
