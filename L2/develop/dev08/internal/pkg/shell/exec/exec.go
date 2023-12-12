package exec

import (
	"dev08/internal/pkg/printer"
	"fmt"
	"os/exec"
)

func Run(args []string) error {
	cmd := exec.Command(args[1], args[2:]...)
	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	printer.Print(string(out))
	return nil
}
