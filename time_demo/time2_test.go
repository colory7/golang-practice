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

func TestTimeZone(txx *testing.T) {
	//time.LoadLocation("Asia/Shanghai")
	//time.LoadLocation("US/Pacific")
	loc, err := time.LoadLocation("Africa/Cairo")
	if err != nil {
		panic(err)
	}
	fmt.Println(loc)
	//time.Local = time.UTC
	time.Local = loc

	t := time.Now()
	fmt.Println(t.Location().String())
	fmt.Println(t.Location())
	fmt.Println("====")
	fmt.Println(t.Local())
	fmt.Println(t.Local().Zone())
	fmt.Println("====")
	fmt.Println(t.UTC())
}

func TestTimeZone2(t *testing.T) {
	layout := "2006-01-02 15:04:05"

	var cstZone = time.FixedZone("aa", 8*3600+120)
	time.Local = cstZone
	fmt.Println("SH: ", time.Now().In(cstZone).Format(layout))

}

//func TestOracleTS(txx *testing.T) {
//	result := bytes.Buffer{}
//	tsFormat := "15:04:05"
//
//	t := time.Now()
//	if t.Hour() > 12 {
//		result.WriteString(oracle_demo.NLS_AM)
//		result.WriteByte(oracle_demo.ASSIC_SPACE)
//	} else {
//		result.WriteString(oracle_demo.NLS_AM)
//		result.WriteByte(oracle_demo.ASSIC_SPACE)
//	}
//	result.WriteString(t.Format(tsFormat))
//
//	fmt.Println(result.String())
//}

func TestOracleTZD(txx *testing.T) {
	t := time.Now()
	zone, _ := t.Local().Zone()
	fmt.Println(zone)
}

func TestOracleTZH(txx *testing.T) {
	t := time.Now()
	result := t.Format("-07")
	fmt.Println(result)

}

func TestOracleTZM(txx *testing.T) {
	t := time.Now()
	result := t.Format("-0700")[3:]
	fmt.Println(result)
}

func TestOracleTZR(txx *testing.T) {
	t := time.Now()
	fmt.Println(t.Location().String())
}

var cstZone = time.FixedZone("CST", 8*3600)

func TestParseTime(txx *testing.T) {
	layout := "2006-01-02 15:04:05"

	t := time.Date(2011, time.Month(3), 12, 15, 30, 20, 0, cstZone)
	fmt.Println(t.Format(layout))
}

func TestParseTime2(txx *testing.T) {
	layout := "200612 15:04:05"
	t, _ := time.Parse(layout, "202111 15:01:02")
	fmt.Println(t)
	fmt.Println(t.Format(layout))
}

func TestParseTime3(txx *testing.T) {
	layout := "2006/1/02"
	t, _ := time.Parse(layout, "2019/5/07")
	fmt.Println(t)
	fmt.Println(t.Format(layout))
}

func TestCreateTime(txx *testing.T) {
	dateLayout := "2006-01-02 15:04:05"
	//offset:=8*3600
	//var zone = time.FixedZone("CST",offset)
	now := time.Now()

	year, month, day := 0, time.Month(0), 0
	hour, min, sec, nsec := 0, 0, 0, 0
	if year == 0 {
		year = now.Year()
	}
	if month == 0 {
		month = now.Month()
	}
	if day == 0 {
		day = 1
	}
	t := time.Date(year, month, day, hour, min, sec, nsec, time.Local)
	fmt.Println(t.Format(dateLayout))
}

func TestFixedZone(t *testing.T) {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	fmt.Println("SH : ", time.Now().In(cstSh).Format("2006-01-02 15:04:05"))
}

func TestAppendInt(t *testing.T) {
	println(string(strconv.Itoa(time.Now().Year())[0]) + "023")
	println(fmt.Sprintf("%d", time.Now().Year()))
}
