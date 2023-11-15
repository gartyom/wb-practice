package task15

import (
	"fmt"
)

var justString string

func Run() {
	fmt.Println()
	fmt.Println("Task 15:")
	s := "фвыощфышгваорфголаjlhasdfioujhkashdf"
	someFunc1(s)
	someFunc2(s)
	someFunc3(s)
}

func someFunc1(s string) {
	fmt.Println("Копируются первые 10 байт:\t", s[:10])
}

func someFunc2(s string) {
	fmt.Println("Копируются первые 10 рун:\t", string([]rune(s)[:10]))
}

func someFunc3(s string) {
	j := 0
	fmt.Print("Копируются первые 10 рун:\t ")
	for _, v := range s {
		if j == 10 {
			break
		}
		j++
		fmt.Print(string(v))
	}
	fmt.Println()
}

/*
В данном случае копируются первые 100 байт строки, а не первые 100 символов
т. к. в Go строки эквивалентны []byte


var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}


func main() {
  someFunc()
}

*/
