package main

import (
	"fmt"
	"strings"
)

func getSum() int {
	var count int
	var sum int
	fmt.Printf("Enter Number Count: ")
	fmt.Scanf("%d", &count)
	var n = make([]int, count)
	fmt.Println("Enter numbers")
	for i := 0; i < count; i++ {
		fmt.Scanf("%d", &n[i])
		sum += n[i]
	}
	return sum
}

func getSumVariadic(n ...int) (sum int) {
	for _, v := range n {
		sum += v
	}
	return
}

func main() {

	// Get sum
	sum := getSum()
	fmt.Printf("Sum is %d\n", sum)
	fmt.Println()

	// Get Sum Variadic
	var count int
	fmt.Printf("Enter Number Count: ")
	fmt.Scanf("%d", &count)
	var n = make([]int, count)
	fmt.Println("Enter numbers")
	for i := 0; i < count; i++ {
		fmt.Scanf("%d", &n[i])
	}
	sum2 := getSumVariadic(n...)
	fmt.Printf("Sum is %d\n", sum2)
	fmt.Println()

	// Anonymous function
	func() {
		fmt.Println("Hello, I'm Anonymus")
	}()

	// Interface Examples
	a := Name("Arya")

	a.Greet() // Prints Goodmorning Arya

	a.ToUpperCase().Greet() // Prints Goodmorning ARYA

}

type Hello interface {
	Greet()
}

type Name string

func (s Name) Greet() {
	fmt.Println("Goodmorning", s)
}

func (s Name) ToUpperCase() Name {
	s = Name(strings.ToUpper(string(s)))
	return s
}
