package app

import (
	"dev05/internal/pkg/args"
	"dev05/internal/pkg/grep"
	"dev05/internal/pkg/matcher"
	"dev05/internal/pkg/qprinter"
	"dev05/internal/pkg/queue"
)

func Run(args *args.Args) error {
	qprinter := qprinter.New(args.Count, args.LineNum)
	queue := queue.New(int(args.Before) + 1)
	matcher := matcher.New(args.IgnoreCase, args.Fixed)
	grep := grep.New(args, queue, qprinter, matcher)
	err := grep.Process()
	if err != nil {
		return err
	}
	return nil
}
