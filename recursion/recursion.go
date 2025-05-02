package main

import "fmt"

func factorial(num int) int {
	// factorial(5) = 5 * factorial(4)
	// factorial(4) = 4 * factorial(3)
	// factorial(3) = 3 * factorial(2)
	// factorial(2) = 2 * factorial(1)
	// factorial(1) = 1 * factorial(0)
	// factorial(0) = 1
	// ---
	// factorial(5) = 1 * 1 * 2 * 3 * 4 * 5 = 120
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {

	fmt.Println(factorial(5))

}
