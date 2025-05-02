package main

import "fmt"

func main() {

	x := 5
	fmt.Println(x)
	fmt.Println(&x)

	changeValue(x)
	fmt.Println(x)
	changeAtLocation(&x)
	fmt.Println(x)

}

func changeValue(x int) {
	// Will not work!
	x = 7
}

func changeAtLocation(x *int) {
	// Should work now!
	*x = 7
}
