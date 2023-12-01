package app

import (
	"dev05/internal/pkg/args"
	"dev05/internal/pkg/file"
	"dev05/internal/pkg/finder"
	"dev05/internal/pkg/printer"
)

func Run(args *args.Args) error {
	file, err := file.New(args.FilePath)
	if err != nil {
		return err
	}

	finder := finder.New(args.After, args.Before, args.IgnoreCase, args.Fixed)
	idxs, err := finder.Find(file, args.Pattern)
	if err != nil {
		return err
	}

	printer := printer.New(args.Count, args.Invert, args.LineNum)
	printer.Print(file, idxs)

	return nil
}
