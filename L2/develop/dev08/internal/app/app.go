package app

import (
	"bufio"
	"dev08/internal/pkg/input"
	"fmt"
	"os"
	"strings"
)

func Run() {

	reader := bufio.NewReader(os.Stdin)

	for true {
		path, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		dirs := strings.Split(path, "/")
		lastDir := dirs[len(dirs)-1]
		fmt.Fprintf(os.Stdout, "%s > ", lastDir)

		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		command = strings.Trim(command, "\n")
		command = strings.Trim(command, " ")

		if err := input.Handle(command); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
