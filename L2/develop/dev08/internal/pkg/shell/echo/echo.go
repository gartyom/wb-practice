package echo

import (
	"dev08/internal/pkg/printer"
)

func Run(args []string) error {
	if len(args) < 2 {
		printer.Print("")
		return nil
	}

	printer.Print(args[1:]...)
	printer.Println("")
	return nil
}
