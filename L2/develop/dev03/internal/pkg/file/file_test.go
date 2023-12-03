package file

import (
	"os"
	"reflect"
	"testing"

	"dev03/internal/pkg/config"
)

func Test_validate(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test.txt")
	if err != nil {
		panic(err) // This should never happen
	}

	// Once all the tests are done. Delete tmp file
	defer os.Remove(tmpFile.Name())

	tests := []struct {
		name     string
		filePath string
		wErr     bool
	}{
		{"File exist", tmpFile.Name(), false},
		{"File not exist", "notadir/file.txt", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hErr := validate(tt.filePath)
			if (hErr != nil) != tt.wErr {
				t.Errorf("Validate():\nwant error: %v\nhave error: %v\n", tt.wErr, hErr)
			}
		})
	}

}

func Test_read(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test.txt")
	if err != nil {
		panic(err)
	}

	os.WriteFile(tmpFile.Name(), []byte("a\nabc d\n"), 0644)
	defer os.Remove(tmpFile.Name())

	tmpFileEmpty, err := os.CreateTemp("", "testEmpty.txt")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFileEmpty.Name())

	tests := []struct {
		name   string
		fp     string
		wFData *[][]string
		wErr   bool
	}{
		{"Default", tmpFile.Name(), &[][]string{{"a"}, {"abc", "d"}}, false},
		{"Empty file", tmpFileEmpty.Name(), nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hFData, hErr := read(tt.fp)

			if (hErr != nil) != tt.wErr {
				t.Errorf("read():\nwant error: %v\nhave error: %v", tt.wErr, hErr)
				return
			}

			if !reflect.DeepEqual(hFData, tt.wFData) {
				t.Errorf("read():\nwant: %v\nhave: %v\n", hFData, tt.wFData)
				return
			}
		})
	}
}

func Test_cmpStrings(t *testing.T) {
	tests := []struct {
		name   string
		first  string
		second string
		wAns   int
	}{
		{"first < second", "a", "b", -1},
		{"first > second", "ac", "ab", 1},
		{"first = second", "abc", "abc", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hAns := cmpStrings(tt.first, tt.second)

			if !reflect.DeepEqual(hAns, tt.wAns) {
				t.Errorf("cmpStrings():\nwant: %v\nhave: %v\n", hAns, tt.wAns)
				return
			}
		})
	}
}

func Test_cmpNumeric(t *testing.T) {
	tests := []struct {
		name   string
		first  string
		second string
		wAns   int
	}{
		{"first < second", "1", "2", -1},
		{"first > second", "123", "13", 1},
		{"first = second", "11.0", "11.0", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hAns := cmpNumeric(tt.first, tt.second)

			if !reflect.DeepEqual(hAns, tt.wAns) {
				t.Errorf("cmpStrings():\nwant: %v\nhave: %v\n", hAns, tt.wAns)
				return
			}
		})
	}
}

func Test_CmpLines(t *testing.T) {
	tests := []struct {
		name string
		i    int
		j    int
		f    *File
		wAns int
	}{
		{"first < second", 0, 1, &File{FileData: [][]string{{"a", "b"}, {"a", "c"}}, Cfg: &config.Config{SortColumn: 1}}, -1},
		{"first > second", 0, 1, &File{FileData: [][]string{{"a", "c", "d"}, {"a", "c"}}, Cfg: &config.Config{SortColumn: 1}}, 1},
		{"first = second", 0, 1, &File{FileData: [][]string{{"a", "b"}, {"a", "b"}}, Cfg: &config.Config{SortColumn: 1}}, 0},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			hAns := tt.f.CmpLines(tt.i, tt.j)

			if !reflect.DeepEqual(hAns, tt.wAns) {
				t.Errorf("cmpStrings():\nwant: %v\nhave: %v\n", hAns, tt.wAns)
				return
			}
		})
	}
}
