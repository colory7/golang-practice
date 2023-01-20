package base_demo

import (
	"fmt"
	"testing"
	"time"
)

func TestTimePointer(xx *testing.T) {
	t := time.Now()

	fmt.Println(t)

	mod(&t)

	fmt.Println(t)
}

func mod(t *time.Time) time.Time {
	return t.AddDate(1, 0, 0)
}

func TestTimePointer2(xx *testing.T) {
	t := time.Now()

	fmt.Println(t)

	mod2(&t)

	fmt.Println(t)
}

func mod2(t *time.Time) {
	*t = t.AddDate(1, 0, 0)
}

func TestTimePointer3(xx *testing.T) {
	t := time.Now()

	fmt.Println(t)

	t2 := mod3(&t)

	fmt.Println(t)
	fmt.Println(t2)
	fmt.Println(&t2)
}

func mod3(t *time.Time) time.Time {
	*t = t.AddDate(1, 0, 0)
	return *t
}

func TestTimePointer4(xx *testing.T) {
	t := time.Now()

	fmt.Println(t)

	mod4(&t)

	fmt.Println(t)
}

func mod4(t *time.Time) time.Time {
	*t = (*t).AddDate(1, 0, 0)
	*t = t.AddDate(1, 0, 0)
	return *t
}

func TestPointer(tx *testing.T) {
	var t1 = &time.Time{}
	fmt.Println(t1)
	fmt.Println(t1 == nil)

	var t2 = new(time.Time)
	fmt.Println(t2)
	fmt.Println(t2 == nil)

	var t3 time.Time
	fmt.Println(t3)
	//fmt.Println(t3 == nil)
	fmt.Println(t3.Year())

	fmt.Println(time.Time{})
	//fmt.Println(time.Time{}==nil)

	var t4 time.Time
	t4 = time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(t4.Year())

	//var t5 *time.Time
	//*t5 = time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)
	//fmt.Println(t5.Year())

	var t6 *time.Time
	tm := time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)
	t6 = &tm
	fmt.Println(t6.Year())

}

func aa() (time.Time, error) {
	return time.Time{}, nil
}
