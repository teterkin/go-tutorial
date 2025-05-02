package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

func (rect Rectangle) area() float64 {
	return rect.height * rect.width
}

func (circ Circle) area() float64 {
	return math.Pi * math.Pow(circ.radius, 2)
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	rect := Rectangle{
		height: 50,
		width:  60,
	}
	circ := Circle{
		radius: 7,
	}
	fmt.Println("Rectangle:")
	fmt.Println("==========")
	fmt.Println("Height:", rect.height)
	fmt.Println("Widtht:", rect.width)
	fmt.Println("Area:", getArea(rect))
	fmt.Println()
	fmt.Println("Circle:")
	fmt.Println("=======")
	fmt.Println("Radius:", circ.radius)
	fmt.Println("Area:", getArea(circ))
}
