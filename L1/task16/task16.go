package task16

import "fmt"

func Run() {
	fmt.Println("Task 16:")
	arr := []int{1, 2, 345, 12, 5, 2, 1, 123, 5, 1, 5, 2, 6, 8, 12}
	fmt.Println(arr)

	l := 0
	r := len(arr) - 1
	quickSort(arr, l, r)

	fmt.Println(arr)
}

func quickSort(a []int, l, r int) {
	if l < r {
		q := partition(a, l, r)
		quickSort(a, l, q)
		quickSort(a, q+1, r)
	}
}

func partition(a []int, l, r int) int {
	v := a[(l+r)/2]
	i, j := l, r

	for i <= j {
		for a[i] < v {
			i++
		}

		for a[j] > v {
			j--
		}

		if i >= j {
			break
		}

		a[i], a[j] = a[j], a[i]
		i++
		j--
	}

	return j
}
