package main

import (
	"fmt"
)

//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func displayGreeting(name string) {
	fmt.Println("Welcome,", name)
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func printMsg() string {
	return "Hope, You are doing great!"
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func add3Numbers(number1, number2, number3 int) int {
	return number1 + number2 + number3
}

// * Write a function that returns any number
func findSquare(num int) int {
	return num * num
}

// * Write a function that returns any two numbers
func smallest2PrimeNumbers() (int, int) {
	return 2, 3
}

func main() {

	displayGreeting("Abdullah")
	fmt.Println(printMsg())
	num1, num2 := smallest2PrimeNumbers()

	//* Add three numbers together using any combination of the existing functions.
	sum3Num := add3Numbers(findSquare(5), num1, num2)

	//  * Print the result
	fmt.Println("Sum of 3 numbers is", sum3Num)

}
