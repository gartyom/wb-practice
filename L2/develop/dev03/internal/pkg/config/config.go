package config

import (
	"errors"
	"flag"
	"os"
)

type Config struct {
	FilePath   string
	SortColumn int
	Numeric    bool
	Reverse    bool
	Unique     bool
}

func Get() (*Config, error) {
	if len(os.Args) < 2 {
		return &Config{}, errors.New("A file path argument is required")
	}

	filePath := os.Args[1]
	os.Args = os.Args[1:]

	sortColumn := flag.Int("k", 1, "sort by column")
	numeric := flag.Bool("n", false, "sort by numeric value")
	reverse := flag.Bool("r", false, "sort in reverse order")
	unique := flag.Bool("u", false, "print unique")

	flag.Parse()

	if *sortColumn <= 0 {
		return &Config{}, errors.New("k should be >= 1")
	}

	cfg := &Config{
		FilePath:   filePath,
		SortColumn: *sortColumn,
		Numeric:    *numeric,
		Reverse:    *reverse,
		Unique:     *unique,
	}

	return cfg, nil
}
