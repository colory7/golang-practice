package time_demo

import (
	"fmt"
	"testing"
	"time"
)

func TestTimestampdiff(t *testing.T) {
	t1, err := time.Parse("2006-01-02", "2022-01-01")
	if err != nil {
		panic(err)
	}
	t2, err := time.Parse("2006-01-02", "2022-12-31")
	if err != nil {
		panic(err)
	}

	d := diffYear(t1, t2)

	fmt.Println(d)
}

func TestTimestampdiff2(t *testing.T) {
	t1, err := time.Parse("2006-01-02", "5020-01-01")
	if err != nil {
		panic(err)
	}
	t2, err := time.Parse("2006-01-02", "2020-12-31")
	if err != nil {
		panic(err)
	}

	d := diffYear(t1, t2)

	fmt.Println(d)
}

func TestTimestampdiff3(t *testing.T) {
	t1, err := time.Parse("2006-01-02", "2020-12-31")
	if err != nil {
		panic(err)
	}
	t2, err := time.Parse("2006-01-02", "5020-01-01")
	if err != nil {
		panic(err)
	}

	d := diffYear(t1, t2)

	fmt.Println(d)
}

func TestSecond(t *testing.T) {
	t1, err := time.Parse("2006-01-02 15:04:05.999999", "2022-04-26 13:00:00")
	if err != nil {
		panic(err)
	}
	t2, err := time.Parse("2006-01-02 15:04:05.999999", "2022-04-26 13:00:00.999999")
	if err != nil {
		panic(err)
	}

	d := (t2.UnixMicro() - t1.UnixMicro()) / 1e6

	fmt.Println(d)
}

func TestSecond2(t *testing.T) {
	t1, err := time.Parse("2006-01-02 15:04:05.999999", "2022-04-26 13:00:00.1")
	if err != nil {
		panic(err)
	}
	t2, err := time.Parse("2006-01-02 15:04:05.999999", "2022-04-26 13:00:00.999999")
	if err != nil {
		panic(err)
	}

	d := (t2.UnixMicro() - t1.UnixMicro()) / 1e6

	fmt.Println(d)
}

func TestSecond3(t *testing.T) {
	t1, err := time.Parse("2006-01-02 15:04:05.999999", "2022-04-26 13:00:00.1")
	if err != nil {
		panic(err)
	}
	t2, err := time.Parse("2006-01-02 15:04:05.999999", "2022-04-26 13:00:01.999999")
	if err != nil {
		panic(err)
	}

	d := (t2.UnixMicro() - t1.UnixMicro()) / 1e6

	fmt.Println(d)
}

func TestSecond4(t *testing.T) {
	t1, err := time.Parse("2006-01-02 15:04:05.999999", "2022-04-26 13:00:01.1")
	if err != nil {
		panic(err)
	}
	t2, err := time.Parse("2006-01-02 15:04:05.999999", "2022-04-26 13:00:00.999999")
	if err != nil {
		panic(err)
	}

	d := (t2.UnixMicro() - t1.UnixMicro()) / 1e6

	fmt.Println(d)
}
