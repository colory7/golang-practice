package enum_demo

import (
	"fmt"
	"testing"
)

type FromCharDateMode int

const (
	FROM_CHAR_DATE_NONE FromCharDateMode = iota
	FROM_CHAR_DATE_GREGORIAN
	FROM_CHAR_DATE_ISOWEEK
)

type KeyWord struct {
	name      string
	len       int
	id        int
	is_digit  bool
	date_mode FromCharDateMode
}

func TestEnum4(t *testing.T) {
	key1 := new(KeyWord)
	key1.name = "aa"
	key1.id = 111
	key1.date_mode = FROM_CHAR_DATE_ISOWEEK

	fmt.Println(key1.name)
	fmt.Println(key1.date_mode)

	key2 := KeyWord{
		name:      "bb",
		len:       3,
		id:        222,
		is_digit:  true,
		date_mode: FROM_CHAR_DATE_GREGORIAN,
	}

	fmt.Println(key2.name)
	fmt.Println(key2.date_mode)

}
