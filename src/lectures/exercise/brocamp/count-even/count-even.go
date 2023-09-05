// Write a program to find the number of even numbers in an array
// Program should accept an array and display the number of even numbers contained in that array
// Eg: Output: Enter the size of an array
// Input: 5
// Output: Enter the values of array
// Input: 11, 20, 34, 50, 33
// Output: Number of even numbers in the given array is 3

package main

import "fmt"

func main() {
	var evenCount, limit uint

	fmt.Printf("Enter limit: ")
	fmt.Scanf("%d", &limit)

	var arr = make([]int, limit)

	fmt.Println("Enter Array elements")
	for _, v := range arr {
		fmt.Scanf("%d", &v)
		if v%2 == 0 {
			evenCount++
		}
	}
	fmt.Println("Number of even number in given array is", evenCount)
}
