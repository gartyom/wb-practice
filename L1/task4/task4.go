package task4

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/google/uuid"
)

func Run() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	fmt.Println("Task 4:")
	numWorkers := 5
	dataChan := make(chan string)
	var wg sync.WaitGroup

	go doWork(dataChan, numWorkers, &wg)

Writer:
	for true {
		select {
		case <-sigs:
			close(dataChan)
			break Writer
		default:
			dataChan <- uuid.New().String()
		}
	}

	wg.Wait()
}

func doWork(dataChan <-chan string, numWorkers int, wg *sync.WaitGroup) {
	for i := numWorkers; i > 0; i-- {
		wg.Add(1)
		go func(i int) {
			for msg := range dataChan {
				fmt.Println(msg, i)
			}

			fmt.Println("Work done ", i)
			defer wg.Done()
		}(i)
	}
}
