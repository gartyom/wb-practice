package task25

import (
	"fmt"
	"time"
)

func Run() {
	fmt.Println()
	fmt.Println("Task 25:")
	go func() {
		sleep(250 * time.Millisecond)
		fmt.Println("1")
	}()

	sleep(500 * time.Millisecond)
	fmt.Println("2")
}

func sleep(t time.Duration) {
	<-time.After(t)
}
