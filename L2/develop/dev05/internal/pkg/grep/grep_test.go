package grep

import (
	"bufio"
	"bytes"
	"dev05/internal/pkg/args"
	"dev05/internal/pkg/matcher"
	"dev05/internal/pkg/qprinter"
	"dev05/internal/pkg/queue"
	"os"
	"testing"
)

func Test_openFile(t *testing.T) {
	tmp, _ := os.CreateTemp("", "tmp")
	wFile, _ := os.Open(tmp.Name())
	defer os.Remove(tmp.Name())
	tests := []struct {
		name  string
		hGrep *Grep
		wErr  bool
		wFile *os.File
	}{
		{"File exist", &Grep{args: &args.Args{FilePath: tmp.Name()}}, false, wFile},
		{"File doesn't exist", &Grep{args: &args.Args{FilePath: "/notadir/notafile"}}, true, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hFile, hErr := tt.hGrep.openFile()

			if (hErr != nil) != tt.wErr {
				t.Errorf("openFile():\nwant error: %v\nhave error: %v\n", tt.wErr, hErr)
			}

			if hFile != tt.wFile && hFile.Name() != tt.wFile.Name() {
				t.Errorf("openFile():\nwant: %v\nhave: %v\n", tt.wFile.Name(), hFile.Name())
			}
		})
	}
}

func Test_validatePattern(t *testing.T) {
	tests := []struct {
		name  string
		wErr  bool
		hGrep *Grep
	}{
		{"Valid pattrern", false, &Grep{args: &args.Args{Pattern: "[a-z]"}, m: &matcher.Matcher{Pipeline: &matcher.Regexp{}}}},
		{"Invalid pattrern", true, &Grep{args: &args.Args{Pattern: "[a-z"}, m: &matcher.Matcher{Pipeline: &matcher.Regexp{}}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hErr := tt.hGrep.validatePattern()

			if (hErr != nil) != tt.wErr {
				t.Errorf("validatePattern():\nwant error: %v\nhave error: %v\n", tt.wErr, hErr)
			}
		})
	}

}

func Test_Process(t *testing.T) {
	tests := []struct {
		name     string
		hGrep    *Grep
		wErr     bool
		hString  string
		wString  string
		hPattern string
	}{
		{"Default", &Grep{
			q:    &queue.Queue{Cap: 1},
			m:    matcher.New(false, false),
			args: &args.Args{},
		}, false, "line 1\nline 512\n", "line 1\nline 512\n", "[0-9]"},

		{"Print before = 2", &Grep{
			q:    &queue.Queue{Cap: 3},
			m:    matcher.New(false, false),
			args: &args.Args{},
		}, false, "line 1\nline 512\nline 6\nline 512\n", "line 1\nline 512\nline 6\n", "6"},

		{"Print after = 2", &Grep{
			q:    &queue.Queue{Cap: 1},
			m:    matcher.New(false, false),
			args: &args.Args{After: 2},
		}, false, "line 1\nline 6\nline 512\n", "line 6\nline 512\n", "6"},

		{"Before is bigger than file", &Grep{
			q:    &queue.Queue{Cap: 10000},
			m:    matcher.New(false, false),
			args: &args.Args{},
		}, false, "line 1\nline 6\nline 512\n", "line 1\nline 6\nline 512\n", "512"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			writer := bufio.NewWriter(buf)
			tt.hGrep.qp = qprinter.New(writer, false, false)

			tmp, _ := os.CreateTemp("", "tmp")
			defer os.Remove(tmp.Name())

			tmp.WriteString(tt.hString)

			tt.hGrep.args.FilePath = tmp.Name()
			tt.hGrep.args.Pattern = tt.hPattern

			hErr := tt.hGrep.Process()
			if (hErr != nil) != tt.wErr {
				t.Errorf("Process():\nwant error: %v\nhave error: %v\n", tt.wErr, tt.wErr)
			}

			if buf.String() != tt.wString {
				t.Errorf("Process():\nwant: %v\nhave: %v\n", tt.wString, buf.String())
			}

		})
	}
}
