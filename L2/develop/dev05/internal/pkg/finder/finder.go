package finder

import (
	"dev05/internal/pkg/file"
	"dev05/internal/pkg/matcher"
)

type Finder struct {
	matcher matcher.Matcher
	Before  *uint
	After   *uint
}

func New(after *uint, before *uint, ignoreCase *bool, fixed *bool) *Finder {
	finder := &Finder{
		After:   after,
		Before:  before,
		matcher: *matcher.New(ignoreCase, fixed),
	}

	return finder
}

func (finder *Finder) Find(file *file.File, pattern string) (*[]int, error) {
	matched := []int{}
	_, err := finder.matcher.Match(pattern, "A")
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(file.Data); i++ {
		eq, _ := finder.matcher.Match(pattern, file.Data[i])

		if eq {
			left, right := finder.getIdxsWindow(file, i, matched)

			for j := left; j < right; j++ {
				matched = append(matched, j)
			}
		}

	}

	return &matched, nil
}

func (finder *Finder) getIdxsWindow(file *file.File, i int, matched []int) (int, int) {
	left := max(i-int(*finder.Before), 0)

	right := min(i+int(*finder.After)+1, len(file.Data))

	lMatched := len(matched)
	if lMatched > 0 {
		left = max(matched[lMatched-1]+1, left)
	}

	return left, right
}
