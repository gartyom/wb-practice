package ps

import (
	"bufio"
	"dev08/internal/pkg/printer"
	"fmt"
	"os"
	"strings"
)

func Run(args []string) error {
	pids, err := os.ReadDir("/proc")
	if err != nil {
		return fmt.Errorf("ps: %w", err)
	}

	printer.Println(fmt.Sprintf("%20s%6s%6s", "Name", "Pid", "PPid"))
	for _, dir := range pids {
		if dir.IsDir() {
			var name, ppid string
			f, err := os.Open("/proc/" + dir.Name() + "/status")
			if err != nil {
				continue
			}

			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				fields := strings.Split(scanner.Text(), ":")
				val := strings.Trim(fields[1], " ")
				val = strings.Trim(fields[1], "\t\r\n")
				switch fields[0] {
				case "Name":
					name = val
				case "PPid":
					ppid = val
				}
			}
			if name != "" && ppid != "" {
				printer.Println(fmt.Sprintf("%25s%6s%6s", name, dir.Name(), ppid))
			}
		}
	}

	return err
}
