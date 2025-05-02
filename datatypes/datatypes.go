package main

import "fmt"

func main() {

	var a int = 5
	var b float32 = 4.32
	const pi float64 = 3.14159265359
	x, y := 14, 15

	fmt.Println(a, b, pi, x, y)

	c, d := 5, 6

	fmt.Println("c + d = ", c+d)
	fmt.Println("c - d = ", c-d)
	fmt.Println("c * d = ", c*d)
	fmt.Println("c / d = ", c/d)
	fmt.Println("c mod d = ", c%d)

	var good bool = true

	evil := false

	fmt.Println(good && good)
	fmt.Println(good && evil)
	fmt.Println(good || evil)
	fmt.Println(!good)
	fmt.Println(!evil)
}
