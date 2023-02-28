// Copyright 2020 The Hubble Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package timeutil

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// FullTimeFormat is the time YYYYMMDDhhmmssfffffffff used to display any unknown timestamp
// type, and always shows the full time zone offset.
const FullTimeFormat = "2006-01-02 15:04:05.999999-07:00:00"

// TimestampWithTZFormat is the time YYYYMMDDhhmmssfffffffff used to display
// timestamps with a time zone offset. The minutes and seconds
// offsets are only added if they are non-zero.
const TimestampWithTZFormat = "2006-01-02 15:04:05.999999-07"

// TimestampWithoutTZFormat is the time YYYYMMDDhhmmssfffffffff used to display
// timestamps without a time zone offset. The minutes and seconds
// offsets are only added if they are non-zero.
const TimestampWithoutTZFormat = "2006-01-02 15:04:05.999999"

// TimeWithTZFormat is the time YYYYMMDDhhmmssfffffffff used to display a time
// with a time zone offset.
const TimeWithTZFormat = "15:04:05.999999-07"

// TimeWithoutTZFormat is the time YYYYMMDDhhmmssfffffffff used to display a time
// without a time zone offset.
const TimeWithoutTZFormat = "15:04:05.999999"

// DateWithMinusFormat is the time YYYYMMDDhhmmssfffffffff used to display a date.
const DateWithMinusFormat = "2006-01-02"
const DateWithSlashFormat = "2006/01/02"
const DateFormat = "20060102"
const DatetimeFormat = "20060102150405"
const DatetimeFormat2 = "2006-01-02 15:04:05"
const DatetimeFormat3 = "2006/01/02 15:04:05"

// TimestampNumWithoutTZFormat is the time YYYYMMDDhhmmssfffffffff used to display
// timestamps without a time zone offset. The minutes and seconds
// offsets are only added if they are non-zero.
const TimestampNumWithoutTZFormat = "20060102150405.999999"

var YYYYMMDDhhmmssfffffffff = "20060102150405999999999"
var YYMM = "0601"
var YYYYMM = "200601"
var hhmmss = "150405"
var hhmmssfffffffff = "150405999999999"

var defaultFormatSlice = []string{
	"2006",
	"01",
	"02",
	"15",
	"04",
	"05",
	"999999999",
}
var dfLen = len(defaultFormatSlice)

// ParseTimeSimple
// golang 中 YYMM [00,68]表示20xx年,[69,99]表示19xx年
// mysql  中 YYMM [00,69]表示20xx年,[70,99]表示19xx年
func ParseTimeSimple(s string) (*time.Time, string, error) {
	var t = new(time.Time)
	var err error

	// YYMM
	// 根据MySQL中的时间格式,此处表示go中的1969，需要转成go中的2069
	if len(s) == 4 {
		year, err := strconv.Atoi(s[0:2])
		if err != nil {
			return nil, "", err
		}
		if year == 69 {
			s = "2069" + s[2:4]
			*t, err = time.Parse(YYYYMM, s)
		} else {
			*t, err = time.Parse(YYMM, s)
		}
		if err != nil {
			return nil, "", err
		}
		return t, YYMM, nil
	}

	customFormat, err := parseFormat(s, YYYYMMDDhhmmssfffffffff)
	if err != nil {
		return nil, "", err
	}
	switch len(customFormat) {
	case 3:
		// HHmmss
		if len(customFormat[0]) <= 2 ||
			len(customFormat[0]) >= 3 && (customFormat[0][2] < '0' || customFormat[0][2] > '9') {
			customFormat, err = parseFormat(s, hhmmss)
			if err != nil {
				return nil, "", err
			}
		}
	case 4:
		// HHmmss.FFFFFFFFF 150405999999999
		customFormat, err = parseFormat(s, hhmmssfffffffff)
		if err != nil {
			return nil, "", err
		}
	}
	format := strings.Join(customFormat, "")
	*t, err = time.Parse(format, s)
	if err != nil {
		return nil, format, err
	}
	return t, format, nil
}

func parseFormat(s, format string) ([]string, error) {
	customFormat := []string{}
	l := len(s)

	fi := 0
	var group bytes.Buffer
	for i := 0; i < l; i++ {
		c := s[i]
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if fi < l {
				group.WriteByte(format[fi])
				fi++
			} else {
				return nil, errors.New("length error")
			}
		case ' ', '-', ':', ',', '.', '/', ';':
			group.WriteByte(c)
			customFormat = append(customFormat, group.String())
			group.Reset()
		default:
			return nil, errors.New("illegal character")
		}
	}
	return append(customFormat, group.String()), nil
}

func parseTimeByOrder2(s string) (*time.Time, error) {
	var t = new(time.Time)
	var err error
	l := len(s)
	customFormat := bytes.Buffer{}

	count := 0
	dfi := 0
	for i := 0; i < l; i++ {
		c := s[i]
		//fmt.Println(string(c))
		if dfi < dfLen {
			if c >= '0' && c <= '9' {
				//fmt.Println(defaultFormatSlice[dfi])
				//fmt.Println(customFormat.String())
				//fmt.Println(count)
				if count < len(defaultFormatSlice[dfi]) {
					count++
				} else {
					customFormat.WriteString(defaultFormatSlice[dfi])
					//fmt.Println(defaultFormatSlice[dfi])
					//fmt.Println(customFormat.String())
					dfi++
					count = 0
				}
			} else if c == ' ' || c == '-' || c == ':' || c == ',' || c == '/' || c == ';' || c == '.' {
				if count != 0 {
					customFormat.WriteString(defaultFormatSlice[dfi][len(defaultFormatSlice[dfi])-count:])
					customFormat.WriteByte(c)
					dfi++
					count = 0
				}
			} else {
				return nil, errors.New("illegal character")
			}
		} else {
			return nil, errors.New("length error")
		}
	}

	fmt.Println(count)
	customFormat.WriteByte('.')
	customFormat.WriteString(defaultFormatSlice[dfi][0:])

	//f := strings.Join(customFormat, "")
	fmt.Println(customFormat.String())
	*t, err = time.Parse(customFormat.String(), s)
	if err != nil {
		return nil, err
	}
	return t, nil
}
