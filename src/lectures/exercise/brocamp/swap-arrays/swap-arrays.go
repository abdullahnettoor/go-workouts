// Write a program to interchange the values of two arrays.
// Program should accept an array from the user, swap the values of two arrays and display it on the console
// Eg: Output: Enter the size of arrays
// Input: 5
// Output: Enter the values of Array 1
// Input: 10, 20, 30, 40, 50
// Output: Enter the values of Array 2
// Input: 15, 25, 35, 45, 55
// Output: Arrays after swapping:
// Array1: 15, 25, 35, 45, 55
// Array2: 10, 20, 30, 40, 50

package main

import "fmt"

func main() {
	var limit int

	fmt.Printf("Enter limit of Arrays : ")
	fmt.Scanf("%d", &limit)

	var arr1 = make([]int, limit)
	var arr2 = make([]int, limit)

	fmt.Println("Enter values of Array 1")
	for i, _ := range arr1 {
		fmt.Scanf("%d", &arr1[i])
	}
	fmt.Println("Enter values of Array 2")
	for i, _ := range arr2 {
		fmt.Scanf("%d", &arr2[i])
	}

	for i := range arr1 {
		temp := arr1[i]
		arr1[i] = arr2[i]
		arr2[i] = temp
	}

	fmt.Println("-- Arrays after Swapping --")
	fmt.Printf("Arrays 1 : ")
	for _, v := range arr1 {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\nArrays 2 : ")
	for _, v := range arr2 {
		fmt.Printf("%d ", v)
	}

}
