package qprinter

import (
	"bufio"
	"bytes"
	"dev05/internal/pkg/queue"
	"os"
	"reflect"
	"testing"
)

func Test_New(t *testing.T) {
	writer := bufio.NewWriter(os.Stdout)
	tests := []struct {
		name      string
		hCount    bool
		hLineNum  bool
		wQprinter *QPrinter
	}{
		{"Default write", false, false, &QPrinter{writer, 0, false, false, write}},
		{"Line number write", false, true, &QPrinter{writer, 0, false, true, writeLineNum}},
		{"Count write", true, false, &QPrinter{writer, 0, true, false, cnt}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hQPrinter := New(writer, tt.hCount, tt.hLineNum)

			if hQPrinter.Writer != tt.wQprinter.Writer ||
				hQPrinter.Counter != tt.wQprinter.Counter ||
				hQPrinter.Count != tt.wQprinter.Count ||
				hQPrinter.LineNum != tt.wQprinter.LineNum ||
				reflect.ValueOf(tt.wQprinter.WriterFunc).Pointer() != reflect.ValueOf(tt.wQprinter.WriterFunc).Pointer() {
				t.Errorf("New():\nwant: %v\nhave: %v", hQPrinter, tt.wQprinter)
			}
		})
	}
}

func Test_WriteFlush(t *testing.T) {
	tests := []struct {
		name    string
		wString string
		qe      queue.QElem
		qp      QPrinter
	}{
		{"Default", "abc\n", queue.QElem{Data: "abc", Print: true, Idx: 1, Next: nil}, QPrinter{nil, 0, false, false, write}},
		{"Line num", "15:test\n", queue.QElem{Data: "test", Print: true, Idx: 15, Next: nil}, QPrinter{nil, 0, false, false, writeLineNum}},
		{"Count", "1\n", queue.QElem{Data: "abc", Print: true, Idx: 1, Next: nil}, QPrinter{nil, 0, true, false, cnt}},
	}

	for _, tt := range tests {
		buf := new(bytes.Buffer)
		writer := bufio.NewWriter(buf)
		t.Run(tt.name, func(t *testing.T) {
			tt.qp.Writer = writer
			tt.qp.Write(tt.qe)
			tt.qp.Flush()

			if buf.String() != tt.wString {
				t.Errorf("Write() & Flush():\nwant: %v\nhave: %v", tt.wString, buf.String())
			}

		})
	}
}
