package oracle_demo

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestToDate(txx *testing.T) {
	dch := "2018-05-06"
	format := "YYYY-MM-DD"
	tm, err := ToDate(dch, format)
	if err != nil {
		panic(err)
	}

	fmt.Println(*tm)
}

func TestToDate2(txx *testing.T) {
	dch := "20180506"
	format := "YYYYMMDD"
	tm, err := ToDate(dch, format)
	if err != nil {
		panic(err)
	}

	fmt.Println(*tm)
}

func TestToDate3(txx *testing.T) {
	dch := "2018-05-06 01:02:05"
	format := "YYYY-MM-DD HH:MI:SS"
	tm, err := ToDate(dch, format)
	if err != nil {
		panic(err)
	}

	if tm != nil {
		fmt.Println(*tm)
	}
	//log.Println(tm)
}

func TestToChar(t *testing.T) {
	tm := time.Date(2017, 02, 27, 20, 20, 20, 20, time.Local)
	fmt.Println(tm)
	format := "YYYY-MM-DD"

	str, err := ToCharByDatetime(tm, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}

func TestToChar2(t *testing.T) {
	tm := time.Date(2017, 02, 27, 9, 10, 30, 50, time.Local)
	fmt.Println(tm)
	format := "YYYY-MM-DD HH:MI:SS"

	str, err := ToCharByDatetime(tm, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}

func TestToNumber(t *testing.T) {
	numParam := "123456"
	format := "999999EEEE"

	numResult, err := ToNumber(numParam, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(numResult)
}

// select to_char(-123.56,'L9,999.999') from dual;
//
//	-￥123.560
func TestToCharByNumDecimal1(t *testing.T) {
	numFloat := -123.56
	format := "L9,999.999"
	s, err := ToCharByNum(numFloat, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}

// select to_char('-123.56','999v999PR') from dual;
func TestToCharByNumDecimal2(t *testing.T) {
	numFloat := -123.56
	format := "999v999PR"
	s, err := ToCharByNum(numFloat, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}

// select to_char('123.56','999999PR') from dual;
func TestToCharByNumDecimal3(t *testing.T) {
	numFloat := 123.56
	format := "999999PR"
	s, err := ToCharByNum(numFloat, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}

// select to_char('321456.789', '000009,99,9.9999') from dual;
//
//	000321,45,6.7890
func TestToCharByNumDecimal4(t *testing.T) {
	numFloat := 321456.789
	format := "000009,99,9.9999"
	s, err := ToCharByNum(numFloat, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}

// select to_char('321456.789', '000009,99,9.99') from dual;
//
//	000321,45,6.79
func TestToCharByNumDecimal5(t *testing.T) {
	numFloat := 321456.789
	format := "000009,99,9.99"
	s, err := ToCharByNum(numFloat, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}

// select to_char('321456.789', 'fm000009,99,9.9999') from dual;
//
//	000321,45,6.789
func TestToCharByNumDecimal6(t *testing.T) {
	numFloat := 321456.789
	format := "fm000009,99,9.9999"
	s, err := ToCharByNum(numFloat, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}

// select to_char('321456.789', '0000000000099.99999999999EEEE') from dual;
//
//	000321,45,6.789
func TestToCharByNumEEEE(t *testing.T) {
	numFloat := 321456.789
	format := "0000000000099.99999999999EEEE"
	s, err := ToCharByNum(numFloat, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}

// select to_char('-321456.789', '9999999V9PR') from dual;
//
// <3214568>
func TestToCharByNumV(t *testing.T) {
	numFloat := -321456.789
	format := "9999999V9PR"
	s, err := ToCharByNum(numFloat, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}

// select to_char('-321456.789', 'L9999999V9PR') from dual;
//
// <3214568>
func TestToCharByNumV2(t *testing.T) {
	numFloat := -321456.789
	format := "L9999999V9PR"
	s, err := ToCharByNum(numFloat, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}

// select to_char('-123.56','L9,999.999') from dual;
func TestToCharByStrDecimal1(t *testing.T) {
	numStr := "-123.56"
	format := "999v999PR"
	s, err := ToCharByStr(numStr, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}

// select to_char('-123.56','999v999PR') from dual;
func TestToCharByStrDecimal2(t *testing.T) {
	numStr := "-123.56"
	format := "999v999PR"
	s, err := ToCharByStr(numStr, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}

// select to_char('123.56','999v999PR') from dual;
func TestToCharByStrDecimal3(t *testing.T) {
	numStr := "123.56"
	format := "999v999PR"
	s, err := ToCharByStr(numStr, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}

// select to_char('321456.789', '000009,99,9.9999') from dual;
//
//	000321,45,6.7890
func TestToCharByStrDecimal4(t *testing.T) {
	numStr := "321456.789"
	format := "000009,99,9.9999"
	s, err := ToCharByStr(numStr, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}

// select to_char('321456.789', '000009,99,9.99') from dual;
//
//	000321,45,6.79
func TestToCharByStrDecimal5(t *testing.T) {
	numStr := "321456.789"
	format := "000009,99,9.99"
	s, err := ToCharByStr(numStr, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}

// select to_char('321456.789', 'fm000009,99,9.9999') from dual;
//
//	000321,45,6.789
func TestToCharByStrDecimal6(t *testing.T) {
	numStr := "321456.789"
	format := "fm000009,99,9.9999"
	s, err := ToCharByStr(numStr, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}

// select to_char('321456.789', '0000000000099.99999999999EEEE') from dual;
//
//	000321,45,6.789
func TestToCharByStrEEEE(t *testing.T) {
	numStr := "321456.789"
	format := "0000000000099.99999999999EEEE"
	s, err := ToCharByStr(numStr, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}

// select to_char('-321456.789', '9999999V9') from dual;
//
// -3214568
func TestToCharByStrV(t *testing.T) {
	numStr := "-321456.789"
	format := "9999999V9"
	s, err := ToCharByStr(numStr, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}

func TestEEEE(t *testing.T) {
	numFloat := 321456.789
	fmt.Println(strconv.FormatFloat(numFloat, 'E', 3, 64))
}

func TestEEEE2(t *testing.T) {

	pre := "321456"
	post := "789"

	fmt.Println(strToEEEEFormat(pre, post))

	pre = "0"
	post = "00789"
	fmt.Println(strToEEEEFormat(pre, post))
}

func strToEEEEFormat(pre string, post string) string {
	result := bytes.Buffer{}
	exp := 0
	if pre == "0" || pre == "" {
		i := 0
		for ; i < len(post); i++ {
			if post[i] != '0' {
				break
			}
		}

		result.WriteByte(post[i])
		result.WriteByte('.')
		result.WriteString(post[i+1:])
		result.WriteByte('E')
		result.WriteByte('-')
		exp = i + 1
	} else {
		result.WriteByte(pre[0])
		result.WriteByte('.')
		result.WriteString(pre[1:])
		result.WriteString(post)
		result.WriteByte('E')
		result.WriteByte('+')
		exp = len(pre) - 1
	}

	if exp < 10 {
		result.WriteByte('0')
	}
	result.WriteString(fmt.Sprint(exp))
	return result.String()
}
