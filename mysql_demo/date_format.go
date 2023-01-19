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

var mismatchError = "mismatch error"
var contentMismatchError = errors.New("format and parameter content mismatch error")
var lengthMismatchError = errors.New("format and parameter length mismatch error")
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
	var t *time.Time
	var err error

	year, month, day := 0, time.Month(0), 0
	hour, min, sec, nsec := 0, 0, 0, 0
	var week *time.Weekday

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
						*week = shortWeekDayNamesInverse[param[pi-3:pi]]
					} else {
						return *t, lengthMismatchError
					}
				// 月份名 缩写
				case 'b':
					pi += 3
					if pi < plen {
						month = shortMonthNamesInverse[param[pi-3:pi]]
					} else {
						return *t, lengthMismatchError
					}
				// 月份 1-12
				case 'c':
					firstChar := param[pi]
					pi++
					if pi < plen {
						secondChar := param[pi]
						if secondChar >= '0' && secondChar <= '9' {
							month = time.Month((firstChar-30)*10 + (secondChar - 30))
						}
					} else {
						return *t, lengthMismatchError
					}
				// 月份的天 01-31
				case 'd':
					pi += 2
					if pi < plen {
						day, err = strconv.Atoi(param[pi-2 : pi])
						if err != nil {
							return *t, err
						}
					} else {
						return *t, lengthMismatchError
					}
				// 月份的天 1-31 序数词后缀
				case 'D':
					firstChar := param[pi]
					pi++
					if pi < plen {
						secondChar := param[pi]
						if secondChar >= '0' && secondChar <= '9' {
							month = time.Month((firstChar-30)*10 + secondChar - 30)
						} else {
							month = time.Month(firstChar - 30)
						}
						// 校验后缀
						pi += 2
						if pi < plen {
							switch month {
							case time.January:
								if param[pi-2:pi] != "st" {
									return *t, errors.New(mismatchError + " with D")
								}
							case time.February:
								if param[pi-2:pi] != "nd" {
									return *t, errors.New(mismatchError + " with D")
								}
							case time.March:
								if param[pi-2:pi] != "rd" {
									return *t, errors.New(mismatchError + " with D")
								}
							default:
								if param[pi-2:pi] != "th" {
									return *t, errors.New(mismatchError + " with D")
								}
							}
						} else {
							return *t, lengthMismatchError
						}
					} else {
						return *t, lengthMismatchError
					}
				// 月份的天 1-31
				case 'e':
					firstChar := param[pi]
					pi++
					if pi < plen {
						secondChar := param[pi]
						if secondChar >= '0' && secondChar <= '9' {
							month = time.Month((firstChar-30)*10 + secondChar - 30)
						} else {
							month = time.Month(firstChar - 30)
						}
					} else {
						month = time.Month(firstChar - 30)
					}
				// 秒的小数部分
				case 'f':
					tmp := bytes.Buffer{}
					for i := 0; i < 9 && pi < plen; i++ {
						pi++
						nextChar := param[pi]
						if nextChar >= '0' && nextChar <= '9' {
							tmp.WriteByte(nextChar)
						} else {
							break
						}
					}
					fracStr := tmp.String()
					frac, err := strconv.Atoi(fracStr)
					nsec = frac * int(math.Pow10(9-len(fracStr)))
					if err != nil {
						return *t, nil
					}
				// 小时 24小时制 00-23
				case 'H':
					pi += 2
					if pi < plen {
						hour, err = strconv.Atoi(param[pi-2 : pi])
						if err != nil {
							return *t, err
						}
					} else {
						return *t, lengthMismatchError
					}
				// 小时 24小时制 0-23
				case 'k':
					firstChar := param[pi]
					pi++
					if pi < plen {
						secondChar := param[pi]
						if secondChar >= '0' && secondChar <= '9' {
							hour = int((firstChar-30)*10 + secondChar - 30)
						} else {
							hour = int(firstChar - 30)
						}
					} else {
						hour = int(firstChar - 30)
					}
				// 小时 12小时制 01-12
				case 'h', 'I':
					pi += 2
					if pi < plen {
						hour, err = strconv.Atoi(param[pi-2 : pi])
						if err != nil {
							return *t, contentMismatchError
						}
					} else {
						return *t, lengthMismatchError
					}
				// 小时 12小时制 1-12
				case 'l':
					firstChar := param[pi]
					pi++
					if pi < plen {
						secondChar := param[pi]
						if secondChar >= '0' && secondChar <= '9' {
							hour = int((firstChar-30)*10 + secondChar - 30)
						} else {
							hour = int(firstChar - 30)
						}
					} else {
						return *t, lengthMismatchError
					}
				// 分钟 00-59
				case 'i':
					pi += 2
					if pi < plen {
						min, err = strconv.Atoi(param[pi-2 : pi])
						if err != nil {
							return *t, contentMismatchError
						}
					} else {
						return *t, lengthMismatchError
					}
				// 年的天 001-366
				case 'j':
					tmp := bytes.Buffer{}
					for i := 0; i < 3 && pi < plen; i++ {
						pi++
						nextChar := param[pi]
						if nextChar >= '0' && nextChar <= '9' {
							tmp.WriteByte(nextChar)
						} else {
							break
						}
					}
					day, err = strconv.Atoi(tmp.String())
					if err != nil {
						return *t, err
					}
					month = time.January
				// 月份 00-12
				case 'm':
					pi += 2
					if pi < plen {
						monthInt, err := strconv.Atoi(param[pi-2 : pi])
						month = time.Month(monthInt)
						if err != nil {
							return *t, contentMismatchError
						}
					} else {
						return *t, lengthMismatchError
					}
				case 'M':
					month, err = matchMonthByFullName(&pi, plen, param)
					if err != nil {
						return *t, contentMismatchError
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
							return *t, contentMismatchError
						}
					} else {
						return *t, lengthMismatchError
					}
				// hh:mm:ss AM或PM 12小时制
				case 'r':
					start := pi
					pi += 11
					if pi < plen {
						hour, err = strconv.Atoi(param[start : start+2])
						if err != nil {
							return *t, err
						}
						if param[start+2] == ':' {
							start += 3
						} else {
							return *t, contentMismatchError
						}
						min, err = strconv.Atoi(param[start : start+2])
						if err != nil {
							return *t, err
						}
						if param[start+2] == ':' {
							start += 3
						} else {
							return *t, contentMismatchError
						}
						sec, err = strconv.Atoi(param[start : start+2])
						if err != nil {
							return *t, err
						}
						if param[start+2] == ' ' {
							start += 3
						} else {
							return *t, contentMismatchError
						}
						meridiem := param[start : start+2]
						if strings.ToUpper(meridiem) == "AM" {
						} else if strings.ToUpper(meridiem) == "PM" {
							isPM = true
						} else {
							return *t, contentMismatchError
						}
					} else {
						return *t, lengthMismatchError
					}
				// 秒 00-59
				case 's', 'S':
					pi += 2
					if pi < plen {
						sec, err = strconv.Atoi(param[pi-2 : pi])
						if err != nil {
							return *t, contentMismatchError
						}
					} else {
						return *t, lengthMismatchError
					}
				// hh:mm:ss 24小时制
				case 'T':
					start := pi
					pi += 8
					if pi < plen {
						hour, err = strconv.Atoi(param[start : start+2])
						if err != nil {
							return *t, err
						}
						if param[start+2] == ':' {
							start += 3
						} else {
							return *t, contentMismatchError
						}
						min, err = strconv.Atoi(param[start : start+2])
						if err != nil {
							return *t, err
						}
						if param[start+2] == ':' {
							start += 3
						} else {
							return *t, contentMismatchError
						}
						sec, err = strconv.Atoi(param[start : start+2])
						if err != nil {
							return *t, err
						}
					} else {
						return *t, lengthMismatchError
					}
				// 周 星期日是一周的第一天 00-53 TODO
				case 'U':
				// 周 星期一是一周的第一天 00-53 TODO
				case 'u':
				// 周 星期日是一周的第一天 01-53 与%X中的年对应 TODO
				case 'V':
				// 周 星期一是一周的第一天 01-53 与%x中的年对应 TODO
				case 'v':

				// 星期几 全名
				case 'W':
					*week, err = matchWeekByFullName(&pi, plen, param)
					if err != nil {
						return *t, err
					}
				// 星期几 数值 0=星期日 ,6=星期六
				case 'w':
					pi++
					if pi < plen {
						ch := param[pi]
						if ch >= '0' && ch <= '9' {
							*week = (time.Weekday)(ch - 30)
						} else {
							return *t, errors.New("out of range error")
						}
					} else {
						return *t, lengthMismatchError
					}
				// 年 星期日是一周的第一天 与%V中的年对应 TODO
				case 'X':
				// 年 星期一是一周的第一天 与%v中的年对应 TODO
				case 'x':
				// 年 4位
				case 'Y':
					pi += 4
					if pi < plen {
						year, err = strconv.Atoi(param[pi-4 : pi])
						if err != nil {
							return *t, err
						}
					} else {
						return *t, lengthMismatchError
					}
				// 年 2位
				case 'y':
					pi += 2
					if pi < plen {
						year, err = strconv.Atoi(param[pi-2 : pi])
						if err != nil {
							return *t, err
						}
					} else {
						return *t, lengthMismatchError
					}
				case '%':
					pi++
				default:
					return *t, illegalCharacterError
				}
				fi++
			}
		default:
			pi++
		}
	}

	// 最后处理
	if isPM && hour > 12 {
		hour += 12
	}

	*t = time.Date(year, month, day, hour, min, sec, nsec, time.UTC)
	return *t, err

}

func matchMonthByFullName(pi *int, plen int, param string) (time.Month, error) {
	found := false
	var month time.Month
	var err error
	*pi++
	switch param[*pi] {
	case 'A', 'a':
		*pi++
		if *pi < plen {
			switch param[*pi] {
			// April
			case 'P', 'p':
				found, err = matchIgnoreCase(pi, plen, param, "ril")
				if found {
					month = time.April
				}
			// August
			case 'U', 'u':
				found, err = matchIgnoreCase(pi, plen, param, "gust")
				if found {
					month = time.August
				}
			default:
				err = contentMismatchError
			}
		} else {
			err = lengthMismatchError
		}
	// December
	case 'D', 'd':
		found, err = matchIgnoreCase(pi, plen, param, "ecember")
		if found {
			month = time.December
		}
	// February
	case 'F', 'f':
		found, err = matchIgnoreCase(pi, plen, param, "ebruary")
		if found {
			month = time.December
		}
	// January June July
	case 'J', 'j':
		found, err = matchIgnoreCase(pi, plen, param, "uly")
		if found {
			month = time.July
		}
		if !found {
			found, err = matchIgnoreCase(pi, plen, param, "une")
			if found {
				month = time.June
			}
		}
		if !found {
			found, err = matchIgnoreCase(pi, plen, param, "anuary")
			if found {
				month = time.January
			}
		}
	// March May
	case 'M', 'm':
		found, err = matchIgnoreCase(pi, plen, param, "ay")
		if found {
			month = time.May
		}
		if !found {
			found, err = matchIgnoreCase(pi, plen, param, "arch")
			if found {
				month = time.March
			}
		}
	// November
	case 'N', 'n':
		found, err = matchIgnoreCase(pi, plen, param, "ovember")
		if found {
			month = time.November
		}
	// October
	case 'O', 'o':
		found, err = matchIgnoreCase(pi, plen, param, "ctober")
		if found {
			month = time.October
		}
	// September
	case 'S', 's':
		found, err = matchIgnoreCase(pi, plen, param, "eptember")
		if found {
			month = time.September
		}
	default:
		return month, lengthMismatchError
	}

	if !found {
		return month, contentMismatchError
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
			return found, illegalCharacterError
		}
	} else {
		return found, lengthMismatchError
	}
	return found, nil
}

func matchWeekByFullName(pi *int, plen int, param string) (time.Weekday, error) {
	//S Sunday Saturday
	//M Monday
	//T Tuesday Thursday
	//W Wednesday
	//F Friday
	var week *time.Weekday
	var err error
	found := false
	*pi++
	if *pi < plen {
		switch param[*pi] {
		case 'F', 'f':
			found, err = matchIgnoreCase(pi, plen, param, "RIDAY")
			if found {
				*week = time.Friday
			}
		case 'M', 'm':
			found, err = matchIgnoreCase(pi, plen, param, "ONDAY")
			if found {
				*week = time.Monday
			}
		case 'S', 's':
			// 2
			*pi++
			if *pi < plen {
				switch param[*pi] {
				case 'A', 'a':
					found, err = matchIgnoreCase(pi, plen, param, "TURDAY")
					if found {
						*week = time.Saturday
					}
				case 'U', 'u':
					found, err = matchIgnoreCase(pi, plen, param, "NDAY")
					if found {
						*week = time.Sunday
					}
				default:
					return *week, illegalCharacterError
				}
			} else {
				return *week, lengthMismatchError
			}
		case 'T', 't':
			// 2
			*pi++
			if *pi < plen {
				switch param[*pi] {
				case 'H', 'h':
					found, err = matchIgnoreCase(pi, plen, param, "URSDAY")
					if found {
						*week = time.Thursday
					}
				case 'U', 'u':
					found, err = matchIgnoreCase(pi, plen, param, "ESDAY")
					if found {
						*week = time.Tuesday
					}
				default:
					return *week, illegalCharacterError
				}
			} else {
				return *week, lengthMismatchError
			}
		case 'W', 'w':
			found, err = matchIgnoreCase(pi, plen, param, "EDNESDAY")
			if found {
				*week = time.Wednesday
			}
		default:
			return *week, illegalCharacterError
		}
	} else {
		return *week, lengthMismatchError
	}
	if !found {
		return *week, contentMismatchError
	}

	return *week, err
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

	dfw := getDaysOfFirstWeek(t, sundayAsFirstDay)

	diff := yearDay - dfw

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
