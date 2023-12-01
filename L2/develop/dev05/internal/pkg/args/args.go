package args

import (
	"errors"
	"flag"
	"os"
)

type Args struct {
	FilePath   string
	Pattern    string
	After      *uint
	Before     *uint
	Count      *bool
	IgnoreCase *bool
	Invert     *bool
	Fixed      *bool
	LineNum    *bool
}

func Get() (*Args, error) {
	l := len(os.Args)
	args := &Args{}
	if l < 3 {
		return args, errors.New("A pattern and file path arguments is required")
	}

	args.FilePath = os.Args[l-1]
	args.Pattern = os.Args[l-2]

	args.After = flag.Uint("A", 0, "print +N lines after")
	args.Before = flag.Uint("B", 0, "print +N lines before")
	context := flag.Uint("C", 0, "print +-N lines (A + B)")
	args.Count = flag.Bool("c", false, "count lines")
	args.IgnoreCase = flag.Bool("i", false, "ignore case")
	args.Invert = flag.Bool("v", false, "invert")
	args.Fixed = flag.Bool("F", false, "fixed strings")
	args.LineNum = flag.Bool("n", false, "print line number")

	flag.Parse()

	if *context > *args.After {
		args.After = context
	}

	if *context > *args.Before {
		args.Before = context
	}

	return args, nil
}
