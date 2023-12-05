package qprinter

import (
	"bufio"
	"dev05/internal/pkg/queue"
	"fmt"
	"os"
)

type QPrinter struct {
	Writer     *bufio.Writer
	Counter    int
	Count      bool
	LineNum    bool
	WriterFunc func(s string, i int)
}

func New(count bool, lineNum bool) *QPrinter {
	qp := &QPrinter{
		Writer:  bufio.NewWriter(os.Stdout),
		Counter: 0,
		Count:   count,
		LineNum: lineNum,
	}

	if lineNum {
		qp.WriterFunc = qp.writeLineNum
	} else {
		qp.WriterFunc = qp.write
	}

	if count {
		qp.WriterFunc = qp.count
	}

	return qp
}

func (p *QPrinter) Write(e queue.QElem) {
	s, i := e.Data, e.Idx
	p.WriterFunc(s, i)
}

func (p *QPrinter) Flush() {
	if p.Count {
		p.Writer.WriteString(fmt.Sprintln(p.Counter))
	}
	p.Writer.Flush()
}

func (p *QPrinter) write(s string, i int) {
	p.Writer.WriteString(s)
	p.Writer.WriteString("\n")
}
func (p *QPrinter) writeLineNum(s string, i int) {
	p.Writer.WriteString(fmt.Sprintf("%d:%s", i, s))
	p.Writer.WriteString("\n")
}
func (p *QPrinter) count(s string, i int) {
	p.Counter += 1
}
