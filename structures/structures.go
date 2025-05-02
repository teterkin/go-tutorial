package main

import "fmt"

type Rectangle struct {
	height float64
	width  float64
}

func (rect Rectangle) area() float64 {
	return rect.height * rect.width
}

func main() {
	rect1 := Rectangle{
		height: 30,
		width:  40,
	}
	fmt.Println("Rectangle:")
	fmt.Println("==========")
	fmt.Println("Height:", rect1.height)
	fmt.Println("Widtht:", rect1.width)
	fmt.Println("Area:", rect1.area())
	fmt.Println()
}
