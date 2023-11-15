package task17

import "fmt"

func Run() {
	fmt.Println()
	fmt.Println("Task 17:")
	arr := []int{1, 2, 13, 14, 20, 22, 102, 104, 105}
	fmt.Println(arr)

	num := 102
	fmt.Println("Find:", num)
	idx, ok := binarySearch(arr, num, 0, len(arr))
	if ok {
		fmt.Println(idx)
	} else {
		fmt.Println("Element not found")
	}
}

func binarySearch(a []int, num, l, r int) (int, bool) {
	for l <= r {
		i := (l + r) / 2

		if a[i] > num {
			return binarySearch(a, num, l, i-1)
		} else if a[i] < num {
			return binarySearch(a, num, i+1, r)
		} else {
			return i, true
		}

	}

	return -1, false
}
