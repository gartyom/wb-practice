package queue

import (
	"reflect"
	"testing"
)

func Test_New(t *testing.T) {
	tests := []struct {
		name   string
		hCap   int
		wQueue *Queue
	}{
		{"Default", 10, &Queue{nil, nil, 0, 0, false, 10}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			hQueue := New(tt.hCap)

			if !reflect.DeepEqual(hQueue, tt.wQueue) {
				t.Errorf("New():\nwant: %v\nhave: %v\n", tt.wQueue, hQueue)
			}
		})
	}
}

func Test_Append(t *testing.T) {
	QElem1 := QElem{"asd", true, 1, nil}
	hQueue1 := &Queue{nil, nil, 0, 0, false, 123}
	wQueue1 := &Queue{&QElem1, &QElem1, 1, 0, false, 123}

	QElem2 := QElem{"d", false, 12, nil}
	QElem3 := QElem1
	QElem3.Next = &QElem2
	hQueue2 := &Queue{&QElem1, &QElem1, 1, 0, false, 2}
	wQueue2 := &Queue{&QElem3, &QElem2, 2, 0, false, 2}

	QElem4 := QElem{"s", true, 3, nil}
	QElem5 := QElem2
	QElem5.Next = &QElem4
	hQueue3 := &Queue{&QElem3, &QElem2, 2, 0, false, 2}
	wQueue3 := &Queue{&QElem5, &QElem4, 2, 0, false, 2}

	tests := []struct {
		name   string
		hS     string
		hI     int
		hP     bool
		hQueue *Queue
		wQueue *Queue
		wEl    *QElem
	}{
		{"Append to empty queue", "asd", 1, true, hQueue1, wQueue1, nil},
		{"Append to non empty queue (len < cap)", "d", 12, false, hQueue2, wQueue2, nil},
		{"Append to non empty queue (len == cap)", "s", 3, true, hQueue3, wQueue3, &QElem{"asd", true, 1, nil}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hEl := tt.hQueue.Append(tt.hS, tt.hI, tt.hP)

			if !reflect.DeepEqual(hEl, tt.wEl) {
				t.Errorf("Append():\nwant el: %v\nhave el: %v\n", hEl, tt.wEl)
			}

			if !reflect.DeepEqual(tt.hQueue, tt.wQueue) {
				t.Errorf("Append():\nwant queue: %v\nhave queue: %v\n", tt.wQueue, tt.hQueue)
			}
		})
	}
}

func Test_SetFlag(t *testing.T) {
	tests := []struct {
		name   string
		hFlag  bool
		hQueue *Queue
		wQueue *Queue
	}{
		{"Set True", true, &Queue{nil, nil, 10, 0, false, 11}, &Queue{nil, nil, 10, 10, true, 11}},
		{"Set False", false, &Queue{nil, nil, 10, 0, true, 11}, &Queue{nil, nil, 10, 10, false, 11}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.hQueue.SetFlag(tt.hFlag)
			if !reflect.DeepEqual(tt.hQueue, tt.wQueue) {
				t.Errorf("SetFlag():\nwant: %v\nhave: %v\n", tt.wQueue, tt.hQueue)
			}
		})
	}
}
