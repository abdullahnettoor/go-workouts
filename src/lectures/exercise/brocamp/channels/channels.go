package main

import (
	"fmt"
	// "time"
)

func main() {

	c := make(chan string)

	fmt.Println("Function Started...")

	go func() {
		fmt.Println("Hello, I'm from goroutine")
		// time.Sleep(2 * time.Second)
		c <- "Ended"
	}()

	fmt.Println("Go routine is on the way")
	v := <-c
	fmt.Println(v)
}
