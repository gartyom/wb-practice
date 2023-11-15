package task22

import (
	"fmt"
	"math/big"
)

func Run() {
	fmt.Println()
	fmt.Println("Task 22:")
}

func add(f, s int64) *big.Int {
	fb := big.NewInt(int64(f))
	sb := big.NewInt(int64(s))

	return fb.Add(fb, sb)
}
