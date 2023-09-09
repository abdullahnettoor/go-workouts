package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("hello")
			fmt.Println("hello")
			fmt.Println("hello")
			time.Sleep(4 * time.Second)
		}
	}()
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("hi")
			time.Sleep(3 * time.Second)
		}
	}()
	time.Sleep(12 * time.Second)

}
