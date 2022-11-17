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
