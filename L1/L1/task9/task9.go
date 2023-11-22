package task9

import (
	"fmt"
	"sync"
)

func Run() {
	fmt.Println()
	fmt.Println("Task 9:")
	nums := make(chan int)
	numsSquare := make(chan int)

	arr := []int{1, 2, 3, 4, 5}

	pipeline(arr, nums, numsSquare)
}

func pipeline(arr []int, nums chan int, numsSquare chan int) {
	var wg sync.WaitGroup
	go toNums(arr, nums)
	go toNumsSquare(nums, numsSquare)
	wg.Add(1)
	go toStdOut(&wg, numsSquare)
	wg.Wait()
}

func toNums(arr []int, nums chan<- int) {
	for _, val := range arr {
		nums <- val
	}
	close(nums)
}

func toNumsSquare(nums <-chan int, numsSquare chan<- int) {
	for msg := range nums {
		numsSquare <- square(msg)
	}
	close(numsSquare)
}

func toStdOut(wg *sync.WaitGroup, numsSquare <-chan int) {
	for msg := range numsSquare {
		stdOut(msg)
	}
	wg.Done()
}

func stdOut(msg int) {
	fmt.Println(msg)
}

func square(num int) int {
	return num * num
}
