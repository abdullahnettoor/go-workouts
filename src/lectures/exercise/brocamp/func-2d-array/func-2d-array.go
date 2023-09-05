// Write a program to add the values of two 2D arrays
// Program should contains 3 functions including the main function
// main()
// Call function getArray()
// Call function addArray()
// Call function displayArray()
// 		getArray()
// Get values to the array
// 		addArray()
// Add array 1 and array 2
// 		displayArray()
// Display the array values

// Eg:
// Enter the size of array
// 2
// Enter the values of array 1
// 1	2
// 3	4
// Enter the values of array 2
// 5	6
// 7	8
// Output:
// Sum of array 1 and array 2:
// 6	8
// 10	12

package main

import "fmt"

func getArray(title string, arr [][]int, limit int) [][]int {
	fmt.Println(title)
	for i := 0; i < limit; i++ {
		arr = append(arr, make([]int, limit))
		for j := 0; j < limit; j++ {
			fmt.Scanf("%d", &arr[i][j])
		}
	}
	return arr
}

func addArray(arr1, arr2 [][]int) {
	for i := range arr1 {
		for j := range arr1 {
			fmt.Printf("%2d\t", arr1[i][j]+arr2[i][j])
		}
		fmt.Println()
	}
}

func displayArray(arr [][]int) {
	for _, row := range arr {
		for _, value := range row {
			fmt.Printf("%2d  ", value)
		}
		fmt.Println()
	}
}

func main() {
	var limit int
	fmt.Println("Enter array limit")
	fmt.Scanf("%d", &limit)

	var (
		arr1 [][]int
		arr2 [][]int
	)

	arr1 = getArray("Enter values of Array 1", arr1, limit)
	arr2 = getArray("Enter values of Array 2", arr2, limit)

	fmt.Println("-- Array 1 --")
	displayArray(arr1)
	fmt.Println("-- Array 2 --")
	displayArray(arr2)

	fmt.Println("After Adding two arrays")
	addArray(arr1, arr2)
}
