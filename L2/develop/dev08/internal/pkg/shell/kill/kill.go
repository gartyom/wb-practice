package kill

import (
	"errors"
	"fmt"
	"strconv"
	"syscall"
)

func Run(args []string) error {
	if len(args) < 2 {
		return errors.New("ps: pid should be specified")
	}

	for _, arg := range args[1:] {
		pid, err := strconv.Atoi(arg)
		if err != nil {
			return err
		}

		err = syscall.Kill(pid, syscall.SIGKILL)
		if err != nil {
			return fmt.Errorf("ps: %w", err)
		}
	}

	return nil
}
