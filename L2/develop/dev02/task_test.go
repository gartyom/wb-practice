package main

import (
	"strings"
	"testing"
)

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
	wErr      bool
}

func TestGetWholeNumber(t *testing.T) {
	testTable := []GetWholeNumberTest{
		{0, []rune("15273"), 15273, 5, false},
		{0, []rune("0"), 0, 1, false},
	}

	for _, test := range testTable {
		hNum, hProcRune, hErr := getWholeNumber(test.ogIn, test.iIn)
		if hProcRune != test.wProcRune || hNum != test.wNum || (hErr != nil) != test.wErr {
			t.Errorf("\nwant %v, %v, %v \nexpected: %v, %v, %v", hProcRune, hErr, hNum, test.wProcRune, test.wErr, test.wNum)
		}
	}
}

type handleRuneTest struct {
	ogIn     []rune
	iIn      int
	resIn    *strings.Builder
	wPocRune int
	wErr     bool
	wBuilder string
}

func TestHandleDefault(t *testing.T) {
	var builder strings.Builder
	testTable := []handleRuneTest{
		{[]rune("asd"), 0, &builder, 1, false, "a"},
		{[]rune("asd"), 1, &builder, 1, false, "s"},
		{[]rune("asd"), 2, &builder, 1, false, "d"},
	}

	for _, test := range testTable {
		test.resIn.Reset()
		hPocRune, hErr := handleDefault(test.ogIn, test.iIn, test.resIn)
		if hPocRune != test.wPocRune || (hErr != nil) != test.wErr {
			t.Errorf("\nhave: %v, %v, %v \nwant: %v, %v, %v", hPocRune, hErr, test.resIn.String(), test.wPocRune, test.wErr, test.wBuilder)
		}
	}

}

func TestHandleDigit(t *testing.T) {
	var builder strings.Builder
	testTable := []handleRuneTest{
		{[]rune("a4"), 1, &builder, 1, false, "aaa"},
		{[]rune("a10"), 1, &builder, 2, false, "aaaaaaaaa"},
	}

	for _, test := range testTable {
		test.resIn.Reset()
		hPocRune, hErr := handleDigit(test.ogIn, test.iIn, test.resIn)
		if hPocRune != test.wPocRune || (hErr != nil) != test.wErr {
			t.Errorf("\nhave: %v, %v, %v \nwant: %v, %v, %v", hPocRune, hErr, test.resIn.String(), test.wPocRune, test.wErr, test.wBuilder)
		}
	}

}

func TestHandleEscape(t *testing.T) {
	testTable := []handleRuneTest{
		{[]rune("\\4"), 0, new(strings.Builder), 2, false, "4"},
		{[]rune("\\\\"), 0, new(strings.Builder), 2, false, "\\"},
		{[]rune("\\"), 0, new(strings.Builder), 0, true, ""},
	}

	for _, test := range testTable {
		hPocRune, hErr := handleEscape(test.ogIn, test.iIn, test.resIn)
		if hPocRune != test.wPocRune || (hErr != nil) != test.wErr {
			t.Errorf("\nhave: %v, %v, %v \nwant: %v, %v, %v", hPocRune, hErr, test.resIn.String(), test.wPocRune, test.wErr, test.wBuilder)
		}
	}

}

type unpackStringTest struct {
	strIn string
	wStr  string
	wErr  bool
}

func TestUnpackString(t *testing.T) {
	testTable := []unpackStringTest{
		{"aa4\\5s", "aaaaa5s", false},
		{"aa1s\\", "", true},
		{"aa4\\\\5s", "aaaaa\\\\\\\\\\s", false},
	}

	for _, test := range testTable {
		hStr, hErr := unpackString(test.strIn)
		if hStr != test.wStr || (hErr != nil) != test.wErr {
			t.Errorf("\nhave: %v, %v \nwant: %v, %v", hStr, hErr, test.wStr, test.wErr)
		}
	}
}
