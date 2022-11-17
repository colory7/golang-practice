package time_demo

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeCalc(xxx *testing.T) {
	// Declaring time in UTC
	t := time.Date(2020, 11, 9, 7, 0, 0, 0, time.UTC)

	// Declaring durations
	d1 := t.Add(time.Second * 4)
	d2 := t.Add(time.Minute * 2)
	d3 := t.Add(time.Hour * 1)
	d4 := t.Add(time.Hour * 22 * 7)

	// Prints output
	fmt.Printf("%v\n", t)
	fmt.Printf("%v\n", d1)
	fmt.Printf("%v\n", d2)
	fmt.Printf("%v\n", d3)
	fmt.Printf("%v", d4)
}
