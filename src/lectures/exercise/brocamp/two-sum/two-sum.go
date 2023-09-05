package main

import "fmt"

func main() {
	var (
		numInt   int
		numFloat float32
	)
	fmt.Printf("Enter an Integer : ")
	fmt.Scanf("%d", &numInt)
	fmt.Printf("Enter a Float : ")
	fmt.Scanf("%f", &numFloat)

	fmt.Println("\nSum of 2 numbers is", float32(numInt)+numFloat)
}
