package tp_demo

import (
	"fmt"
	"testing"
	"time"
)

func TestTruncate(xx *testing.T) {
	// Defining t for Truncate method
	t := time.Date(2047, 47, 96, 123, 98, 81, 999434, time.UTC)

	// Defining duration
	d := (2 * time.Hour)

	// Calling Truncate() method
	trunc := t.Truncate(d)

	// Prints output
	fmt.Println(trunc)
}

func TestTruncate2(xx *testing.T) {
	// Defining t for Truncate method
	t := time.Date(2020, 12, 31, 23, 56, 45, 999434, time.UTC)

	fmt.Println(t)

	// Defining duration
	d := (2 * 24 * time.Hour)

	// Calling Truncate() method
	trunc := t.Truncate(d)

	// Prints output
	fmt.Println(trunc)
}
