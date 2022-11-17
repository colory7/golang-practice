package number_demo

import (
	"fmt"
	"strconv"
	"testing"
)

//float保留小数
func Test(t *testing.T) {
	f := 3.2568
	var nf, error = strconv.ParseFloat(fmt.Sprintf("%.6f", f), 64)
	if error != nil {
		panic(error)
	}
	fmt.Println(nf)
}

//float转字符串
func TPFunctionFloat64Format(a float64, fmtStr string) string {
	return fmt.Sprintf(fmtStr, a)
}
