package matcher

import (
	"reflect"
	"testing"
)

func Test_New(t *testing.T) {
	tests := []struct {
		name         string
		hIghnoreCase bool
		hFixed       bool
		wMatcher     *Matcher
	}{
		{"Defalut", false, false, &Matcher{&Regexp{}}},
		{"Fixed", false, true, &Matcher{&Fixed{}}},
		{"Ignore case", true, false, &Matcher{&Lower{&Regexp{}}}},
		{"Ignore case & Fixed", true, true, &Matcher{&Lower{&Fixed{}}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hMathcer := New(tt.hIghnoreCase, tt.hFixed)

			if !reflect.DeepEqual(hMathcer, tt.wMatcher) {
				t.Errorf("New():\nwant: %v\nhave: %v\n", tt.wMatcher, hMathcer)
			}
		})
	}
}

func Test_Match(t *testing.T) {
	tests := []struct {
		name     string
		hPattern string
		hString  string
		wErr     bool
		wAns     bool
		hMatcher *Matcher
	}{
		{"Regexp equal", "[a-z]", "g", false, true, &Matcher{&Regexp{}}},
		{"Fixed equal", "[a-z]", "[a-z]", false, true, &Matcher{&Fixed{}}},
		{"Regexp not equal", "A", "a", false, false, &Matcher{&Regexp{}}},
		{"Fixed not equal", "ab", "abc", false, false, &Matcher{&Fixed{}}},
		{"Regexp invalid pattern", "[a", "abc", true, false, &Matcher{&Regexp{}}},
		{"Ignore case equal", "ab", "AB", false, true, &Matcher{&Lower{&Regexp{}}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hAns, hErr := tt.hMatcher.Match(tt.hPattern, tt.hString)

			if (hErr != nil) != tt.wErr {
				t.Errorf("New():\nwant error: %v\nhave error: %v\n", tt.wErr, hErr)
			}

			if !reflect.DeepEqual(hAns, tt.wAns) {
				t.Errorf("New():\nwant: %v\nhave: %v\n", tt.wAns, hAns)
			}
		})
	}
}
