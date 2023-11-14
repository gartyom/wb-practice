package task8

import (
	"fmt"
	"strconv"
)

func Run() {
	fmt.Println("Task 8:")
	var number int64 = 32
	newNumber := setBit(number, 63)

	fmt.Println(strconv.FormatInt(number, 2))
	fmt.Println(strconv.FormatInt(newNumber, 2))
	fmt.Println(number)
	fmt.Println(newNumber)
}

func setBit(n int64, pos uint) int64 {
	return n | (1 << pos)
}
