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
	"strings"
	"time"
)

// FullTimeFormat is the time format used to display any unknown timestamp
// type, and always shows the full time zone offset.
const FullTimeFormat = "2006-01-02 15:04:05.999999-07:00:00"

// TimestampWithTZFormat is the time format used to display
// timestamps with a time zone offset. The minutes and seconds
// offsets are only added if they are non-zero.
const TimestampWithTZFormat = "2006-01-02 15:04:05.999999-07"

// TimestampWithoutTZFormat is the time format used to display
// timestamps without a time zone offset. The minutes and seconds
// offsets are only added if they are non-zero.
const TimestampWithoutTZFormat = "2006-01-02 15:04:05.999999"

// TimeWithTZFormat is the time format used to display a time
// with a time zone offset.
const TimeWithTZFormat = "15:04:05.999999-07"

// TimeWithoutTZFormat is the time format used to display a time
// without a time zone offset.
const TimeWithoutTZFormat = "15:04:05.999999"

// DateWithMinusFormat is the time format used to display a date.
const DateWithMinusFormat = "2006-01-02"
const DateWithSlashFormat = "2006/01/02"
const DateFormat = "20060102"
const DatetimeFormat = "20060102150405"
const DatetimeFormat2 = "2006-01-02 15:04:05"
const DatetimeFormat3 = "2006/01/02 15:04:05"

// TimestampNumWithoutTZFormat is the time format used to display
// timestamps without a time zone offset. The minutes and seconds
// offsets are only added if they are non-zero.
const TimestampNumWithoutTZFormat = "20060102150405.999999"

var format = "20060102150405999999999"
var formatLen = len(format)

var defaultFormat = []string{
	"2006",
	"01",
	"02",
	"15",
	"04",
	"05",
	"999999999",
}
var dfLen = len(defaultFormat)

func parseTimeSimple(s string) (*time.Time, error) {
	var t = new(time.Time)
	var err error
	l := len(s)
	customFormat := []string{}

	fi := 0
	for i := 0; i < l; i++ {
		c := s[i]
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if fi < formatLen {
				customFormat = append(customFormat, string(format[fi]))
				fi++
			} else {
				return nil, errors.New("length error")
			}
		case ' ', '-', ':', ',', '.', '/', ';':
			customFormat = append(customFormat, string(c))
		default:
			return nil, errors.New("illegal character")
		}
	}

	*t, err = time.Parse(strings.Join(customFormat, ""), s)
	if err != nil {
		return nil, err
	}
	return t, nil
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
				//fmt.Println(defaultFormat[dfi])
				//fmt.Println(customFormat.String())
				//fmt.Println(count)
				if count < len(defaultFormat[dfi]) {
					count++
				} else {
					customFormat.WriteString(defaultFormat[dfi])
					//fmt.Println(defaultFormat[dfi])
					//fmt.Println(customFormat.String())
					dfi++
					count = 0
				}
			} else if c == ' ' || c == '-' || c == ':' || c == ',' || c == '/' || c == ';' || c == '.' {
				if count != 0 {
					customFormat.WriteString(defaultFormat[dfi][len(defaultFormat[dfi])-count:])
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
	customFormat.WriteString(defaultFormat[dfi][0:])

	//f := strings.Join(customFormat, "")
	fmt.Println(customFormat.String())
	*t, err = time.Parse(customFormat.String(), s)
	if err != nil {
		return nil, err
	}
	return t, nil
}
