package errs

import (
	"fmt"
	"os"
)

func Check(err error) {
	if err != nil {
		exitGracefully(err)
	}
}

func exitGracefully(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}
