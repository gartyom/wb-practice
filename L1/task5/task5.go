package task5

import (
	"context"
	"fmt"
	"time"
)

func Run() {
	fmt.Println("Task 5:")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		return
	default:
		doWork(ctx)
	}

}

func doWork(ctx context.Context) {
	for true {
		fmt.Println("123123123")
		time.Sleep(1 * time.Second)
	}
}
