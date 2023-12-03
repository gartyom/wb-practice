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

	var i, j int
	var qe queue.QElem
	n := int(g.args.After)
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

		if qe.Print {
			g.qp.Write(qe)
		}

		if eq {
			g.q.SetFlag(!g.args.Invert)
			j = n
		}

	}

	for i := 0; i < g.q.Length; i++ {
		qe = g.q.Append("", 0, false)
		if qe.Print {
			g.qp.Write(qe)
		}
	}

	g.qp.Flush()

	return nil
}
