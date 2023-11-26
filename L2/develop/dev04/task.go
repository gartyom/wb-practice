package main

import (
	"fmt"
	"slices"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	array := []string{"пятка", "пятка", "тяпка", "пятак", "листок", "слиток"}

	m := findAnagrams(&array)
	printMap(m)
}

func findAnagrams(words *[]string) *map[string]*[]string {
	anagramMap := make(map[string]*[]string)
	if len(*words) < 1 {
		return &anagramMap
	}

	for _, word := range *words {
		if len(word) == 1 {
			continue
		}
		word := strings.ToLower(word)
		sortedWord := sortString(word)
		v, ok := anagramMap[sortedWord]

		if !ok {
			anagramMap[sortedWord] = &[]string{word}
			continue
		}

		if !slices.Contains(*v, word) {
			*v = append(*v, word)
		}
	}

	resultMap := make(map[string]*[]string)
	for k, v := range anagramMap {
		firstEl := (*v)[0]
		slices.Sort(*v)
		resultMap[firstEl] = v

		delete(anagramMap, k)
	}

	return &resultMap
}

func sortString(s string) string {
	r := []rune(s)
	slices.Sort(r)
	return string(r)
}

func printMap(m *map[string]*[]string) {
	for k, v := range *m {
		fmt.Println(k, *v)
	}
}
