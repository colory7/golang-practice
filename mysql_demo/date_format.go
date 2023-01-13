package mysql_demo

import (
	"errors"
	"fmt"
	"golang_practice/oracle_demo/builtins"
	"strconv"
	"strings"
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
					_, w := week(t, 'V')
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
					y, _ := week(t, 'X')
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
