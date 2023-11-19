package task22

import (
	"fmt"
	"math"
	"math/big"
)

func Run() {
	fmt.Println()
	fmt.Println("Task 22:")
	x := big.NewInt(int64(math.Pow(2, 21)))
	y := big.NewInt(int64(math.Pow(3, 22)))
	z := big.NewInt(0)
	z.Add(x, y)
	fmt.Println("2^21 + 3^22", z)
	z.Sub(x, y)
	fmt.Println("2^21 - 3^22", z)
	z.Mul(x, y)
	fmt.Println("2^21 * 3^22", z)
	z.Div(x, y)
	fmt.Println("2^21 / 3^22", z)
}
