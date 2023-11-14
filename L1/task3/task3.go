package task3

import (
	"fmt"
	"sync"
)

func Run() {
	fmt.Println("Task 3:")
	array := []int{2, 4, 6, 8, 10}
	var sum int
	sum = sumOfSquaresChan(&array)
	fmt.Println("Chan: ", sum)
	sum = sumOfSquaresMutex(&array)
	fmt.Println("Mutex:", sum)
}

func sumOfSquaresChan(array *[]int) int {
	ch := make(chan int)
	for _, el := range *array {
		go func(el int, ch chan<- int) {
			ch <- calculateSquare(el)
		}(el, ch)
	}

	sum := 0
	for range *array {
		sum += <-ch
	}

	return sum
}

func sumOfSquaresMutex(array *[]int) int {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	sum := 0
	for _, el := range *array {
		wg.Add(1)
		go func(el int, mutex *sync.Mutex) {
			result := calculateSquare(el)
			mutex.Lock()
			sum += result
			mutex.Unlock()
			defer wg.Done()

		}(el, &mutex)
	}
	wg.Wait()
	return sum
}

func calculateSquare(el int) int {
	return el * el
}
