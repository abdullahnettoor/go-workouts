//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	x, y float32
}

func findArea(rec Rectangle) float32 {
	return rec.x * rec.y
}

func main() {

	var rec = Rectangle{5.5, 6}
	fmt.Println("Length :", rec.x, "Breadth :", rec.y, "\nArea =", findArea(rec))

	rec.x, rec.y = rec.x*2, rec.y*2
	fmt.Println("Length :", rec.x, "Breadth :", rec.y, "\nArea =", findArea(rec))

}
