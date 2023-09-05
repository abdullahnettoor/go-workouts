// Write a program to accept an array and display it on the console using functions
// Program should contain 3 functions including main() function
// main()
// Declare an array
// Call function getArray()
// Call function displayArray()
// 		getArray()
// Get values to the array
// 		displayArray()
// Display the array values

package main

import "fmt"

func getArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Scanf("%d", &arr[i])
	}

}

func displayArray(title string, arr []int) {
	fmt.Println(title)
	for _, v := range arr {
		fmt.Printf("%d  ", v)
	}
}

func main() {
	var limit int
	fmt.Println("Enter the limit of Array")
	fmt.Scanf("%d", &limit)
	var myArr = make([]int, limit)

	fmt.Println("Enter Array Values")
	getArray(myArr)

	displayArray("Displaying array..", myArr)
}
