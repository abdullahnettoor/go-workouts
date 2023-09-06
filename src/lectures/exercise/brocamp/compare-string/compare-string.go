package main

import (
	"fmt"
	"strings"
)

func main() {
	s1, s2 := "hi", "Hi"
	isEqual := s1 == s2

	fmt.Println(s1, "&", s2, isEqual)

	s2 = "hi"
	isEqual = s1 == s2
	fmt.Println(s1, "&", s2, isEqual)

	res := strings.Compare(s1, s2)
	fmt.Println(s1, "&", s2, res) //prints 0 when equal

	s2 = "Hi"
	res = strings.Compare(s1, s2)
	fmt.Println(s1, "&", s2, res) //prints 1 when greater

	s1 = "F"
	s2 = "H"
	res = strings.Compare(s1, s2)
	fmt.Println(s1, "&", s2, res) //prints -1 when lesser
}
