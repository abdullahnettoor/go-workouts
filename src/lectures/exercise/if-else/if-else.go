//--Summary:
//  The existing program is used to restrict access to a resource
//  based on day of the week and user role. Currently, the program
//  allows any user to access the resource. Implement the functionality
//  needed to grant resource access using any combination of `if`, `else if`,
//  and `else`.
//
//--Requirements:
//* Use the accessGranted() and accessDenied() functions to display
//  informational messages
//* Access at any time: Admin, Manager
//* Access weekends: Contractor
//* Access weekdays: Member
//* Access Mondays, Wednesdays, and Fridays: Guest

package main

import "fmt"

// Days of the week
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// User roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func accessGranted() {
	fmt.Println("Access Granted")
}

func accessDenied() {
	fmt.Println("Access Denied")
}

func isWeekEnd(day int) bool {
	return day > 4
}

func main() {
	// The day and role. Change these to check your work.
	today, role := Tuesday, Guest

	if role == Admin || role == Manager {
		accessGranted()
	} else if role == Contractor && isWeekEnd(today) {
		accessGranted()
	} else if role == Member && !isWeekEnd(today) {
		accessGranted()
	} else if role == Guest && (today < 5 && today%2 == 0) {
		accessGranted()
	} else {
		accessDenied()
	}

}