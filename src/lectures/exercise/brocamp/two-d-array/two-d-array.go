// Write a program to add to two dimensional arrays
// Program should accept two 2D arrays and display its sum
// Eg: Output: Enter the size of arrays
// Input: 3
// Output: Enter the values of array 1
// Input:
// 1 2 3
// 4 5 6
// 7 8 9
// Output: Enter the values of array 2
// Input:
// 10 20 30
// 40 50 60
// 70 80 90
// Output: Sum of 2 arrays is:
// 11 22 33
// 44 55 66
// 77 88 99

package main

import "fmt"

func main() {
	var limit int

	fmt.Printf("Enter Size of Array: ")
	fmt.Scanf("%d", &limit)

	fmt.Println("Enter elements of Array 1 ")
	var arr1 [][]int
	for i := 0; i < limit; i++ {
		arr1 = append(arr1, make([]int, limit))
		for j := 0; j < limit; j++ {
			fmt.Scanf("%d", &arr1[i][j])
		}
	}

	fmt.Println("Enter elements of Array 2")
	var arr2 [][]int
	for i := 0; i < limit; i++ {
		arr2 = append(arr2, make([]int, limit))
		for j := 0; j < limit; j++ {
			fmt.Scanf("%d", &arr2[i][j])
		}
	}

	fmt.Println("--- Array 1 ---")
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			fmt.Printf("%2d\t", arr1[i][j])
		}
		fmt.Println()
	}

	fmt.Println("--- Array 2 ---")
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			fmt.Printf("%2d\t", arr2[i][j])
		}
		fmt.Println()
	}

	fmt.Println("Sum of two 2D arrays")
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			fmt.Printf("%2d\t", arr1[i][j]+arr2[i][j])
		}
		fmt.Println()
	}

}
