package oracle_demo

import (
	"fmt"
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
	numFmtDesc := parseNumFormat(f)

	str := fmt.Sprintf("%#v\n", numFmtDesc)
	fmt.Println(str)
}

func TestParseNumParam(t *testing.T) {
	num := "-36.25e+97"
	numParamDesc := parseNumParam(num)
	str := fmt.Sprintf("%#v\n", numParamDesc)
	fmt.Println(str)
}

func TestParseNum(t *testing.T) {
	f := "99"
	num := "-36.25"

	result := parseNum(f, num)
	fmt.Println(result)
}

func TestParseNum2(t *testing.T) {
	f := "99999999999999999999"
	num := "-36.25e+3"

	result := parseNum(f, num)
	fmt.Println(result)
}

func TestParseNum2Err(t *testing.T) {
	f := "9"
	num := "-36.25e+3"

	result := parseNum(f, num)
	fmt.Println(result)
}

func TestParseNum3(t *testing.T) {
	f := "99EEEE"
	num := "-36.25e+3"

	result := parseNum(f, num)
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

	result := parseDchByStr(param, format)
	fmt.Println(result)

}

func TestParseDchByStrY_YYY2(t *testing.T) {
	param := "2,013"
	format := "Y,YYY"

	result := parseDchByStr(param, format)
	fmt.Println(result)

	param = "213"
	format = "YYY"

	result = parseDchByStr(param, format)
	fmt.Println(result)

	param = "13"
	format = "YY"

	result = parseDchByStr(param, format)
	fmt.Println(result)

	param = "3"
	format = "Y"

	result = parseDchByStr(param, format)
	fmt.Println(result)

}

func TestParseDchByStrADY_YYY(t *testing.T) {
	param := "公元 2023"
	format := "A.D. yyyy"

	result := parseDchByStr(param, format)
	fmt.Println(result)
}

func TestParseDchByStr2(t *testing.T) {
	param := "公元 2023-10-26 01:30:56"
	format := "A.D. yyyy-MM-dd hh:mi:ss"

	result := parseDchByStr(param, format)
	fmt.Println(result)
}

func TestParseDchByTime(t *testing.T) {

}

func TestCentry(t *testing.T) {

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
	fmt.Println(toJulianDayNumber(2022, 12, 19))
	fmt.Println(toJulianDayNumber(2023, 10, 29))
}
