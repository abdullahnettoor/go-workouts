// Write an object oriented program to store and display the values of a 2D array
// Program should contains 3 functions including the main function
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

type TwoDArray struct {
	arr [][]int
}

func (twoD *TwoDArray) GetArray() {
	var size int
	fmt.Printf("Enter Array Size : ")
	fmt.Scanf("%d", &size)

	fmt.Println("Enter Array Values")
	twoD.arr = make([][]int, size)
	for i := 0; i < size; i++ {
		twoD.arr[i] = make([]int, size)
		for j := 0; j < size; j++ {
			fmt.Scanf("%d", &twoD.arr[i][j])
		}
	}
}

func (twoD *TwoDArray) DisplayArray() {
	fmt.Println("Displaying Array...")
	for _, row := range twoD.arr {
		for _, value := range row {
			fmt.Printf("%d\t", value)
		}
		fmt.Println()
	}
}

func main() {
	arr := TwoDArray{}

	arr.GetArray()
	arr.DisplayArray()

}
