package main

import "fmt"

func main() {

	var studentCount [10]int

	for i := 0; i < 10; i++ {

		studentCount[i] = i + 1
		fmt.Println(studentCount[i])

	}

	evenNum := [5]int{0, 2, 4, 6, 8}
	for i := range 5 {
		fmt.Println(evenNum[i])
	}

	for _, value := range evenNum {
		fmt.Println(value)
	}

	newArr := []int{5, 4, 3, 2, 1}
	arrSliced := newArr[3:5]
	fmt.Println(arrSliced)
	fmt.Println(newArr[3:])

	fmt.Println(newArr[:5])
	fmt.Println(newArr[0:])

	fmt.Println(newArr[:3])

	// Empty slice
	arr2 := make([]int, 5, 10)
	fmt.Println(arr2)

	copy(arr2, newArr)
	fmt.Println(arr2)

	arr3 := append(newArr, 3, 2, -1)
	fmt.Println(arr3)

}
