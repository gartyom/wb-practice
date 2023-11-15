package task18

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Run() {
	fmt.Println()
	fmt.Println("Task 18:")
	counterTest()
}

type someStruct struct {
	counter atomic.Uint32
}

func counterTest() {
	counter := &someStruct{}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				counter.counter.Add(1)
			}
		}()
	}

	wg.Wait()
	fmt.Println(counter.counter.Load())
}
