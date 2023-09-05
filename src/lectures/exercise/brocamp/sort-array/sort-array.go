// Write a program to sort an array in descending order
// Program should accept and array, sort the array values in descending order and display it
// Eg: Output: Enter the size of an array
// Input: 5
// Output: Enter the values of array
// Input: 20, 10, 50, 30, 40
// Output: Sorted array:
// 50, 40, 30, 20, 10

package main

import (
	"fmt"
	"sort"
)

func main() {
	var limit uint

	fmt.Println("Enter the size of an array")
	fmt.Scanf("%d", &limit)

	var arr = make([]int, limit)

	fmt.Println("Enter values of Array")
	for i := range arr {
		var v int
		fmt.Scanf("%d", &v)
		arr[i] = v
	}
	fmt.Println(arr)

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	fmt.Println("Array after sorting Descending")
	fmt.Println(arr)
}
