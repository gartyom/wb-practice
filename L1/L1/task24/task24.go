package task24

import (
	"fmt"
	"math"
)

func Run() {
	fmt.Println()
	fmt.Println("Task 24:")

	p1 := NewPoint(18, 1)
	p2 := NewPoint(4, 13)

	fmt.Println(calcDistance(p1, p2))
}

func calcDistance(p1 *Point, p2 *Point) float64 {
	x := math.Pow((p1.GetX() - p2.GetX()), 2)
	y := math.Pow((p1.GetY() - p2.GetY()), 2)
	return math.Sqrt(x + y)
}

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}
