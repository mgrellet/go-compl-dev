package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {

	tr := triangle{
		base:   2.2,
		height: 4.3,
	}
	sq := square{
		sideLength: 2.9,
	}
	printArea(tr)
	printArea(sq)

}

func printArea(s shape) {
	fmt.Println("The area is: ", s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}