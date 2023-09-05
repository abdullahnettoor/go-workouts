// Write a program to find the sum of all the odd numbers for a given limit
// Program should accept an input as limit from the user and display the sum of all the odd numbers within that limit
//*For example if the input limit is 10 then the result is 1+3+5+7+9 = 25
// Output: Enter a limit
// Input: 10
// Output: Sum of odd numbers = 25

package main

import "fmt"

func main() {
	var limit, sum int

	fmt.Println("Enter a limit to find sum of Odd numbers")
	fmt.Scanf("%d", &limit)

	for i := 1; i <= limit; i++ {
		if i%2 == 1 {
			sum += i
		}
	}
	fmt.Println("Sum is", sum)

}
