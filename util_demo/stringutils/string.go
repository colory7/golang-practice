package stringutils

import (
	"bytes"
	"errors"
	"unicode/utf8"
)

// Substring 按照开始索引和结束索引截取字符串
// @Param source
// @Param start 字符开始索引
// @Param count 字符个数
func SubstringByCount(src string, start, count int) (string, error) {
	if count <= 0 {
		return "", nil
	}

	if start >= 0 {
		return substring(src, start, start+count)
	} else {
		start = utf8.RuneCountInString(src) + start
		return substring(src, start, start+count)
	}

}

func Substring(src string, start int, end int) (string, error) {
	if start < 0 {
		return "", errors.New("index error")
	}
	if end < 0 {
		return "", errors.New("index error")
	}

	if start >= end {
		return "", errors.New("start index is greater than or equal to end index error")
	}

	return substring(src, start, end)
}

// substring 按照开始索引和结束索引截取字符串
// @Param source 字符串
// @Param start 字符开始索引，从0开始
// @Param end 字符结束索引，从0开始
func substring(src string, start, end int) (string, error) {
	var ret bytes.Buffer
	var pos = 0
	for _, ch := range src {
		if pos < start {
			pos++
			continue
		}
		if pos >= end {
			break
		}
		ret.WriteRune(ch)
		pos++
	}
	return ret.String(), nil
}
