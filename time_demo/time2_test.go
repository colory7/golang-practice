package time_demo

import (
	"fmt"
	"math"
	"strconv"
	"testing"
	"time"
)

func TestSecondFloat(xx *testing.T) {
	t := time.Now()
	fmt.Println(t.UnixMicro())
	fmt.Println(t.Nanosecond())
	// 毫秒 3位
	fmt.Println(t.Nanosecond() / 1e6)
	fmt.Println(t.Nanosecond() / 1e3 / 1e3)

	// 微秒 6位
	fmt.Println(t.Nanosecond() / 1e3)
	fmt.Println(t.Nanosecond() % 1e3)

	fmt.Println(math.Floor(2.666555))
	fmt.Println(2.666555 - math.Floor(2.666555))

	fmt.Println(2.666555*1e6 - math.Floor(2.666555)*1e6)

	f := 2.666555777
	t2 := t.Add(time.Duration(f*1e6) * time.Microsecond)
	//2666555
	//1668477270992898
	fmt.Println(t2.UnixMicro())
}

func TestDate(xx *testing.T) {
	//9996-12-31 23:59:59
	y := 9996
	m := 12
	d := 31

	t0 := time.Date(y, time.Month(m), d, 23, 59, 59, 999_999_999, time.UTC)

	t3 := t0.AddDate(0, 0, 59)
	fmt.Println(t3)

	t4 := t0.AddDate(3, 0, 0)
	fmt.Println(t4)
}

func TestDate2(xx *testing.T) {
	//9996-12-31 23:59:59
	y := 9996
	m := 12
	d := 31

	mInterval := 2
	yInterval := 0

	//////////////////////////////////////
	mCalculated := m + mInterval
	mFinal := mCalculated % 12
	yFinal := y + yInterval + mCalculated/12

	t0 := time.Date(y, time.Month(m), d, 23, 59, 59, 999_999_999, time.UTC)
	fmt.Println(t0)

	t := time.Date(y, time.Month(mCalculated), d, 23, 59, 59, 999_999_999, time.UTC)
	fmt.Println(t)

	if mCalculated > 12 {
		m = mCalculated % 12
	}

	t2 := time.Date(yFinal, time.Month(mCalculated), d, 23, 59, 59, 999_999_999, time.UTC)
	fmt.Println(t2)

	t3 := time.Date(yFinal, time.Month(mFinal), d, 23, 59, 59, 999_999_999, time.UTC)
	fmt.Println(t3)
}

func TestAddMonth(xx *testing.T) {
	t0 := time.Date(9996, 12, 31, 23, 59, 59, 999_999_999, time.UTC)
	t1 := addMonth(&t0, 0, 2)
	fmt.Println(t1)

}

func TestRoundMicrosecond(tx *testing.T) {
	t1 := time.Date(9996, 12, 31, 23, 59, 59, 999_999_500, time.UTC)

	nsec := t1.Nanosecond()
	fmt.Println(nsec)

	ns := time.Nanosecond * time.Duration(nsec%1e3)
	fmt.Println(ns)
	if ns > 499 {
		t1 = t1.Add(time.Microsecond - ns)
	} else {
		t1 = t1.Add(-ns)
	}

	fmt.Println(t1)
}

func TestRoundMicrosecond2(tx *testing.T) {
	//纳秒长度不能超过9位,否则丢失精度
	t1 := time.Date(9996, 12, 31, 23, 59, 59, 999_999_500, time.UTC)
	RoundToMicrosecond(&t1)
	fmt.Println(t1)

	t2 := time.Date(9996, 12, 31, 23, 59, 59, 999_999_499, time.UTC)
	RoundToMicrosecond(&t2)
	fmt.Println(t2)
}

func RoundToMicrosecond(t *time.Time) {
	nsec := t.Nanosecond()
	ns := time.Nanosecond * time.Duration(nsec%1e3)
	fmt.Println(ns)
	if ns > 499 {
		*t = t.Add(time.Microsecond - ns)
	} else {
		*t = t.Add(-ns)
	}
}

func addMonth(t *time.Time, yInterval int, mInterval int) time.Time {
	y := t.Year()
	m := t.Month()
	d := t.Day()

	mCalculated := m + time.Month(mInterval)
	mFinal := mCalculated % 12
	yFinal := y + yInterval + int(mCalculated/12)

	switch mFinal {
	case 4, 6, 9, 11:
		if d > 30 {
			d = 30
		}
		break
	case 2:
		if yFinal%4 == 0 && d > 29 {
			d = 29
		} else {
			d = 28
		}
		break
	}

	return time.Date(yFinal, mFinal, d, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
}

func TestWeek(t *testing.T) {
	weekDay := time.Now().Weekday()
	fmt.Println(weekDay)
	fmt.Println(int(weekDay))
	fmt.Println(strconv.Itoa(int(weekDay)))
}

func TestMonth(t *testing.T) {
	t2 := time.Now()
	fmt.Println(t2.Month() < 10)
}

func TestNanoSecond(t *testing.T) {
	fmt.Println(time.Now().Second())
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().Nanosecond())
}

func TestHour24(t *testing.T) {
	fmt.Println(time.Now().Hour())
}

func TestCalendarWeekOfYear(txx *testing.T) {
	fmt.Println(time.Now().Weekday())
	fmt.Println(time.Now().ISOWeek())
}

func TestISO(t *testing.T) {
	ti := time.Now()
	y, _ := ti.ISOWeek()
	fmt.Println(strconv.Itoa(y)[1:])
	fmt.Println(strconv.Itoa(y)[2:])
	fmt.Println(strconv.Itoa(y)[3:])
}

func TestBC(t *testing.T) {
	time1 := time.Unix(-1480390585, 0)
	fmt.Println(time1.Year())
	fmt.Println(time.Now().Day())
}

func TestInsertStr(t *testing.T) {
	p := "green"
	index := 2
	q := p[:index] + "HI" + p[index:]
	fmt.Println(p, q)
}
