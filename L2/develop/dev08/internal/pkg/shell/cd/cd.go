package cd

import (
	"errors"
	"fmt"
	"os"
)

func Run(args []string) error {
	if len(args) > 2 {
		return fmt.Errorf("cd: %w", errors.New("too many arguments"))
	}

	if len(args) == 1 {
		hDir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("cd: %w", err)
		}
		err = os.Chdir(hDir)
		if err != nil {
			return fmt.Errorf("cd: %w", err)
		}
	} else {
		err := os.Chdir(args[1])
		if err != nil {
			return fmt.Errorf("cd: %w", err)
		}
	}

	return nil
}
