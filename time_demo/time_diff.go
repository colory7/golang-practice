package time_demo

import "time"

func diffYear(t1 time.Time, t2 time.Time) int {
	d := t2.Year() - t1.Year()
	var t2Align time.Time
	if d > 0 {
		t2Align = t2.AddDate(-d, 0, 0)
		if t2Align.Before(t1) {
			d -= 1
		}
	} else if d < 0 {
		t2Align = t2.AddDate(-d, 0, 0)
		if t2Align.After(t1) {
			d += 1
		}
	}
	return d
}

func diffMonth(t1 time.Time, t2 time.Time) int {
	year1, month1, day1 := t1.Date()
	year2, month2, day2 := t2.Date()

	var d = (year2-year1)*12 + int(month2-month1)

	t1Trunc := time.Date(0, 0, day1, t1.Hour(), t1.Minute(), t1.Second(), t1.Nanosecond(), time.UTC)
	t2Trunc := time.Date(0, 0, day2, t2.Hour(), t2.Minute(), t2.Second(), t2.Nanosecond(), time.UTC)
	t1AfterT2 := t1Trunc.After(t2Trunc)

	if !t1Trunc.Equal(t2Trunc) {
		if d > 0 {
			if t1AfterT2 {
				d -= 1
			}
		} else if d < 0 {
			if !t1AfterT2 {
				d += 1
			}
		}
	}

	return d
}
