// Write a program to identify whether a string is a palindrome or not
// A string is a palindrome if it reads the same backward or forward eg: MALAYALAM
// Program should accept a string and display whether the string is a palindrome or not
// Eg: Output: Enter a string
// Input: MALAYALAM
// Output: Entered string is a palindrome
// Eg 2: Output: Enter a string
// Input: HELLO
// Output: Entered string is not a palindromeˀˀ

package main

import "fmt"

func main() {
	var word string
	isPalindrome := true
	var l int

	fmt.Println("Enter word to check Palindrome")
	fmt.Scanln(&word)

	var strLen = len(word)

	if strLen%2 == 0 {
		l = strLen / 2
	} else {
		l = strLen/2 + 1
	}

	for i := 0; i < l; i++ {
		if word[i] != word[strLen-1-i] {
			isPalindrome = false
		}
	}

	if isPalindrome {
		fmt.Println("Word", word, "is Palindrome")
	} else {
		fmt.Println("Word", word, "is not Palindrome")
	}

}
