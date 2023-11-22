package task20

import (
	"fmt"
	"strings"
)

func Run() {
	fmt.Println()
	fmt.Println("Task 20:")

	str := "snow dog sun"
	fmt.Println(str)

	rev := reverse(str)
	fmt.Println(rev)
}

func reverse(str string) string {
	array := strings.Split(str, " ")

	i := 0
	j := len(array) - 1
	for i < j {
		array[i], array[j] = array[j], array[i]
		i++
		j--
	}

	return strings.Join(array, " ")
}
