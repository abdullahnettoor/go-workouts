//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

type IterateRunes func(string)

func iterateLines(lines []string, callback IterateRunes) {
	for _, string := range lines {
		callback(string)
	}
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	var letters, digits, spaces, punct = 0, 0, 0, 0

	iterate := func(s string) {
		for _, r := range s {
			if unicode.IsLetter(r) {
				letters++
			} else if unicode.IsDigit(r) {
				digits++
			} else if unicode.IsSpace(r) {
				spaces++
			} else if unicode.IsPunct(r) {
				punct++
			}
		}
	}

	iterateLines(lines, iterate)

	fmt.Printf("%2d Letters\n", letters)
	fmt.Printf("%2d Digits\n", digits)
	fmt.Printf("%2d Spaces\n", spaces)
	fmt.Printf("%2d Punctuations\n", punct)
}
