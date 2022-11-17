package number_demo

import (
	"fmt"
	"github.com/shopspring/decimal"
	"reflect"
	"testing"
)

// https://blog.csdn.net/t949500898/article/details/125168112
func TestIncorrectFloat(t *testing.T) {
	d := 1129.6
	fmt.Println((d * 100)) //输出：112959.99999999999

	var d2 float64 = 1129.6
	fmt.Println((d2 * 100)) //输出：112959.99999999999

	m1 := 8.2
	m2 := 3.8
	fmt.Println(m1 - m2) // 期望是4.4，结果打印出了4.399999999999999
}

func TestIncorrectFloat2(t *testing.T) {
	var v1 = decimal.NewFromFloat(0.1)
	var resp1 float64
	var resp2 string
	resp1, _ = v1.Float64() // note： 注意这边第二个参数并不是ok，而是是否准确
	resp2 = v1.String()

	fmt.Printf("resp1 = %v, type = %v\n", resp1, reflect.TypeOf(resp1).String())
	// output: resp1 = 0.1, type = float64
	fmt.Printf("resp2 = %v, type = %v\n", resp2, reflect.TypeOf(resp2).String())
	// output: resp2 = 0.1, type = string
}
