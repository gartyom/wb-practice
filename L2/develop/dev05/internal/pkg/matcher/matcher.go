package matcher

import (
	"regexp"
	"strings"
)

type Matcher struct {
	pipline Pipeline
}

type Pipeline interface {
	Process(pattern string, s string) (bool, error)
	SetNext(p Pipeline)
}

func New(ignoreCase bool, fixed bool) *Matcher {
	var first Pipeline

	first = &Regexp{}
	if fixed {
		first = &Fixed{}
	}

	if ignoreCase {
		first = &Lower{next: first}
	}

	return &Matcher{first}
}

func (m *Matcher) Match(pattern string, s string) (bool, error) {
	return m.pipline.Process(pattern, s)
}

type Lower struct {
	next Pipeline
}

func (l *Lower) Process(pattern string, s string) (bool, error) {
	pattern = strings.ToLower(pattern)
	s = strings.ToLower(s)
	return l.next.Process(pattern, s)
}
func (l *Lower) SetNext(p Pipeline) {
	l.next = p
}

type Fixed struct {
}

func (f *Fixed) Process(pattern string, s string) (bool, error) {
	return pattern == s, nil
}
func (l *Fixed) SetNext(p Pipeline) {
}

type Regexp struct {
}

func (r *Regexp) Process(pattern string, s string) (bool, error) {
	return regexp.MatchString(pattern, s)
}
func (l *Regexp) SetNext(p Pipeline) {
}
