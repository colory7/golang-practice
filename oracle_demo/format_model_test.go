package oracle_demo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"strconv"
	"strings"
	"testing"
	"time"
)

//NumberFormatModel
//DatetimeFormatModel

func TestMatchNumberFormatModel(t *testing.T) {
	// 1.找到模式,最大匹配
	// 2.根据模式替换字符串

	//param:="111"
	format := "999"
	parseNumFormat(format)
}

func TestAscii(t *testing.T) {
	theme := "狙击 start"
	for i := 0; i < len(theme); i++ {
		fmt.Printf("ascii:%c %d\n", theme[i], theme[i])
	}
}

func TestAscii2(t *testing.T) {
	theme := "狙击 start"
	for i := 0; i < len(theme); i++ {
		f := theme[i]
		if f >= 32 && f <= 127 {
			fmt.Println((string)(f))
		}
	}
}

func TestE(t *testing.T) {
	fmt.Printf("%f\n", 1.345e2*10) //e9 就是小数点向右移动9位

	fmt.Printf("%.10f\n", 12344e-9) //e-9就是小数点向左移动9位，%.10f表示精确到小数点后10位

	fmt.Printf("%.2e\n", 12312312321312123123123.0)
	fmt.Printf("%.2E", 12312312321312123123123.0)
	fmt.Println("%")
}

func TestE2(t *testing.T) {
	preDec := ""
	postDec := "10"
	eNum := 1.345e2
	v := fmt.Sprintf("%"+preDec+"."+postDec+"f", eNum)
	fmt.Println(v)

}

func TestStrSearch(t *testing.T) {
	fmt.Println(strings.Index("widuu", "i")) //1
	fmt.Println(strings.Index("widuu", "u")) //3
}

func TestV(t *testing.T) {
	//f := "99V999"
	//preV := "2"
	postV := "3"

	param := "12"
	paramNum, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}
	postVNum, _ := strconv.Atoi(postV)
	if err != nil {
		panic(err)
	}

	result := paramNum * int(math.Pow10(postVNum))

	fmt.Println(result)

}

func TestToString(t *testing.T) {
	var numParam NumParamDesc
	numParam.sign = plus
	numParam.preDec = "36"
	numParam.postDec = "25"
	numParam.eSign = empty
	numParam.eExponent = 12
	fmt.Println(numParam.string())
}

func TestParseNumFmt(t *testing.T) {
	f := "99EEEE"
	numFmtDesc, err := parseNumFormat(f)
	assert.NoError(t, err)

	str := fmt.Sprintf("%#v\n", numFmtDesc)
	fmt.Println(str)
}

func TestParseNumParam(t *testing.T) {
	num := "-36.25e+97"
	numParamDesc, err := parseNumParam(num)
	assert.NoError(t, err)
	str := fmt.Sprintf("%#v\n", numParamDesc)
	fmt.Println(str)
}

func TestParseNum(t *testing.T) {
	f := "99"
	num := "-36.25"

	result, err := toNumber(f, num)
	assert.NoError(t, err)
	assert.Equal(t, "-36", result)
	fmt.Println(result)
}

func TestParseNum2(t *testing.T) {
	f := "99999999999999999999"
	num := "-36.25e+3"

	result, err := toNumber(f, num)
	assert.NoError(t, err)
	fmt.Println(result)
}

func TestParseNum2Err(t *testing.T) {
	f := "9"
	num := "-36.25e+3"

	result, err := toNumber(f, num)
	assert.NoError(t, err)
	fmt.Println(result)
}

func TestParseNum3(t *testing.T) {
	f := "99EEEE"
	num := "-36.25e+3"

	result, err := toNumber(f, num)
	assert.NoError(t, err)
	fmt.Println(result)
}

func TestParseDchByStrAD(t *testing.T) {
	param := "公元"
	format := "AD"

	//fmt.Println(format[2:4])

	parseDchByStr(param, format)
}

func TestParseDchByStrA_D_(t *testing.T) {
	param := "公元"
	format := "A.D."

	//fmt.Println(format[2:4])

	parseDchByStr(param, format)
}

func TestParseDchByStrAM(t *testing.T) {
	param := "上午"
	format := "AM"

	//fmt.Println(format[2:4])

	parseDchByStr(param, format)
}

func TestParseDchByStrA_M_(t *testing.T) {
	param := "上午"
	format := "A.M."

	//fmt.Println(format[2:4])

	parseDchByStr(param, format)
}

func TestParseDchByStrY_YYY(t *testing.T) {
	param := "2013"
	format := "yyyy"

	result, err := parseDchByStr(param, format)
	assert.NoError(t, err)
	fmt.Println(result)

}

func TestParseDchByStrY_YYY2(t *testing.T) {
	param := "2,013"
	format := "Y,YYY"

	result, err := parseDchByStr(param, format)
	assert.NoError(t, err)
	fmt.Println(result)

	param = "213"
	format = "YYY"

	result, err = parseDchByStr(param, format)
	assert.NoError(t, err)
	fmt.Println(result)

	param = "13"
	format = "YY"

	result, err = parseDchByStr(param, format)
	assert.NoError(t, err)
	fmt.Println(result)

	param = "3"
	format = "Y"

	result, err = parseDchByStr(param, format)
	assert.NoError(t, err)
	fmt.Println(result)

}

func TestParseDchByStrADY_YYY(t *testing.T) {
	param := "公元 2023"
	format := "A.D. yyyy"

	result, err := parseDchByStr(param, format)
	assert.NoError(t, err)
	fmt.Println(result)
}

func TestParseDchByStr2(t *testing.T) {
	param := "公元 2023-10-26 01:30:56"
	format := "A.D. yyyy-MM-dd hh:mi:ss"

	result, err := parseDchByStr(param, format)
	assert.NoError(t, err)
	fmt.Println(result)
}

func TestCentury(t *testing.T) {

	ti := time.Now()
	fmt.Println(ti.Year())
	fmt.Println((ti.Year()-1)/100 + 1)
	fmt.Println((1900-1)/100 + 1)
	fmt.Println((1901-1)/100 + 1)
	fmt.Println((2000-1)/100 + 1)

	fmt.Println("====")
	fmt.Println((ti.Year() + 99) / 100)
	fmt.Println((1900 + 99) / 100)
	fmt.Println((1901 + 99) / 100)
	fmt.Println((2000 + 99) / 100)
	fmt.Println((2001 + 99) / 100)
}
func TestChar(t *testing.T) {
	param := "公元 2023"
	fmt.Println(param[0:3])
	fmt.Println(param[0:6])
	fmt.Println(param[3:6])
	fmt.Println(param[0:8])
	fmt.Println(param[0:10])
	fmt.Println(param[0:11])
	fmt.Println(param[3:11])
	fmt.Println("====")
	fmt.Println(param[4:11])
	fmt.Println(param[5:11])
	fmt.Println(param[6:11])
	fmt.Println(param[7:11])
	fmt.Println(param[8:11])
}

func TestJulian(t *testing.T) {
	fmt.Println(ToJulian(2022, 12, 19))
	fmt.Println(ToJulian(2023, 10, 29))
}

func TestToRoman(t *testing.T) {
	b := ToRoman(4278)
	fmt.Println(b.String())
}

func TestParseDchByTimeEmpty(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")

	format := ""
	actual, err := ParseDchFmtByTime(ti, format)
	assert.NoError(t, err)
	expected := ""

	if actual != expected {
		t.Errorf("actual %q expected %q", actual, expected)
	}
}

func TestParseDchByTimeSkipChars(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	ti, _ := time.Parse(layout, "2021-01-10 15:01:02")

	format := "./- ,,,:;,,"
	actual, err := ParseDchFmtByTime(ti, format)
	assert.NoError(t, err)
	expected := format + "fs"

	// 是否符合我们的预期
	assert.Equal(t, expected, actual)

}

func TestParseDchByTimeA(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")

	format := "AM"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "A.M."
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "AD"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "A.D."
	fmt.Println(ParseDchFmtByTime(ti, format))

}

func TestParseDchByTimeAErr(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")

	format := "A"
	actual, err := ParseDchFmtByTime(ti, format)
	expected := empty_str
	assert.Equal(t, expected, actual)
	assert.Error(t, err)
}

func TestParseDchByTimeAErr2(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")

	format := "A."
	actual, err := ParseDchFmtByTime(ti, format)
	expected := empty_str
	assert.Equal(t, expected, actual)
	assert.Error(t, err)
}

func TestParseDchByTimeAErr3(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")

	format := "AX"
	actual, err := ParseDchFmtByTime(ti, format)
	expected := empty_str
	assert.Equal(t, expected, actual)
	assert.Error(t, err)
}

func TestParseDchByTimeAErr4(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")

	format := "A.D"

	actual, err := ParseDchFmtByTime(ti, format)
	expected := empty_str
	assert.Equal(t, expected, actual)
	assert.Error(t, err)
}

func TestParseDchByTimeB(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")

	format := "BC"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "B.C."
	fmt.Println(ParseDchFmtByTime(ti, format))
}

func TestParseDchByTimeAB(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")

	format := "ADBC"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "BCAD"
	fmt.Println(ParseDchFmtByTime(ti, format))
}

func TestParseDchByTimeC(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")

	format := "CC"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "C"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "BCADCC"
	fmt.Println(ParseDchFmtByTime(ti, format))
}

func TestParseDchByTimeD(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")

	var format = ""
	format = "DAY"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "DD"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "DDD"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = NLS_DL
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "DL"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = NLS_DS
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "DS"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "DY"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "D"
	fmt.Println(ParseDchFmtByTime(ti, format))
}
func TestParseDchByTimeF(t *testing.T) {
	layout := "2006-01-02 15:04:05.000000000"
	ti, _ := time.Parse(layout, "2031-01-10 15:01:02.789321456")

	var format = ""
	format = "FXFM"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "FF1"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "FF2"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "FF3"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "FF4"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "FF5"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "FF6"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "FF7"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "FF8"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "FF9"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "FF1-FF2-FF9-FF8"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "FF"
	actual, err := ParseDchFmtByTime(ti, format)
	expected := empty_str
	assert.Equal(t, expected, actual)
	assert.Error(t, err)

	format = "F"
	actual, err = ParseDchFmtByTime(ti, format)
	expected = empty_str
	assert.Equal(t, expected, actual)
	assert.Error(t, err)

	format = "F0"
	actual, err = ParseDchFmtByTime(ti, format)
	expected = empty_str
	assert.Equal(t, expected, actual)
	assert.Error(t, err)

	format = "FFD"
	actual, err = ParseDchFmtByTime(ti, format)
	expected = empty_str
	assert.Equal(t, expected, actual)
	assert.Error(t, err)
}

func TestParseDchByTimeH(t *testing.T) {
	layout := "2006-01-02 15:04:05.000000000"
	ti, _ := time.Parse(layout, "2031-01-10 15:01:02.789321456")

	var format = ""
	format = "HH"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "H24"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "H12"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "H24-H12"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "H2"
	actual, err := ParseDchFmtByTime(ti, format)
	expected := empty_str
	assert.Equal(t, expected, actual)
	assert.Error(t, err)

	format = "H2D"
	actual, err = ParseDchFmtByTime(ti, format)
	expected = empty_str
	assert.Equal(t, expected, actual)
	assert.Error(t, err)

	format = "H1"
	actual, err = ParseDchFmtByTime(ti, format)
	expected = empty_str
	assert.Equal(t, expected, actual)
	assert.Error(t, err)

	format = "H13"
	actual, err = ParseDchFmtByTime(ti, format)
	expected = empty_str
	assert.Equal(t, expected, actual)
	assert.Error(t, err)
}

func TestParseDchByTimeH2(t *testing.T) {
	layout := "2006-01-02 15:04:05.000000000"
	ti, _ := time.Parse(layout, "2031-01-10 02:01:02.789321456")

	var format = ""
	format = "HH"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "H24"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "H12"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "H24-H12"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "H2"
	actual, err := ParseDchFmtByTime(ti, format)
	expected := empty_str
	assert.Equal(t, expected, actual)
	assert.Error(t, err)

	format = "H2D"
	actual, err = ParseDchFmtByTime(ti, format)
	expected = empty_str
	assert.Equal(t, expected, actual)
	assert.Error(t, err)

	format = "H1"
	actual, err = ParseDchFmtByTime(ti, format)
	expected = empty_str
	assert.Equal(t, expected, actual)
	assert.Error(t, err)

	format = "H13"
	actual, err = ParseDchFmtByTime(ti, format)
	expected = empty_str
	assert.Equal(t, expected, actual)
	assert.Error(t, err)
}

func TestParseDchByTimeI(t *testing.T) {
	layout := "2006-01-02 15:04:05.000000000"
	ti, _ := time.Parse(layout, "2031-01-10 02:01:02.789321456")
	format := "I"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "IYYY"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "IYY"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "IY"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "IW"
	fmt.Println(ParseDchFmtByTime(ti, format))
}

func TestParseDchByTimeJ(t *testing.T) {
	layout := "2006-01-02 15:04:05.000000000"
	ti, _ := time.Parse(layout, "2031-01-10 02:01:02.789321456")
	format := "J"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "JD"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "DJD"
	fmt.Println(ParseDchFmtByTime(ti, format))
}

func TestParseDchByTimeM(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")

	format := "MI"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "MM"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "MON"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "MONTH"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "MONTHMON"
	fmt.Println(ParseDchFmtByTime(ti, format))

}

func TestParseDchByTimeM2(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")

	format := "MMYYYY"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "MM\"健康\"YYYY"
	fmt.Println(ParseDchFmtByTime(ti, format))
}

func TestParseDchByTimePM(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")

	format := "PM"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "P.M."
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "P.M.-A.M....PM" //FIXME 少1个点没有输出
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "P.M"
	fmt.Println(ParseDchFmtByTime(ti, format))
	actual, err := ParseDchFmtByTime(ti, format)
	expected := empty_str
	assert.Equal(t, expected, actual)
	assert.Error(t, err)

	format = "P.MX"
	fmt.Println(ParseDchFmtByTime(ti, format))
	actual, err = ParseDchFmtByTime(ti, format)
	expected = empty_str
	assert.Equal(t, expected, actual)
	assert.Error(t, err)
}

func TestParseDchByTimeQ(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	ti, _ := time.Parse(layout, "2031-04-01 15:01:02")

	format := "Q"
	fmt.Println(ParseDchFmtByTime(ti, format))
}

func TestParseDchByTimeR(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	ti, _ := time.Parse(layout, "2031-04-01 15:01:02")

	format := "RR"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "RM"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "RRRR"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "R"
	actual, err := ParseDchFmtByTime(ti, format)
	expected := empty_str
	assert.Equal(t, expected, actual)
	assert.Error(t, err)

	format = "RT"
	actual, err = ParseDchFmtByTime(ti, format)
	expected = empty_str
	assert.Equal(t, expected, actual)
	assert.Error(t, err)

	format = "RRR"
	actual, err = ParseDchFmtByTime(ti, format)
	expected = empty_str
	assert.Equal(t, expected, actual)
	assert.Error(t, err)

	format = "RRR."
	actual, err = ParseDchFmtByTime(ti, format)
	expected = empty_str
	assert.Equal(t, expected, actual)
	assert.Error(t, err)

	format = "RRR-"
	actual, err = ParseDchFmtByTime(ti, format)
	expected = empty_str
	assert.Equal(t, expected, actual)
	assert.Error(t, err)

	format = "RRRU"
	actual, err = ParseDchFmtByTime(ti, format)
	expected = empty_str
	assert.Equal(t, expected, actual)
	assert.Error(t, err)

}

func TestParseDchByTimeS(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	ti, _ := time.Parse(layout, "2031-02-21 15:01:02")

	format := "SSSSS"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "SS"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "SP"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "SYEAR"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "SYYYY"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "SSSSSSSSS"
	fmt.Println(ParseDchFmtByTime(ti, format))

}

func TestParseDchByTimeT(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	ti, _ := time.Parse(layout, "2031-02-21 15:01:02")

	format := "TS"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "TZD"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "TZH"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "TZM"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "TZR"
	fmt.Println(ParseDchFmtByTime(ti, format))
}

func TestParseDchByTimeW(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	ti, _ := time.Parse(layout, "2031-02-21 15:01:02")

	format := "W"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "WW"
	fmt.Println(ParseDchFmtByTime(ti, format))
}

func TestParseDchByTimeX(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")

	format := "X"
	fmt.Println(ParseDchFmtByTime(ti, format))
}

func TestParseDchByTimeY(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")

	format := "Y"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "YY"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "YYY"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "YYYY"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "Y,YYY"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "YEAR"
	fmt.Println(ParseDchFmtByTime(ti, format))

	format = "YEARYYYY"
	fmt.Println(ParseDchFmtByTime(ti, format))

}

func TestParseDchByTimeYYYYMMDD(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	ti, _ := time.Parse(layout, "2031-01-10 15:01:02")

	format := "YYYY----MM--DD"
	fmt.Println(ParseDchFmtByTime(ti, format))
}
