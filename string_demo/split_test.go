package string_demo

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"
)

func TestSplit(t *testing.T) {
	srcStr1 := "abc:def:k:g"
	desStr1 := strings.Split(srcStr1, ":")
	fmt.Printf("ret:%s\n", desStr1)

	srcStr2 := "a b c,def,k,g"
	desStr2 := strings.FieldsFunc(srcStr2, splitFunc)
	fmt.Printf("ret:%s\n", desStr2)
}

func splitFunc(r rune) bool {
	return r == ' ' || r == ','
}

func TestSplit2(t *testing.T) {
	src := "10:18$%$%$%#25.9999999999999999999999999999999299999"
	out := regexp.MustCompile("^[0-9]+").FindAllString(src, 6)

	fmt.Printf("ret:%s\n", out)

}

func Test3(t *testing.T) {
	str := `3223-08-03` // `9x_xx:995`
	// 使用命名分组，一次匹配多个值
	re := regexp.MustCompile(`(?P<year>[1-9]{4})(?P<month>[1-9]{2})(?P<day>[0-9]{2})`)
	match := re.FindStringSubmatch(str)
	groupNames := re.SubexpNames()
	fmt.Printf("%v, %v, %d, %d\n", match, groupNames, len(match), len(groupNames))

	result := make(map[string]string)
	if len(match) == len(groupNames) {
		// 转换为map
		for i, name := range groupNames {
			if i != 0 && name != "" { // 第一个分组为空（也就是整个匹配）
				result[name] = match[i]
			}
		}
	}
}

func Test6(t *testing.T) {
	str := "123abcdefghijk123"
	// 找到字符串中的123
	regx, ok := regexp.Compile("123") // 创建正则表达式规则对象

	str1 := regx.FindAllString(str, -1) // -1 代表找到所有的

	fmt.Println(ok)

	fmt.Println(str1) //  str1 = [123 123]
}

func Test7(t *testing.T) {
	str := "10:18$%$%$%#25.9999999999999999999999999999999299999"
	regx, ok := regexp.Compile("[0-9\\.]+") // 创建正则表达式规则对象
	str1 := regx.FindAllString(str, -1)     // -1 代表找到所有的
	fmt.Println(ok)
	fmt.Println(str1) //  str1 = [123 123]
}

func parse_time(s string) (*time.Time, error) {
	yyyymmddhhmmssFmt := "20060102150405"
	yyyymmddFmt := "20060102"
	hhmmssFmt := "150405"

	regx, err := regexp.Compile("[0-9]+")
	if err != nil {
		return nil, err
	}

	arr := regx.FindAllString(s, 6)
	n := len(arr)

	var t = time.Time{}
	if n == 6 {
		yyyymmddhhmmss := strings.Join(arr[0:6], "")
		t, err = time.Parse(yyyymmddhhmmssFmt, yyyymmddhhmmss)
	} else if n == 3 {
		if len(arr[0]) > 2 {
			yymmdd := strings.Join(arr[0:3], "")
			t, err = time.Parse(yyyymmddFmt, yymmdd)
		} else {
			hhmmss := strings.Join(arr[3:6], "")
			t, err = time.Parse(hhmmssFmt, hhmmss)
		}
	} else {
		l := len(s)
		if l == 6 {
			t, err = time.Parse(yyyymmddhhmmssFmt, s)
		} else if l == 3 {
			if len(arr[0]) > 2 {
				yymmdd := strings.Join(arr[0:3], "")
				t, err = time.Parse(yyyymmddFmt, yymmdd)
			} else {
				hhmmss := strings.Join(arr[3:6], "")
				t, err = time.Parse(hhmmssFmt, hhmmss)
			}
		}
	}

	return &t, err
}

func TestParseTime(xxx *testing.T) {
	yyyymmddhhmmssFmt := "20060102150405"
	yyyymmddhhmmss := "00001011080706"
	t, err := time.Parse(yyyymmddhhmmssFmt, yyyymmddhhmmss)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t.Hour())
}

func TestParseTime2(xxx *testing.T) {
	yyyymmddhhmmssFmt := "20060102150405"
	yyyymmddhhmmss := "00010203040706"
	t, err := time.Parse(yyyymmddhhmmssFmt, yyyymmddhhmmss)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(int8(t.Month()))
	fmt.Printf("%d\n", uint8(t.Month()))
	fmt.Println("====")
	fmt.Println(t.Day())
	fmt.Println(t.Hour())
	fmt.Println(t.Minute())
	fmt.Println(t.Second())
}

func TestParseTime3(xxx *testing.T) {
	yyyymmddhhmmssFmt := "\"2006-01-02 15:04:05\""
	yyyymmddhhmmss := "\"2022-02-28 10:10:10\""
	t, err := time.Parse(yyyymmddhhmmssFmt, yyyymmddhhmmss)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t.Hour())
}

func TestParseDuration(t *testing.T) {
	d, err := time.ParseDuration("2006-01-02 15:04:05")
	if err != nil {

	}

	fmt.Println(d.Hours())
}
