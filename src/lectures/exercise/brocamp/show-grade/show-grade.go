// Write a program to show the grade obtained by a student after he/she enters their total mark percentage.
// Program should accept an input from the user and display their grade as follows
// Mark
// Grade
// > 90		- A
// 80-89	- B
// 70-79	- C
// 60-69	- D
// 50-59	- E
// < 50		- Failed

// Total mark - float

package main

import "fmt"

func main() {
	var mark float32
	var grade string
	fmt.Println("Enter mark to know your GRADE")
	fmt.Scanf("%f", &mark)

	if mark > 90 {
		grade = "A"
	} else if mark >= 80 {
		grade = "B"
	} else if mark >= 70 {
		grade = "C"
	} else if mark >= 60 {
		grade = "D"
	} else if mark >= 50 {
		grade = "E"
	} else {
		grade = "Failed"
	}

	fmt.Println("Your Grade :", grade)
}
