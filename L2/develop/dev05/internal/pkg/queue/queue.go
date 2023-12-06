package queue

type QElem struct {
	Data  string
	Print bool
	Idx   int
	Next  *QElem
}

func newQelem(s string, i int, p bool) QElem {
	return QElem{Data: s, Print: p, Idx: i, Next: nil}
}

type Queue struct {
	First *QElem

	Last *QElem

	Length int

	// FlagCounter is a counter for printing strings with
	// before (-B) argument.
	FlagCounter int
	Flag        bool

	Cap int
}

func New(Cap int) *Queue {

	return &Queue{
		First:  nil,
		Last:   nil,
		Cap:    Cap,
		Length: 0,
	}
}

// Append new element and return first element from queue if
// capacity is reached else return nil
func (q *Queue) Append(s string, i int, p bool) *QElem {
	if q.Length < q.Cap {
		q.appendNew(s, i, p)
		return nil
	}

	// Put first element to the end so we dont create new QElem.
	if q.Cap > 1 {
		q.Last.Next = q.First
		q.First = q.First.Next
		q.Last = q.Last.Next

		q.Last.Next = nil
	}

	// Copy first element to a variable.
	fCopy := *q.Last

	// Modify last element.
	q.Last.Data = s
	q.Last.Print = p
	q.Last.Idx = i

	if q.FlagCounter > 0 && fCopy.Idx > 0 {
		fCopy.Print = q.Flag
		q.FlagCounter -= 1
	}

	return &fCopy
}

func (q *Queue) appendNew(s string, i int, p bool) {
	e := newQelem(s, i, p)
	if q.Length == 0 {
		q.First = &e
		q.Last = q.First
	} else {
		q.Last.Next = &e
		q.Last = q.Last.Next
	}
	q.Length += 1
}

func (q *Queue) SetFlag(flag bool) {
	q.FlagCounter = q.Length
	q.Flag = flag
}
