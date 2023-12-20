package date

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Validate(date string) bool {
	// date has yyyy-mm-dd format
	d := strings.Split(date, "-")
	if len(d) < 3 {
		return false
	}

	year, err1 := strconv.Atoi(d[0])
	month, err2 := strconv.Atoi(d[1])
	day, err3 := strconv.Atoi(d[2])

	if err1 != nil || err2 != nil || err3 != nil {
		return false
	}

	if month < 1 || month > 12 ||
		day < 1 || day > 31 ||
		year < 1 {
		return false
	}

	if month%2 == 0 {
		if month == 2 {
			if year%4 == 0 {
				if day > 29 {
					return false
				}
			} else {
				if day > 28 {
					return false
				}
			}
		}
		if day > 30 {
			return false
		}
	}
	return true
}

func GetMonthPeriod(date string) (string, string, error) {
	// date has yyyy-mm-dd format
	d := strings.Split(date, "-")
	if len(d) < 3 {
		return "", "", errors.New("not valid date")
	}

	year, _ := strconv.Atoi(d[0])
	month, _ := strconv.Atoi(d[1])
	var m string
	if month < 10 {
		m = fmt.Sprintf("0%d", month)
	} else {
		m = fmt.Sprint(month)
	}

	return fmt.Sprintf("%d-%s-01", year, m), fmt.Sprintf("%d-%s-31", year, m), nil
}
