package number_demo

import (
	"fmt"
	"math"
	"strconv"
	"testing"
	"time"
)

//TestDecimalPlacesPerformance 小数保留性能测试
func TestDecimalPlacesPerformance(t *testing.T) {

	max := 922337203685477580.1234567890
	fmt.Printf("%v\n", max*11)

	pi := math.Pi

	//第1种方式功能测试
	b1 := math.Round(pi*1e15) / 1e15
	fmt.Printf("结果1  :%v\n", b1)

	//第2种方式功能测试
	b2, _ := strconv.ParseFloat(strconv.FormatFloat(pi, 'f', 15, 64), 64)
	fmt.Printf("结果2  :%v\n", b2)

	//第3种方式功能测试
	b3, _ := strconv.ParseFloat(fmt.Sprintf("%.15f", pi), 64)
	fmt.Printf("结果3  :%v\n", b3)

	fmt.Printf("----------------------------------------------------------------------\n")

	//循环次数
	testTimes := int(1e7)

	//第1种方式耗时
	startTime := time.Now().UnixNano() / 1e6
	for i := 0; i < testTimes; i++ {
		_ = math.Round(pi*1e15) / 1e15
	}
	fmt.Printf("耗时1  :%v\n", time.Now().UnixNano()/1e6-startTime)

	//第2种方式耗时
	startTime = time.Now().UnixNano() / 1e6
	for i := 0; i < testTimes; i++ {
		_, _ = strconv.ParseFloat(strconv.FormatFloat(pi, 'f', 15, 64), 64)
	}
	fmt.Printf("耗时2  :%v\n", time.Now().UnixNano()/1e6-startTime)

	//第3种方式耗时
	startTime = time.Now().UnixNano() / 1e6
	for i := 0; i < testTimes; i++ {
		_, _ = strconv.ParseFloat(fmt.Sprintf("%.15f", pi), 64)
	}
	fmt.Printf("耗时3  :%v\n", time.Now().UnixNano()/1e6-startTime)
}

func TestFloat(t *testing.T) {
	fmt.Println(math.Round(231.678*1e0) / 1e0)
	fmt.Println(math.Round(231.678*1e1) / 1e1)
	fmt.Println(math.Round(231.678*1e2) / 1e2)
	fmt.Println(math.Round(231.678*1e15) / 1e15)
}

func TestFloat2(t *testing.T) {
	fmt.Println(math.Abs(-19))                // 取绝对值
	fmt.Println(math.Ceil(3.14))              // 向下取整
	fmt.Println(math.Floor(3.14))             // 向上取整
	fmt.Println(math.Round(3.3478))           // 就近取整
	fmt.Println(math.Round(3.5478*100) / 100) // 保留小数点后2位
	fmt.Println(math.Mod(11, 3))              // 取余数
	fmt.Println(math.Pow(2, 5))               // 计算次方，如：2的5次方
	fmt.Println(math.Pow10(3))                // 计算10次方，如：10的3次方
	fmt.Println(math.Max(1, 2))               // 两个值，取较大的值
	fmt.Println(math.Min(1, 2))               // 两个值，取较小的值
}

func TestFloat3(t *testing.T) {
	res_1 := math.Round(36.98)
	res_2 := math.Round(-100.98)
	res_3 := math.Round(math.Inf(-1))
	res_4 := math.Round(math.NaN())
	res_5 := math.Round(math.Inf(1))

	// Displaying the result
	fmt.Printf("Result 1:%.1f", res_1)
	fmt.Printf("\nResult 2:%.1f", res_2)
	fmt.Printf("\nResult 3:%.1f", res_3)
	fmt.Printf("\nResult 4:%.1f", res_4)
	fmt.Printf("\nResult 5:%.1f", res_5)
}

func TestFloat4(t *testing.T) {

	fmt.Println(math.Round(312.5678))
	fmt.Println(math.Round(989.9678))
	fmt.Println(math.Round(989.9678*100) / 100)

	fmt.Println("=======================")
	fmt.Println(math.Floor(989.9678))
	fmt.Println(math.Floor(989.9678*100) / 100)
}

func TestFloat5(t *testing.T) {
	fmt.Println(truncate(989.9678, -5))
	fmt.Println(truncate(989.9678, -4))
	fmt.Println(truncate(989.9678, -3))
	fmt.Println(truncate(989.9678, -2))
	fmt.Println(truncate(989.9678, -1))

	fmt.Println("==============================")

	fmt.Println(truncate(989.9678, 0))
	fmt.Println(truncate(989.9678, 1))
	fmt.Println(truncate(989.9678, 2))
	fmt.Println(truncate(989.9678, 3))
	fmt.Println(truncate(989.9678, 4))
	fmt.Println(truncate(989.9678, 5))
}

func truncate(num float64, n int) float64 {
	base10 := math.Pow10(n)

	return math.Floor(num*base10) / base10
}
