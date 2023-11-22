package task5

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

func Run() {
	fmt.Println()
	fmt.Println("Task 5:")
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	ch := make(chan string)

	go Writer(ctx, ch)

	var wg sync.WaitGroup
	wg.Add(1)
	go Reader(ch, &wg)

	wg.Wait()

}

func Writer(ctx context.Context, ch chan string) {
	for true {
		select {
		case <-ctx.Done():
			close(ch)
			return
		default:
			Write(ch)
		}
	}
}

func Reader(ch chan string, wg *sync.WaitGroup) {
	for msg := range ch {
		fmt.Println(msg)
	}
	wg.Done()
}

func Write(ch chan<- string) {
	ch <- uuid.New().String()
	time.Sleep(100 * time.Millisecond)
}
