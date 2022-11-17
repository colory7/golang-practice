package time_demo

import (
	"fmt"
	"testing"
	"time"
)

func TestFormat6(t *testing.T) {
	// now
	now := time.Now()
	// year
	year := now.Year()
	// month
	month := now.Month()
	// day
	day := now.Day()
	// hour
	hour := now.Hour()
	// minute
	minute := now.Minute()
	// second
	second := now.Second()

	// fmt.Printf 输出当前详细时间
	fmt.Printf("直接获取当前详细时间: %d-%d-%d %d:%d:%d\n", year, month, day, hour, minute, second)

	// Format 进行当前时间格式化
	fmt.Println("1.特定格式时间: ", now.Format("2006-01-02 15:04:05"))
	fmt.Println("2.特定格式时间: ", now.Format("2006/01/02 15:04:05"))
	fmt.Println("3.特定格式时间: ", now.Format("2006/01/02"))
	fmt.Println("4.特定格式时间: ", now.Format("15:04:05"))
}
