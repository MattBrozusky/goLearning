package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	mySquare := square{
		sideLength: 5.5,
	}

	myTriangle := triangle{
		height: 10,
		base:   3.5,
	}

	printArea(mySquare)
	printArea(myTriangle)
}

func (t triangle) getArea() float64 {
	return t.base * t.height * 0.5
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
