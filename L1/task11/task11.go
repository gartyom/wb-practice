package task11

import "fmt"

func Run() {
	fmt.Println("Task 11:")
	first := []int{1, 14, 12, 22, 5, 1123, 56}
	second := []int{1, 12, 53, 5, 11, 0, 7, -1}
	fmt.Println(makeIntersection(first, second))
}

func makeIntersection(first []int, second []int) map[int]int {
	intersection := make(map[int]int)
	for _, val := range first {
		intersection[val] += 1
	}

	for _, val := range second {
		intersection[val] += 1
	}

	for key, val := range intersection {
		if val == 1 {
			delete(intersection, key)
		}
	}

	return intersection
}
