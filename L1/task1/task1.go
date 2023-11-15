package task1

import "fmt"

type Human struct {
}

func (h *Human) Walk() {
	fmt.Println("Walk")
}

type Action struct {
	Human
}

func Run() {
	fmt.Println()
	fmt.Println("Task 1:")
	h := &Action{}
	h.Walk()
}
