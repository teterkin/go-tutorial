package main

import "fmt"

func main() {

	const pi float64 = 3.14159265359
	x := 5
	isbool := true
	var name string = "Alex Teterkin"

	fmt.Printf(" %f \n", pi)
	fmt.Printf(" %.3f \n", pi)
	fmt.Printf(" %T \n", isbool)
	fmt.Printf(" %T \n", name)
	fmt.Printf(" %t \n", isbool)
	fmt.Printf(" %d \n", x)
	// binary
	fmt.Printf(" %b \n", 25)
	// key code
	fmt.Printf(" %c \n", 33)
	// hex number
	fmt.Printf(" %x \n", 15)
	// scentific notation
	fmt.Printf(" %e \n", pi)
}
