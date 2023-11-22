package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры
использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

// Стратегия — это поведенческий паттерн, выносит набор алгоритмов в
// собственные классы и делает их взаимозаменимыми. Другие объекты содержат
// ссылку на объект-стратегию и делегируют ей работу. Программа может подменить
// этот объект другим, если требуется иной способ решения задачи.

type Strategy func(*[]int) *[]int

func QuickSort(array *[]int) *[]int {
	fmt.Println("quicksorted array")
	return array
}

func MergeSort(array *[]int) *[]int {
	fmt.Println("mergesorted array")
	return array
}

type cache struct {
	array    *[]int
	sortFunc Strategy
}

func (cch *cache) sort() {
	cch.array = cch.sortFunc(cch.array)
}

func strategyExample() {
	cch := &cache{
		array: &[]int{},
	}

	cch.sortFunc = QuickSort
	cch.sort()
	cch.sortFunc = MergeSort
	cch.sort()
}

// Так же Strategy может быть интерфейсом с методом sort, а quickSort &
// mergeSort могут быть реализованы как структуры с методом sort

// +
//  Реализует принцип открытости/закрытости.
//  Изолирует код и данные алгоритмов от остальных классов.
//  Замена алгоритмов на лету.

// -
// Усложняет программу за счёт дополнительных классов/интерфейсов
