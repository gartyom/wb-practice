package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую
повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать
ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	str := "a10sф3ыd\\42asf"
	res, err := unpackString(str)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
}

func unpackString(str string) (string, error) {
	og := []rune(str)
	var result strings.Builder

	for i := 0; i < len(og); {
		runesProcessed, err := handleRune(og, i, &result)
		if err != nil {
			return "", err
		}
		i += runesProcessed
	}
	return result.String(), nil
}

func handleRune(og []rune, i int, result *strings.Builder) (int, error) {
	switch {
	case isDigit(og[i]):
		return handleDigit(og, i, result)
	case isEscape(og[i]):
		return handleEscape(og, i, result)
	default:
		return handleDefault(og, i, result)
	}
}

func handleDigit(og []rune, i int, result *strings.Builder) (int, error) {
	if i < 1 {
		return 0, errors.New("Not a valid string")
	}

	num, runesProcessed, err := getWholeNumber(i, og)
	if err != nil {
		return 0, err
	}

	for j := 0; j < num-1; j++ {
		_, err := result.WriteRune(og[i-1])
		if err != nil {
			return 0, err
		}
	}

	return runesProcessed, nil
}

func handleEscape(og []rune, i int, result *strings.Builder) (int, error) {
	if i == len(og)-1 {
		return 0, errors.New("Unknown escape sequence")
	}

	if !(isDigit(og[i+1]) || isEscape(og[i+1])) {
		return 0, errors.New("Unknown escape sequence")
	}

	_, err := result.WriteRune(og[i+1])
	if err != nil {
		return 0, err
	}

	return 2, nil
}

func handleDefault(og []rune, i int, result *strings.Builder) (int, error) {
	_, err := result.WriteRune(og[i])
	return 1, err
}

func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func isEscape(r rune) bool {
	return '\\' == r
}

func getWholeNumber(i int, og []rune) (int, int, error) {
	var result strings.Builder
	result.WriteRune(og[i]) // already checked

	runesProcessed := 1

	for i+runesProcessed < len(og) {
		if !isDigit(og[i+runesProcessed]) {
			break
		}
		result.WriteRune(og[i+runesProcessed])
		runesProcessed++
	}

	number, err := strconv.Atoi(result.String())
	if err != nil {
		return 0, 0, err
	}
	return number, runesProcessed, nil
}
