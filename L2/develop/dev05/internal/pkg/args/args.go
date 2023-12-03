package args

import (
	"errors"
	"flag"
	"os"
)

type Args struct {
	FilePath   string
	Pattern    string
	After      uint
	Before     uint
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
}

func Get() (*Args, error) {
	l := len(os.Args)
	if l < 3 {
		return &Args{}, errors.New("A pattern and file path arguments is required")
	}

	FilePath := os.Args[l-1]
	Pattern := os.Args[l-2]

	After := flag.Uint("A", 0, "print +N lines after")
	Before := flag.Uint("B", 0, "print +N lines before")
	Context := flag.Uint("C", 0, "print +-N lines (A + B)")
	Count := flag.Bool("c", false, "count lines")
	IgnoreCase := flag.Bool("i", false, "ignore case")
	Invert := flag.Bool("v", false, "invert")
	Fixed := flag.Bool("F", false, "fixed strings")
	LineNum := flag.Bool("n", false, "print line number")

	flag.Parse()

	if (*Context) > (*After) {
		After = Context
	}

	if (*Context) > (*Before) {
		Before = Context
	}

	args := &Args{
		FilePath,
		Pattern,
		*After,
		*Before,
		*Count,
		*IgnoreCase,
		*Invert,
		*Fixed,
		*LineNum,
	}

	return args, nil
}
