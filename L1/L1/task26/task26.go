package task26

import (
	"fmt"
	"unicode"
)

func Run() {
	fmt.Println()
	fmt.Println("Task 26:")

	str1 := "aA"
	str2 := "asdjbop13rmфыа"
	fmt.Println("str1: ", str1)
	fmt.Println("str2: ", str2)
	fmt.Println("str1 has only unique characters: ", uniqueCharacters(str1))
	fmt.Println("str2 has only unique characters: ", uniqueCharacters(str2))
}

func uniqueCharacters(str string) bool {
	tmp := make(map[rune]bool)
	for _, char := range str {
		char = unicode.ToLower(char)
		if _, ok := tmp[char]; ok == true {
			return false
		}
		tmp[char] = true
	}
	return true
}
