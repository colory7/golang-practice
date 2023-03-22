package tmp

import (
	"errors"
	"time"
)

/*
WeekWithModel
This function returns the week number for date. The two-argument form of WEEK() enables you
to specify whether the week starts on Sunday or Monday and whether the return value should
be in the range from 0 to 53 or from 1 to 53. If the mode argument is omitted, the value of the
default_week_format system variable is used.
mode = 0,从本年的第一个星期日开始，是第一周。前面的计算为第0周
mode = 1,假如1月1日到第一个周一的天数超过3天，则计算为本年的第一周。否则为第0周
mode = 2,从本年的第一个星期日开始，是第一周。前面的计算为上年度的第5x周
mode = 3,假如1月1日到第一个周一的天数超过3天，则计算为本年的第一周。否则为上年度的第5x周
mode = 4,假如1月1日到第一个周日的天数超过3天，则计算为本年的第一周。否则为第0周
mode = 5,从本年的第一个星期一开始，是第一周。前面的计算为第0周。
mode = 6,假如1月1日到第一个周日的天数超过3天，则计算为本年的第一周。否则为上年度的第5x周
mode = 7,从本年的第一个星期一开始，是第一周。前面的计算为上年度的第5x周
*/
func WeekWithModel(date string, mode int) (int, error) {
	if mode < 0 || mode > 7 {
		return -1, errors.New("mode must be in the range from 0 to 7")
	}
	var week int = 0
	//日期格式为YYYY-MM-DD,字符串转换为日期
	t, _ := time.Parse("2006-01-02", date)
	switch mode {

	//从本年的第一个星期日开始，是第一周。前面的计算为第0周
	case 0:
		d := t.YearDay() - int(t.Weekday()) - 1 //从本年的第一天开始，到本周的第一天的天数
		if d < 0 {
			week = 0
		} else {
			week = d/7 + 1
		}

	//假如1月1日到第一个周一的天数超过3天，则计算为本年的第一周。否则为第0周
	case 1:
		weekDay := int(t.Weekday())
		//如果是周日，则移到一周最后一天
		if weekDay == 0 {
			weekDay = 7
		}
		d := t.YearDay() - weekDay
		if d%7 > 3 {
			week = 1
		}
		week += d/7 + 1

	//从本年的第一个星期日开始，是第一周。前面的计算为上年度的第5x周
	case 2:
		d := t.YearDay() - int(t.Weekday()) - 1
		if d <= 0 {
			week = 52 + d/7
		} else {
			week = d/7 + 1
		}

	//假如1月1日到第一个周一的天数超过3天，则计算为本年的第一周。否则为上年度的第5x周
	case 3:
		_, week = t.ISOWeek()

	//假如1月1日到第一个周日的天数超过3天，则计算为本年的第一周。否则为第0周
	case 4:
		d := t.YearDay() - int(t.Weekday()) - 1
		if d%7 > 3 {
			week = 1
		}
		week += d/7 + 1

	//从本年的第一个星期一开始，是第一周。前面的计算为第0周
	case 5:
		weekDay := int(t.Weekday())
		//如果是周日，则移到一周最后一天
		if weekDay == 0 {
			weekDay = 7
		}
		d := t.YearDay() - weekDay
		if d < 0 {
			week = 0
		} else {
			week = d/7 + 1
		}

	//假如1月1日到第一个周日的天数超过3天，则计算为本年的第一周。否则为上年度的第5x周
	case 6:
		d := t.YearDay() - int(t.Weekday()) - 1
		if d%7 > 3 {
			week = 1
		}
		week += d/7 + 1
		if week > 52 {
			week = 1
		}

	//从本年的第一个星期一开始，是第一周。前面的计算为上年度的第5x周
	case 7:
		weekDay := int(t.Weekday())

		//如果是周日，则移到一周最后一天
		if weekDay == 0 {
			weekDay = 7
		}
		d := t.YearDay() - weekDay
		if d < 0 {
			//计算t去年的周数
			lastT := time.Date(t.Year()-1, 12, 31, 0, 0, 0, 0, time.Local)
			lastYearWeek, _ := WeekWithModel(lastT.Format("2006-01-02"), 7)
			week = lastYearWeek + d/7
		} else {
			week = d/7 + 1
		}

	}

	return week, nil

}

func Week(date string) (int, error) {
	return WeekWithModel(date, 0)
}
