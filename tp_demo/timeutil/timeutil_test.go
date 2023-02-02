package timeutil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSuite(t *testing.T) {
	tests := []struct {
		i         int
		ch        string
		exception bool
	}{
		{1, "20220810", false},
		{1, "2022-08-11", false},
		{1, "2022/08/10", false},
		{1, "20170623113939", false},
		{1, "2017-06-23 11:39:39", false},
		{1, "2017/06/23 11:39:39", false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			tm, _, err := ParseTime2(test.ch)
			if test.exception {
				assert.Error(t, err)
				//fmt.Println(err)
			} else {
				if err != nil {
					assert.NoError(t, err)
				}
				fmt.Println(tm)
			}
		})
	}
}

func TestParseTime(t *testing.T) {
	s := "2006-01-02 15:04:05.999939999"
	f := "2006-01-02 15:04:05.999999999"
	tm, err := time.Parse(f, s)
	if err != nil {
		panic(err)
	}
	fmt.Println(tm)
}

func TestParseTime2(t *testing.T) {
	tests := []struct {
		i         int
		ch        string
		exception bool
	}{
		{1, "20091130", false},
		{1, "2002 01 06 00 00 00", false},
		{1, "23020106\"fsfs\"", true},
		{1, "23020106", false},
		{1, "2002-01-06a00:00:00", true},
		{1, "2002-01-06 00:00:00", false},
		{1, "20:00:00", false},
		{1, "2002-01-06 00:00:00.356789231", false},
		{1, "2002-01-06 00:00:00.1234567891", true},
		{1, "2017-06-23", false},
		{1, "2017-06-23 11:39:39.123456789", false},
		{1, "2019/01/02", false},
		{1, "2019/1/02", false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			tm, err := parseTimeSimple(test.ch)
			if test.exception {
				assert.Error(t, err)
				fmt.Println("error:", err)
			} else {
				if err != nil {
					assert.NoError(t, err)
				}
				fmt.Println(tm)
			}
		})
	}
}

func TestParseTime3(t *testing.T) {
	tests := []struct {
		i         int
		ch        string
		exception bool
	}{
		{1, "20091130", false},
		{1, "2002 01 06 00 00 00", false},
		{1, "23020106\"fsfs\"", true},
		{1, "2002-01-06a00:00:00", true},
		{1, "2002-01-06 00:00:00", false},
		{1, "2002-01-06 00:00:00.356789231", false},
		{1, "2002-01-06 00:00:00.1234567891", true},
		{1, "2017-06-23", false},
		{1, "2017-06-23 11:39:39.123456789", false},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			tm, err := parseTimeByOrder2(test.ch)
			if test.exception {
				assert.Error(t, err)
				fmt.Println("error:", err)
			} else {
				if err != nil {
					assert.NoError(t, err)
				}
				fmt.Println(tm)
			}
		})
	}
}

func TestParseTime4(t *testing.T) {
	s := "3026-05-03    080617.  999939999"
	tm, err := parseTimeByOrder2(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(tm)
}
