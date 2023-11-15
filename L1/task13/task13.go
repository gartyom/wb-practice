package task13

import "fmt"

func Run() {
	fmt.Println("Task 13:")
	array := []int{1, 2}
	fmt.Println("array: ", array)
	Swap1(array, 0, 1)
	fmt.Println("swap 1:", array)

	Swap2(array, 0, 1)
	fmt.Println("swap 2:", array)

	Swap3(array, 0, 1)
	fmt.Println("swap 3:", array)
}

func Swap1(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

// может появится int overflow, если значения достаточно большие
func Swap2(array []int, i, j int) {
	array[i] += array[j]
	array[j] = array[i] - array[j]
	array[i] -= array[j]
}

// Перестановка не сработает, если оба индекса указывают на один элемент
func Swap3(array []int, i, j int) {
	array[i] ^= array[j]
	array[j] ^= array[i]
	array[i] ^= array[j]
}
