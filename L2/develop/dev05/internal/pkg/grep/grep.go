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
	fp := g.args.FilePath
	_, err := os.Stat(fp)
	if err != nil {
		return err
	}

	in, _ := os.Open(fp)
	defer in.Close()

	_, err = g.m.Match(g.args.Pattern, "A")
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

	// Print print all remaining elements
	qe = g.q.First
	for qe != nil {
		if qe.Idx > 0 {
			if g.q.FlagCounter > 0 {
				qe.Print = g.q.Flag
				g.q.FlagCounter -= 1
			}
			if qe.Print {
				g.qp.Write(*qe)
			}
		} else {
			g.q.FlagCounter -= 1
		}
		qe = qe.Next
	}

	g.qp.Flush()

	return nil
}
