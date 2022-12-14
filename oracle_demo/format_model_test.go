package oracle_demo

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

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
	f := "99EEEE"
	num := "-36.25e+12"

	result := parseNum(f, num)
	fmt.Println(result)
}
