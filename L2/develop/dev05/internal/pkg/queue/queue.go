package queue

type QElem struct {
	Data  string
	Print bool
	Idx   int
}

func newQelem(s string, i int, p bool) QElem {
	return QElem{s, p, i}
}

type Queue struct {
	Data   []QElem
	Length int
}

func New(Length int) *Queue {
	array := make([]QElem, Length, Length)
	return &Queue{
		Data:   array,
		Length: Length,
	}
}

func (q *Queue) Append(s string, i int, p bool) QElem {
	first := q.Data[0]
	for i := 0; i < q.Length-1; i++ {
		q.Data[i] = q.Data[i+1]
	}

	q.Data[q.Length-1] = newQelem(s, i, p)

	return first

}

func (q *Queue) SetFlag(flag bool) {
	for i := 0; i < q.Length; i++ {
		if q.Data[i].Idx != 0 {
			q.Data[i].Print = flag
		}
	}
}
