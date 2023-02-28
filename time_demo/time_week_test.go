package time_demo

import (
	"fmt"
	"testing"
	"time"
)

func TestWeek2(t *testing.T) {
	t1, _ := time.Parse(time.RFC3339, "2023-02-20T22:08:41+00:00")
	fmt.Println(t1.ISOWeek())
	fmt.Println(int(t1.Weekday()))

	t1, _ = time.Parse(time.RFC3339, "2023-02-21T22:08:41+00:00")
	fmt.Println(t1.ISOWeek())
	fmt.Println(int(t1.Weekday()))

	t1, _ = time.Parse(time.RFC3339, "2023-02-26T22:08:41+00:00")
	fmt.Println(t1.ISOWeek())
	fmt.Println(int(t1.Weekday()))

	t1, _ = time.Parse(time.RFC3339, "2023-02-27T22:08:41+00:00")
	fmt.Println(t1.ISOWeek())
	fmt.Println(int(t1.Weekday()))

}
