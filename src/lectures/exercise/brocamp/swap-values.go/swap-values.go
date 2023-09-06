package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Enter two numbers")
	fmt.Printf("a: ")
	fmt.Scanf("%d", &a)
	fmt.Printf("b: ")
	fmt.Scanf("%d", &b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Println("After Swapping")
	fmt.Println("a: ", a, "   b: ", b)
}
