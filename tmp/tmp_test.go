package tmp

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	fmt.Println(runtime.GOMAXPROCS(0))
}

func Test2(tx *testing.T) {
	t := time.Date(2018, 0, 59, 0, 0, 0, 0, time.UTC)
	fmt.Println(t)

	t = time.Date(2020, 0, 59, 0, 0, 0, 0, time.UTC)
	fmt.Println(t)

	t = time.Date(2018, 0, 60, 0, 0, 0, 0, time.UTC)
	fmt.Println(t)

	t = time.Date(2020, 0, 60, 0, 0, 0, 0, time.UTC)
	fmt.Println(t)

	fmt.Println()
	t = time.Date(2018, 1, 59, 0, 0, 0, 0, time.UTC)
	fmt.Println(t)

	t = time.Date(2020, 1, 59, 0, 0, 0, 0, time.UTC)
	fmt.Println(t)

	t = time.Date(2018, 1, 60, 0, 0, 0, 0, time.UTC)
	fmt.Println(t)

	t = time.Date(2020, 1, 60, 0, 0, 0, 0, time.UTC)
	fmt.Println(t)
}
