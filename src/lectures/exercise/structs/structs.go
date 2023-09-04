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
	a, b Cordinate
}

type Cordinate struct {
	x, y float32
}

func width(rec Rectangle) float32 {
	return (rec.b.x - rec.a.x)
}

func length(rec Rectangle) float32 {
	return (rec.a.y - rec.b.y)
}

func findArea(rec Rectangle) float32 {
	return length(rec) * width(rec)
}

func findPerimeter(rec Rectangle) float32 {
	return 2 * (length(rec) + width(rec))
}

func divider() {
	fmt.Println("----------------------------")
}

func main() {

	var rec = Rectangle{
		a: Cordinate{0, 7},
		b: Cordinate{5, 0},
	}

	divider()
	fmt.Println("Length :", length(rec), "Breadth :", width(rec), "\nArea =", findArea(rec), "\nPerimeter =", findPerimeter(rec))
	divider()

	rec.a.y *= 2
	rec.b.x *= 2

	fmt.Println("Length :", length(rec), "Breadth :", width(rec), "\nArea =", findArea(rec), "\nPerimeter =", findPerimeter(rec))
	divider()

}
