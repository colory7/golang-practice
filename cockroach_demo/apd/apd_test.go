package apd

import (
	"fmt"
	"github.com/cockroachdb/apd/v3"
	"math"
	"testing"
)

func TestApd(t *testing.T) {
	d := apd.New(27, 0)
	three := apd.New(3, 0)
	c := apd.BaseContext.WithPrecision(5)
	for {
		res, err := c.Quo(d, d, three)
		fmt.Printf("d: %7s, inexact: %5v, err: %v\n", d, res.Inexact(), err)
		if err != nil {
			return
		}
		if res.Inexact() {
			return
		}
	}

}

func TestDecimal2(t *testing.T) {
	d := &apd.Decimal{}
	d.SetFloat64(220.36002030)

	fmt.Println(d)
	fmt.Println(d.Coeff)
	fmt.Println(d.Form)
	fmt.Println(d.Negative)
	fmt.Println(d.Exponent)
}

func TestDecimal3(t *testing.T) {

	n := 229.36002030
	d := &apd.Decimal{}
	d.SetFloat64(n)

	fmt.Println(d)
	fmt.Println(d.Coeff)
	fmt.Println(d.Form)
	fmt.Println(d.Negative)
	fmt.Println(d.Exponent)

	fmt.Println("=======")

	apd.NewBigInt(d.Coeff.Int64())

	fmt.Println(d)
	fmt.Println(d.Coeff)
	fmt.Println(d.Form)
	fmt.Println(d.Negative)
	fmt.Println(d.Exponent)

	fmt.Println("=======")
	e := int32(-3)

	d2 := truncateDecimal(*d, e)

	fmt.Println("====d2===")
	if d2 != nil {
		fmt.Println(d2)
		fmt.Println(d2.Coeff)
		fmt.Println(d2.Form)
		fmt.Println(d2.Negative)
		fmt.Println(d2.Exponent)
		fmt.Println(d2.String())
	}
}

func truncateDecimal(d apd.Decimal, e int32) *apd.Decimal {
	var c *apd.BigInt

	if e > 0 {
		if -d.Exponent > e {
			mv := -d.Exponent - e
			base10 := apd.NewBigInt(int64(math.Pow10(int(mv))))
			c = d.Coeff.Quo(&d.Coeff, base10)
			return &apd.Decimal{
				d.Form,
				d.Negative,
				-e,
				*c,
			}
		} else {
			return &d
		}
	} else if e == 0 {
		mv := -d.Exponent + e
		base10 := apd.NewBigInt(int64(math.Pow10(int(mv))))
		c = d.Coeff.Quo(&d.Coeff, base10)
		return &apd.Decimal{
			d.Form,
			d.Negative,
			-e,
			*c,
		}
	} else if e < 0 {
		if -e >= (int32(len(d.Coeff.String())) + d.Exponent) {
			return apd.New(0, 0)
		} else {
			mv := -d.Exponent - e
			c = d.Coeff.Quo(&d.Coeff, apd.NewBigInt(int64(math.Pow10(int(mv)))))
			c.Mul(c, apd.NewBigInt(int64(math.Pow10(int(-e)))))

			return &apd.Decimal{
				d.Form,
				d.Negative,
				0,
				*c,
			}
		}
	}

	return nil
}

func TestNum(t *testing.T) {
	fmt.Printf("%f\n", 1.345e9) //e9 就是小数点向右移动9位

	fmt.Printf("%.10f\n", 12344e-9) //e-9就是小数点向左移动9位，%.10f表示精确到小数点后10位

	fmt.Printf("%.2e\n", 12312312321312123123123.0)
	fmt.Printf("%.2E", 12312312321312123123123.0)
	fmt.Println("%")
}

func TestApdError(t *testing.T) {
	n1 := 229.36002030
	d1 := &apd.Decimal{}
	d1.SetFloat64(n1)

	n2 := 20210101160700.333666
	d2 := &apd.Decimal{}
	d2.SetFloat64(n2)

	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d2.String())
	fmt.Println(20210101160700.333666)
}

func TestApdError2(t *testing.T) {
	n1 := 229.36002030
	d1 := &apd.Decimal{}
	d1.SetFloat64(n1)

	n2 := 20210101160700.333666
	d2 := &apd.Decimal{}
	d2.SetFloat64(n2)

	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d2.String())
	fmt.Println(20210101160700.333666)
}

func TestApdCorrect(t *testing.T) {
	d1 := apd.Decimal{}
	d1.Coeff.SetString("20210101160700333666", 10)

	fmt.Println(d1)
	fmt.Println(d1.Text('f'))
	fmt.Println(d1.Text('e'))
	fmt.Println(d1.String())
	fmt.Println("===============================")

	d2 := apd.Decimal{}
	d2.SetString("20210101160700333666")

	fmt.Println(d2)
	fmt.Println(d2.Text('f'))
	fmt.Println(d2.Text('e'))
	fmt.Println(d2.String())
	fmt.Println("===============================")

	d3 := apd.Decimal{
		Exponent: 1,
	}
	d3.Coeff.SetString("20210101160700333666", 10)

	fmt.Println(d3)
	fmt.Println(d3.Text('f'))
	fmt.Println(d3.Text('e'))
	fmt.Println(d3.Text('G'))
	fmt.Println(d3.String())
	i, err := d3.Int64()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
	fmt.Println("===============================")

}
