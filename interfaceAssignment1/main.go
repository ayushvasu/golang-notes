package main

import "fmt"

type Shape interface {
	getArea() float64
}

type Triangle struct {
	base   float64
	height float64
}

type Square struct {
	sideLength float64
}

func main() {

	t := Triangle{10.5, 12}
	s := Square{10.5}

	printArea(t)
	printArea(s)

}

func printArea(s Shape) {
	fmt.Println(s.getArea())
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t Triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
