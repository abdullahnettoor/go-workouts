// Income tax is calculated as per the following table
// Annual Income
// Tax percentage
// Up to 2.5 Lakhs
// No Tax
// Above 2.5 Lakhs to 5 Lakhs
// 5%
// Above 5 Lakhs to 10 Lakhs
// 20%
// Above 10 Lakhs to 50 Lakhs
// 30%

// Write a program to find out the income tax amount of a person.
// Program should accept annual income of a person
// Output the amount of tax he has to pay

// Eg 1:
// Enter the annual income
// 495000
// Income tax amount = 24750.00

// Eg 2:
// Enter the annual income
// 500000
// Income tax amount = 25000.00

package main

import "fmt"

func main() {
	var income int
	var tax float32
	fmt.Println("Enter your annual income")
	fmt.Scanf("%d", &income)

	if income > 1000000 {
		tax = float32(income) * 30 / 100
	} else if income > 500000 {
		tax = float32(income) * 20 / 100
	} else if income > 250000 {
		tax = float32(income) * 5 / 100
	} else {
		fmt.Println("You don't have pay any income tax")
	}

	if tax != 0 {
		fmt.Println("Your Income Tax is", tax)
	}
}
