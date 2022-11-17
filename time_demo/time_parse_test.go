package time_demo

import (
	"fmt"
	"github.com/cockroachdb/apd/v3"
	"github.com/shopspring/decimal"
	"math"
	"math/big"
	"strconv"
	"strings"
	"testing"
)

func TestTimeParse(t *testing.T) {
	f1 := 20210101160700.333666
	t1 := strconv.FormatFloat(f1, 'f', 10, 64)
	fmt.Println(t1)
}

func TestApdParse(t *testing.T) {
	apd.NewBigInt(apd.New(1, 9).Coeff.Int64())
}

func TestDecimal(t *testing.T) {
	f1 := 20210101160700.333666
	t1 := strconv.FormatFloat(f1, 'f', 10, 64)
	fmt.Println(t1)

	fmt.Println("================")
	decimalValue := decimal.NewFromFloat(f1)
	decimalValue = decimalValue.Mul(decimal.NewFromInt(100))

	res, _ := decimalValue.Float64()
	fmt.Println(res)
}

func TestDecimal2(t *testing.T) {
	num, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", 19.90), 64)
	fmt.Println(num)
	fmt.Println(num * 100)

	fmt.Println("================")
	num, _ = strconv.ParseFloat(fmt.Sprintf("%.8f", 19.90), 64)
	fmt.Println(num)

	decimalValue := decimal.NewFromFloat(num)
	decimalValue = decimalValue.Mul(decimal.NewFromInt(100))

	res, _ := decimalValue.Float64()
	fmt.Println(res)
}

func TestDecimal3(t *testing.T) {
	f1 := 20210101160700.333666
	num, _ := strconv.ParseFloat(fmt.Sprintf("%.6f", f1), 64)
	fmt.Println(num)

	num, _ = strconv.ParseFloat(fmt.Sprintf("%v", f1), 64)
	fmt.Println(num)

	pi := 3.14159
	// 按数值本身的格式输出
	variant := fmt.Sprintf("%v %v %v", "月球基地", pi, true)
	fmt.Println(variant)

	// 按数值本身的格式输出
	variant = fmt.Sprintf("%v", f1)
	fmt.Println(variant)

}

func TestParse(t *testing.T) {
	f1 := 20210101160700.333666
	f := strconv.FormatFloat(f1, 'f', 24, 64)
	fmt.Println(f)

	fmt.Println(math.Trunc(f1))
}

func TestParse2(t *testing.T) {
	f := fmt.Sprint(5.03)
	i := fmt.Sprint(5)
	fmt.Println("float:", f, "\nint:", i)
}

func TestApdDecimal(t *testing.T) {
	f := float64(20210101160700.333666)

	d := apd.Decimal{
		0,
		false,
		-6,
		*apd.NewBigInt(int64(f * 1e6)),
	}

	fmt.Println(d.String())

	i, err := d.Int64()
	if err != nil {
		panic(err)
	}

	fmt.Println(i)
}

func TestApd(t *testing.T) {
	//d := apd.Decimal{}
	//
	//DecimalCtx := &apd.Context{
	//	Precision:   20,
	//	Rounding:    apd.RoundHalfUp,
	//	MaxExponent: 2000,
	//	MinExponent: -2000,
	//	// Don't error on invalid operation, return NaN instead.
	//	Traps: apd.DefaultTraps &^ apd.InvalidOperation,
	//}
	//
	//ctx := DecimalCtx.WithPrecision(6)

}

func TestFloatMax(t *testing.T) {
	var float32Max = math.MaxFloat32
	var float64Max = math.MaxFloat64
	fmt.Print("Float32Max = ", float32Max, " Float64Max = ", float64Max)
}

func TestDecimal4(t *testing.T) {
	a := new(big.Float).SetFloat64(20210101160700.333666)
	b := new(big.Float).SetFloat64(1000_1000)
	c := new(big.Float).Mul(a, b)
	println(Trunc(c.Text('g', 6), 6))
}

func Trunc(a string, prec int) string {
	newn := strings.Split(a, ".")

	if prec <= 0 {
		return newn[0]
	}

	if len(newn) < 2 || prec >= len(newn[1]) {
		return a
	}

	return newn[0] + "." + newn[1][:prec]
}

func TestString(t *testing.T) {
	f1 := 20210101160700.333666
	fmt.Println(f1)

	s := fmt.Sprintf("%v", f1)
	fmt.Println(s)

	decimalNum := decimal.NewFromFloat(f1)
	fmt.Println(decimalNum.String())

	decimalNum = decimal.NewFromFloat(f1 * 1000_000)
	fmt.Println(decimalNum.String())
}

func Test31(t *testing.T) {
	d := decimal.NewFromFloat(20210101160700.333666)
	fmt.Println(d)
}
