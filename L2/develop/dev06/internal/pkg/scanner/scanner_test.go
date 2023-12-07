package scanner

import (
	"bufio"
	"bytes"
	"testing"
)

func Test_Scan(t *testing.T) {
	tests := []struct {
		name string
		hBuf []byte
		wOut []string
	}{
		{"Default", []byte("hello world 1\nhello world 2\n"), []string{"hello world 1", "hello world 2"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			scanner := bufio.NewScanner(bytes.NewReader(tt.hBuf))
			out := Scan(scanner)

			var i int
			for hOut := range out {
				if hOut != tt.wOut[i] {
					t.Errorf("Scan():\nwant: %v\nhave: %v\n", tt.wOut[i], hOut)
				}
				i += 1
			}
		})
	}
}
