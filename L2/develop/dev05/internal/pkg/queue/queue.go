package queue

type QElem struct {
	Data  string // string to be printed.
	Print bool   // flag whether or not we should print this element.
	Idx   int    // string index.
	Next  *QElem // next element.
}

func newQelem(s string, i int, p bool) QElem {
	return QElem{Data: s, Print: p, Idx: i, Next: nil}
}

type Queue struct {
	// Pointer to first queue element.
	First *QElem

	// Pointer to last queue element.
	Last *QElem

	// Current length.
	Length int

	// FlagCounter is a counter for printing strings with
	// before (-B) argument.
	FlagCounter int
	Flag        bool

	// Max queue capacity. Length cannot be greater than this.
	Cap int
}

func New(Cap int) *Queue {
	first := newQelem("", 0, false)
	last := &first

	return &Queue{
		First:  &first,
		Last:   last,
		Cap:    Cap,
		Length: 1,
	}
}

// Append new element and return first element from queue if
// capacity is reached else return nil
func (q *Queue) Append(s string, i int, p bool) *QElem {
	if q.Length < q.Cap {
		q.appendNew(s, i, p)
		return nil
	}

	// If length is equal to capacity and cap > 1 put first
	// element to the end so we dont allocate new memory.
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
	q.Last.Next = &e
	q.Last = &e
	q.Length += 1
}

func (q *Queue) SetFlag(flag bool) {
	q.FlagCounter = q.Length
	q.Flag = flag
}
