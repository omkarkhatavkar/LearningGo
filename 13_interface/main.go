package main

import (
	"fmt"
)

func main() {
	c := Circle{radius: 22}
	fmt.Println(getArea(c))
}

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Recatangle struct {
	height, width float64
}

func (c Circle) area() float64 {
	return (3.14 * c.radius * c.radius)
}

func (c Recatangle) area() float64 {
	return c.height * c.height
}

func getArea(s Shape) float64 {
	return s.area()

}
