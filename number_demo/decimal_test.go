package number_demo

import (
	"fmt"
	"github.com/shopspring/decimal"
	"testing"
)

func TestDecimal2(t *testing.T) {
	a := decimal.NewFromFloat(1136.1)
	b := a.Mul(decimal.NewFromInt(100))
	fmt.Println(b) // 正确输出 113610

	c := decimal.NewFromFloat(1.7)
	fmt.Println(a.Sub(c)) // 正确输出 1134.4
	fmt.Println(b.Sub(c)) // 正确输出 113608.3
}

func TestDecimal3(t *testing.T) {
	a := decimal.NewFromFloat(1.52)
	b := decimal.NewFromFloat(0.02)

	// 加减乘除运算
	c := a.Add(b) // 1.52 + 0.02 = 1.54
	d := a.Sub(b) // 1.52 - 0.02 = 1.5
	e := a.Mul(b) // 1.52 * 0.02 = 0.0304
	f := a.Div(b) // 1.52 / 0.02 = 76
	fmt.Println(a, b, c, d, e, f)

	// 对于保留小数的处理
	pi := decimal.NewFromFloat(3.1415926535897932384626)
	pi1 := pi.Round(3)    // 对pi值四舍五入保留3位小数
	fmt.Println(pi1)      // 3.142
	pi2 := pi.Truncate(3) // 对pi值保留3位小数之后直接舍弃
	fmt.Println(pi2)      // 3.141
}
