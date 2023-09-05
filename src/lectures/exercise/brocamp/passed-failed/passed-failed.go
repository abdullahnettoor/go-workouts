// Write a program to check whether a student has passed or failed in a subject after
// he or she enters their mark (pass mark for a subject is 50 out of 100).
// Program should accept an input from the user and output a message as “Passed” or “Failed”
// Variable
// Data type
// mark float

package main

import "fmt"

func main() {
	var mark float32
	fmt.Println("Enter mark to know Passed or Failed")
	fmt.Scanf("%f", &mark)

	if mark < 50 {
		fmt.Println("FAILED")
	} else {
		fmt.Println("PASSED")
	}
}
