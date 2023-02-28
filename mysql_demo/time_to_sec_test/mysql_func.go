package mysql_func_demo

import (
	"errors"
	"math/bits"
	"strconv"
	"strings"
	"unicode/utf8"
)

func Elt(n int, str []string) (string, error) {
	if n < 1 || n > len(str) {
		return "", errors.New("array out of bounds")
	}
	return str[n-1], nil
}

func MakeSet(index int, strSet []string) (string, error) {
	if index < 0 {
		return "", errors.New("not support")
	}
	numBitsLen := bits.Len64(uint64(index))
	strSetLen := len(strSet)
	var ret []string

	for i := 0; i < numBitsLen && i < strSetLen; i++ {
		if (index&(1<<i)) != 0 && strSet[i] != "" {
			ret = append(ret, strSet[i])
		}
	}
	return strings.Join(ret, ","), nil
}

// Oct 十进制转换八进制
func Oct[T int | int64 | float32 | float64 | string](t T) (string, error) {
	var ret string
	switch x := any(t).(type) {
	case int:
		if x < 0 {
			return "", errors.New("not support")
		}
		ret = strconv.FormatInt(int64(x), 8)
	case int64:
		if x < 0 {
			return "", errors.New("not support")
		}
		ret = strconv.FormatInt(x, 8)
	case float32:
		if x < 0 {
			return "", errors.New("not support")
		}
		ret = strconv.FormatInt(int64(x), 8)
	case float64:
		if x < 0 {
			return "", errors.New("not support")
		}
		ret = strconv.FormatInt(int64(x), 8)
	case string:
		f, err := strconv.ParseFloat(x, 64)
		if err != nil {
			return "", err
		}
		if f < 0 {
			return "", errors.New("not support")
		}
		ret = strconv.FormatInt(int64(f), 8)
	}

	return ret, nil
}

// Ord 返回字符串首个字符的utf8编码
func Ord(str string) int {
	if str == "" {
		return 0
	}
	code := 0
	r, size := utf8.DecodeRuneInString(str)
	if size == 1 {
		code = int(r)
	} else {
		bytes := []byte(string(r))
		for _, b := range bytes {
			code = code<<8 + int(b)
		}
	}
	return code
}

// FindInSet 对逗号分隔的字符串分隔，返回第1个完全相等的下标
// @Param strs 用逗号分隔字符串
// @Return 索引从1开始，不存在就返回0
func FindInSet(s, strs string) int {
	hasComma := strings.Index(strs, ",")
	if hasComma >= 0 {
		strArr := strings.Split(strs, ",")
		strArrLen := len(strArr)
		for i := 0; i < strArrLen; i++ {
			if s == strArr[i] {
				return i + 1
			}
		}
		return 0
	} else {
		if s == strs {
			return 1
		} else {
			return 0
		}
	}
}

// SubstringIndex
// @Param s 带匹配字符串
// @Param sep 区分大小写
// @Param count
//
//	0 返回空字符串
//	负数 从右向左 匹配 取右半部分字符串
//	正数 从左向右 匹配 取左半部分字符串
//
// @Return 不包含分隔符
func SubstringIndex2(s, sep string, count int) string {
	if count == 0 {
		return ""
	}

	var ret []string
	sepIndex := strings.Index(s, sep)
	if sepIndex >= 0 {
		strArr := strings.Split(s, sep)
		strArrLen := len(strArr)
		if count > 0 {
			if count >= strArrLen {
				return s
			} else {
				for i := 0; i < count; i++ {
					ret = append(ret, strArr[i])
				}
				return strings.Join(ret, sep)
			}
		} else {
			if -count >= strArrLen {
				return s
			} else {
				for i := strArrLen + count; i < strArrLen; i++ {
					ret = append(ret, strArr[i])
				}
				return strings.Join(ret, sep)
			}
		}
	} else {
		return s
	}
}
