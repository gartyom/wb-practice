package file

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strings"

	"dev03/internal/pkg/config"
)

type File struct {
	FileData [][]string
	Cfg      *config.Config
}

func NewFile(cfg *config.Config) (*File, error) {
	err := validate(cfg.FilePath)
	if err != nil {
		return nil, err
	}

	fData, err := read(cfg.FilePath)
	if err != nil {
		return nil, err
	}

	return &File{FileData: *fData, Cfg: cfg}, nil
}

func (f *File) Less(i, j int) (answer bool) {
	cmp := f.CmpLines(i, j)

	if cmp == -1 {
		return false
	}

	if cmp == 1 {
		return true
	}

	return true // doesn't matter what to return since lines are equal
}

func (f *File) Len() int {
	return len(f.FileData)
}

func (f *File) Swap(i, j int) {
	f.FileData[i], f.FileData[j] = f.FileData[j], f.FileData[i]
}

func (f *File) CmpLines(i, j int) int {
	first := f.FileData[i]
	second := f.FileData[j]

	for k := f.Cfg.SortColumn - 1; k < len(first); k++ {
		if k < len(second) {
			result := cmpWords(first[k], second[k], f.Cfg.Numeric)
			if result != 0 {
				return result
			}
		} else {
			return 1
		}
	}

	return 0
}

func (f *File) Print() {
	i, l := 0, len(f.FileData)

	increment := func(i int) int {
		i += 1
		return i
	}
	decrement := func(i int) int {
		i -= 1
		return i
	}
	next := increment

	if f.Cfg.Reverse {
		i, l = l-1, i-1
		next = decrement

	}

	for i != l {
		str := strings.Join(f.FileData[i], " ")

		j := next(i)

		if f.Cfg.Unique {
			if j != l {
				if str == strings.Join(f.FileData[j], " ") {
					i = j // skip printing current line if next line the same
					continue
				}
			}
		}

		fmt.Fprintln(os.Stdout, str)
		i = j
	}
}

func Sort(f *File) (*File, error) {
	sort.Sort(f)
	return f, nil
}

func validate(fp string) error {
	_, err := os.Stat(fp)
	return err
}

func read(fp string) (*[][]string, error) {

	fileBytes, err := os.ReadFile(fp)
	if err != nil {
		return nil, err // should never happen
	}

	fileString := string(fileBytes)
	var fData [][]string
	strs := strings.Split(fileString, "\n")
	for _, line := range strs {
		if len(line) == 0 {
			continue
		}

		str := strings.Split(line, " ")
		fData = append(fData, str)
	}

	if len(fData) < 2 {
		fmt.Println(len(fData))
		return nil, errors.New("File must contain at least 2 lines")
	}

	return &fData, nil
}

func cmpWords(first string, second string, numeric bool) int {
	if numeric {
		return cmpNumeric(first, second)
	}

	return cmpStrings(first, second)
}

func cmpStrings(first string, second string) int {
	if first < second {
		return -1
	}

	if first > second {
		return 1
	}
	return 0
}

func cmpNumeric(first string, second string) int {
	fnum, _, err := big.ParseFloat(first, 10, 256, big.ToNearestEven)
	if err != nil {
		return cmpStrings(first, second)
	}
	snum, _, err := big.ParseFloat(second, 10, 256, big.ToNearestEven)
	if err != nil {
		return cmpStrings(first, second)
	}

	cmp := fnum.Cmp(snum)

	return cmp
}
