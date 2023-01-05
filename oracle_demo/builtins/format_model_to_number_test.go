package builtins

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestToNumber2(t *testing.T) {
	numParam := "34,50"
	format := "999,99"

	numResult, err := ToNumber(numParam, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(numResult)
}

func TestToNumber3(t *testing.T) {
	numParam := "12,4,548"
	format := "99G9G99B9"

	numResult, err := ToNumber(numParam, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(numResult)
}

func TestToNumber4(t *testing.T) {
	numParam := "12,4,548.689"
	format := "99G9G99B9.999"

	numResult, err := ToNumber(numParam, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(strconv.FormatFloat(numResult, 'f', -1, 64))
	fmt.Println(numResult)
}

func TestToNumberNegative(t *testing.T) {
	numParam := "12,4,548.689"
	format := "9G9G99B9.999"

	numResult, err := ToNumber(numParam, format)
	assert.Error(t, err)
	fmt.Println(numResult)
}

func TestToNumber6(t *testing.T) {
	numParam := "12,4,548.689"
	format := "999G9G99B9D999"

	numResult, err := ToNumber(numParam, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(numResult)
}

func TestToNumber7(t *testing.T) {
	numParam := "12,4,548.689"
	format := "RN"

	numResult, err := ToNumber(numParam, format)
	// Oracle to_number函数不支持
	assert.Error(t, err)
	fmt.Println(numResult)
}

func TestToNumberXPositive(t *testing.T) {
	numParam := "12,4,548"
	format := "XXXXXX"

	numResult, err := ToNumber(numParam, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(strconv.FormatFloat(numResult, 'f', -1, 64))
	//fmt.Printf("%f \n")
	fmt.Println(numResult)
}

func TestToNumberXPositive2(t *testing.T) {
	numParam := "1ABC3EDEF"
	format := "XXXXXXxXXXX"

	numResult, err := ToNumber(numParam, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(strconv.FormatFloat(numResult, 'f', -1, 64))
	//fmt.Printf("%f \n")
	fmt.Println(numResult)
}

func TestToNumberXPositive3(t *testing.T) {
	numParam := "11"
	format := "XXXXXX"

	numResult, err := ToNumber(numParam, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(numResult)
}

func TestToNumberXPositive4(t *testing.T) {
	numParam := "124548"
	format := "XXXXXX"

	numResult, err := ToNumber(numParam, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(numResult)
}

func TestToNumberXPositive5(t *testing.T) {
	numParam := ",48"
	format := "XXXXXX"

	numResult, err := ToNumber(numParam, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(numResult)
}

func TestToNumberXPositive6(t *testing.T) {
	numParam := "48,"
	format := "XXXXXX"

	numResult, err := ToNumber(numParam, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(numResult)
}

func TestToNumberXNegative(t *testing.T) {
	numParam := "12,4,548.689"
	format := "XXXXXXXXXXXXXXXX"

	numResult, err := ToNumber(numParam, format)
	// Oracle to_number函数不支持
	assert.Error(t, err)
	fmt.Println(numResult)
}

func TestToNumberXNegative2(t *testing.T) {
	numParam := "12,4,"
	format := "XXXXXXXXXXXXXXXX"

	numResult, err := ToNumber(numParam, format)
	// Oracle to_number函数不支持
	assert.Error(t, err)
	fmt.Println(numResult)
}

func TestToNumberXNegative3(t *testing.T) {
	numParam := "124548"
	format := "XXXXXXA"

	numResult, err := ToNumber(numParam, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(numResult)
}

func TestToNumberXNegative4(t *testing.T) {
	numParam := "1a"
	format := "x"

	numResult, err := ToNumber(numParam, format)
	assert.Error(t, err)
	fmt.Println(numResult)
}

func TestToNumberTM(t *testing.T) {
	numParam := "12,4,548.689"
	format := "TM"

	numResult, err := ToNumber(numParam, format)
	// Oracle to_number函数不支持
	assert.Error(t, err)
	fmt.Println(numResult)
}

func TestToNumberTMNegative(t *testing.T) {
	numParam := "12,4,548.689"
	format := "TM0"

	numResult, err := ToNumber(numParam, format)
	assert.Error(t, err)
	fmt.Println(numResult)
}

func TestToNumberGD(t *testing.T) {
	numParam := "12,4,548.327"
	format := "99G9G999.999"

	numResult, err := ToNumber(numParam, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(numResult)

	aa := 124548.327
	fmt.Println(aa)
}

func TestToNumberNegative2(t *testing.T) {
	numParam := "12,4,548.689"
	format := "aa"

	numResult, err := ToNumber(numParam, format)
	assert.Error(t, err)
	fmt.Println(numResult)
}

func TestToNumberNegative3(t *testing.T) {
	numParam := "12,4,548.689"
	format := "90"

	numResult, err := ToNumber(numParam, format)
	assert.Error(t, err)
	fmt.Println(numResult)
}

func TestToNumberDecPositive(t *testing.T) {
	numParam := ".48"
	format := ".99"

	numResult, err := ToNumber(numParam, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(numResult)
}

func TestToNumberDecPositive2(t *testing.T) {
	numParam := "48."
	format := "99"

	numResult, err := ToNumber(numParam, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(numResult)
}

func TestToNumberCommaPositive(t *testing.T) {
	numParam := "48,"
	format := "99,"

	numResult, err := ToNumber(numParam, format)
	// 与Oracle不同，数值参数不允许逗号结尾
	assert.Error(t, err)
	fmt.Println(numResult)
}

func TestToNumberCommaNegative(t *testing.T) {
	numParam := ",48"
	format := ",99"

	numResult, err := ToNumber(numParam, format)
	assert.Error(t, err)
	fmt.Println(numResult)
}
