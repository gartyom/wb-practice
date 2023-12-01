package printer

import (
	"bufio"
	"dev05/internal/pkg/file"
	"fmt"
	"os"
	"slices"
)

type Printer struct {
	Count    *bool
	Inverted *bool
	LineNum  *bool
}

func New(count *bool, invert *bool, lineNum *bool) *Printer {
	return &Printer{
		count,
		invert,
		lineNum,
	}
}

func (p *Printer) Print(f *file.File, idxs *[]int) {
	writer := bufio.NewWriter(os.Stdout)
	if *p.Inverted {
		idxs = p.Invert(idxs, len(f.Data))
	}

	if *p.Count {
		writer.WriteString(fmt.Sprintln(len(*idxs)))
		writer.Flush()
		return
	}

	for _, v := range *idxs {
		if *p.LineNum {
			writer.WriteString(fmt.Sprintf("%d:", v+1))
		}
		writer.WriteString(f.Data[v])
		writer.WriteString("\n")
	}

	writer.Flush()
}

func (p *Printer) Invert(idxs *[]int, fLen int) *[]int {
	inverted := make([]int, 0)

	for i := 0; i < fLen; i++ {
		if slices.Contains(*idxs, i) {
			l := len(*idxs)
			*idxs = (*idxs)[1:l]
			continue
		}
		inverted = append(inverted, i)
	}

	return &inverted
}
