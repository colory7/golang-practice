package oracle_demo

//
//import (
//	"bytes"
//	"errors"
//	"fmt"
//	"github.com/stretchr/testify/assert"
//	"math"
//	"strconv"
//	"strings"
//	"testing"
//	"time"
//)
//
////NumberFormatModel
////DatetimeFormatModel
//
//func TestMatchNumberFormatModel(t *testing.T) {
//	// 1.找到模式,最大匹配
//	// 2.根据模式替换字符串
//
//	//param:="111"
//	format := "999"
//	parseNumFormat(format)
//}
//
//func TestAscii(t *testing.T) {
//	theme := "狙击 start"
//	for i := 0; i < len(theme); i++ {
//		fmt.Printf("ascii:%c %d\n", theme[i], theme[i])
//	}
//}
//
//func TestAscii2(t *testing.T) {
//	theme := "狙击 start"
//	for i := 0; i < len(theme); i++ {
//		f := theme[i]
//		if f >= 32 && f <= 127 {
//			fmt.Println((string)(f))
//		}
//	}
//}
//
//func TestE(t *testing.T) {
//	fmt.Printf("%f\n", 1.345e2*10) //e9 就是小数点向右移动9位
//
//	fmt.Printf("%.10f\n", 12344e-9) //e-9就是小数点向左移动9位，%.10f表示精确到小数点后10位
//
//	fmt.Printf("%.2e\n", 12312312321312123123123.0)
//	fmt.Printf("%.2E", 12312312321312123123123.0)
//	fmt.Println("%")
//}
//
//func TestE2(t *testing.T) {
//	preDec := ""
//	postDec := "10"
//	eNum := 1.345e2
//	v := fmt.Sprintf("%"+preDec+"."+postDec+"f", eNum)
//	fmt.Println(v)
//
//}
//
//func TestStrSearch(t *testing.T) {
//	fmt.Println(strings.Index("widuu", "i")) //1
//	fmt.Println(strings.Index("widuu", "u")) //3
//}
//
//func TestV(t *testing.T) {
//	//f := "99V999"
//	//preV := "2"
//	postV := "3"
//
//	param := "12"
//	paramNum, err := strconv.Atoi(param)
//	if err != nil {
//		panic(err)
//	}
//	postVNum, _ := strconv.Atoi(postV)
//	if err != nil {
//		panic(err)
//	}
//
//	result := paramNum * int(math.Pow10(postVNum))
//
//	fmt.Println(result)
//
//}
//
//func TestToString(t *testing.T) {
//	var numParam NumParamDesc
//	numParam.sign = plus
//	numParam.preDec = "36"
//	numParam.postDec = "25"
//	numParam.eSign = empty
//	numParam.eExponent = 12
//	fmt.Println(numParam.string())
//}
//
//func TestParseNumFmt(t *testing.T) {
//	f := "99EEEE"
//	numFmtDesc, err := parseNumFormat(f)
//	assert.NoError(t, err)
//
//	str := fmt.Sprintf("%#v\n", numFmtDesc)
//	fmt.Println(str)
//}
//
//func TestParseNumParam(t *testing.T) {
//	num := "-36.25e+97"
//	numParamDesc, err := parseNumParam(num)
//	assert.NoError(t, err)
//	str := fmt.Sprintf("%#v\n", numParamDesc)
//	fmt.Println(str)
//}
//
//func TestParseNum(t *testing.T) {
//	f := "99"
//	num := "-36.25"
//
//	result, err := ToNumber(f, num)
//	assert.NoError(t, err)
//	assert.Equal(t, "-36", result)
//	fmt.Println(result)
//}
//
//func TestParseNum2(t *testing.T) {
//	f := "99999999999999999999"
//	num := "-36.25e+3"
//
//	result, err := ToNumber(f, num)
//	assert.NoError(t, err)
//	fmt.Println(result)
//}
//
//func TestParseNum2Err(t *testing.T) {
//	f := "9"
//	num := "-36.25e+3"
//
//	result, err := ToNumber(f, num)
//	assert.NoError(t, err)
//	fmt.Println(result)
//}
//
//func TestParseNum3(t *testing.T) {
//	f := "99EEEE"
//	num := "-36.25e+3"
//
//	result, err := ToNumber(f, num)
//	assert.NoError(t, err)
//	fmt.Println(result)
//}
//
//func TestParseDchByStrAD(t *testing.T) {
//	param := "公元"
//	format := "AD"
//
//	//fmt.Println(format[2:4])
//
//	parseDchByStr(param, format)
//}
//
//func TestParseDchByStrA_D_(t *testing.T) {
//	param := "公元"
//	format := "A.D."
//
//	//fmt.Println(format[2:4])
//
//	parseDchByStr(param, format)
//}
//
//func TestParseDchByStrAM(t *testing.T) {
//	param := "上午"
//	format := "AM"
//
//	//fmt.Println(format[2:4])
//
//	parseDchByStr(param, format)
//}
//
//func TestParseDchByStrA_M_(t *testing.T) {
//	param := "上午"
//	format := "A.M."
//
//	//fmt.Println(format[2:4])
//
//	parseDchByStr(param, format)
//}
//
//func TestParseDchByStrY_YYY(t *testing.T) {
//	param := "2013"
//	format := "yyyy"
//
//	result, err := parseDchByStr(param, format)
//	assert.NoError(t, err)
//	fmt.Println(result)
//
//}
//
//func TestParseDchByStrY_YYY2(t *testing.T) {
//	param := "2,013"
//	format := "Y,YYY"
//
//	result, err := parseDchByStr(param, format)
//	assert.NoError(t, err)
//	fmt.Println(result)
//
//	param = "213"
//	format = "YYY"
//
//	result, err = parseDchByStr(param, format)
//	assert.NoError(t, err)
//	fmt.Println(result)
//
//	param = "13"
//	format = "YY"
//
//	result, err = parseDchByStr(param, format)
//	assert.NoError(t, err)
//	fmt.Println(result)
//
//	param = "3"
//	format = "Y"
//
//	result, err = parseDchByStr(param, format)
//	assert.NoError(t, err)
//	fmt.Println(result)
//
//}
//
//func TestParseDchByStrADY_YYY(t *testing.T) {
//	param := "公元 2023"
//	format := "A.D. yyyy"
//
//	result, err := parseDchByStr(param, format)
//	assert.NoError(t, err)
//	fmt.Println(result)
//}
//
//func TestParseDchByStr2(t *testing.T) {
//	param := "公元 2023-10-26 01:30:56"
//	format := "A.D. yyyy-MM-dd hh:mi:ss"
//
//	result, err := parseDchByStr(param, format)
//	assert.NoError(t, err)
//	fmt.Println(result)
//}
//
//func TestCentury(t *testing.T) {
//
//	ti := time.Now()
//	fmt.Println(ti.Year())
//	fmt.Println((ti.Year()-1)/100 + 1)
//	fmt.Println((1900-1)/100 + 1)
//	fmt.Println((1901-1)/100 + 1)
//	fmt.Println((2000-1)/100 + 1)
//
//	fmt.Println("====")
//	fmt.Println((ti.Year() + 99) / 100)
//	fmt.Println((1900 + 99) / 100)
//	fmt.Println((1901 + 99) / 100)
//	fmt.Println((2000 + 99) / 100)
//	fmt.Println((2001 + 99) / 100)
//}
//func TestChar(t *testing.T) {
//	param := "公元 2023"
//	fmt.Println(param[0:3])
//	fmt.Println(param[0:6])
//	fmt.Println(param[3:6])
//	fmt.Println(param[0:8])
//	fmt.Println(param[0:10])
//	fmt.Println(param[0:11])
//	fmt.Println(param[3:11])
//	fmt.Println("====")
//	fmt.Println(param[4:11])
//	fmt.Println(param[5:11])
//	fmt.Println(param[6:11])
//	fmt.Println(param[7:11])
//	fmt.Println(param[8:11])
//}
//
//func TestJulian(t *testing.T) {
//	fmt.Println(toJulian(2022, 12, 19))
//	fmt.Println(toJulian(2023, 10, 29))
//}
//
//func TestToRoman(t *testing.T) {
//	b := toRoman(4278)
//	fmt.Println(b.String())
//}
//
//func TestParseDchByTimeEmpty(t *testing.T) {
//	layout := "2006-01-02 15:04:05"
//	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")
//
//	format := ""
//	actual, err := parseFmt(ti, format)
//	assert.NoError(t, err)
//	expected := ""
//
//	if actual != expected {
//		t.Errorf("actual %q expected %q", actual, expected)
//	}
//}
//
//func TestParseDchByTimeSkipChars(t *testing.T) {
//	layout := "2006-01-02 15:04:05"
//	ti, _ := time.Parse(layout, "2021-01-10 15:01:02")
//
//	format := "./- ,,,:;,,"
//	actual, err := parseFmt(ti, format)
//	assert.NoError(t, err)
//	expected := format + "fs"
//
//	// 是否符合我们的预期
//	assert.Equal(t, expected, actual)
//
//}
//
//func TestParseDchByTimeA(t *testing.T) {
//	layout := "2006-01-02 15:04:05"
//	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")
//
//	format := "AM"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "A.M."
//	fmt.Println(parseFmt(ti, format))
//
//	format = "AD"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "A.D."
//	fmt.Println(parseFmt(ti, format))
//
//}
//
//func TestParseDchByTimeAErr(t *testing.T) {
//	layout := "2006-01-02 15:04:05"
//	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")
//
//	format := "A"
//	actual, err := parseFmt(ti, format)
//	expected := empty_str
//	assert.Equal(t, expected, actual)
//	assert.Error(t, err)
//}
//
//func TestParseDchByTimeAErr2(t *testing.T) {
//	layout := "2006-01-02 15:04:05"
//	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")
//
//	format := "A."
//	actual, err := parseFmt(ti, format)
//	expected := empty_str
//	assert.Equal(t, expected, actual)
//	assert.Error(t, err)
//}
//
//func TestParseDchByTimeAErr3(t *testing.T) {
//	layout := "2006-01-02 15:04:05"
//	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")
//
//	format := "AX"
//	actual, err := parseFmt(ti, format)
//	expected := empty_str
//	assert.Equal(t, expected, actual)
//	assert.Error(t, err)
//}
//
//func TestParseDchByTimeAErr4(t *testing.T) {
//	layout := "2006-01-02 15:04:05"
//	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")
//
//	format := "A.D"
//
//	actual, err := parseFmt(ti, format)
//	expected := empty_str
//	assert.Equal(t, expected, actual)
//	assert.Error(t, err)
//}
//
//func TestParseDchByTimeB(t *testing.T) {
//	layout := "2006-01-02 15:04:05"
//	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")
//
//	format := "BC"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "B.C."
//	fmt.Println(parseFmt(ti, format))
//}
//
//func TestParseDchByTimeAB(t *testing.T) {
//	layout := "2006-01-02 15:04:05"
//	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")
//
//	format := "ADBC"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "BCAD"
//	fmt.Println(parseFmt(ti, format))
//}
//
//func TestParseDchByTimeC(t *testing.T) {
//	layout := "2006-01-02 15:04:05"
//	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")
//
//	format := "CC"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "C"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "BCADCC"
//	fmt.Println(parseFmt(ti, format))
//}
//
//func TestParseDchByTimeD(t *testing.T) {
//	layout := "2006-01-02 15:04:05"
//	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")
//
//	var format = ""
//	format = "DAY"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "DD"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "DDD"
//	fmt.Println(parseFmt(ti, format))
//
//	format = NLS_DL
//	fmt.Println(parseFmt(ti, format))
//
//	format = "DL"
//	fmt.Println(parseFmt(ti, format))
//
//	format = NLS_DS
//	fmt.Println(parseFmt(ti, format))
//
//	format = "DS"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "DY"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "D"
//	fmt.Println(parseFmt(ti, format))
//}
//func TestParseDchByTimeF(t *testing.T) {
//	layout := "2006-01-02 15:04:05.000000000"
//	ti, _ := time.Parse(layout, "2031-01-10 15:01:02.789321456")
//
//	var format = ""
//	format = "FXFM"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "FF1"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "FF2"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "FF3"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "FF4"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "FF5"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "FF6"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "FF7"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "FF8"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "FF9"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "FF1-FF2-FF9-FF8"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "FF"
//	actual, err := parseFmt(ti, format)
//	expected := empty_str
//	assert.Equal(t, expected, actual)
//	assert.Error(t, err)
//
//	format = "F"
//	actual, err = parseFmt(ti, format)
//	expected = empty_str
//	assert.Equal(t, expected, actual)
//	assert.Error(t, err)
//
//	format = "F0"
//	actual, err = parseFmt(ti, format)
//	expected = empty_str
//	assert.Equal(t, expected, actual)
//	assert.Error(t, err)
//
//	format = "FFD"
//	actual, err = parseFmt(ti, format)
//	expected = empty_str
//	assert.Equal(t, expected, actual)
//	assert.Error(t, err)
//}
//
//func TestParseDchByTimeH(t *testing.T) {
//	layout := "2006-01-02 15:04:05.000000000"
//	ti, _ := time.Parse(layout, "2031-01-10 15:01:02.789321456")
//
//	var format = ""
//	format = "HH"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "H24"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "H12"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "H24-H12"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "H2"
//	actual, err := parseFmt(ti, format)
//	expected := empty_str
//	assert.Equal(t, expected, actual)
//	assert.Error(t, err)
//
//	format = "H2D"
//	actual, err = parseFmt(ti, format)
//	expected = empty_str
//	assert.Equal(t, expected, actual)
//	assert.Error(t, err)
//
//	format = "H1"
//	actual, err = parseFmt(ti, format)
//	expected = empty_str
//	assert.Equal(t, expected, actual)
//	assert.Error(t, err)
//
//	format = "H13"
//	actual, err = parseFmt(ti, format)
//	expected = empty_str
//	assert.Equal(t, expected, actual)
//	assert.Error(t, err)
//}
//
//func TestParseDchByTimeH2(t *testing.T) {
//	layout := "2006-01-02 15:04:05.000000000"
//	ti, _ := time.Parse(layout, "2031-01-10 02:01:02.789321456")
//
//	var format = ""
//	format = "HH"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "H24"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "H12"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "H24-H12"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "H2"
//	actual, err := parseFmt(ti, format)
//	expected := empty_str
//	assert.Equal(t, expected, actual)
//	assert.Error(t, err)
//
//	format = "H2D"
//	actual, err = parseFmt(ti, format)
//	expected = empty_str
//	assert.Equal(t, expected, actual)
//	assert.Error(t, err)
//
//	format = "H1"
//	actual, err = parseFmt(ti, format)
//	expected = empty_str
//	assert.Equal(t, expected, actual)
//	assert.Error(t, err)
//
//	format = "H13"
//	actual, err = parseFmt(ti, format)
//	expected = empty_str
//	assert.Equal(t, expected, actual)
//	assert.Error(t, err)
//}
//
//func TestParseDchByTimeI(t *testing.T) {
//	layout := "2006-01-02 15:04:05.000000000"
//	ti, _ := time.Parse(layout, "2031-01-10 02:01:02.789321456")
//	format := "I"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "IYYY"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "IYY"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "IY"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "IW"
//	fmt.Println(parseFmt(ti, format))
//}
//
//func TestParseDchByTimeJ(t *testing.T) {
//	layout := "2006-01-02 15:04:05.000000000"
//	ti, _ := time.Parse(layout, "2031-01-10 02:01:02.789321456")
//	format := "J"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "JD"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "DJD"
//	fmt.Println(parseFmt(ti, format))
//}
//
//func TestParseDchByTimeM(t *testing.T) {
//	layout := "2006-01-02 15:04:05"
//	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")
//
//	format := "MI"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "MM"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "MON"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "MONTH"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "MONTHMON"
//	fmt.Println(parseFmt(ti, format))
//
//}
//
//func TestParseDchByTimeM2(t *testing.T) {
//	layout := "2006-01-02 15:04:05"
//	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")
//
//	format := "MMYYYY"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "MM\"健康\"YYYY"
//	fmt.Println(parseFmt(ti, format))
//}
//
//func TestParseDchByTimePM(t *testing.T) {
//	layout := "2006-01-02 15:04:05"
//	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")
//
//	format := "PM"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "P.M."
//	fmt.Println(parseFmt(ti, format))
//
//	format = "P.M.-A.M....PM" //FIXME 少1个点没有输出
//	fmt.Println(parseFmt(ti, format))
//
//	format = "P.M"
//	fmt.Println(parseFmt(ti, format))
//	actual, err := parseFmt(ti, format)
//	expected := empty_str
//	assert.Equal(t, expected, actual)
//	assert.Error(t, err)
//
//	format = "P.MX"
//	fmt.Println(parseFmt(ti, format))
//	actual, err = parseFmt(ti, format)
//	expected = empty_str
//	assert.Equal(t, expected, actual)
//	assert.Error(t, err)
//}
//
//func TestParseDchByTimeQ(t *testing.T) {
//	layout := "2006-01-02 15:04:05"
//	ti, _ := time.Parse(layout, "2031-04-01 15:01:02")
//
//	format := "Q"
//	fmt.Println(parseFmt(ti, format))
//}
//
//func TestParseDchByTimeR(t *testing.T) {
//	layout := "2006-01-02 15:04:05"
//	ti, _ := time.Parse(layout, "2031-04-01 15:01:02")
//
//	format := "RR"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "RM"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "RRRR"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "R"
//	actual, err := parseFmt(ti, format)
//	expected := empty_str
//	assert.Equal(t, expected, actual)
//	assert.Error(t, err)
//
//	format = "RT"
//	actual, err = parseFmt(ti, format)
//	expected = empty_str
//	assert.Equal(t, expected, actual)
//	assert.Error(t, err)
//
//	format = "RRR"
//	actual, err = parseFmt(ti, format)
//	expected = empty_str
//	assert.Equal(t, expected, actual)
//	assert.Error(t, err)
//
//	format = "RRR."
//	actual, err = parseFmt(ti, format)
//	expected = empty_str
//	assert.Equal(t, expected, actual)
//	assert.Error(t, err)
//
//	format = "RRR-"
//	actual, err = parseFmt(ti, format)
//	expected = empty_str
//	assert.Equal(t, expected, actual)
//	assert.Error(t, err)
//
//	format = "RRRU"
//	actual, err = parseFmt(ti, format)
//	expected = empty_str
//	assert.Equal(t, expected, actual)
//	assert.Error(t, err)
//
//}
//
//func TestParseDchByTimeS(t *testing.T) {
//	layout := "2006-01-02 15:04:05"
//	ti, _ := time.Parse(layout, "2031-02-21 15:01:02")
//
//	format := "SSSSS"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "SS"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "SP"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "SYEAR"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "SYYYY"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "SSSSSSSSS"
//	fmt.Println(parseFmt(ti, format))
//
//}
//
//func TestParseDchByTimeT(t *testing.T) {
//	layout := "2006-01-02 15:04:05"
//	ti, _ := time.Parse(layout, "2031-02-21 15:01:02")
//
//	format := "TS"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "TZD"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "TZH"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "TZM"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "TZR"
//	fmt.Println(parseFmt(ti, format))
//}
//
//func TestParseDchByTimeW(t *testing.T) {
//	layout := "2006-01-02 15:04:05"
//	ti, _ := time.Parse(layout, "2031-02-21 15:01:02")
//
//	format := "W"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "WW"
//	fmt.Println(parseFmt(ti, format))
//}
//
//func TestParseDchByTimeX(t *testing.T) {
//	layout := "2006-01-02 15:04:05"
//	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")
//
//	format := "X"
//	fmt.Println(parseFmt(ti, format))
//}
//
//func TestParseDchByTimeY(t *testing.T) {
//	layout := "2006-01-02 15:04:05"
//	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")
//
//	format := "Y"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "YY"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "YYY"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "YYYY"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "Y,YYY"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "YEAR"
//	fmt.Println(parseFmt(ti, format))
//
//	format = "YEARYYYY"
//	fmt.Println(parseFmt(ti, format))
//
//}
//
//func TestParseDchByTimeYYYYMMDD(t *testing.T) {
//	layout := "2006-01-02 15:04:05"
//	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")
//
//	format := "YYYY----MM--DD"
//	fmt.Println(parseFmt(ti, format))
//}
//
//func parseDchByStr(param string, format string) (string, error) {
//	//var keywordGroup = make([]keyword, 4)
//
//	result := bytes.Buffer{}
//	flen := len(format)
//
//	pi := 0
//	for fi := 0; fi < flen; {
//		// 截取一个字符
//		c := format[fi]
//		if c >= 32 && c <= 127 {
//			fmt.Println("c: " + (string)(c))
//
//			frest := flen - fi
//
//			// 不区分大小写 FIXME
//			toUpper(&c)
//
//			// 匹配关键词并存储
//			switch c {
//			// 跳过字符
//			case '-', '/', ',', '.', ';', ':', ' ':
//				fmt.Println("skip: " + string(c))
//				fmt.Println(fi)
//				pc := param[pi]
//				switch pc {
//				case '-', '/', ',', '.', ';', ':', ' ':
//					fmt.Println("TODO: skip char " + string(pc))
//					fi++
//					pi++
//				default:
//					return empty_str, errors.New("不匹配的字符: " + string(pc))
//				}
//			case '"':
//				var skipWord bytes.Buffer
//				for ; fi < flen; fi++ {
//					if '"' == format[fi] {
//						break
//					} else {
//						skipWord.WriteByte(format[fi])
//					}
//				}
//				result.Write(skipWord.Bytes())
//				// 日期类型参数 输出
//				// 字符串类型参数 只做匹配
//			case 'A':
//				fi++
//				followingOneChar := format[fi]
//				switch followingOneChar {
//				case '.':
//					fi++
//					j := fi + 2
//					if j <= len(format) {
//						followingChars := format[fi:j]
//						if "D." == followingChars {
//							// DCH A.D.
//							pe := pi + len(NLS_AD)
//							v := param[pi:pe]
//							if v == NLS_AD {
//								fmt.Println("TODO: " + NLS_AD)
//								result.WriteString(NLS_AD)
//								fi = j
//								pi = pe
//							} else {
//								return empty_str, errors.New("语法错误,参数与 A.D. 格式不匹配")
//							}
//						} else if "M." == followingChars {
//							// DCH A.M.
//							pe := pi + len(NLS_AD)
//							v := param[pi:pe]
//							if v == NLS_AM {
//								fmt.Println("TODO: " + NLS_AM)
//								result.WriteString(NLS_AM)
//								fi = j
//								pi = pe
//
//							} else {
//								return empty_str, errors.New("语法错误,参数与 A.M. 格式不匹配")
//							}
//						} else {
//							return empty_str, errors.New(dch_fmt_mismatch_err + "A.")
//						}
//					} else {
//						return empty_str, errors.New(dch_fmt_mismatch_err + "A.")
//					}
//				case 'D':
//					// DCH AD
//					fi++
//					pe := pi + len(NLS_AD)
//					v := param[pi:pe]
//					if v == NLS_AD {
//						fmt.Println("TODO: " + NLS_AD)
//						result.WriteString(NLS_AD)
//						pi = pe
//					} else {
//						return empty_str, errors.New("语法错误,参数与 AD 格式不匹配")
//					}
//				case 'M':
//					// DCH AM
//					fi++
//					pe := pi + len(NLS_AD)
//					v := param[pi:pe]
//					if v == NLS_AM {
//						fmt.Println("TODO: " + NLS_AM)
//						result.WriteString(NLS_AM)
//						pi = pe
//					} else {
//						return empty_str, errors.New("语法错误,参数与 AM 格式不匹配")
//					}
//				default:
//					return empty_str, errors.New(dch_fmt_mismatch_err + "A")
//				}
//				// 同上
//			case 'B':
//				fi++
//				followingOneChar := format[fi]
//				switch followingOneChar {
//				case 'C':
//					result.WriteString(NLS_BC)
//				case '.':
//					fe := fi + 4
//					followingChars := format[fi:fe]
//					if ".C." == followingChars {
//						result.WriteString(NLS_BC)
//					}
//					fi = fe
//				default:
//					return empty_str, errors.New(dch_fmt_mismatch_err + "B")
//				}
//				// 只适用于时间类型参数
//			case 'C':
//				fi++
//				followingOneChar := format[fi]
//				switch followingOneChar {
//				case 'C':
//					year := time.Now().Year()
//					result.WriteString(strconv.Itoa((year + 99) / 100))
//				default:
//					return empty_str, errors.New(dch_fmt_mismatch_err + "C")
//				}
//				// 字符串类型参数的 D 周中的日和julia 冲突
//				// 时间类型参数的 D
//			case 'D':
//				fi++
//
//				// DAY 同 DY
//				if frest >= 2 && format[fi:fi+2] == "AY" || frest >= 1 && format[fi] == 'Y' {
//					weekDay := time.Now().Weekday()
//					result.WriteString(NLS_WEEKS[weekDay])
//				} else if frest >= 1 && format[fi] == 'D' {
//					day := time.Now().Day()
//					result.WriteString(strconv.Itoa(day))
//				} else if frest >= 1 && format[fi] == 'L' {
//					// TODO 插入格式 continue
//					//result.WriteString(parseDchByStr("???", NLS_DL))
//
//				} else if frest >= 1 && format[fi] == 'S' {
//					// TODO 插入格式 continue
//					//result.WriteString(parseDchByStr("???", NLS_DS))
//
//				} else {
//					weekDay := time.Now().Weekday()
//					result.WriteString(strconv.Itoa(int(weekDay)))
//				}
//
//			case 'F':
//				fi++
//				followingOneChar := format[fi]
//				switch followingOneChar {
//				case 'X':
//					//keywordGroup = append(keywordGroup, DCH_FX)
//				case 'M':
//					//FIXME
//				default:
//					followingTwoChars := format[fi : fi+3]
//					fi = fi + 3
//					switch followingTwoChars {
//					case "F1":
//						//keywordGroup = append(keywordGroup, DCH_FF1)
//					case "F2":
//						//keywordGroup = append(keywordGroup, DCH_FF2)
//					case "F3":
//						//keywordGroup = append(keywordGroup, DCH_FF3)
//					case "F4":
//						//keywordGroup = append(keywordGroup, DCH_FF4)
//					case "F5":
//						//keywordGroup = append(keywordGroup, DCH_FF5)
//					case "F6":
//						//keywordGroup = append(keywordGroup, DCH_FF6)
//					case "F7":
//						//keywordGroup = append(keywordGroup, DCH_FF7)
//					case "F8":
//						//keywordGroup = append(keywordGroup, DCH_FF8)
//					case "F9":
//						//keywordGroup = append(keywordGroup, DCH_FF9)
//					default:
//						return empty_str, errors.New(dch_fmt_mismatch_err + "F")
//					}
//				}
//			case 'H':
//				fi++
//				followingOneChar := format[fi]
//				switch followingOneChar {
//				case 'H':
//					//keywordGroup = append(keywordGroup, DCH_HH)
//				default:
//					followingThreeChars := format[fi : fi+4]
//					fi = fi + 4
//
//					switch followingThreeChars {
//					case "H24":
//						//keywordGroup = append(keywordGroup, DCH_HH24)
//					case "H12":
//						//keywordGroup = append(keywordGroup, DCH_HH12)
//					default:
//						return empty_str, errors.New(dch_fmt_mismatch_err + "H")
//					}
//				}
//			case 'I':
//				fi++
//				followingOneChar := format[fi]
//				switch followingOneChar {
//				case 'W':
//					//keywordGroup = append(keywordGroup, DCH_IW)
//				case 'Y':
//					followingTwoChars := format[fi : fi+3]
//					if "YY" == followingTwoChars {
//						//keywordGroup = append(keywordGroup, DCH_IYYY)
//						fi = fi + 3
//					}
//
//					followingOneChar = followingTwoChars[0]
//					if 'Y' == followingOneChar {
//						//keywordGroup = append(keywordGroup, DCH_IYY)
//						fi = fi + 2
//					}
//					//keywordGroup = append(keywordGroup, DCH_IY)
//				default:
//					return empty_str, errors.New(dch_fmt_mismatch_err + "I")
//				}
//				// 匹配单个字符
//				//keywordGroup = append(keywordGroup, DCH_I)
//			case 'J':
//				t := time.Now()
//				result.WriteString(strconv.Itoa(toJulian(t.Year(), int(t.Month()), t.Day())))
//			case 'M':
//				t := time.Now()
//				fi++
//				if fi <= flen && format[fi] == 'I' {
//					// DCH MI
//					result.WriteString(strconv.Itoa(t.Minute()))
//				} else if fi <= flen && format[fi] == 'M' {
//					// DCH MM
//					if t.Month() < 10 {
//						result.WriteString("0")
//					}
//					result.WriteString(strconv.Itoa(int(t.Month())))
//				} else if fi <= flen && format[fi] == 'O' {
//					fi++
//					if fi <= flen && format[fi] == 'N' {
//						fe := fi + 2
//						if fi <= flen && format[fi:fe] == "TH" {
//							// DCH MONTH
//							fi = fe
//						} else {
//							// DCH MON
//						}
//						result.WriteString(NLS_MONTHS[t.Month()])
//					} else {
//						return empty_str, errors.New(dch_fmt_mismatch_err + "MO")
//					}
//				} else {
//					return empty_str, errors.New(dch_fmt_mismatch_err + "M")
//				}
//			case 'P':
//				fi++
//				if fi < flen {
//					if 'M' == format[fi] {
//						//keywordGroup = append(keywordGroup, DCH_PM)
//					} else {
//						start := fi
//						fi += 3
//						if fi < flen {
//							followingThreeChars := format[start:fi]
//							if ".M." == followingThreeChars {
//								//keywordGroup = append(keywordGroup, DCH_P_M)
//							} else {
//								return empty_str, errors.New(dch_fmt_mismatch_err + "P")
//							}
//						} else {
//							return empty_str, errors.New(dch_fmt_mismatch_err + "P")
//						}
//					}
//				} else {
//					return empty_str, errors.New(dch_fmt_mismatch_err + "P")
//				}
//			case 'Q':
//				t := time.Now()
//				result.WriteString(strconv.Itoa(int(t.Month()+2) / 3))
//			case 'R':
//				fi++
//				if 'M' == format[fi] {
//					//keywordGroup = append(keywordGroup, DCH_RM)
//				} else {
//					return empty_str, errors.New(dch_fmt_mismatch_err + "R")
//				}
//			case 'S':
//				start := fi
//				fi += 4
//				followingFourChars := format[start:fi]
//				if "SSSS" == followingFourChars {
//					//keywordGroup = append(keywordGroup, DCH_SSSSS)
//				} else if "SSS" == followingFourChars[0:3] {
//					//keywordGroup = append(keywordGroup, DCH_SSSS)
//				} else if "S" == followingFourChars[0:3] {
//					//keywordGroup = append(keywordGroup, DCH_SS)
//				} else {
//					return empty_str, errors.New(dch_fmt_mismatch_err + "S")
//				}
//			case 'T':
//				start := fi
//				fi += 2
//				followingTwoChars := format[start:fi]
//				if "ZH" == followingTwoChars {
//					//keywordGroup = append(keywordGroup, DCH_TZH)
//				} else if "ZM" == followingTwoChars {
//					//keywordGroup = append(keywordGroup, DCH_TZM)
//				} else if 'Z' == followingTwoChars[0] {
//					//keywordGroup = append(keywordGroup, DCH_TZM)
//				} else {
//					return empty_str, errors.New(dch_fmt_mismatch_err + "T")
//				}
//			case 'W':
//				fi++
//				if format[fi] == 'W' {
//					fi++
//
//				} else {
//
//				}
//			case 'Y':
//				fi++
//				start := fi
//				frest--
//				if frest >= 4 {
//					fi += 4
//				} else {
//					fi += frest
//				}
//				followingNChars := format[start:fi]
//				missed := true
//
//				if frest >= 4 && (",YYY" == followingNChars[0:4] || ",yyy" == followingNChars[0:4]) {
//					// DCH Y,YYY
//					pe := pi + 5
//					v := param[pi:pe]
//					result.WriteString(v)
//					missed = false
//				}
//				if missed && frest >= 3 && ("YYY" == followingNChars[0:3] || "yyy" == followingNChars[0:3]) {
//					// DCH YYYY
//					pe := pi + 4
//					v := param[pi:pe]
//					result.WriteString(v)
//					missed = false
//				}
//				if missed && frest >= 2 && ("YY" == followingNChars[0:2] || "yy" == followingNChars[0:4]) {
//					// DCH YYY
//					pe := pi + 3
//					v := param[pi:pe]
//					result.WriteString(v)
//					missed = false
//				}
//				if missed && frest >= 1 && ("Y" == followingNChars[0:1] || "y" == followingNChars[0:4]) {
//					// DCH YY
//					pe := pi + 2
//					v := param[pi:pe]
//					result.WriteString(v)
//					missed = false
//				}
//
//				if missed {
//					// DCH Y
//					pe := pi + 1
//					v := param[pi:pe]
//					result.WriteString(v)
//				}
//
//			default:
//				return empty_str, errors.New(out_keyword_range_err)
//			}
//		} else {
//			return empty_str, errors.New(out_ascii_range_err)
//		}
//	}
//
//	return result.String(), nil
//}
//
//func ToTimestamp() {
//
//}
//
//func ToTimestampTimeZone() {
//
//}
