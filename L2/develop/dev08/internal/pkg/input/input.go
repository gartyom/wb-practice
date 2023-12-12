package input

import (
	"dev08/internal/pkg/printer"
	"dev08/internal/pkg/shell/cd"
	"dev08/internal/pkg/shell/echo"
	"dev08/internal/pkg/shell/exec"
	"dev08/internal/pkg/shell/kill"
	"dev08/internal/pkg/shell/ps"
	"dev08/internal/pkg/shell/pwd"
	"errors"
	"fmt"
	"os"
	"strings"
	"syscall"
)

func delegate(args []string) error {
	switch args[0] {
	case "exit":
		os.Exit(0)
	case "cd":
		return cd.Run(args)
	case "pwd":
		return pwd.Run(args)
	case "echo":
		return echo.Run(args)
	case "kill":
		return kill.Run(args)
	case "exec":
		return exec.Run(args)
	case "ps":
		return ps.Run(args)
	case "fork":
		ret, _, err := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
		if err != 0 {
			os.Exit(1)
		}

		if ret == 0 {
			pid := fmt.Sprint(os.Getpid())
			printer.Println(pid)
			err := delegate(args[1:])
			if err != nil {
				return err
			}
			printer.Println("Done: " + pid)
			os.Exit(0)
		}

		return nil

	default:
		return errors.New(fmt.Sprintf("command not found: %s", args[0]))
	}
	return nil
}

func Handle(cmd string) error {
	pipe := strings.Split(cmd, "|")
	for _, cmd := range pipe {

		cmd = strings.Trim(cmd, " ")
		args := strings.Split(cmd, " ")

		err := delegate(args)
		if err != nil {
			return err
		}
	}

	return nil
}
