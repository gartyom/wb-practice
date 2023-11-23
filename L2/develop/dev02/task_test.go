package main

import (
	"errors"
	"strings"
	"testing"
)

func assertErrors(err1 error, err2 error) bool {
	if err1 == nil && err2 != nil {
		return false
	}

	if err1 != nil && err2 == nil {
		return false
	}

	return true

}

func TestIsDigit(t *testing.T) {
	testTable := map[rune]bool{
		'a': false,
		'2': true,
	}

	for key, val := range testTable {
		if want := isDigit(key); want != val {
			t.Errorf("Output %v not equal to expected %v", want, val)
		}
	}
}

func TestIsEscape(t *testing.T) {
	testTable := map[rune]bool{
		'a':  false,
		'\\': true,
	}

	for key, val := range testTable {
		if want := isEscape(key); want != val {
			t.Errorf("\nhave %v \nwant %v", val, want)
		}
	}
}

type GetWholeNumberTest struct {
	ogIn      int
	iIn       []rune
	wNum      int
	wProcRune int
	wErr      error
}

func TestGetWholeNumber(t *testing.T) {
	testTable := []GetWholeNumberTest{
		{0, []rune("15273"), 15273, 5, nil},
		{0, []rune("0"), 0, 1, nil},
	}

	for _, test := range testTable {
		hNum, hProcRune, hErr := getWholeNumber(test.ogIn, test.iIn)
		if hProcRune != test.wProcRune || hNum != test.wNum || !assertErrors(hErr, test.wErr) {
			t.Errorf("\nwant %v, %v, %v \nexpected: %v, %v, %v", hProcRune, hErr, hNum, test.wProcRune, test.wErr, test.wNum)
		}
	}
}

type handleRuneTest struct {
	ogIn     []rune
	iIn      int
	resIn    *strings.Builder
	wPocRune int
	wErr     error
	wBuilder string
}

func TestHandleDefault(t *testing.T) {
	var builder strings.Builder
	testTable := []handleRuneTest{
		{[]rune("asd"), 0, &builder, 1, nil, "a"},
		{[]rune("asd"), 1, &builder, 1, nil, "s"},
		{[]rune("asd"), 2, &builder, 1, nil, "d"},
	}

	for _, test := range testTable {
		test.resIn.Reset()
		hPocRune, hErr := handleDefault(test.ogIn, test.iIn, test.resIn)
		if hPocRune != test.wPocRune || !assertErrors(hErr, test.wErr) || test.wBuilder != test.resIn.String() {
			t.Errorf("\nhave: %v, %v, %v \nwant: %v, %v, %v", hPocRune, hErr, test.resIn.String(), test.wPocRune, test.wErr, test.wBuilder)
		}
	}

}

func TestHandleDigit(t *testing.T) {
	var builder strings.Builder
	testTable := []handleRuneTest{
		{[]rune("a4"), 1, &builder, 1, nil, "aaa"},
		{[]rune("a10"), 1, &builder, 2, nil, "aaaaaaaaa"},
	}

	for _, test := range testTable {
		test.resIn.Reset()
		hPocRune, hErr := handleDigit(test.ogIn, test.iIn, test.resIn)
		if hPocRune != test.wPocRune || !assertErrors(hErr, test.wErr) || test.wBuilder != test.resIn.String() {
			t.Errorf("\nhave: %v, %v, %v \nwant: %v, %v, %v", hPocRune, hErr, test.resIn.String(), test.wPocRune, test.wErr, test.wBuilder)
		}
	}

}

func TestHandleEscape(t *testing.T) {
	testTable := []handleRuneTest{
		{[]rune("\\4"), 0, new(strings.Builder), 2, nil, "4"},
		{[]rune("\\\\"), 0, new(strings.Builder), 2, nil, "\\"},
		{[]rune("\\"), 0, new(strings.Builder), 0, errors.New(""), ""},
	}

	for _, test := range testTable {
		hPocRune, hErr := handleEscape(test.ogIn, test.iIn, test.resIn)
		if hPocRune != test.wPocRune || !assertErrors(hErr, test.wErr) || test.wBuilder != test.resIn.String() {
			t.Errorf("\nhave: %v, %v, %v \nwant: %v, %v, %v", hPocRune, hErr, test.resIn.String(), test.wPocRune, test.wErr, test.wBuilder)
		}
	}

}

type unpackStringTest struct {
	strIn  string
	wStr   string
	wError error
}

func TestUnpackString(t *testing.T) {
	testTable := []unpackStringTest{
		{"aa4\\5s", "aaaaa5s", nil},
		{"aa1s\\", "", errors.New("")},
		{"aa4\\\\5s", "aaaaa\\\\\\\\\\s", nil},
	}

	for _, test := range testTable {
		hStr, hErr := unpackString(test.strIn)
		if hStr != test.wStr || !assertErrors(hErr, test.wError) {
			t.Errorf("\nhave: %v, %v \nwant: %v, %v", hStr, hErr, test.wStr, test.wError)
		}
	}
}
