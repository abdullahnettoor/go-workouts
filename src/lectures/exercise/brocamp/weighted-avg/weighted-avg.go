// Grades are computed using a weighted average. Suppose that the written test counts 70%,  lab exams 20% and assignments 10%.
// If Arun has a score of
// Written test = 81
// Lab exams = 68
// Assignments = 92
// Arun’s overall grade = (81x70)/100 + (68x20)/100 + (92x10)/100 = 79.5
//  Write a program to find the grade of a student during his academic year.
// Program should accept the scores for written test, lab exams and assignments
// Output the grade of a student (using weighted average)
// Eg:
// Enter the marks scored by the students
// Written test = 55
// Lab exams = 73
// Assignments = 87
// Grade of the student is 61.8

package main

import "fmt"

func main() {
	var written, lab, assignments, avg float32
	fmt.Println("Enter the mark obtained")
	fmt.Printf("Written Test\t: ")
	fmt.Scanf("%f", &written)
	fmt.Printf("Lab Exams\t: ")
	fmt.Scanf("%f", &lab)
	fmt.Printf("Assignments\t: ")
	fmt.Scanf("%f", &assignments)
	avg = (written * 70 / 100) + (lab * 20 / 100) + (assignments * 10 / 100)
	fmt.Println("Grade of the student is", avg)
}
