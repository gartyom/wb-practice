package file

import (
	"errors"
	"math/big"
	"os"
	"sort"
	"strings"

	"github.com/gartyom/wb-practice/L2/develop/dev03/internal/pkg/config"
)

func cmpStrings(first string, second string) int {
	if first < second {
		return -1
	}

	if first > second {
		return 1
	}
	return 0
}

func cmp(first string, second string, numeric bool) int {
	if numeric {
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

	return cmpStrings(first, second)
}

type File struct {
	FileData [][]string
	Cfg      *config.Config
}

func (f *File) Less(i, j int) (answer bool) {
	first := f.FileData[i]
	second := f.FileData[j]

	for k := f.Cfg.SortColumn - 1; k < len(first); k++ {
		if k < len(second) {
			cmp := cmp(first[k], second[k], f.Cfg.Numeric)
			switch cmp {
			case -1:
				return true
			case 1:
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func (f *File) Len() int {
	return len(f.FileData)
}

func (f *File) Swap(i, j int) {
	f.FileData[i], f.FileData[j] = f.FileData[j], f.FileData[i]
}

func NewFile(cfg *config.Config) (*File, error) {
	err := validate(cfg.FilePath)
	if err != nil {
		return nil, err
	}

	fData, err := read(cfg)
	if err != nil {
		return nil, err
	}

	return &File{FileData: fData, Cfg: cfg}, nil
}

func validate(fp string) error {
	_, err := os.Stat(fp)
	return err
}

func read(cfg *config.Config) ([][]string, error) {

	fileBytes, err := os.ReadFile(cfg.FilePath)
	if err != nil {
		return nil, err // should never happen
	}

	// Converting bytes to string
	fileString := string(fileBytes)
	// 	Splitting each string into two parts
	var fData [][]string
	for _, line := range strings.Split(fileString, "\n") {
		str := strings.Split(line, " ")
		fData = append(fData, str)
	}

	return fData, nil
}

func Sort(f *File) (*File, error) {
	sort.Sort(f)
	return f, nil
}

func splitString(str string, sep rune, k int) ([]string, error) {
	if k <= 0 {
		return nil, errors.New("k should be >= 1")
	}

	if k == 1 {
		return []string{"", str}, nil
	}

	var first strings.Builder
	var second strings.Builder
	k--
	for i, r := range str {
		if r == rune(sep) {
			k--
		}
		if k == 0 {
			first.WriteString(str[0:i])
			second.WriteString(str[i+1:])
			return []string{first.String(), second.String()}, nil
		}

	}

	return []string{str, ""}, nil
}
