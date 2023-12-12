package pwd

import (
	"dev08/internal/pkg/printer"
	"fmt"
	"os"
)

func Run(args []string) error {
	path, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("pwd: %w", err)
	}

	printer.Println(path)
	return nil
}
