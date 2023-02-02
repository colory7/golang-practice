package number_demo

import (
	"fmt"
	amountdecimal "github.com/jishulangcom/go-amount-decimal"
	"math"
	"strconv"
	"testing"
)

func TestLosePrecision(t *testing.T) {
	f1 := 1129.6
	fmt.Println((f1 * 100)) //输出：112959.99999999999

	var f2 float64 = 1129.6
	fmt.Println((f2 * 100)) //输出：112959.99999999999

	m1 := 8.2
	m2 := 3.8
	fmt.Println(m1 - m2) // 期望是4.4，结果打印出了4.399999999999999
}

func TestLosePrecision2(t *testing.T) {
	f1 := 1129.6
	var v1, _ = amountdecimal.New(f1).Mul(100).ToString(nil)
	fmt.Println(v1) // 输出：112960.0

	var f2 float64 = 1129.6
	var v2, _ = amountdecimal.New(f2).Mul(100).ToString(nil)
	fmt.Println(v2) // 输出：112960.0

	m1 := 8.2
	m2 := 3.8
	var v3, _ = amountdecimal.New(m1).Sub(m2).ToString(nil)
	fmt.Println(v3) // 输出：4.4
}

func TestLosePrecision3(t *testing.T) {
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)

	f := 22342353676555895867765676574.54
	fmt.Println(f)
	fmt.Println(f + 1)
	fmt.Println(strconv.FormatFloat(f, 'f', -1, 64))
	fmt.Println(strconv.FormatFloat(f+1, 'f', -1, 64))

}
