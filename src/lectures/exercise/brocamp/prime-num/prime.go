// Write a program to check whether a given number is prime or not
// Program should accept an input from the user and display whether the number is prime or not
// Eg: Output: Enter a number
// Input: 7
// Output: Entered number is a Prime number

package main

import "fmt"

func main() {
	var num int
	isPrime := false
	fmt.Printf("Enter a number :")
	fmt.Scanf("%d", &num)

	for i := 2; i < num/2; i++ {
		if num%i == 1 {
			isPrime = true
		}
	}

	if isPrime && num > 1 {
		fmt.Println(num, "is Prime")
	} else {
		fmt.Println(num, "is Not Prime")
	}

}
