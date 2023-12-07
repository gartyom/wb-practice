package printer

import (
	"bufio"
	"bytes"
	"reflect"
	"testing"
)

func Test_Print(t *testing.T) {
	tests := []struct {
		name     string
		hStrings []string
		wBuf     []byte
	}{
		{"Default", []string{"line 123", "line 345"}, []byte("line 123\nline 345\n")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := make(chan string)
			var hBuf bytes.Buffer
			go func(in chan string) {
				for _, data := range tt.hStrings {
					in <- data
				}
				close(in)
			}(in)
			out := Print(in, bufio.NewWriter(&hBuf))
			<-out

			if !reflect.DeepEqual(hBuf.Bytes(), tt.wBuf) {
				t.Errorf("Print()\nwant: %v\nhave: %v\n", tt.wBuf, hBuf.Bytes())
			}
		})
	}
}
