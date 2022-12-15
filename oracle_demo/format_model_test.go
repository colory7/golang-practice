package oracle_demo

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

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
