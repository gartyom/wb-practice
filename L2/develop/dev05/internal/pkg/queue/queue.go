package queue

type QElem struct {
	Data  string
	Print bool
	Idx   int
	next  *QElem
}

func newQelem(s string, i int, p bool) QElem {
	return QElem{Data: s, Print: p, Idx: i, next: nil}
}

type Queue struct {
	First     *QElem
	Last      *QElem
	Length    int
	FlagCount int
	Flag      bool
}

func New(Length int) *Queue {
	first := newQelem("", 0, false)
	last := &first
	for i := 0; i < Length-1; i++ {
		e := newQelem("", 0, false)
		last.next = &e
		last = &e
	}

	return &Queue{
		First:  &first,
		Last:   last,
		Length: Length,
	}
}

func (q *Queue) Append(s string, i int, p bool) QElem {
	if q.Length > 1 {
		first := q.First

		q.First = q.First.next
		q.Last.next = first
		q.Last = q.Last.next

		q.Last.next = nil
	}

	fCopy := *q.Last
	q.Last.Data = s
	q.Last.Print = p
	q.Last.Idx = i
	if q.FlagCount > 0 {
		if fCopy.Idx > 0 {
			fCopy.Print = q.Flag
		}
		q.FlagCount -= 1
	}
	return fCopy

}

func (q *Queue) SetFlag(flag bool) {
	q.FlagCount = q.Length
	q.Flag = flag
}
