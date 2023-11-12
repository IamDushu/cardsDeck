package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {

	triangle := triangle{
		height: 10,
		base:   5,
	}

	square := square{
		sideLength: 10.5,
	}

	// fmt.Println(triangle.getArea())
	// fmt.Println(square.getArea())

	printArea(triangle)
	printArea(square)

}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	area := s.getArea()
	fmt.Println("The area is", area)
}
