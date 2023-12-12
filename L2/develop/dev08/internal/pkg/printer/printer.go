package printer

import (
	"fmt"
	"os"
)

func Println(strings ...string) {
	for _, str := range strings {
		fmt.Fprintln(os.Stdout, str)
	}
}

func Print(strings ...string) {
	for _, str := range strings[:len(strings)-1] {
		fmt.Fprint(os.Stdout, str)
		fmt.Fprint(os.Stdout, " ")
	}
	fmt.Fprint(os.Stdout, strings[len(strings)-1])
}
