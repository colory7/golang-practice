package base_demo

import (
	"fmt"
	"math"
	"strings"
	"testing"
	"time"
)

func TestNumber(xx *testing.T) {
	fmt.Println(math.Trunc(222.99999999))
	fmt.Println(math.Floor(223.99999999))

	t := time.Date(2111, 11, 11, 11, 11, 11, 111111111, time.Local)
	fmt.Println(t)
	fmt.Println(time.Now().Nanosecond() / 1e3)
	fmt.Println(t.Add(time.Duration(222333444) * time.Nanosecond))
	fmt.Println(t.Add(time.Duration(222333) * time.Nanosecond))
	fmt.Println(t.Add(time.Duration(222333*1000) * time.Nanosecond))

	aa := strings.SplitN("1,2", ",", 2)
	fmt.Println(aa[0])
	fmt.Println("==")
	fmt.Println(aa[1])

}

func TestPopMicrosecond(t *testing.T) {
	fmt.Println(popMicrosecond(0))
	fmt.Println(popMicrosecond(1))
	fmt.Println(popMicrosecond(12))
	fmt.Println(popMicrosecond(123))
	fmt.Println(popMicrosecond(1234))
	fmt.Println(popMicrosecond(12345))
	fmt.Println(popMicrosecond(123456))
	fmt.Println(popMicrosecond(1234567))

	fmt.Println("====")
	fmt.Println(popMicrosecond(-0))
	fmt.Println(popMicrosecond(-1))
	fmt.Println(popMicrosecond(-12))
	fmt.Println(popMicrosecond(-123))
	fmt.Println(popMicrosecond(-1234))
	fmt.Println(popMicrosecond(-12345))
	fmt.Println(popMicrosecond(-123456))
	fmt.Println(popMicrosecond(-1234567))
}

// 兼容MySQL 中函数 TIMESTAMPDIFF(SECOND,t1,t2)
// 如果秒的差值小于0 且 微秒部分的差值大于0，则最后结果+1
// 如果秒的差值大于0 且 微秒部分的差值小于0，则最后结果-1
func repairTime(d *int64, diff int64) {
	if *d > 0 {
		if diff < 0 {
			*d -= 1
		}
	} else if *d < 0 {
		if diff > 0 {
			*d += 1
		}
	}
}

func popMicrosecond(i int) int {
	negative := false
	if i < 0 {
		i = -i
		negative = true
	}

	if i > 0 && i < 1e1 {
		i = i * 1e5
	} else if i > 1e1 && i < 1e2 {
		i = i * 1e4
	} else if i > 1e2 && i < 1e3 {
		i = i * 1e3
	} else if i > 1e3 && i < 1e4 {
		i = i * 1e2
	} else if i > 1e4 && i < 1e5 {
		i = i * 1e1
	} else {
		for i > 1e6 {
			i /= 10
		}
	}
	if negative {
		i = -i
	}
	return i
}

func TestFormat(t *testing.T) {
	fmt.Println(fmt.Sprintf("%04d\n", 21))
	fmt.Println(fmt.Sprintf("%04s\n", "21"))
	fmt.Println(fmt.Sprintf("%04s\n", "021"))
	fmt.Println(fmt.Sprintf("%01s\n", "3"))
	fmt.Println(fmt.Sprintf("%01s\n", ""))
	fmt.Println(fmt.Sprintf("%01s\n", "0"))
	fmt.Println(fmt.Sprintf("%00s\n", "68"))
	fmt.Println(fmt.Sprintf("%0s\n", "69"))
}
