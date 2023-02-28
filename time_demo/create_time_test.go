package time_demo

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateTime2(tx *testing.T) {
	year := 0
	day := 399
	t1 := time.Date(year, 0, day, 22, 0, 3, 123456789789, time.UTC)
	fmt.Println(t1)
}
