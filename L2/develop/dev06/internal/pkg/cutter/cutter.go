package cutter

import (
	"dev06/internal/pkg/args"
	"strings"
)

func Cut(in <-chan string, args *args.Args) <-chan string {
	out := make(chan string)
	go func() {
		for data := range in {
			splitted := strings.Split(data, args.Delimiter)
			if args.Separated && len(splitted) < 2 {
				continue
			}

			left := args.Fields[0] - 1
			if left >= len(splitted) {
				continue
			}

			right := min(args.Fields[1], len(splitted))
			if right == 0 {
				out <- splitted[left]
				continue
			}

			out <- strings.Join(splitted[left:right], args.Delimiter)

		}
		close(out)
	}()

	return out
}
