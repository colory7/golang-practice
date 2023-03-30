package time_demo

import (
	"fmt"
	"testing"
	"time"
)

func TestUnixtime(txx *testing.T) {
	ts := 1344887103
	t := time.Date(0, 0, 0, 0, 0, ts, 0, time.UTC)
	fmt.Println(t)
}

func TestUnixtime3(txx *testing.T) {
	local := time.Local
	loc, _ := time.LoadLocation("UTC")
	time.Local = loc

	ts := 1344887103
	t := time.Unix(int64(ts), 0)
	fmt.Println(t)

	TestUnixtime(nil)
	time.Local = local
	TestUnixtime(nil)
}

func TestUnixtime4(t *testing.T) {
	unixtime()
}

func unixtime() {
	ts := 1344887103
	t := time.Unix(int64(ts), 0)
	fmt.Println(t)
}

func TestUnixtime5(txx *testing.T) {
	local := time.Local
	loc, _ := time.LoadLocation("UTC")
	time.Local = loc

	unixtime()

	time.Local = local

	unixtime()
}

func TestUnix(t *testing.T) {
	fmt.Println(Unix(-1, 0))
	fmt.Println(Unix(0, 0))
}

func TestUnix2(txx *testing.T) {
	ts := 1344887103
	t := time.Unix(int64(ts), 0)
	fmt.Println(t.In(time.UTC))
}

func Unix(sec int64, nsec int64) time.Time {
	local := time.Local
	loc, _ := time.LoadLocation("UTC")
	time.Local = loc
	t := time.Unix(sec, nsec)
	time.Local = local
	return t
}

func TestNano(t *testing.T) {
	nano := time.Now().Nanosecond()
	fmt.Println(nano)
}
