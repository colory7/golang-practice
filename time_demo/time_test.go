package time_demo

import (
	"fmt"
	"regexp"
	"testing"
	"time"
)

const (
	Microsecond = 1
	Millisecond = 1000 * Microsecond
	Second      = 1000 * Millisecond
	Minute      = 60 * Second
	Hour        = 60 * Minute
	Day         = 24 * Hour
	Week        = 7 * Day
)

func TestTime(t *testing.T) {
	local, _ := time.LoadLocation("Asia/Shanghai")
	showTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-11-07 11:34:00", local)
	fmt.Printf("showTime=%v, type=%T,\n", showTime, showTime)
	showTime, _ = time.ParseInLocation("2006-01-02", "2021-11-07", local)
	fmt.Printf("showTime=%v, type=%T,\n", showTime, showTime)
	showTime, _ = time.ParseInLocation("2006-01-02 15:04:05", "2021-11-07", local)
	fmt.Printf("showTime=%v, type=%T,\n", showTime, showTime)
}

func Test2(t *testing.T) {
	now := time.Now()
	timestr := now.Format("2006-01-02 15:04:05") //输出 2020-07-21 10:12:13
	fmt.Println(timestr)
}

func Test3(t *testing.T) {
	eg, err := time.Parse("2006-01-02 15:04:05", "2022-10-19 15:15:43")
	if err != nil {
		panic(err)
	}
	fmt.Println(eg.Hour())

	fmt.Println(eg)
	fmt.Println(eg.Unix())
	fmt.Println(eg.UnixNano())
	fmt.Println(eg.Local().Unix())
}

func Test4(t *testing.T) {
	eg, err := time.Parse("2006-01-02 15:04:05 -0700 MST", "2019-08-29 16:48:21 +0800 CST")
	if err != nil {
		panic(err)
	}
	fmt.Println(eg)
	fmt.Println(eg.Unix())
	fmt.Println(eg.UnixNano())
	fmt.Println(eg.Local().Unix())
}

func Test5(t *testing.T) {
	eg, err := time.ParseInLocation("2006-01-02T15:04", "2020-01-02T15:04", time.Local)
	if err != nil {
		panic(err)
	}
	fmt.Println(eg)
	fmt.Println(eg.Unix())
	fmt.Println(eg.UnixNano())
	fmt.Println(eg.Local().Unix())
}

func Test6(t *testing.T) {
	eg, err := time.Parse("2006-01-02 15:04:05", "15:15:43")
	if err != nil {
		panic(err)
	}
	fmt.Println(eg.Hour())

	fmt.Println(eg)
	fmt.Println(eg.Unix())
	fmt.Println(eg.UnixNano())
	fmt.Println(eg.Local().Unix())
}

func Test7(t *testing.T) {
	eg, err := time.Parse("15:04:05", "335:15:43")
	if err != nil {
		panic(err)
	}
	fmt.Println(eg.Hour())

	fmt.Println(eg)
	fmt.Println(eg.Unix())
	fmt.Println(eg.UnixNano())
	fmt.Println(eg.Local().Unix())
}

func Test8(t *testing.T) {
	eg, err := time.Parse("15:04:05", "10:18:25.9999999999999999999999999999999299999")
	if err != nil {
		panic(err)
	}
	fmt.Println(eg.Hour())

	fmt.Println(eg)
	fmt.Println(eg.Unix())
	fmt.Println(eg.UnixNano())
	fmt.Println(eg.Local().Unix())

}

func Test9(t *testing.T) {
	eg, err := time.Parse("15:04:05", "99991231235959")
	if err != nil {
		panic(err)
	}
	fmt.Println(eg.Month())
	fmt.Println(eg.Day())
	fmt.Println(eg.Hour())

}

func Test10(t *testing.T) {
	var date int64 = 1257894000
	tm := time.Unix(date, 0)
	fmt.Println(tm.Hour())
}

func Test11(t *testing.T) {
	//9999 12 31 23 59 59
	var date int64 = 99991231235959
	tm := time.UnixMilli(date)
	fmt.Println(tm.Hour())
}

func Test12(t *testing.T) {
	eg, err := time.Parse("15:04:05.999", "15:27:36.567")
	if err != nil {
		panic(err)
	}

	fmt.Println(uint(eg.Year()))
	fmt.Println(uint(eg.Month()))
	fmt.Println(uint(eg.Day()))
	fmt.Println(eg.YearDay())

	fmt.Println(eg)

	beforeDay := eg.AddDate(0, 0, -1)
	fmt.Println(beforeDay.Day())
	fmt.Println(eg.UnixNano())
	fmt.Println(eg.Unix())
	fmt.Println(eg.Nanosecond())

}

func Test13(t *testing.T) {
	eg, err := time.Parse("2006-01-02", "2020-01-02")
	if err != nil {
		panic(err)
	}

	fmt.Println(uint(eg.Year()))
	fmt.Println(uint(eg.Month()))
	fmt.Println(uint(eg.Day()))
	fmt.Println(eg.YearDay())

	fmt.Println(eg)

	fmt.Println(eg.Hour())
	fmt.Println(eg.Minute())
	fmt.Println(eg.Second())
	fmt.Println(eg.Unix())
	fmt.Println(eg.UnixMilli())
	fmt.Println(eg.UnixMicro())
	fmt.Println(eg.UnixNano())
}

func Test14(t *testing.T) {
	eg, err := time.Parse("2006-01-02 15:04:05.999999", "2020-01-02 15:27:36.996998")
	if err != nil {
		panic(err)
	}

	fmt.Println(uint(eg.Year()))
	fmt.Println(uint(eg.Month()))
	fmt.Println(uint(eg.Day()))
	fmt.Println(eg.YearDay())

	fmt.Println(eg)

	fmt.Println(eg.Hour())
	fmt.Println(eg.Minute())
	fmt.Println(eg.Second())
	fmt.Println("--")
	fmt.Println(eg.Unix())
	fmt.Println(eg.UnixMilli())
	fmt.Println(eg.UnixMicro())
	fmt.Println(eg.UnixNano())

	fmt.Println("--")
	fmt.Println(eg.Nanosecond())
	fmt.Println(eg.Nanosecond() / 1e3)
	fmt.Println(eg.Nanosecond() / 1e6)
}

func TestDiff(t *testing.T) {
	t1, err := time.Parse("2006-01-02 15:04:05.999999", "2021-03-30 14:00:00.999999")
	if err != nil {
		panic(err)
	}

	t2, err := time.Parse("2006-01-02 15:04:05.999999", "2020-01-30 14:00:00.000001")
	if err != nil {
		panic(err)
	}

	fmt.Println(t2.Nanosecond() - t1.Nanosecond())
	fmt.Println(t2.UnixMicro() - t1.UnixMicro())
	fmt.Println(t2.UnixMilli() - t1.UnixMilli())

	fmt.Println(t2.Sub(t1).Microseconds())
	fmt.Println(t2.Sub(t1).Milliseconds())
}

func TestDiffMonth(t *testing.T) {
	t1, err := time.Parse("2006-01-02 15:04:05.999999", "2023-03-07 13:00:00")
	if err != nil {
		panic(err)
	}

	t2, err := time.Parse("2006-01-02 15:04:05.999999", "2022-07-05 13:00:00")
	if err != nil {
		panic(err)
	}

	d := diffMonth(t1, t2)
	fmt.Println(d)
}

func TestDiffMonth2(t *testing.T) {
	t1, err := time.Parse("2006-01-02 15:04:05.999999", "2022-07-05 13:00:00")
	if err != nil {
		panic(err)
	}

	t2, err := time.Parse("2006-01-02 15:04:05.999999", "2023-03-07 13:00:00")
	if err != nil {
		panic(err)
	}

	d := diffMonth(t1, t2)
	fmt.Println(d)
}

func TestDiffMonth3(t *testing.T) {
	t1, err := time.Parse("2006-01-02 15:04:05.999999", "2022-07-05 13:00:00.999999")
	if err != nil {
		panic(err)
	}

	t2, err := time.Parse("2006-01-02 15:04:05.999999", "2023-03-07 13:00:00.1")
	if err != nil {
		panic(err)
	}

	d := diffMonth(t1, t2)
	fmt.Println(d)
}

func TestDiffMonth4(t *testing.T) {
	t1, err := time.Parse("2006-01-02 15:04:05.999999", "2022-07-05 13:00:00.999999")
	if err != nil {
		panic(err)
	}

	t2, err := time.Parse("2006-01-02 15:04:05.999999", "2023-03-01 13:00:00.1")
	if err != nil {
		panic(err)
	}

	d := diffMonth(t1, t2)
	fmt.Println(d)
}

func TestDiffMonth5(t *testing.T) {
	t1, err := time.Parse("2006-01-02 15:04:05.999999", "2023-03-07 13:00:00.1")
	if err != nil {
		panic(err)
	}

	t2, err := time.Parse("2006-01-02 15:04:05.999999", "2022-07-05 13:00:00.999999")
	if err != nil {
		panic(err)
	}

	d := diffMonth(t1, t2)
	fmt.Println(d)
}

func TestDiffMonth6(t *testing.T) {
	t1, err := time.Parse("2006-01-02 15:04:05.999999", "2023-03-07 13:00:00.1")
	if err != nil {
		panic(err)
	}

	t2, err := time.Parse("2006-01-02 15:04:05.999999", "2022-07-10 13:00:00.999999")
	if err != nil {
		panic(err)
	}

	d := diffMonth(t1, t2)
	fmt.Println(d)
}

func TestDiffMonth7(t *testing.T) {
	t1, err := time.Parse("2006-01-02 15:04:05.999999", "2023-03-07 13:00:00.1")
	if err != nil {
		panic(err)
	}

	t2, err := time.Parse("2006-01-02 15:04:05.999999", "2022-07-05 13:00:00.999999")
	if err != nil {
		panic(err)
	}

	d := diffMonth(t1, t2)
	fmt.Println(d)
}

func TestDiffMonth8(t *testing.T) {
	t1, err := time.Parse("2006-01-02 15:04:05.999999", "2022-03-07 13:00:00")
	if err != nil {
		panic(err)
	}

	t2, err := time.Parse("2006-01-02 15:04:05.999999", "2022-07-05 13:00:00")
	if err != nil {
		panic(err)
	}

	d := diffMonth(t1, t2)
	fmt.Println(d)
}

func TestDiffMonth9(t *testing.T) {
	t1, err := time.Parse("2006-01-02 15:04:05.999999", "2022-07-05 13:00:00")
	if err != nil {
		panic(err)
	}

	t2, err := time.Parse("2006-01-02 15:04:05.999999", "2022-03-07 13:00:00")
	if err != nil {
		panic(err)
	}

	d := diffMonth(t1, t2)
	fmt.Println(d)
}

func TestDiffMonth10(t *testing.T) {
	t1, err := time.Parse("2006-01-02", "9020-01-01")
	if err != nil {
		panic(err)
	}

	t2, err := time.Parse("2006-01-02", "2020-01-01")
	if err != nil {
		panic(err)
	}

	d := diffMonth(t1, t2)
	fmt.Println(d)
}

func TestDiffMonth11(t *testing.T) {
	t1, err := time.Parse("2006-01-02", "2020-01-01")
	if err != nil {
		panic(err)
	}

	t2, err := time.Parse("2006-01-02", "9020-01-01")
	if err != nil {
		panic(err)
	}

	d := diffMonth(t1, t2)
	fmt.Println(d)
}

func TestDiffMicrosecond(t *testing.T) {
	t1, err := time.Parse("2006-01-02", "2020-12-31")
	if err != nil {
		panic(err)
	}

	t2, err := time.Parse("2006-01-02", "2385-01-01")
	if err != nil {
		panic(err)
	}

	d := t2.UnixMicro() - t1.UnixMicro()
	fmt.Println(d)

	// 11486793600000000
}

func TestDiffMicrosecond2(t *testing.T) {
	t1, err := time.Parse("2006-01-02 15:04:05.999999", "2022-04-26 13:00:00.333666999")
	if err != nil {
		panic(err)
	}

	t2, err := time.Parse("2006-01-02 15:04:05.999999", "2022-04-26 14:00:00.333")
	if err != nil {
		panic(err)
	}

	d := t2.UnixMicro() - t1.UnixMicro()
	fmt.Println(t1.UnixMicro())
	fmt.Println(t2.UnixMicro())
	fmt.Println(d)

	// 1 1486 7936 0000 0000
}

func TestDiffMicrosecond3(t *testing.T) {
	t1, err := time.Parse("2006-01-02 15:04:05.999999", "2022-04-26 13:00:00.333666999")
	if err != nil {
		panic(err)
	}

	t2, err := time.Parse("2006-01-02 15:04:05.999999", "2022-04-26 13:00:00.333")
	if err != nil {
		panic(err)
	}

	d := t2.UnixMicro() - t1.UnixMicro()
	fmt.Println(t1.UnixMicro())
	fmt.Println(t2.UnixMicro())
	fmt.Println(d)

	// 1 1486 7936 0000 0000
}

func TestDatediff(t *testing.T) {
	t1, err := time.Parse("2006-01-02 15:04:05.999999", "2022-01-01 13:00:00")
	if err != nil {
		panic(err)
	}

	t2, err := time.Parse("2006-01-02 15:04:05.999999", "2021-12-31 14:00:00")
	if err != nil {
		panic(err)
	}

	d := (t2.UnixMicro() - t1.UnixMicro()) / Day

	fmt.Println(t2.UnixMicro() - t1.UnixMicro())
	fmt.Println(t1.UnixMicro())
	fmt.Println(t2.UnixMicro())
	fmt.Println(d)
}

func TestDatediff2(t *testing.T) {
	t1, err := time.Parse("2006-01-02", "2020-12-31")
	if err != nil {
		panic(err)
	}

	t2, err := time.Parse("2006-01-02", "2385-01-01")
	if err != nil {
		panic(err)
	}

	d := (t2.UnixMicro() - t1.UnixMicro()) / Day

	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t2.UnixMicro() - t1.UnixMicro())
	fmt.Println(t1.UnixMicro())
	fmt.Println(t2.UnixMicro())
	fmt.Println(d)
}

func TestDatediff3(t *testing.T) {
	t1, err := time.Parse("2006-01-02 15:04:05.999999", "2020-12-31 00:00:00")
	if err != nil {
		panic(err)
	}

	t2, err := time.Parse("2006-01-02 15:04:05.999999", "2385-01-01 00:00:00")
	if err != nil {
		panic(err)
	}

	d := (t2.UnixMicro() - t1.UnixMicro()) / Day

	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t2.UnixMicro() - t1.UnixMicro())
	fmt.Println(t1.UnixMicro())
	fmt.Println(t2.UnixMicro())
	fmt.Println(d)
}

func TestTimestamp(t *testing.T) {
	timeUnix := time.Now().Unix()
	fmt.Println(timeUnix)
}

func TestRegex(t *testing.T) {
	hhmmss := "[]:[]:[]"
	r, _ := regexp.Compile(hhmmss)
	b := r.MatchString("15:27:36")
	fmt.Println(b) //结果为true
}

func TestInverseMicrosecond(t *testing.T) {
	s1, m1 := InverseMySQLMicrosecond(98)
	s2, m2 := InverseMySQLMicrosecond(908992)
	s3, m3 := InverseMySQLMicrosecond(908992124)

	fmt.Println(s1, " ", m1)
	fmt.Println(s2, " ", m2)
	fmt.Println(s3, " ", m3)

	fmt.Println("==")
	fmt.Println(98 / 1000_000)
	fmt.Println(98 % 1000_000)

	fmt.Println(908992 / 1000_000)
	fmt.Println(908992 % 1000_000)

	fmt.Println(908992124 / 1000_000)
	fmt.Println(908992124 % 1000_000)

	fmt.Println("==")
	fmt.Println(-98 / 1000_000)
	fmt.Println(-98 % 1000_000)

	fmt.Println(-908992 / 1000_000)
	fmt.Println(-908992 % 1000_000)

	fmt.Println(-908992124 / 1000_000)
	fmt.Println(-908992124 % 1000_000)

}

// 示例: 78123456,在MySQL中的date_sub函数的 _microsecond后缀类型的数据， 78表示78秒，123456表示 毫秒
func InverseMySQLMicrosecond(i int64) (int64, int64) {
	var sec, mic int64
	if i != 0 {
		sec = i / 1e6
		mic = i % 1e6
	} else {
		return 0, 0
	}

	//mysql中的毫秒部分要反转补0
	negative := false
	if mic < 0 {
		mic = -mic
		negative = true
	}

	if mic < 1e1 {
		mic = mic * 1e5
	} else if mic > 1e1 && mic < 1e2 {
		mic = mic * 1e4
	} else if mic > 1e2 && mic < 1e3 {
		mic = mic * 1e3
	} else if mic > 1e3 && mic < 1e4 {
		mic = mic * 1e2
	} else if mic > 1e4 && mic < 1e5 {
		mic = mic * 1e1
	}

	if negative {
		mic = -mic
	}
	return sec, mic
}
