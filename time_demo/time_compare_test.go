package time_demo

import (
	"fmt"
	"testing"
	"time"
)

func TestCompare(t *testing.T) {
	time1 := "2015-03-20 08:50:29"
	time2 := "2015-03-21 09:04:25"
	//先把时间字符串格式化成相同的时间类型
	t1, _ := time.Parse("2006-01-02 15:04:05", time1)
	t2, _ := time.Parse("2006-01-02 15:04:05", time2)
	fmt.Println(t1.After(t2))
	fmt.Println(t1.Before(t2))
}

func TestCompare2(t *testing.T) {
	time1 := "2015-03-20 08:50:29"
	time2 := "2015-03-20 08:50:29"
	//先把时间字符串格式化成相同的时间类型
	t1, _ := time.Parse("2006-01-02 15:04:05", time1)
	t2, _ := time.Parse("2006-01-02 15:04:05", time2)
	fmt.Println(t1.After(t2))
	fmt.Println(t1.Before(t2))
	fmt.Println(t1.Equal(t2))
}
