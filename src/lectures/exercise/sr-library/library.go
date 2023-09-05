//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type list []string
type MemberId int
type BookId int

const (
	CheckedIn  = true
	CheckedOut = false
)

type Member struct {
	id         MemberId
	name       string
	booksTaken []Book
}

type Book struct {
	id       BookId
	name     string
	checkIn  time.Time
	checkOut time.Time
	status   bool
}

func checkOutBook(book *Book, member *Member) {
	book.checkOut = time.Now()
	book.status = CheckedOut
	member.booksTaken = append(member.booksTaken, *book)
}

func checkInBook(book *Book, member *Member) {
	book.checkIn = time.Now()
	book.status = CheckedIn
	member.booksTaken = append(member.booksTaken, *book)
}

func main() {

	fmt.Printf("%T", time.Now())

	members := make(map[MemberId]Member)
	mammu := Member{0, "mammu", []}
	members[0] = mammu 
	books := make(map[BookId]Book)

}
