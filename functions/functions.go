package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func addemup(args ...int) int {
	sum := 0
	for _, value := range args {
		sum += value
	}
	return sum
}

func main() {

	fmt.Println("5 + 4 =", add(5, 4))
	fmt.Println(addemup(10, 20, 30, 40, 50))

}
