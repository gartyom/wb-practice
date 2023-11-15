package task2

import (
	"fmt"
	"sync"
)

func Run() {
	fmt.Println()
	fmt.Println("Task 2:")
	array := []int{2, 4, 6, 8, 10}
	calculateSquaresWaitGroup(&array)
	calculateSquaresChannels(&array)
}

func calculateSquaresWaitGroup(array *[]int) {
	var wg sync.WaitGroup
	for _, el := range *array {
		wg.Add(1)
		go func(el int) {
			result := calculateSquare(el)
			fmt.Printf("wg: %d\n", result)
			defer wg.Done()
		}(el)
	}
	wg.Wait()
}

func calculateSquaresChannels(array *[]int) {
	ch := make(chan int)
	for _, el := range *array {
		go func(el int) {
			ch <- calculateSquare(el)
		}(el)
	}

	for range *array {
		fmt.Printf("chan: %d\n", <-ch)
	}
}

func calculateSquare(el int) int {
	return el * el
}
