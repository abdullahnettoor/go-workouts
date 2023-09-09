package main

import (
	"fmt"
)

func main() {
	marks := map[string]int{"English": 25, "Maths": 50}
	marks["Coding"] = 49

	fmt.Println(marks)

	_, ok := marks["English"]
	if ok {
		marks["English"] += 20
	}

	fmt.Println(marks)

	delete(marks, "Maths")
	fmt.Println(marks)

}
