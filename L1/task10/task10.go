package task10

import "fmt"

func Run() {
	fmt.Println("Task 10:")
	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 1, 1178}

	fmt.Println(makeGroups(arr))
}

func makeGroups(arr []float32) map[int][]float32 {
	groups := make(map[int][]float32)
	for _, val := range arr {
		r := int(val / 10)
		key := r * 10
		groups[key] = append(groups[key], val)
	}

	return groups
}
