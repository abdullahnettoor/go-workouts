package main

import "fmt"

func getSum() int {
	var count int
	var sum int
	fmt.Printf("Enter Number Count: ")
	fmt.Scanf("%d", count)
	var n = make([]int, count)
	fmt.Println("Enter numbers")
	for i := 0; i < count; i++ {
		fmt.Scanf("%d", &n[i])

		sum += n[i]
	}

	return sum
}

func getSumVariadic(n ...int) (sum int) {
	for _, v := range n {
		sum += v
	}
	return
}

func main() {

	sum := getSum()
	fmt.Printf("Sum is %d", sum)

	var count int
	fmt.Printf("Enter Number Count: ")
	fmt.Scanf("%d", count)
	var n = make([]int, count)
	fmt.Println("Enter numbers")
	for i := 0; i < count; i++ {
		fmt.Scanf("%d", &n[i])
	}
	sum2 := getSumVariadic(n...)
	fmt.Printf("Sum is %d", sum2)

	func() {
		fmt.Println("Hello, I'm Anonymus")
	}()

}
