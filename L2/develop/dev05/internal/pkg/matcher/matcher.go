package matcher

import (
	"regexp"
	"strings"
)

type Matcher struct {
	IgnoreCase *bool
	Fixed      *bool
}

func New(ignoreCase *bool, fixed *bool) *Matcher {
	return &Matcher{ignoreCase, fixed}
}

func (m *Matcher) Match(pattern string, s string) (bool, error) {
	if *m.IgnoreCase {
		pattern = strings.ToLower(pattern)
		s = strings.ToLower(s)
	}

	if *m.Fixed {
		return matchFixed(pattern, s)
	}

	return match(pattern, s)
}

func match(pattern string, s string) (bool, error) {
	return regexp.MatchString(pattern, s)
}

func matchFixed(pattern string, s string) (bool, error) {
	return pattern == s, nil
}
