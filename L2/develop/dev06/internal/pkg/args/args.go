package args

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type Args struct {
	Fields    FieldRange
	Delimiter string
	Separated bool
}

type FieldRange [2]int

func (this *FieldRange) Set(value string) error {
	borders := strings.Split(value, "-")
	if len(borders) < 1 || len(borders) > 2 {
		return errors.New("invalid field range")
	}

	left, err := strconv.Atoi(borders[0])
	if err != nil {
		return err
	}

	if len(borders) < 2 {
		(*this)[0] = left
		return nil
	}

	right, err := strconv.Atoi(borders[1])
	if err != nil {
		return err
	}

	if left < 0 || right < 0 {
		return errors.New("invelid field range")
	}

	(*this)[0] = left
	(*this)[1] = right

	return nil
}

func (this *FieldRange) String() string {
	return fmt.Sprint(*this)
}

func New() (*Args, error) {
	var fields FieldRange
	flag.Var(&fields, "f", "fields")

	delimiter := flag.String("d", " ", "change delimiter, default is one space")
	separated := flag.Bool("s", false, "print only separated strings")

	flag.Parse()
	if fields[0] == 0 {
		return nil, errors.New("invalid field range")
	}

	a := Args{
		Fields:    fields,
		Delimiter: *delimiter,
		Separated: *separated,
	}
	return &a, nil
}
