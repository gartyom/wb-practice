package task6

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Run() {
	fmt.Println()
	fmt.Println("Task 6:")
	var wg sync.WaitGroup
	signalChannel(&wg)
	closeChannel(&wg)
	cancelContext(&wg)
	wg.Wait()
}

func signalChannel(wg *sync.WaitGroup) {
	ch := make(chan string, 1)
	wg.Add(1)
	go func(ch chan string) {
		fmt.Println("signalChannel")
		defer fmt.Println("signalChannel gorutine stopped")
		defer wg.Done()

		for true {
			select {
			case <-ch:
				return
			}
		}
	}(ch)

	time.Sleep(100 * time.Millisecond)
	ch <- "1"
}

func closeChannel(wg *sync.WaitGroup) {
	ch := make(chan string)
	wg.Add(1)
	go func(ch chan string) {
		fmt.Println("closeChannel")
		defer fmt.Println("closeChannel gorutine stopped")
		defer wg.Done()

		for true {
			_, ok := <-ch
			if ok == false {
				return
			}
		}

	}(ch)

	time.Sleep(100 * time.Millisecond)
	close(ch)
}

func cancelContext(wg *sync.WaitGroup) {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func(ctx context.Context) {
		fmt.Println("cancelContext")
		defer fmt.Println("cancelContext gorutine stopped")
		defer wg.Done()

		for true {
			select {
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	time.Sleep(100 * time.Millisecond)
	cancel()
}
