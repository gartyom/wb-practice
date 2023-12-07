package scanner

import "bufio"

func Scan(scanner *bufio.Scanner) <-chan string {
	out := make(chan string)
	go func() {
		for scanner.Scan() {
			out <- scanner.Text()
		}
		close(out)
	}()

	return out
}
