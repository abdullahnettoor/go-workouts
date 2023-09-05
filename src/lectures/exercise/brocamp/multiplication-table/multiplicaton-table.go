package main

import "fmt"

func main() {
	var num int
	fmt.Printf("Enter a number: ")
	fmt.Scanf("%d", &num)
	fmt.Println("Multiplication Table of", num)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%2d x %d = %2d\n", i, num, i*num)
	}
}
