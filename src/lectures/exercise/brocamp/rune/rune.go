package main

import "fmt"

func main() {
	var char rune
	fmt.Println("Enter a character")
	fmt.Scanf("%c", &char)
	fmt.Printf("You have entered: %c\n", char)
}
