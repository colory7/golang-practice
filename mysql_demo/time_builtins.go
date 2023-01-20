package mysql_demo

import (
	"bytes"
	"errors"
	"fmt"
	"golang_practice/oracle_demo/builtins"
	"math"
	"strconv"
	"strings"
	"time"
)

var mismatchTip = "mismatch error "
var mismatchError = errors.New(mismatchTip)
var lengthMismatchError = errors.New("format and parameter length mismatch")
var illegalCharacterError = errors.New("illegal character")

var longWeekDayNames = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

var shortWeekDayNames = []string{
	"Sun",
	"Mon",
	"Tue",
	"Wed",
	"Thu",
	"Fri",
	"Sat",
}

var shortWeekDayNamesInverse = map[string]time.Weekday{
	"Sun": time.Sunday,
	"Mon": time.Monday,
	"Tue": time.Tuesday,
	"Wed": time.Wednesday,
	"Thu": time.Thursday,
	"Fri": time.Friday,
	"Sat": time.Saturday,
}

var shortMonthNames = []string{
	"",
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

var shortMonthNamesInverse = map[string]time.Month{
	"Jan": time.January,
	"Feb": time.February,
	"Mar": time.March,
	"Apr": time.April,
	"May": time.May,
	"Jun": time.June,
	"Jul": time.July,
	"Aug": time.August,
	"Sep": time.September,
	"Oct": time.October,
	"Nov": time.November,
	"Dec": time.December,
}

var longMonthNames = []string{
	"",
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

	flen := len(format)
	fli := flen - 1
	for i := 0; i < flen; i++ {
		switch format[i] {
		case '%':
			if i < fli {
				switch format[i+1] {
				// 星期几 缩写
				case 'a':
					add(shortWeekDayNames[t.Weekday()])
				// 月份名 缩写
				case 'b':
					add(shortMonthNames[t.Month()])
				// 月份 1-12
				case 'c':
					add(strconv.FormatUint(uint64(t.Month()), 10))
				// 月份的天 01-31
				// NB: MySQL中是00-31
				case 'd':
					add(fmt.Sprintf("%02d", t.Day()))
				// 月份的天 序数词后缀
				case 'D':
					add(builtins.NumToWithOrdinalSuffix(t.Day()))
				// 月份的天 1-31
				// NB: MySQL中是0-31
				case 'e':
					add(strconv.FormatUint(uint64(t.Day()), 10))
				// 秒的小数部分
				case 'f':
					add(fmt.Sprintf("%06d", t.Nanosecond()/1000))
				// 小时 24小时制 00-23
				case 'H':
					add(fmt.Sprintf("%02d", t.Hour()))
				// 小时 12小时制 0-23
				case 'k':
					add(strconv.FormatUint(uint64(t.Hour()), 10))
				// 小时 12小时制 01-12
				case 'h', 'I':
					h := t.Hour()
					if h == 0 {
						add("12")
					} else if h <= 12 {
						add(fmt.Sprintf("%02d", h))
					} else {
						add(fmt.Sprintf("%02d", h-12))
					}
				// 小时 12小时制 1-12
				case 'l':
					h := t.Hour()
					if h == 0 {
						add("12")
					} else if h <= 12 {
						add(strconv.FormatUint(uint64(h), 10))
					} else {
						add(strconv.FormatUint(uint64(h-12), 10))
					}
				// 分钟 00-59
				case 'i':
					add(fmt.Sprintf("%02d", t.Minute()))
				// 年的天 001-366
				case 'j':
					add(fmt.Sprintf("%03d", t.YearDay()))
				// 月份 00-12
				case 'm':
					add(fmt.Sprintf("%02d", t.Month()))
				// 月份名 全名
				case 'M':
					add(longMonthNames[t.Month()])
				// 上午 下午 AM 或 PM
				case 'p':
					if t.Hour() < 12 {
						add("AM")
					} else {
						add("PM")
					}
				// hh:mm:ss AM或PM 12小时制
				case 'r':
					s, err := DateFormat(t, "%I:%H:%S %p")
					if err != nil {
						return "", err
					}
					add(s)
				// 秒
				case 's', 'S':
					add(fmt.Sprintf("%02d", t.Second()))
				// hh:mm:ss 24小时制
				case 'T':
					add(fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second()))
				// 周 星期日是一周的第一天 00-53
				case 'U':
					_, w := yearWeek(t, 'U')
					add(w)
				// 周 星期一是一周的第一天 00-53
				case 'u':
					_, w := yearWeek(t, 'u')
					add(w)
				// 周 星期日是一周的第一天 01-53 与%X中的年对应
				case 'V':
					_, w := yearWeek(t, 'V')
					add(w)
				// 周 星期一是一周的第一天 01-53 与%x中的年对应
				case 'v':
					_, isoWeek := t.ISOWeek()
					if isoWeek < 1e1 {
						add("0")
					}
					add(strconv.FormatUint(uint64(isoWeek), 10))
				// 星期几 全名
				case 'W':
					add(longWeekDayNames[t.Weekday()])
				// 星期几 数值 0=星期日 ,6=星期六
				case 'w':
					add(strconv.FormatUint(uint64(t.Weekday()), 10))
				// 年 星期日是一周的第一天 与%V中的年对应
				case 'X':
					y, _ := yearWeek(t, 'X')
					add(y)
				// 年 星期一是一周的第一天 与%v中的年对应
				case 'x':
					isoYear, _ := t.ISOWeek()
					if isoYear < 1e1 {
						add("0")
					}
					add(strconv.FormatUint(uint64(isoYear), 10))
				// 年 4位
				case 'Y':
					add(fmt.Sprintf("%04d", t.Year()))
				// 年 2位
				case 'y':
					add(fmt.Sprintf("%02d", t.Year()%100))
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

func StrToDate(param string, f string) (time.Time, error) {
	var t time.Time
	var err error

	year, month, day := 0, time.Month(0), 0
	hour, min, sec, nsec := 0, 0, 0, 0
	weekDay := time.Weekday(-1)

	yearX := 0
	weeksV := 0

	format := []rune(f)

	pi := 0
	plen := len(param)
	fi := 0
	flen := len(format)
	fli := flen - 1
	isPM := false

	for ; fi < flen; fi++ {
		switch format[fi] {
		case '%':
			if fi < fli {
				switch format[fi+1] {
				// 星期几 缩写
				case 'a':
					pi += 3
					if pi < plen {
						weekDay = shortWeekDayNamesInverse[FirstUpper(param[pi-3:pi])]
					} else {
						return t, lengthMismatchError
					}
				// 月份名 缩写
				case 'b':
					pi += 3
					if pi < plen {
						month = shortMonthNamesInverse[FirstUpper(param[pi-3:pi])]
					} else {
						return t, lengthMismatchError
					}
				// 月份 1-12
				case 'c':
					m, err := parseNChars(param, &pi, plen, 2, false)
					if err != nil {
						return t, err
					}
					month = time.Month(m)
				// 月份的天 01-31
				case 'd':
					day, err = parseNChars(param, &pi, plen, 2, true)
					if err != nil {
						return t, err
					}
				// 月份的天 1-31 序数词后缀
				case 'D':
					firstChar := param[pi]
					if pi < plen {
						secondChar := param[pi]
						if secondChar >= '0' && secondChar <= '9' {
							month = time.Month((firstChar-48)*10 + secondChar - 48)
						} else {
							month = time.Month(firstChar - 48)
						}
						// 校验后缀
						pi += 2
						if pi < plen {
							switch month {
							case time.January:
								if param[pi-2:pi] != "st" {
									return t, errors.New(mismatchTip + " with D")
								}
							case time.February:
								if param[pi-2:pi] != "nd" {
									return t, errors.New(mismatchTip + " with D")
								}
							case time.March:
								if param[pi-2:pi] != "rd" {
									return t, errors.New(mismatchTip + " with D")
								}
							default:
								if param[pi-2:pi] != "th" {
									return t, errors.New(mismatchTip + " with D")
								}
							}
						} else {
							return t, lengthMismatchError
						}
					} else {
						return t, lengthMismatchError
					}
					pi++
				// 月份的天 1-31
				case 'e':
					day, err = parseNChars(param, &pi, plen, 2, false)
					if err != nil {
						return t, err
					}
				// 秒的小数部分
				case 'f':
					frac, err := parseNChars(param, &pi, plen, 2, true)
					if err != nil {
						return t, err
					}
					nsec = frac * int(math.Pow10(9-len(strconv.Itoa(frac))))
					if err != nil {
						return t, nil
					}
				// 小时 24小时制 00-23
				case 'H':
					hour, err = parseNChars(param, &pi, plen, 2, true)
					if err != nil {
						return t, err
					}
				// 小时 24小时制 0-23
				case 'k':
					hour, err = parseNChars(param, &pi, plen, 2, false)
					if err != nil {
						return t, err
					}
				// 小时 12小时制 01-12
				case 'h', 'I':
					hour, err = parseNChars(param, &pi, plen, 2, true)
					if err != nil {
						return t, err
					}
				// 小时 12小时制 1-12
				case 'l':
					hour, err = parseNChars(param, &pi, plen, 2, false)
					if err != nil {
						return t, err
					}
				// 分钟 00-59
				case 'i':
					min, err = parseNChars(param, &pi, plen, 2, true)
					if err != nil {
						return t, err
					}
				// 年的天 001-366
				case 'j':
					day, err = parseNChars(param, &pi, plen, 3, false)
					if err != nil {
						return t, err
					}
					month = time.January
				// 月份 00-12
				case 'm':
					m, err := parseNChars(param, &pi, plen, 2, true)
					if err != nil {
						return t, err
					}
					month = time.Month(m)
				case 'M':
					month, err = matchMonthByFullName(&pi, plen, param)
					if err != nil {
						return t, mismatchError
					}
				// 上午 下午 AM 或 PM
				case 'p':
					pi += 2
					if pi < plen {
						chs := param[pi-2 : pi]
						if strings.ToUpper(chs) == "AM" {
						} else if strings.ToUpper(chs) == "PM" {
							isPM = true
						} else {
							return t, mismatchError
						}
					} else {
						return t, lengthMismatchError
					}
				// hh:mm:ss AM或PM 12小时制
				case 'r':
					start := pi
					pi += 11
					if pi < plen {
						hour, err = strconv.Atoi(param[start : start+2])
						if err != nil {
							return t, err
						}
						if param[start+2] == ':' {
							start += 3
						} else {
							return t, mismatchError
						}
						min, err = strconv.Atoi(param[start : start+2])
						if err != nil {
							return t, err
						}
						if param[start+2] == ':' {
							start += 3
						} else {
							return t, mismatchError
						}
						sec, err = strconv.Atoi(param[start : start+2])
						if err != nil {
							return t, err
						}
						if param[start+2] == ' ' {
							start += 3
						} else {
							return t, mismatchError
						}
						meridiem := param[start : start+2]
						if strings.ToUpper(meridiem) == "AM" {
						} else if strings.ToUpper(meridiem) == "PM" {
							isPM = true
						} else {
							return t, mismatchError
						}
					} else {
						return t, lengthMismatchError
					}
				// 秒 00-59
				case 's', 'S':
					sec, err = parseNChars(param, &pi, plen, 2, true)
					if err != nil {
						return t, err
					}
				// hh:mm:ss 24小时制
				case 'T':
					start := pi
					pi += 8
					if pi < plen {
						hour, err = strconv.Atoi(param[start : start+2])
						if err != nil {
							return t, err
						}
						if param[start+2] == ':' {
							start += 3
						} else {
							return t, mismatchError
						}
						min, err = strconv.Atoi(param[start : start+2])
						if err != nil {
							return t, err
						}
						if param[start+2] == ':' {
							start += 3
						} else {
							return t, mismatchError
						}
						sec, err = strconv.Atoi(param[start : start+2])
						if err != nil {
							return t, err
						}
					} else {
						return t, lengthMismatchError
					}
				// 周 星期日是一周的第一天 00-53
				// NB: MySQL中不支持
				case 'U':
					return t, errors.New("not support")
				// 周 星期一是一周的第一天 00-53
				// NB: MySQL中不支持
				case 'u':
					return t, errors.New("not support")
				// 周 星期日是一周的第一天 01-53 与%X中的年对应
				// X V + 星期几
				case 'V':
					weeksV, err = parseNChars(param, &pi, plen, 2, false)
					if err != nil {
						return t, err
					}
				// 周 星期一是一周的第一天 01-53 与%x中的年对应 TODO
				// x v + 星期几
				case 'v':
				// 星期几 全名
				case 'W':
					weekDay, err = matchWeekByFullName(&pi, plen, param)
					if err != nil {
						return t, err
					}
				// 星期几 数值 0=星期日 ,6=星期六
				case 'w':
					w, err := parseNChars(param, &pi, plen, 1, true)
					if err != nil {
						return t, err
					}
					weekDay = time.Weekday(w)
				// 年 星期日是一周的第一天 与%V中的年对应
				// X V + 星期几
				case 'X':
					yearX, err = parseNChars(param, &pi, plen, 4, true)
					if err != nil {
						return t, err
					}
				// 年 星期一是一周的第一天 与%v中的年对应 TODO
				// x v + 星期几
				case 'x':
				// 年 4位
				case 'Y':
					year, err = parseNChars(param, &pi, plen, 4, true)
					if err != nil {
						return t, err
					}
				// 年 2位
				case 'y':
					year, err = parseNChars(param, &pi, plen, 2, true)
					if year <= 69 {
						year = 2000 + year
					} else {
						year = 1900 + year
					}
					if err != nil {
						return t, err
					}
				case '%':
					pi++
				default:
					return t, illegalCharacterError
				}
				fi++
			}
		default:
			if rune(param[pi]) != format[fi] {
				return t, mismatchError
			}
			pi++
		}
	}

	// TODO 格式重复问题待解决

	// 最后处理部分
	if isPM && hour > 12 {
		hour += 12
	}

	// X V + 星期 确定年月日
	if yearX != 0 || weeksV != 0 || weekDay != -1 {
		if yearX != 0 && weeksV != 0 && weekDay != -1 {
			tXV, err := yearWeekInverse(yearX, weeksV, weekDay, true)
			if err != nil {
				return t, err
			}
			year = tXV.Year()
			month = tXV.Month()
			day = tXV.Day()
		} else {
			return t, errors.New("format error")
		}
	}

	return time.Date(year, month, day, hour, min, sec, nsec, time.UTC), err
}

func parseNChars(param string, pi *int, plen int, n int, strict bool) (int, error) {
	var num int
	var err error
	tmp := bytes.Buffer{}
	if *pi < plen {
		i := 0
		for ; i < n && *pi < plen; i++ {
			nextChar := param[*pi]
			if nextChar >= '0' && nextChar <= '9' {
				tmp.WriteByte(nextChar)
			} else {
				if strict {
					errors.New("out of range error")
				} else {
					break
				}
			}
			*pi++
		}
		if i != n && strict {
			return num, lengthMismatchError
		}
		num, err = strconv.Atoi(tmp.String())
	} else {
		return num, lengthMismatchError
	}
	return num, err
}
func matchMonthByFullName(pi *int, plen int, param string) (time.Month, error) {
	found := false
	var month time.Month
	var err error
	switch param[*pi] {
	case 'A', 'a':
		found, err = matchIgnoreCase(pi, plen, param, "APRIL")
		if found {
			month = time.April
		}
		if !found {
			found, err = matchIgnoreCase(pi, plen, param, "AUGUST")
			if found {
				month = time.August
			}
		}
	// December
	case 'D', 'd':
		found, err = matchIgnoreCase(pi, plen, param, "DECEMBER")
		if found {
			month = time.December
		}
	// February
	case 'F', 'f':
		found, err = matchIgnoreCase(pi, plen, param, "FEBRUARY")
		if found {
			month = time.February
		}
	// January June July
	case 'J', 'j':
		found, err = matchIgnoreCase(pi, plen, param, "JULY")
		if found {
			month = time.July
		}
		if !found {
			found, err = matchIgnoreCase(pi, plen, param, "JUNE")
			if found {
				month = time.June
			}
		}
		if !found {
			found, err = matchIgnoreCase(pi, plen, param, "JANUARY")
			if found {
				month = time.January
			}
		}
	// March May
	case 'M', 'm':
		found, err = matchIgnoreCase(pi, plen, param, "MAY")
		if found {
			month = time.May
		}
		if !found {
			found, err = matchIgnoreCase(pi, plen, param, "MARCH")
			if found {
				month = time.March
			}
		}
	// November
	case 'N', 'n':
		found, err = matchIgnoreCase(pi, plen, param, "NOVEMBER")
		if found {
			month = time.November
		}
	// October
	case 'O', 'o':
		found, err = matchIgnoreCase(pi, plen, param, "OCTOBER")
		if found {
			month = time.October
		}
	// September
	case 'S', 's':
		found, err = matchIgnoreCase(pi, plen, param, "SEPTEMBER")
		if found {
			month = time.September
		}
	default:
		return month, lengthMismatchError
	}

	if !found {
		return month, mismatchError
	}
	return month, err
}

func matchIgnoreCase(pi *int, plen int, param string, match string) (bool, error) {
	found := false
	n := len(match)
	*pi += n
	if *pi < plen {
		if strings.ToUpper(param[*pi-n:*pi]) == match {
			found = true
		} else {
			*pi -= n
			return found, illegalCharacterError
		}
	} else {
		return found, lengthMismatchError
	}
	return found, nil
}

// S Sunday Saturday
// M Monday
// T Tuesday Thursday
// W Wednesday
// F Friday
func matchWeekByFullName(pi *int, plen int, param string) (time.Weekday, error) {
	var week time.Weekday = -1
	var err error
	found := false
	if *pi < plen {
		switch param[*pi] {
		case 'F', 'f':
			found, err = matchIgnoreCase(pi, plen, param, "RIDAY")
			if found {
				week = time.Friday
			}
		case 'M', 'm':
			found, err = matchIgnoreCase(pi, plen, param, "ONDAY")
			if found {
				week = time.Monday
			}
		case 'S', 's':
			// 2
			if *pi < plen {
				switch param[*pi] {
				case 'A', 'a':
					found, err = matchIgnoreCase(pi, plen, param, "TURDAY")
					if found {
						week = time.Saturday
					}
				case 'U', 'u':
					found, err = matchIgnoreCase(pi, plen, param, "NDAY")
					if found {
						week = time.Sunday
					}
				default:
					return week, illegalCharacterError
				}
			} else {
				return week, lengthMismatchError
			}
			*pi++
		case 'T', 't':
			// 2
			if *pi < plen {
				switch param[*pi] {
				case 'H', 'h':
					found, err = matchIgnoreCase(pi, plen, param, "URSDAY")
					if found {
						week = time.Thursday
					}
				case 'U', 'u':
					found, err = matchIgnoreCase(pi, plen, param, "ESDAY")
					if found {
						week = time.Tuesday
					}
				default:
					return week, illegalCharacterError
				}
			} else {
				return week, lengthMismatchError
			}
			*pi++
		case 'W', 'w':
			found, err = matchIgnoreCase(pi, plen, param, "EDNESDAY")
			if found {
				week = time.Wednesday
			}
		default:
			return week, illegalCharacterError
		}
	} else {
		return week, lengthMismatchError
	}
	if !found {
		return week, mismatchError
	}

	return week, err
}

func yearWeek(t time.Time, mode byte) (string, string) {
	var sundayAsFirstDay bool
	switch mode {
	case 'U', 'V', 'X':
		sundayAsFirstDay = true
	case 'u', 'v', 'x':
		sundayAsFirstDay = false
	default:
		panic(illegalCharacterError)
	}

	y := t.Year()
	w := 0
	yearDay := t.YearDay()

	daysOfFirstWeek := getDaysOfFirstWeek(t, sundayAsFirstDay)

	diff := yearDay - daysOfFirstWeek

	if diff <= 0 {
		if mode == 'V' || mode == 'v' || mode == 'X' || mode == 'x' {
			return yearWeek(t.AddDate(0, 0, -yearDay), mode)
		}
	} else if diff < 7 {
		w = 1
	} else if diff >= 7 {
		w = diff / 7
		if diff%7 != 0 {
			w++
		}
	}
	if mode == 'u' && daysOfFirstWeek >= 4 {
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

func yearWeekInverse(year int, weeks int, weekDay time.Weekday, sundayAsFirstDay bool) (time.Time, error) {
	firstDay := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	if !sundayAsFirstDay {
		// TODO
	}

	return firstDay.AddDate(0, 0, -int(firstDay.Weekday())+(weeks)*7+int(weekDay)), nil
}

// FirstUpper 字符串首字母大写
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}
