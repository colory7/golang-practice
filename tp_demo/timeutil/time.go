package timeutil

import (
	"strings"
	"time"
)

// Copyright 2016 The Hubble Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

// LibPQTimePrefix is the prefix lib/pq prints time-type datatypes with.
const LibPQTimePrefix = "0000-01-01"

// Now returns the current UTC time.
func Now() time.Time {
	return time.Now().UTC()
}

// Since returns the time elapsed since t.
// It is shorthand for Now().Sub(t), but more efficient.
func Since(t time.Time) time.Duration {
	return time.Since(t)
}

// Until returns the duration until t.
// It is shorthand for t.Sub(Now()), but more efficient.
func Until(t time.Time) time.Duration {
	return time.Until(t)
}

// UnixEpoch represents the Unix epoch, January 1, 1970 UTC.
var UnixEpoch = time.Unix(0, 0).UTC()

// FromUnixMicros returns the UTC time.Time corresponding to the given Unix
// time, usec microseconds since UnixEpoch. In Go's current time.Time
// implementation, all possible values for us can be represented as a time.Time.
func FromUnixMicros(us int64) time.Time {
	return time.Unix(us/1e6, (us%1e6)*1e3).UTC()
}

// ToUnixMicros returns t as the number of microseconds elapsed since UnixEpoch.
// Fractional microseconds are rounded, half up, using time.Round. Similar to
// time.Time.UnixNano, the result is undefined if the Unix time in microseconds
// cannot be represented by an int64.
func ToUnixMicros(t time.Time) int64 {
	return t.Unix()*1e6 + int64(t.Round(time.Microsecond).Nanosecond())/1e3
}

// Unix wraps time.Unix ensuring that the result is in UTC instead of Local.
func Unix(sec, nsec int64) time.Time {
	return time.Unix(sec, nsec).UTC()
}

// ReplaceLibPQTimePrefix replaces unparsable lib/pq dates used for timestamps
// (0000-01-01) with timestamps that can be parsed by date libraries.
func ReplaceLibPQTimePrefix(s string) string {
	if strings.HasPrefix(s, LibPQTimePrefix) {
		return "1970-01-01" + s[len(LibPQTimePrefix):]
	}
	return s
}

// TODO To simplify the date function, the current date YYYYMMDDhhmmssfffffffff support is limited
// Currently, only the following time formats are supported:
// timeutil.DateWithMinusFormat
// timeutil.TimestampWithoutTZFormat
// timeutil.TimeWithoutTZFormat
// timeutil.TimestampNumWithoutTZFormat
func ParseTime(s string) (time.Time, string, error) {
	var t time.Time
	var err error
	var format string
	if strings.ContainsAny(s, "-") {
		if strings.ContainsAny(s, ":") {
			t, err = time.Parse(TimestampWithoutTZFormat, s)
			format = TimestampWithoutTZFormat
		} else {
			t, err = time.Parse(DateWithMinusFormat, s)
			format = DateWithMinusFormat
		}
	} else if strings.ContainsAny(s, ":") {
		t, err = time.Parse(TimeWithoutTZFormat, s)
		format = TimeWithoutTZFormat
	} else {
		t, err = time.Parse(TimestampNumWithoutTZFormat, s)
		format = TimestampNumWithoutTZFormat
	}

	return t, format, err
}

func ParseTime2(s string) (time.Time, string, error) {
	var t time.Time
	var err error
	var format string

	t, err = time.Parse(TimestampWithoutTZFormat, s)
	format = TimestampWithoutTZFormat

	return t, format, err
}
