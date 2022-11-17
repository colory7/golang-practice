package time_demo

import (
	"fmt"
	"github.com/knz/strtime"
	"strconv"
	"testing"
	"time"
)

func TestFormat(xxx *testing.T) {
	start := "2016-10-20"
	t, _ := strtime.Strptime(start, "%Y-%m-%d")
	end, _ := strtime.Strftime(t, "%Y-%m-%d")
	fmt.Println(start, end)

}

func TestFloat(t *testing.T) {
	dt1 := strconv.FormatFloat(2021010000000000000000011600000700.333666, 'f', 0, 64)
	fmt.Println(dt1)
}

func TestTimeParse33(xx *testing.T) {
	t, err := time.Parse("20060102150405.999999", "20210101160700.33366699876698")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
}
