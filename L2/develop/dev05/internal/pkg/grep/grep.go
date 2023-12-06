package grep

import (
	"bufio"
	"dev05/internal/pkg/args"
	"dev05/internal/pkg/matcher"
	"dev05/internal/pkg/qprinter"
	"dev05/internal/pkg/queue"
	"os"
)

type Grep struct {
	q    *queue.Queue
	qp   *qprinter.QPrinter
	m    *matcher.Matcher
	args *args.Args
}

func New(args *args.Args, q *queue.Queue, qp *qprinter.QPrinter, m *matcher.Matcher) *Grep {
	return &Grep{
		q:    q,
		qp:   qp,
		m:    m,
		args: args,
	}
}

func (g *Grep) Process() error {
	in, err := g.openFile()
	if err != nil {
		return err
	}

	err = g.validatePattern()
	if err != nil {
		return err
	}

	// i - line counter. j - counter for After(-A) flag
	var i, j int

	n := int(g.args.After)

	var qe *queue.QElem
	pattern := g.args.Pattern
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		i += 1
		text := scanner.Text()
		eq, _ := g.m.Match(pattern, text)
		if j > 0 {
			qe = g.q.Append(text, i, !g.args.Invert)
			j--
		} else {
			qe = g.q.Append(text, i, g.args.Invert)
		}

		if qe != nil && qe.Print {
			g.qp.Write(*qe)
		}

		if eq {
			g.q.SetFlag(!g.args.Invert)
			j = n
		}

	}

	g.printRemainingElements()
	g.qp.Flush()

	return nil
}

func (g *Grep) openFile() (*os.File, error) {
	fp := g.args.FilePath
	_, err := os.Stat(fp)
	if err != nil {
		return nil, err
	}

	in, _ := os.Open(fp)

	return in, nil
}

func (g *Grep) validatePattern() error {
	_, err := g.m.Match(g.args.Pattern, "A")
	return err
}

func (g *Grep) printRemainingElements() {
	qe := g.q.First
	for qe != nil {
		if g.q.FlagCounter > 0 {
			qe.Print = g.q.Flag
			g.q.FlagCounter -= 1
		}
		if qe.Print {
			g.qp.Write(*qe)
		}
		qe = qe.Next
	}
}
