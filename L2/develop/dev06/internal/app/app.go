package app

import (
	"bufio"
	"dev06/internal/pkg/args"
	"dev06/internal/pkg/cutter"
	"dev06/internal/pkg/printer"
	"dev06/internal/pkg/scanner"
	"os"
)

func Run(args *args.Args) {
	sc := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	scOut := scanner.Scan(sc)
	ctOut := cutter.Cut(scOut, args)
	pOut := printer.Print(ctOut, writer)

	<-pOut
}
