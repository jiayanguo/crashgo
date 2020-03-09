package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	base float64
	height float64
}
type square struct {
	sideLength float64
}

func (t triangle) getArea() float64{
	return 0.5*t.base*t.height
}

func (s square) getArea() float64 {
	return s.sideLength*s.sideLength
}

func main() {
	t := triangle {
		base : 5.0,
		height : 4.0,
	}

	s := square {
		sideLength : 4.0,
	}

	fmt.Println("triangle ", t, "erea ", t.getArea())
	fmt.Println("squre ", s, "erea ", s.getArea())
}