// Write a program to multiply the adjacent values of an array and store it in an another array
// Program should accept an array
// Multiply the adjacent values
// Store the result into another array
// Eg:
// Enter the array limit
// 5
// Enter the values of array
// 1	2	3	4	5
// Output
// 2	6	12	20

package main

import "fmt"

func main() {
	var limit int

	fmt.Println("Enter limit of Array")
	fmt.Scanf("%d", &limit)

	var arr = make([]int, limit)
	var newArr = make([]int, limit-1)

	fmt.Println("Enter Array Values")
	for i := range arr {
		fmt.Scanf("%d", &arr[i])
	}

	for i := range newArr {
		newArr[i] = arr[i] * arr[i+1]
	}

	fmt.Println("Array after multiplying adjacents")
	for _, value := range newArr {
		fmt.Printf("%d  ", value)
	}
}
