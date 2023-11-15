package task19

import "fmt"

func Run() {
	fmt.Println()
	fmt.Println("Task 19:")
	str := "главрыба"
	fmt.Println(str)
	fmt.Println(reverse(str))
}

func reverse(str string) string {
	r := []rune(str)
	i := 0
	j := len(r) - 1
	for i < j {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}

	return string(r)
}
