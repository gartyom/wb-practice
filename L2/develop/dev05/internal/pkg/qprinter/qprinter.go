package qprinter

import (
	"bufio"
	"dev05/internal/pkg/queue"
	"fmt"
)

type QPrinter struct {
	Writer     *bufio.Writer
	Counter    int
	Count      bool
	LineNum    bool
	WriterFunc func(p *QPrinter, s string, i int)
}

func New(writer *bufio.Writer, count bool, lineNum bool) *QPrinter {
	qp := &QPrinter{
		Writer:  writer,
		Counter: 0,
		Count:   count,
		LineNum: lineNum,
	}

	if lineNum {
		qp.WriterFunc = writeLineNum
	} else {
		qp.WriterFunc = write
	}

	if count {
		qp.WriterFunc = cnt
	}

	return qp
}

func (p *QPrinter) Write(e queue.QElem) {
	s, i := e.Data, e.Idx
	p.WriterFunc(p, s, i)
}

func (p *QPrinter) Flush() {
	if p.Count {
		p.Writer.WriteString(fmt.Sprintln(p.Counter))
	}
	p.Writer.Flush()
}

func write(p *QPrinter, s string, i int) {
	p.Writer.WriteString(s)
	p.Writer.WriteString("\n")
}
func writeLineNum(p *QPrinter, s string, i int) {
	p.Writer.WriteString(fmt.Sprintf("%d:%s", i, s))
	p.Writer.WriteString("\n")
}
func cnt(p *QPrinter, s string, i int) {
	p.Counter += 1
}
