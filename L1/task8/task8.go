package task8

import (
	"fmt"
	"strconv"
)

func Run() {
	fmt.Println()
	fmt.Println("Task 8:")
	var number int64 = -16123
	newNumber := setBit(number, 4)

	fmt.Println(strconv.FormatInt(number, 2), number)
	fmt.Println(strconv.FormatInt(newNumber, 2), newNumber)
}

func setBit(n int64, pos uint) int64 {
	if pos == 63 {
		return -n
	}

	if pos >= 0 && pos < 63 {
		if n < 0 {
			return -(-n ^ (1 << pos))
		}
		return n ^ (1 << pos)
	}

	return n
}
