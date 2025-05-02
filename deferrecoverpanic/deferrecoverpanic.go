package main

import "fmt"

func firstRun() { fmt.Println("I executed First!") }

func secondRun() { fmt.Println("I executed Second!") }

func lastRun() { fmt.Println("I executed Last!") }

func div(num1, num2 int) int {
	defer func() {
		fmt.Println("Devision is done!")
		if rec := recover(); rec != nil {
			fmt.Println(rec)
		}
	}()
	return num1 / num2
}

func panicDemo() {
	defer func() {
		fmt.Println(recover())
		fmt.Println("Cleaning up...")
		fmt.Println("Done.")
	}()

	panic("Alarm!")
}

func main() {

	defer lastRun()
	firstRun()
	secondRun()

	fmt.Println(div(6, 2))
	fmt.Println(div(3, 0))

	panicDemo()

}
