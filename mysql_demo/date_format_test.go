package mysql_demo

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang_practice/oracle_demo/builtins"
	"strconv"
	"strings"
	"testing"
	"time"
)

var longDayNames = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

var shortDayNames = []string{
	"Sun",
	"Mon",
	"Tue",
	"Wed",
	"Thu",
	"Fri",
	"Sat",
}

var shortMonthNames = []string{
	"---",
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

var longMonthNames = []string{
	"---",
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

func DateFormat(t time.Time, f string) (string, error) {
	var result []string
	format := []rune(f)

	add := func(str string) {
		result = append(result, str)
	}
	for i := 0; i < len(format); i++ {
		switch format[i] {
		case '%':
			if i < len(format)-1 {
				switch format[i+1] {
				case 'a':
					add(shortDayNames[t.Weekday()])
				case 'b':
					add(shortMonthNames[t.Month()])
				case 'c':
					add(strconv.FormatUint(uint64(t.Month()), 10))
				case 'C':
					add(fmt.Sprintf("%02d", t.Year()/100))
				case 'd':
					add(fmt.Sprintf("%02d", t.Day()))
				case 'D':
					add(builtins.NumToWithOrdinalSuffix(t.Day()))
				case 'e':
					add(strconv.FormatUint(uint64(t.Day()), 10))
				case 'f':
					add(fmt.Sprintf("%06d", t.Nanosecond()/1000))
				case 'H':
					add(fmt.Sprintf("%02d", t.Hour()))
				case 'k':
					add(strconv.FormatUint(uint64(t.Hour()), 10))
				case 'h', 'I':
					h := t.Hour()
					if h == 0 {
						add("12")
					} else if h <= 12 {
						add(fmt.Sprintf("%02d", h))
					} else {
						add(fmt.Sprintf("%02d", h-12))
					}
				case 'l':
					h := t.Hour()
					if h == 0 {
						add("12")
					} else if h <= 12 {
						add(strconv.FormatUint(uint64(h), 10))
					} else {
						add(strconv.FormatUint(uint64(h-12), 10))
					}
				case 'i':
					add(fmt.Sprintf("%02d", t.Minute()))
				case 'j':
					add(fmt.Sprintf("%03d", t.YearDay()))
				case 'm':
					add(fmt.Sprintf("%02d", t.Month()))
				case 'M':
					add(longMonthNames[t.Month()])
				case 'p':
					if t.Hour() < 12 {
						add("AM")
					} else {
						add("PM")
					}
				case 'r':
					s, err := DateFormat(t, "%I:%H:%S %p")
					if err != nil {
						return "", err
					}
					add(s)
				case 's', 'S':
					add(fmt.Sprintf("%02d", t.Second()))
				case 'T':
					add(fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second()))
				case 'U':
					_, w := week(t, 'U')
					add(w)
				case 'u':
					_, w := week(t, 'u')
					add(w)
				case 'V':
					_, w := weekV(t, true)
					add(w)
				case 'v':
					_, isoWeek := t.ISOWeek()
					if isoWeek < 1e1 {
						add("0")
					}
					add(strconv.FormatUint(uint64(isoWeek), 10))
				case 'W':
					add(longDayNames[t.Weekday()])
				case 'w':
					add(strconv.FormatUint(uint64(t.Weekday()), 10))
				case 'X':
					y, _ := weekV(t, true)
					add(y)
				case 'x':
					isoYear, _ := t.ISOWeek()
					if isoYear < 1e1 {
						add("0")
					}
					add(strconv.FormatUint(uint64(isoYear), 10))
				case 'y':
					add(fmt.Sprintf("%02d", t.Year()%100))
				case 'Y':
					add(fmt.Sprintf("%04d", t.Year()))
				case '%':
					add("%")
				default:
					return "", fmt.Errorf("invalid format code: %c", format[i+1])
				}
				i += 1
			}
		default:
			add(string(format[i]))
		}
	}
	return strings.Join(result, ""), nil
}

func TestDateFormat(t *testing.T) {
	tests := []struct {
		i         int
		ch        string
		format    string
		expected  string
		exception bool
	}{
		{1, "2019-12-28 00:00:00", "%u %v %x - %U %V %X", "52 52 2019 - 51 51 2019", false},
		{1, "2019-12-29 00:00:00", "%u %v %x - %U %V %X", "52 52 2019 - 52 52 2019", false},
		{1, "2019-12-30 00:00:00", "%u %v %x - %U %V %X", "53 01 2020 - 52 52 2019", false},
		{1, "2019-12-31 00:00:00", "%u %v %x - %U %V %X", "53 01 2020 - 52 52 2019", false},
		{1, "2020-01-01 00:00:00", "%u %v %x - %U %V %X", "01 01 2020 - 00 52 2019", false},
		{1, "2020-01-02 00:00:00", "%u %v %x - %U %V %X", "01 01 2020 - 00 52 2019", false},
		{1, "2020-01-03 00:00:00", "%u %v %x - %U %V %X", "01 01 2020 - 00 52 2019", false},
		{1, "2020-01-04 00:00:00", "%u %v %x - %U %V %X", "01 01 2020 - 00 52 2019", false},
		{1, "2020-01-05 00:00:00", "%u %v %x - %U %V %X", "01 01 2020 - 01 01 2020", false},
		{1, "2020-01-06 00:00:00", "%u %v %x - %U %V %X", "02 02 2020 - 01 01 2020", false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			tm, err := time.Parse("2006-01-02 15:04:05", test.ch)
			if err != nil {
				panic(err)
			}
			actual, err := DateFormat(tm, test.format)
			if test.exception {
				assert.Error(t, err)
				//fmt.Println(err)
			} else {
				if err != nil {
					assert.NoError(t, err)
				}
				fmt.Println("actual: " + actual)
				if actual != test.expected {
					t.Fail()
				}
			}
		})
	}
}

func TestDateFormat2(t *testing.T) {
	tests := []struct {
		i         int
		ch        string
		format    string
		expected  string
		exception bool
	}{
		{1, "1998-12-27 00:00:00", "%u %v %x - %U %V %X", "52 52 1998 - 52 52 1998", false},
		{1, "1998-12-28 00:00:00", "%u %v %x - %U %V %X", "53 53 1998 - 52 52 1998", false},
		{1, "1998-12-29 00:00:00", "%u %v %x - %U %V %X", "53 53 1998 - 52 52 1998", false},
		{1, "1998-12-30 00:00:00", "%u %v %x - %U %V %X", "53 53 1998 - 52 52 1998", false},
		{1, "1998-12-31 00:00:00", "%u %v %x - %U %V %X", "53 53 1998 - 52 52 1998", false},
		{1, "1999-01-01 00:00:00", "%u %v %x - %U %V %X", "00 53 1998 - 00 52 1998", false},
		{1, "1999-01-02 00:00:00", "%u %v %x - %U %V %X", "00 53 1998 - 00 52 1998", false},
		{1, "1999-01-03 00:00:00", "%u %v %x - %U %V %X", "00 53 1998 - 01 01 1999", false},
		{1, "1999-01-04 00:00:00", "%u %v %x - %U %V %X", "01 01 1999 - 01 01 1999", false},
		{1, "1999-01-05 00:00:00", "%u %v %x - %U %V %X", "01 01 1999 - 01 01 1999", false},
		{1, "1999-01-06 00:00:00", "%u %v %x - %U %V %X", "01 01 1999 - 01 01 1999", false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			tm, err := time.Parse("2006-01-02 15:04:05", test.ch)
			if err != nil {
				panic(err)
			}
			actual, err := DateFormat(tm, test.format)
			if test.exception {
				assert.Error(t, err)
				//fmt.Println(err)
			} else {
				if err != nil {
					assert.NoError(t, err)
				}
				fmt.Println("actual: " + actual)
				if actual != test.expected {
					fmt.Println("expected: " + test.expected)
					t.Fail()
				}
			}
		})
	}
}

func TestDateFormat3(t *testing.T) {
	tests := []struct {
		i         int
		ch        string
		format    string
		expected  string
		exception bool
	}{
		{1, "2000-12-22 00:00:00", "%u %v %x - %U %V %X", "51 51 2000 - 51 51 2000", false},
		{1, "2000-12-23 00:00:00", "%u %v %x - %U %V %X", "51 51 2000 - 51 51 2000", false},
		{1, "2000-12-24 00:00:00", "%u %v %x - %U %V %X", "51 51 2000 - 52 52 2000", false},
		{1, "2000-12-25 00:00:00", "%u %v %x - %U %V %X", "52 52 2000 - 52 52 2000", false},
		{1, "2000-12-26 00:00:00", "%u %v %x - %U %V %X", "52 52 2000 - 52 52 2000", false},
		{1, "2000-12-27 00:00:00", "%u %v %x - %U %V %X", "52 52 2000 - 52 52 2000", false},
		{1, "2000-12-28 00:00:00", "%u %v %x - %U %V %X", "52 52 2000 - 52 52 2000", false},
		{1, "2000-12-29 00:00:00", "%u %v %x - %U %V %X", "52 52 2000 - 52 52 2000", false},
		{1, "2000-12-30 00:00:00", "%u %v %x - %U %V %X", "52 52 2000 - 52 52 2000", false},
		{1, "2000-12-31 00:00:00", "%u %v %x - %U %V %X", "52 52 2000 - 53 53 2000", false},
		{1, "2001-01-01 00:00:00", "%u %v %x - %U %V %X", "01 01 2001 - 00 53 2000", false},
		{1, "2001-01-02 00:00:00", "%u %v %x - %U %V %X", "01 01 2001 - 00 53 2000", false},
		{1, "2001-01-03 00:00:00", "%u %v %x - %U %V %X", "01 01 2001 - 00 53 2000", false},
		{1, "2001-01-04 00:00:00", "%u %v %x - %U %V %X", "01 01 2001 - 00 53 2000", false},
		{1, "2001-01-05 00:00:00", "%u %v %x - %U %V %X", "01 01 2001 - 00 53 2000", false},
		{1, "2001-01-06 00:00:00", "%u %v %x - %U %V %X", "01 01 2001 - 00 53 2000", false},
		{1, "2001-01-07 00:00:00", "%u %v %x - %U %V %X", "01 01 2001 - 01 01 2001", false},
		{1, "2001-01-08 00:00:00", "%u %v %x - %U %V %X", "02 02 2001 - 01 01 2001", false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			tm, err := time.Parse("2006-01-02 15:04:05", test.ch)
			if err != nil {
				panic(err)
			}
			actual, err := DateFormat(tm, test.format)
			if test.exception {
				assert.Error(t, err)
				//fmt.Println(err)
			} else {
				if err != nil {
					assert.NoError(t, err)
				}
				fmt.Println("actual:   " + actual)
				if actual != test.expected {
					fmt.Println("expected: " + test.expected)
					t.Fail()
				}
			}
		})
	}
}

func TestDateFormat4(t *testing.T) {
	tests := []struct {
		i         int
		ch        string
		format    string
		expected  string
		exception bool
	}{
		{1, "2001-12-27 00:00:00", "%u %v %x - %U %V %X %W %w", "52 52 2001 - 51 51 2001 Thursday 4", false},
		{1, "2001-12-28 00:00:00", "%u %v %x - %U %V %X %W %w", "52 52 2001 - 51 51 2001 Friday 5", false},
		{1, "2001-12-29 00:00:00", "%u %v %x - %U %V %X %W %w", "52 52 2001 - 51 51 2001 Saturday 6", false},
		{1, "2001-12-30 00:00:00", "%u %v %x - %U %V %X %W %w", "52 52 2001 - 52 52 2001 Sunday 0", false},
		{1, "2001-12-31 00:00:00", "%u %v %x - %U %V %X %W %w", "53 01 2002 - 52 52 2001 Monday 1", false},
		{1, "2002-01-01 00:00:00", "%u %v %x - %U %V %X %W %w", "01 01 2002 - 00 52 2001 Tuesday 2", false},
		{1, "2002-01-02 00:00:00", "%u %v %x - %U %V %X %W %w", "01 01 2002 - 00 52 2001 Wednesday 3", false},
		{1, "2002-01-03 00:00:00", "%u %v %x - %U %V %X %W %w", "01 01 2002 - 00 52 2001 Thursday 4", false},
		{1, "2002-01-04 00:00:00", "%u %v %x - %U %V %X %W %w", "01 01 2002 - 00 52 2001 Friday 5", false},
		{1, "2002-01-05 00:00:00", "%u %v %x - %U %V %X %W %w", "01 01 2002 - 00 52 2001 Saturday 6", false},
		{1, "2002-01-06 00:00:00", "%u %v %x - %U %V %X %W %w", "01 01 2002 - 01 01 2002 Sunday 0", false},
		{1, "2002-01-07 00:00:00", "%u %v %x - %U %V %X %W %w", "02 02 2002 - 01 01 2002 Monday 1", false},
		{1, "2002-01-08 00:00:00", "%u %v %x - %U %V %X %W %w", "02 02 2002 - 01 01 2002 Tuesday 2", false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			tm, err := time.Parse("2006-01-02 15:04:05", test.ch)
			if err != nil {
				panic(err)
			}
			actual, err := DateFormat(tm, test.format)
			if test.exception {
				assert.Error(t, err)
				//fmt.Println(err)
			} else {
				if err != nil {
					assert.NoError(t, err)
				}

				if actual != test.expected {
					fmt.Println("actual  : " + actual)
					fmt.Println("expected: " + test.expected)
					t.Fail()
				}
			}
		})
	}
}

func TestDateFormat5(t *testing.T) {
	tests := []struct {
		i         int
		ch        string
		format    string
		expected  string
		exception bool
	}{
		{1, "2001-12-27 00:00:00", "%a - %b - %c, - %D, - %d, - %e, - %f, - %H, - %h, - %I, - %i, - %j, - %k, - %l, - %M, - %m, - %p, - %r, - %S, - %s, - %T, - %U, - %u, - %V, - %v, - %W, - %w, - %X, - %x, - %Y, - %y - %%", "Thu - Dec - 12, - 27th, - 27, - 27, - 000000, - 00, - 12, - 12, - 00, - 361, - 0, - 12, - December, - 12, - AM, - 12:00:00 AM, - 00, - 00, - 00:00:00, - 51, - 52, - 51, - 52, - Thursday, - 4, - 2001, - 2001, - 2001, - 01 - %", false},
		{1, "2001-12-28 00:00:00", "%a - %b - %c, - %D, - %d, - %e, - %f, - %H, - %h, - %I, - %i, - %j, - %k, - %l, - %M, - %m, - %p, - %r, - %S, - %s, - %T, - %U, - %u, - %V, - %v, - %W, - %w, - %X, - %x, - %Y, - %y - %%", "Fri - Dec - 12, - 28th, - 28, - 28, - 000000, - 00, - 12, - 12, - 00, - 362, - 0, - 12, - December, - 12, - AM, - 12:00:00 AM, - 00, - 00, - 00:00:00, - 51, - 52, - 51, - 52, - Friday, - 5, - 2001, - 2001, - 2001, - 01 - %", false},
		{1, "2001-12-29 00:00:00", "%a - %b - %c, - %D, - %d, - %e, - %f, - %H, - %h, - %I, - %i, - %j, - %k, - %l, - %M, - %m, - %p, - %r, - %S, - %s, - %T, - %U, - %u, - %V, - %v, - %W, - %w, - %X, - %x, - %Y, - %y - %%", "Sat - Dec - 12, - 29th, - 29, - 29, - 000000, - 00, - 12, - 12, - 00, - 363, - 0, - 12, - December, - 12, - AM, - 12:00:00 AM, - 00, - 00, - 00:00:00, - 51, - 52, - 51, - 52, - Saturday, - 6, - 2001, - 2001, - 2001, - 01 - %", false},
		{1, "2001-12-30 00:00:00", "%a - %b - %c, - %D, - %d, - %e, - %f, - %H, - %h, - %I, - %i, - %j, - %k, - %l, - %M, - %m, - %p, - %r, - %S, - %s, - %T, - %U, - %u, - %V, - %v, - %W, - %w, - %X, - %x, - %Y, - %y - %%", "Sun - Dec - 12, - 30th, - 30, - 30, - 000000, - 00, - 12, - 12, - 00, - 364, - 0, - 12, - December, - 12, - AM, - 12:00:00 AM, - 00, - 00, - 00:00:00, - 52, - 52, - 52, - 52, - Sunday, - 0, - 2001, - 2001, - 2001, - 01 - %", false},
		{1, "2001-12-31 00:00:00", "%a - %b - %c, - %D, - %d, - %e, - %f, - %H, - %h, - %I, - %i, - %j, - %k, - %l, - %M, - %m, - %p, - %r, - %S, - %s, - %T, - %U, - %u, - %V, - %v, - %W, - %w, - %X, - %x, - %Y, - %y - %%", "Mon - Dec - 12, - 31st, - 31, - 31, - 000000, - 00, - 12, - 12, - 00, - 365, - 0, - 12, - December, - 12, - AM, - 12:00:00 AM, - 00, - 00, - 00:00:00, - 52, - 53, - 52, - 01, - Monday, - 1, - 2001, - 2002, - 2001, - 01 - %", false},
		{1, "2002-01-01 00:00:00", "%a - %b - %c, - %D, - %d, - %e, - %f, - %H, - %h, - %I, - %i, - %j, - %k, - %l, - %M, - %m, - %p, - %r, - %S, - %s, - %T, - %U, - %u, - %V, - %v, - %W, - %w, - %X, - %x, - %Y, - %y - %%", "Tue - Jan - 1, - 1st, - 01, - 1, - 000000, - 00, - 12, - 12, - 00, - 001, - 0, - 12, - January, - 01, - AM, - 12:00:00 AM, - 00, - 00, - 00:00:00, - 00, - 01, - 52, - 01, - Tuesday, - 2, - 2001, - 2002, - 2002, - 02 - %", false},
		{1, "2002-01-02 00:00:00", "%a - %b - %c, - %D, - %d, - %e, - %f, - %H, - %h, - %I, - %i, - %j, - %k, - %l, - %M, - %m, - %p, - %r, - %S, - %s, - %T, - %U, - %u, - %V, - %v, - %W, - %w, - %X, - %x, - %Y, - %y - %%", "Wed - Jan - 1, - 2nd, - 02, - 2, - 000000, - 00, - 12, - 12, - 00, - 002, - 0, - 12, - January, - 01, - AM, - 12:00:00 AM, - 00, - 00, - 00:00:00, - 00, - 01, - 52, - 01, - Wednesday, - 3, - 2001, - 2002, - 2002, - 02 - %", false},
		{1, "2002-01-03 00:00:00", "%a - %b - %c, - %D, - %d, - %e, - %f, - %H, - %h, - %I, - %i, - %j, - %k, - %l, - %M, - %m, - %p, - %r, - %S, - %s, - %T, - %U, - %u, - %V, - %v, - %W, - %w, - %X, - %x, - %Y, - %y - %%", "Thu - Jan - 1, - 3rd, - 03, - 3, - 000000, - 00, - 12, - 12, - 00, - 003, - 0, - 12, - January, - 01, - AM, - 12:00:00 AM, - 00, - 00, - 00:00:00, - 00, - 01, - 52, - 01, - Thursday, - 4, - 2001, - 2002, - 2002, - 02 - %", false},
		{1, "2002-01-04 00:00:00", "%a - %b - %c, - %D, - %d, - %e, - %f, - %H, - %h, - %I, - %i, - %j, - %k, - %l, - %M, - %m, - %p, - %r, - %S, - %s, - %T, - %U, - %u, - %V, - %v, - %W, - %w, - %X, - %x, - %Y, - %y - %%", "Fri - Jan - 1, - 4th, - 04, - 4, - 000000, - 00, - 12, - 12, - 00, - 004, - 0, - 12, - January, - 01, - AM, - 12:00:00 AM, - 00, - 00, - 00:00:00, - 00, - 01, - 52, - 01, - Friday, - 5, - 2001, - 2002, - 2002, - 02 - %", false},
		{1, "2002-01-05 00:00:00", "%a - %b - %c, - %D, - %d, - %e, - %f, - %H, - %h, - %I, - %i, - %j, - %k, - %l, - %M, - %m, - %p, - %r, - %S, - %s, - %T, - %U, - %u, - %V, - %v, - %W, - %w, - %X, - %x, - %Y, - %y - %%", "Sat - Jan - 1, - 5th, - 05, - 5, - 000000, - 00, - 12, - 12, - 00, - 005, - 0, - 12, - January, - 01, - AM, - 12:00:00 AM, - 00, - 00, - 00:00:00, - 00, - 01, - 52, - 01, - Saturday, - 6, - 2001, - 2002, - 2002, - 02 - %", false},
		{1, "2002-01-06 00:00:00", "%a - %b - %c, - %D, - %d, - %e, - %f, - %H, - %h, - %I, - %i, - %j, - %k, - %l, - %M, - %m, - %p, - %r, - %S, - %s, - %T, - %U, - %u, - %V, - %v, - %W, - %w, - %X, - %x, - %Y, - %y - %%", "Sun - Jan - 1, - 6th, - 06, - 6, - 000000, - 00, - 12, - 12, - 00, - 006, - 0, - 12, - January, - 01, - AM, - 12:00:00 AM, - 00, - 00, - 00:00:00, - 01, - 01, - 01, - 01, - Sunday, - 0, - 2002, - 2002, - 2002, - 02 - %", false},
		{1, "2002-01-07 00:00:00", "%a - %b - %c, - %D, - %d, - %e, - %f, - %H, - %h, - %I, - %i, - %j, - %k, - %l, - %M, - %m, - %p, - %r, - %S, - %s, - %T, - %U, - %u, - %V, - %v, - %W, - %w, - %X, - %x, - %Y, - %y - %%", "Mon - Jan - 1, - 7th, - 07, - 7, - 000000, - 00, - 12, - 12, - 00, - 007, - 0, - 12, - January, - 01, - AM, - 12:00:00 AM, - 00, - 00, - 00:00:00, - 01, - 02, - 01, - 02, - Monday, - 1, - 2002, - 2002, - 2002, - 02 - %", false},
		{1, "2002-01-08 00:00:00", "%a - %b - %c, - %D, - %d, - %e, - %f, - %H, - %h, - %I, - %i, - %j, - %k, - %l, - %M, - %m, - %p, - %r, - %S, - %s, - %T, - %U, - %u, - %V, - %v, - %W, - %w, - %X, - %x, - %Y, - %y - %%", "Tue - Jan - 1, - 8th, - 08, - 8, - 000000, - 00, - 12, - 12, - 00, - 008, - 0, - 12, - January, - 01, - AM, - 12:00:00 AM, - 00, - 00, - 00:00:00, - 01, - 02, - 01, - 02, - Tuesday, - 2, - 2002, - 2002, - 2002, - 02 - %", false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			tm, err := time.Parse("2006-01-02 15:04:05", test.ch)
			if err != nil {
				panic(err)
			}
			actual, err := DateFormat(tm, test.format)
			if test.exception {
				assert.Error(t, err)
				//fmt.Println(err)
			} else {
				if err != nil {
					assert.NoError(t, err)
				}

				if actual != test.expected {
					fmt.Println("actual  : " + actual)
					fmt.Println("expected: " + test.expected)
					t.Fail()
				}
			}
		})
	}
}

func TestSuite(t *testing.T) {
	tests := []struct {
		i         int
		ch        string
		exception bool
	}{
		// 周六 52 52 2019 - 51 51 2019
		{1, "2019-12-28T00:00:00+00:00", false},
		// 周日 52 52 2019 - 52 52 2019
		{2, "2019-12-29T00:00:00+00:00", false},
		// 周一 53 01 2020 - 52 52 2019
		{3, "2019-12-30T00:00:00+00:00", false},
		// 周二 53 01 2020 - 52 52 2019
		{4, "2019-12-31T00:00:00+00:00", false},

		// 周三 01 01 2020 - 00 52 2019
		{5, "2020-01-01T00:00:00+00:00", false},
		// 周四 01 01 2020 - 00 52 2019
		{6, "2020-01-02T00:00:00+00:00", false},
		// 周五 01 01 2020 - 00 52 2019
		{7, "2020-01-03T00:00:00+00:00", false},
		// 周六 01 01 2020 - 00 52 2019
		{8, "2020-01-04T00:00:00+00:00", false},
		// 周日 01 01 2020 - 01 01 2020
		{9, "2020-01-05T00:00:00+00:00", false},
		// 周一 02 02 2020 - 01 01 2020
		{10, "2020-01-06T00:00:00+00:00", false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			tm, err := time.Parse(time.RFC3339, test.ch)
			if err != nil {
				panic(err)
			}

			w, err := weekU(tm, 'U')
			if test.exception {
				assert.Error(t, err)
				//fmt.Println(err)
			} else {
				if err != nil {
					assert.NoError(t, err)
				}
				fmt.Println(w)
			}
		})
	}
}

func TestSuite3(t *testing.T) {
	tests := []struct {
		i         int
		ch        string
		exception bool
	}{
		// 周六 52 52 2019 - 51 51 2019
		{1, "2019-12-28T00:00:00+00:00", false},
		// 周日 52 52 2019 - 52 52 2019
		{2, "2019-12-29T00:00:00+00:00", false},
		// 周一 53 01 2020 - 52 52 2019
		{3, "2019-12-30T00:00:00+00:00", false},
		// 周二 53 01 2020 - 52 52 2019
		{4, "2019-12-31T00:00:00+00:00", false},

		// 周三 01 01 2020 - 00 52 2019
		{5, "2020-01-01T00:00:00+00:00", false},
		// 周四 01 01 2020 - 00 52 2019
		{6, "2020-01-02T00:00:00+00:00", false},
		// 周五 01 01 2020 - 00 52 2019
		{7, "2020-01-03T00:00:00+00:00", false},
		// 周六 01 01 2020 - 00 52 2019
		{8, "2020-01-04T00:00:00+00:00", false},
		// 周日 01 01 2020 - 01 01 2020
		{9, "2020-01-05T00:00:00+00:00", false},
		// 周一 02 02 2020 - 01 01 2020
		{10, "2020-01-06T00:00:00+00:00", false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			tm, err := time.Parse(time.RFC3339, test.ch)
			if err != nil {
				panic(err)
			}

			isoYear, isoWeek := tm.ISOWeek()
			if test.exception {
				assert.Error(t, err)
				//fmt.Println(err)
			} else {
				if err != nil {
					assert.NoError(t, err)
				}
				fmt.Println(isoYear)
				fmt.Println(isoWeek)
			}
		})
	}
}

//func weekU(t time.Time, sundayAsFirstDay bool) (string, error) {
//	yearDay := t.YearDay()
//
//	var weekDay = 0
//	// 整周的倍数
//	// 如果周日作为一周的第一天,则第一周为01
//	// 如果周一作为一周的第一天,则第一周为00
//	// 第一周的天数,有可能不足7天
//	var weeks = 0
//	if sundayAsFirstDay {
//		weekDay = int(t.Weekday())
//	} else {
//		weeks = 1
//		if t.Weekday() == time.Sunday {
//			weekDay = 6
//		} else {
//			weekDay--
//		}
//	}
//
//	// 除去当前周的天数
//	daysOfExceptCurrWeek := yearDay - weekDay
//
//	if daysOfExceptCurrWeek > 0 {
//		weeks = daysOfExceptCurrWeek/7 + 1
//	} else if daysOfExceptCurrWeek == 0 {
//		weeks = daysOfExceptCurrWeek / 7
//	}
//
//	ret := strings.Builder{}
//	if weeks < 10 {
//		ret.WriteByte('0')
//	}
//	ret.WriteString(strconv.FormatUint(uint64(weeks), 10))
//	return ret.String(), nil
//}

func TestSuiteWeekV(t *testing.T) {
	tests := []struct {
		i         int
		ch        string
		exception bool
	}{
		{1, "2019-12-28T00:00:00+00:00", false},
		{2, "2019-12-29T00:00:00+00:00", false},
		{3, "2019-12-30T00:00:00+00:00", false},
		{4, "2019-12-31T00:00:00+00:00", false},
		{5, "2020-01-01T00:00:00+00:00", false},
		{6, "2020-01-02T00:00:00+00:00", false},
		{7, "2020-01-03T00:00:00+00:00", false},
		{8, "2020-01-04T00:00:00+00:00", false},
		{9, "2020-01-05T00:00:00+00:00", false},
		{10, "2020-01-06T00:00:00+00:00", false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			tm, err := time.Parse(time.RFC3339, test.ch)
			if err != nil {
				panic(err)
			}

			y, w := weekV(tm, true)
			if test.exception {
				assert.Error(t, err)
				//fmt.Println(err)
			} else {
				if err != nil {
					assert.NoError(t, err)
				}
				fmt.Printf("y: %s, w: %s\n", y, w)
			}
		})
	}
}

func TestSuiteWeekV2(t *testing.T) {
	tests := []struct {
		i         int
		ch        string
		exception bool
	}{
		{1, "1998-12-27T00:00:00+00:00", false},
		{1, "1998-12-28T00:00:00+00:00", false},
		{2, "1998-12-29T00:00:00+00:00", false},
		{3, "1998-12-30T00:00:00+00:00", false},
		{4, "1998-12-31T00:00:00+00:00", false},
		{5, "1999-01-01T00:00:00+00:00", false},
		{6, "1999-01-02T00:00:00+00:00", false},
		{7, "1999-01-03T00:00:00+00:00", false},
		{8, "1999-01-04T00:00:00+00:00", false},
		{9, "1999-01-05T00:00:00+00:00", false},
		{10, "1999-01-06T00:00:00+00:00", false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			tm, err := time.Parse(time.RFC3339, test.ch)
			if err != nil {
				panic(err)
			}

			y, w := weekV(tm, true)
			if test.exception {
				assert.Error(t, err)
				//fmt.Println(err)
			} else {
				if err != nil {
					assert.NoError(t, err)
				}
				fmt.Printf("y: %s, w: %s\n", y, w)
			}
		})
	}
}

func TestSuiteWeekV3(t *testing.T) {
	tests := []struct {
		i         int
		ch        string
		exception bool
	}{
		{1, "2000-12-22T00:00:00+00:00", false},
		{1, "2000-12-23T00:00:00+00:00", false},
		{1, "2000-12-24T00:00:00+00:00", false},
		{1, "2000-12-25T00:00:00+00:00", false},
		{1, "2000-12-26T00:00:00+00:00", false},
		{1, "2000-12-27T00:00:00+00:00", false},
		{2, "2000-12-28T00:00:00+00:00", false},
		{3, "2000-12-29T00:00:00+00:00", false},
		{4, "2000-12-30T00:00:00+00:00", false},
		{5, "2000-12-31T00:00:00+00:00", false},
		{6, "2001-01-01T00:00:00+00:00", false},
		{7, "2001-01-02T00:00:00+00:00", false},
		{8, "2001-01-03T00:00:00+00:00", false},
		{9, "2001-01-04T00:00:00+00:00", false},
		{10, "2001-01-05T00:00:00+00:00", false},
		{11, "2001-01-06T00:00:00+00:00", false},
		{12, "2001-01-07T00:00:00+00:00", false},
		{13, "2001-01-08T00:00:00+00:00", false},

		//{13, "2001-01-02T00:00:00+00:00", false},
		//{13, "2001-01-06T00:00:00+00:00", false},
		//{13, "2001-01-07T00:00:00+00:00", false},
		//{13, "2001-01-08T00:00:00+00:00", false},
		//{13, "2001-01-09T00:00:00+00:00", false},
		//{13, "2001-01-13T00:00:00+00:00", false},
		//{13, "2001-01-14T00:00:00+00:00", false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			tm, err := time.Parse(time.RFC3339, test.ch)
			if err != nil {
				panic(err)
			}

			y, w := weekV(tm, true)
			if test.exception {
				assert.Error(t, err)
				//fmt.Println(err)
			} else {
				if err != nil {
					assert.NoError(t, err)
				}
				fmt.Printf("y: %s, w: %s\n", y, w)
			}
		})
	}
}

func TestSuiteWeekV4(t *testing.T) {
	tests := []struct {
		i         int
		ch        string
		exception bool
	}{
		{2, "2019-01-06T00:00:00+00:00", false},
		{2, "2019-12-29T00:00:00+00:00", false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			tm, err := time.Parse(time.RFC3339, test.ch)
			if err != nil {
				panic(err)
			}

			y, w := weekV(tm, true)
			if test.exception {
				assert.Error(t, err)
				//fmt.Println(err)
			} else {
				if err != nil {
					assert.NoError(t, err)
				}
				fmt.Printf("y: %s, w: %s\n", y, w)
			}
		})
	}
}

func TestWeek(tt *testing.T) {
	t1, _ := time.Parse(time.RFC3339, "2022-12-29T22:08:41+00:00")
	fmt.Println(t1.ISOWeek())

	t2, _ := time.Parse(time.RFC3339, "2022-12-30T22:08:41+00:00")
	fmt.Println(t2.ISOWeek())

	t3, _ := time.Parse(time.RFC3339, "2022-12-31T22:08:41+00:00")
	fmt.Println(t3.ISOWeek())

	t4, _ := time.Parse(time.RFC3339, "2023-01-01T22:08:41+00:00")
	fmt.Println(t4.ISOWeek())

	t5, _ := time.Parse(time.RFC3339, "2023-01-02T22:08:41+00:00")
	fmt.Println(t5.ISOWeek())

	t6, _ := time.Parse(time.RFC3339, "2023-01-03T22:08:41+00:00")
	fmt.Println(t6.ISOWeek())

	fmt.Println("====")
	t11, _ := time.Parse(time.RFC3339, "2023-12-29T22:08:41+00:00")
	fmt.Println(t11.ISOWeek())

	t12, _ := time.Parse(time.RFC3339, "2023-12-30T22:08:41+00:00")
	fmt.Println(t12.ISOWeek())

	t13, _ := time.Parse(time.RFC3339, "2023-12-31T22:08:41+00:00")
	fmt.Println(t13.ISOWeek())

	t14, _ := time.Parse(time.RFC3339, "2024-01-01T22:08:41+00:00")
	fmt.Println(t14.ISOWeek())

	t15, _ := time.Parse(time.RFC3339, "2024-01-02T22:08:41+00:00")
	fmt.Println(t15.ISOWeek())

	t16, _ := time.Parse(time.RFC3339, "2024-01-03T22:08:41+00:00")
	fmt.Println(t16.ISOWeek())

}

func TestISOWeek(t *testing.T) {
	t1, err := time.Parse(time.RFC3339, "2019-12-29T22:08:41+00:00")
	if err != nil {
		panic(err)
	}
	isoYear, isoWeek := t1.ISOWeek()
	fmt.Println(isoYear)
	fmt.Println(isoWeek)
	fmt.Println(t1.Weekday())
	fmt.Println(int(t1.Weekday()))

	// 1-52
	d1 := (isoWeek*7 - 1) / 7
	fmt.Println(d1)

}

func TestWeek3(txx *testing.T) {
	sundayAsFirstDay := true
	t1, _ := time.Parse(time.RFC3339, "2019-12-28T22:08:41+00:00")
	weekV(t1, sundayAsFirstDay)
	t1, _ = time.Parse(time.RFC3339, "2019-12-29T22:08:41+00:00")
	weekV(t1, sundayAsFirstDay)
	t1, _ = time.Parse(time.RFC3339, "2019-12-30T22:08:41+00:00")
	weekV(t1, sundayAsFirstDay)
	t1, _ = time.Parse(time.RFC3339, "2019-12-31T22:08:41+00:00")
	weekV(t1, sundayAsFirstDay)
	t1, _ = time.Parse(time.RFC3339, "2020-01-01T22:08:41+00:00")
	weekV(t1, sundayAsFirstDay)
	t1, _ = time.Parse(time.RFC3339, "2020-01-02T22:08:41+00:00")
	weekV(t1, sundayAsFirstDay)
	t1, _ = time.Parse(time.RFC3339, "2020-01-03T22:08:41+00:00")
	weekV(t1, sundayAsFirstDay)
	t1, _ = time.Parse(time.RFC3339, "2020-01-04T22:08:41+00:00")
	weekV(t1, sundayAsFirstDay)
	t1, _ = time.Parse(time.RFC3339, "2020-01-05T22:08:41+00:00")
	weekV(t1, sundayAsFirstDay)
	t1, _ = time.Parse(time.RFC3339, "2020-01-06T22:08:41+00:00")
	weekV(t1, sundayAsFirstDay)

	fmt.Println("====")
	t1, _ = time.Parse(time.RFC3339, "1999-01-01T22:08:41+00:00")
	weekV(t1, sundayAsFirstDay)
	fmt.Println(t1.ISOWeek())

	fmt.Println("====")
	t1, _ = time.Parse(time.RFC3339, "2002-01-06T22:08:41+00:00")
	weekV(t1, sundayAsFirstDay)
	fmt.Println(t1.ISOWeek())

}

func weekV(t time.Time, sundayAsFirstDay bool) (string, string) {
	y := t.Year()
	w := 0
	yearDay := t.YearDay()

	dfw := getDaysOfFirstWeek(t, sundayAsFirstDay)

	diff := yearDay - dfw

	if diff <= 0 {
		return weekV(t.AddDate(0, 0, -yearDay), sundayAsFirstDay)
	} else if diff < 7 {
		w = 1
	} else if diff >= 7 {
		w = diff / 7
		if diff%7 != 0 {
			w++
		}
	}
	return fmt.Sprintf("%04d", y), fmt.Sprintf("%02d", w)
}

func weekU(t time.Time, mode byte) (string, error) {
	var sundayAsFirstDay bool
	switch mode {
	case 'U':
		sundayAsFirstDay = true
	case 'u':
		sundayAsFirstDay = false
	default:
		panic(errors.New("illegal character"))
	}

	w := 0
	yearDay := t.YearDay()

	dfw := getDaysOfFirstWeek(t, sundayAsFirstDay)

	diff := yearDay - dfw

	if diff <= 0 {
	} else if diff < 7 {
		w = 1
	} else if diff >= 7 {
		w = diff / 7
		if diff%7 != 0 {
			w++
		}
	}

	if mode == 'u' {
		w++
	}
	return fmt.Sprintf("%02d", w), nil
}

func week(t time.Time, mode byte) (string, string) {
	var sundayAsFirstDay bool
	switch mode {
	case 'U', 'V', 'X':
		sundayAsFirstDay = true
	case 'u', 'v', 'x':
		sundayAsFirstDay = false
	default:
		panic(errors.New("illegal character"))
	}

	y := t.Year()
	w := 0
	yearDay := t.YearDay()

	dfw := getDaysOfFirstWeek(t, sundayAsFirstDay)

	diff := yearDay - dfw

	if diff <= 0 {
		if mode == 'V' || mode == 'v' || mode == 'X' || mode == 'x' {
			return week(t.AddDate(0, 0, -yearDay), mode)
		}
	} else if diff < 7 {
		w = 1
	} else if diff >= 7 {
		w = diff / 7
		if diff%7 != 0 {
			w++
		}
	}
	if mode == 'u' && dfw >= 4 {
		w++
	}
	return fmt.Sprintf("%04d", y), fmt.Sprintf("%02d", w)
}

func getDaysOfFirstWeek(t time.Time, sundayAsFirstDay bool) int {
	yearDay := t.YearDay()

	var weekDay = int(t.Weekday())

	if !sundayAsFirstDay {
		if weekDay == 0 {
			weekDay = 6
		} else {
			weekDay--
		}
	}

	diff := yearDay - (weekDay + 1)
	var daysOfFirstWeek = 0
	if diff < 0 {
		daysOfFirstWeek = diff + 7
	} else {
		daysOfFirstWeek = diff % 7
	}

	return daysOfFirstWeek
}

func TestA(txx *testing.T) {
	t1, _ := time.Parse(time.RFC3339, "2020-01-25T22:08:41+00:00")
	t2 := t1.AddDate(0, 0, -t1.YearDay())
	fmt.Println(t2)

	fmt.Println(builtins.NumToOrdinalWord(t1.Day()))
	fmt.Println(builtins.NumToCardinalWord(t1.Day()))
	fmt.Println(builtins.NumToWithOrdinalSuffix(t1.Day()))
}
