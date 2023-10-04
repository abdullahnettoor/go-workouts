package main

import "fmt"

func main() {

	arr := []int{5, 1, 3, 7, 8, 10, 6, 7, 3, 10}

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				arr = append(arr[:j], arr[j+1:]...)
			}
		}
	}

	fmt.Println(arr)

}
