package time_demo

import (
	"gitee.com/go-package/carbon"
	"testing"
)

func TestCarbon(t *testing.T) {
	// 今天此刻
	carbon.Now().ToDateTimeString() // 2020-08-05 13:14:15
	// 今天日期
	carbon.Now().ToDateString() // 2020-08-05
	// 今天时间
	carbon.Now().ToTimeString() // 13:14:15
	// 今天秒级时间戳
	carbon.Now().Timestamp()           // 1596604455
	carbon.Now().TimestampWithSecond() // 1596604455
	// 今天毫秒级时间戳
	carbon.Now().TimestampWithMillisecond() // 1596604455000
	// 今天微秒级时间戳
	carbon.Now().TimestampWithMicrosecond() // 1596604455000000
	// 今天纳秒级时间戳
	carbon.Now().TimestampWithNanosecond() // 1596604455000000000
	// 指定时区的今天此刻
	carbon.SetTimezone(carbon.NewYork).Now().ToDateTimeString() // 2020-08-05 01:14:15
	// 昨天此刻
	carbon.Yesterday().ToDateTimeString() // 2020-08-04 13:14:15
	// 昨天日期
	carbon.Yesterday().ToDateString() // 2020-08-04
	// 昨天时间
	carbon.Yesterday().ToTimeString() // 13:14:15
	// 昨天秒级时间戳
	carbon.Yesterday().Timestamp()           // 1596518055
	carbon.Yesterday().TimestampWithSecond() // 1596518055
	// 明天毫秒级时间戳
	carbon.Yesterday().TimestampWithMillisecond() // 1596518055000
	// 明天微秒级时间戳
	carbon.Yesterday().TimestampWithMicrosecond() // 1596518055000000
	// 明天纳秒级时间戳
	carbon.Yesterday().TimestampWithNanosecond() // 1596518055000000000
	// 指定时区的昨天此刻
	carbon.SetTimezone(carbon.NewYork).Yesterday().ToDateTimeString() // 2020-08-04 01:14:15
	// 指定日期的昨天此刻
	carbon.Parse("2021-01-28 13:14:15").Yesterday().ToDateTimeString() // 2021-01-27 13:14:15
	// 明天此刻
	carbon.Tomorrow().ToDateTimeString() // 2020-08-06 13:14:15
	// 明天日期
	carbon.Tomorrow().ToDateString() // 2020-08-06
	// 明天时间
	carbon.Tomorrow().ToTimeString() // 13:14:15
	// 明天秒级时间戳
	carbon.Tomorrow().Timestamp()           // 1596690855
	carbon.Tomorrow().TimestampWithSecond() // 1596690855
	// 明天毫秒级时间戳
	carbon.Tomorrow().TimestampWithMillisecond() // 1596690855000
	// 明天微秒级时间戳
	carbon.Tomorrow().TimestampWithMicrosecond() // 1596690855000000
	// 明天纳秒级时间戳
	carbon.Tomorrow().TimestampWithNanosecond() // 1596690855000000000
	// 指定时区的明天此刻
	carbon.SetTimezone(carbon.NewYork).Tomorrow().ToDateTimeString() // 2020-08-06 01:14:15
	// 指定日期的明天此刻
	carbon.Parse("2021-01-28 13:14:15").Tomorrow().ToDateTimeString() // 2021-01-29 13:14:15
}
