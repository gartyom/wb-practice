package task23

import "fmt"

func Run() {
	fmt.Println()
	fmt.Println("Task 23:")
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(removeElementSaveOrder(s, 4))
	fmt.Println(removeElement(s, 4))
}

func removeElementSaveOrder(s []int, i int) []int {
	result := append([]int{}, s...)
	return append(result[:i], result[i+1:]...)
}
func removeElement(s []int, i int) []int {
	result := append([]int{}, s...)
	result[i] = result[len(s)-1]
	return result[:len(result)-1]
}
