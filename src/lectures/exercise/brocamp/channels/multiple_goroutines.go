package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the function is done.
	fmt.Printf("Worker %d is working\n", id)
}

func main() {
	var wg sync.WaitGroup

	numWorkers := 5
	wg.Add(numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg)
	}

	wg.Wait() // Wait for all workers to finish.
	fmt.Println("All workers have finished.")

	nums := []int{5, 4, 6, 7, 10}

	nums2 := nums[1:5]

	fmt.Println(nums2)

}
