package base_demo

import (
	"fmt"
	"testing"
)

func TestScience(t *testing.T) {
	fmt.Printf("%f\n", 1.345e9)     //e9 就是小数点向右移动9位
	fmt.Printf("%.10f\n", 12344e-9) //e-9就是小数点向左移动9位，%.10f表示精确到小数点后10位
	fmt.Printf("%.2e\n", 12312312321312123123123.0)
	fmt.Printf("%.2E", 12312312321312123123123.0)
	fmt.Println("%")

	fmt.Printf("%.10e\n", 1.345e9)
}

func TestScience2(t *testing.T) {
	s := fmt.Sprintf("%.6E", 1237823123.0)
	fmt.Println(s)
}
