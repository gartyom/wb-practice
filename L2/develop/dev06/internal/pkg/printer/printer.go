package printer

import (
	"bufio"
)

func Print(in <-chan string, writer *bufio.Writer) <-chan struct{} {
	out := make(chan struct{})
	go func() {
		for data := range in {
			writer.WriteString(data)
			writer.WriteString("\n")
		}

		writer.Flush()
		out <- struct{}{}
	}()

	return out
}
