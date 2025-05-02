package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {

		fmt.Println(i)

	}

	// while loop in go
	k := 1
	for k <= 10 {

		fmt.Printf("%2.0d \n", k)
		k++
	}

	// loop inside the loop (nested loop)
	for l := 1; l < 10; l++ {
		for m := 1; m < l; m++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
