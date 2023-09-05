// 1
// 2	3
// 4	5	6
// 7	8	9	10

package main

import "fmt"

func main() {
	limit := 1
	num := 1

	for {
		for j := 1; j <= limit; j++ {
			fmt.Printf("%d ", num)
			num++
		}
		fmt.Println()
		if limit == 4 {
			break
		}
		limit++
	}
}
