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

	doWork(ctx)

}

func doWork(ctx context.Context) {
	if ctx.Err() == context.DeadlineExceeded {
		return
	}
}
