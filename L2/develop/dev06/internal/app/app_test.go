package app

import (
	"bufio"
	"dev06/internal/pkg/args"
	"os"
	"testing"
)

func Test_Run(t *testing.T) {
	tests := []struct {
		name    string
		args    *args.Args
		hStdin  string
		wStdout []string
	}{
		{"One field", &args.Args{Fields: [2]int{2, 0}, Delimiter: " "}, "word1 word2 word3\nword4 word5 word6\nword8", []string{"word2", "word5"}},
		{"Full range", &args.Args{Fields: [2]int{2, 3}, Delimiter: " "}, "word1 word2 word3\nword4 word5 word6\nword8", []string{"word2 word3", "word5 word6"}},
		{"Change delimiter", &args.Args{Fields: [2]int{2, 3}, Delimiter: ","}, "word1,word2,word3\nword4,word5,word6\nword8", []string{"word2,word3", "word5,word6"}},
		{"Separated", &args.Args{Fields: [2]int{2, 3}, Delimiter: ",", Separated: true}, "word1,word2,word3\nword", []string{"word2,word3"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in, err := os.CreateTemp("", "tmp")
			if err != nil {
				panic(err)
			}
			defer os.Remove(in.Name())
			if _, err := in.Write([]byte(tt.hStdin)); err != nil {
				panic(err)
			}
			if _, err = in.Seek(0, 0); err != nil {
				panic(err)
			}
			out, err := os.CreateTemp("", "tmp")
			if err != nil {
				panic(err)
			}
			defer os.Remove(out.Name())

			oldStdin := os.Stdin
			oldStdout := os.Stdout
			defer func() { os.Stdin = oldStdin }()
			os.Stdin = in
			os.Stdout = out

			Run(tt.args)

			os.Stdout.Close()
			os.Stdout = oldStdout

			newf, err := os.Open(out.Name())
			if err != nil {
				panic(err)
			}
			sc := bufio.NewScanner(newf)

			var i int
			for sc.Scan() {
				hStdout := sc.Text()
				if hStdout != tt.wStdout[i] {
					t.Errorf("Run():\nwant: %v\nhave: %v\n", tt.wStdout[i], hStdout)
				}
				i += 1
			}

			if i != len(tt.wStdout) {
				t.Errorf("Run():\nwant: %v\nhave: %v\n", tt.wStdout[i:], "")
			}

		})
	}
}
