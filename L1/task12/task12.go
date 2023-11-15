package task12

import (
	"fmt"
	"strings"
)

func Run() {
	fmt.Println("Task 12:")
	str := "cat, cat, dog, cat, tree"
	subs := makeSubsets(str)
	fmt.Println(subs)
}

func makeSubsets(str string) [][]string {
	strArray := strings.Split(str, ",")
	set := makeSet(strArray)

	var subsets [][]string

	p := len(set) - 1 // количество элементов в собственном множестве
	for p > 0 {

		for i := 0; i < len(set); i++ { // проходим все элементы
			var subset []string
			for j := i; j < i+p; j++ { // составляем собственное множество
				index := j
				if j >= len(set) {
					index = j - len(set)
				}
				subset = append(subset, set[index])
			}

			subsets = append(subsets, subset)
		}

		p -= 1
	}

	subsets = append(subsets, make([]string, 0))

	return subsets
}

func makeSet(array []string) []string {
	helper := map[string]int{}
	set := []string{}
	for _, val := range array {
		val = strings.TrimSpace(val)
		if el := helper[val]; el == 0 {
			helper[val] = 1
			set = append(set, val)
		}
	}

	return set
}
