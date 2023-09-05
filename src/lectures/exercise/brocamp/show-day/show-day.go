// Using the ‘switch case’ write a program to accept
// an input number from the user and output the day as follows.

// Input | Output
// 1 - Sunday
// 2 - Monday
// 3 - Tuesday
// 4 - Wednesday
// 5 - Thursday
// 6 - Friday
// 7 - Saturday
// Any other input - Invalid Entry

package main

import "fmt"

const (
	Sunday    = 1
	Monday    = 2
	Tuesday   = 3
	Wednesday = 4
	Thursday  = 5
	Friday    = 6
	Saturday  = 7
)

func main() {
	var day int
	fmt.Println("Enter a number to find day")
	fmt.Scanf("%d", &day)

	switch day {
	case Sunday:
		fmt.Println("It's Sunday")
	case Monday:
		fmt.Println("It's Monday")
	case Tuesday:
		fmt.Println("It's Tuesday")
	case Wednesday:
		fmt.Println("It's Wednesday")
	case Thursday:
		fmt.Println("It's Thursday")
	case Friday:
		fmt.Println("It's Friday")
	case Saturday:
		fmt.Println("It's Saturday")
	default:
		fmt.Println("Invalid Entry")
	}
}
