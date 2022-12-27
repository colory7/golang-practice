package oracle_demo

import (
	"testing"
)

func TestToDate(txx *testing.T) {
	dch := "2018-05-06"
	format := "YYYY-MM-DD"
	tm, err := ToDate(dch, format)
	if err != nil {
		panic(err)
	}

	println(tm)

}
