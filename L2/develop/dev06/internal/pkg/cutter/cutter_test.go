package cutter

import (
	"dev06/internal/pkg/args"
	"testing"
)

func Test_Cut(t *testing.T) {
	tests := []struct {
		name  string
		hArgs *args.Args
		hIn   []string
		wOut  []string
	}{
		{"One field", &args.Args{Fields: [2]int{1, 0}, Delimiter: "-", Separated: false}, []string{"line-1-some-text", "line-2", "line3"}, []string{"line", "line", "line3"}},
		{"Full range", &args.Args{Fields: [2]int{1, 3}, Delimiter: "-", Separated: false}, []string{"line-1-some-text", "line-2", "line3"}, []string{"line-1-some", "line-2", "line3"}},
		{"Separated", &args.Args{Fields: [2]int{1, 3}, Delimiter: "-", Separated: true}, []string{"line-1-some-text", "line-2", "line3"}, []string{"line-1-some", "line-2"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := make(chan string)
			go func() {
				for _, data := range tt.hIn {
					in <- data
				}
				close(in)
			}()

			out := Cut(in, tt.hArgs)

			var i int
			for hOut := range out {
				if hOut != tt.wOut[i] {
					t.Errorf("Cut():\nwant: %v\nhave: %v\n", tt.wOut[i], hOut)
				}

				i += 1
			}
		})
	}
}
