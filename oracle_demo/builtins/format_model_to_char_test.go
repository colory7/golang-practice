package builtins

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSuiteToCharByStr(t *testing.T) {
	tests := []struct {
		i         int
		ch        string
		format    string
		exception bool
	}{
		{1, "3450", "999,99", false},
		{1, "3450", "99999", false},
		{1, "3450", "99999,", false},
		{1, "3450", "99999", false},
		{1, "3450", "99999", false},
		{1, "124548", "99G9G999", false},
		{1, "1245.48", "99G9G9D99", false},
		{1, "1245", "$99999", false},
		{1, "124548-", "99G9G999S", true},
		{1, "12,4,548-", "99G9G999S", true},
		{1, "12,4,548", "99G9G999", true},
		{1, "12,4,5.48", "99G9G9D99", true},
		{1, "34,50", "99999", true},
		{1, "3450", ",99999", true},
		{1, "34,50", "99999,", true},
		{1, "34,50", ",99999", true},
		{1, "34,50", "999,99", true},
		{1, "34,50", "99999", true},
		{1, "3450", "999,99", true},
		{1, "3450", "99", true},
		{1, "3450", "9G9", true},
		{1, "12,4,548-", "99G9G999S", true},
		{1, "12,4,5.4,8", "99G9G9D9G9", true},
		{1, "34,50", "999,99", true},
		{1, "34,50", "999G99", true},
		{1, "34,50", "99G9G99", true},
		{1, "34,50", "99999", true},
		{1, "34,50", "99999,", true},
		{1, "34,50", ",99999", true},
		{1, "34,50", "99999", true},
		{1, "-12,4,548", "99G9G999S", true},
		{1, "-12,4,548", "S99G9G999S", true},
		{1, "12,4,5.4,8", "99G9G9D9G9", true},
		{1, "-12,4,548", "S99G9G999", true},
		{1, "23.54", "99.99", false},
		{1, "23.54", "99D99", false},
		{1, "12454.8", "99G999D9S", false},
		{1, "12454.8", "99G9G99D9S", false},
		{1, "124548", "99G9G999S", false},
		{1, "124548", "99G9G999", false},
		{1, "12,454.8-", "99G999D9S", true},
		{1, "12,4,54.8-", "99G9G99D9S", true},
		{1, "12,4,548-", "99G9G999S", true},
		{1, "12,4,548", "99G9G999", true},
		{1, "23.5.4", "99.9.9", true},
		{1, "23.54", "99.9.9", true},
		{1, "23.5.4", "99D9D9", true},
		{1, "23.54", "99D9D9", true},
		{1, "1", "0", false},
		{1, "1322526", "0099000", false},
		{1, "1322526", "0099000", false},
		{1, "1322526", "9999099", false},
		{1, "1322526", "99990999999", false},
		{1, "1322526", "999090909999999009", false},
		{1, "1322526", "99909090999009", false},
		{1, "1322526", "000000000000000", false},
		{1, "1322526", "90009000009", false},
		{1, "1322526", "900000000009", false},
		{1, "1322526", "00000000009", false},
		{1, "1322526", "00", false},
		{1, "1322526", "99", false},
		{1, "1322526", "99.9", false},
		{1, "1322526", "99990999999", false},
		{1, "1322526", "9999999999", false},
		{1, "1322526", "99999,99999", false},
		{1, "1322526", "9999", true},
		{1, "1322526", "966", true},
		{1, "1322526", "66", true},
		{1, "1322526", "6", true},
		{1, "1322526", "666666666", true},
		{1, "123.45", "B999.99", false},
		{1, "12345.678", "$999999.999", false},
		{1, "123.45", "L999.99", false},
		{1, "123.45", "U999.99", false},
		{1, "123.45", "U99999.990", false},
		{1, "356", "C999", false},
		{1, "¥123.45", "S999.99", true},
		{1, "$12345.678", "$999999.999", true},
		{1, "$12345.678", "$$999999.999", true},
		{1, "$12345.678", "$999999.999", true},
		{1, "$12345.678", "$99.999", true},
		{1, "$12345.678", "9$9.999", true},
		{1, "$12345.678", "99.999$", true},
		{1, "$12345.678", "99.999", true},
		{1, "￥123.45", "B999.99", true},
		{1, "485", "9999MI", false},
		{1, "-485", "9999MI", false},
		{1, "485-", "9999MI", false},
		{1, "-485", "9999MI", false},
		{1, "-485", "9999MI", false},
		{1, "-485", "MI9999", false},
		{1, "-485", "99MI99", false},
		{1, "-485", "99MI99", false},
		{1, "485", "9999MIMI", false},
		{1, "485", "999PR", false},
		{1, "-485", "999PR", false},
		{1, "485", "PR999", true},
		{1, "-485", "PR999", true},
		{1, "-1234567890", "9999999999S", false},
		{1, "+1234567890", "9999999999S", false},
		{1, "258-", "999S", true},
		{1, "-258", "9S99", true},
		{1, "-258", "9SS99", true},
		{1, "-258", "S9SS99", true},
		{1, "1", "S", true},
		{1, "1", "SS", true},
		{1, "-258", "FMSB999", false},
		{1, "258", "FMSL999", false},
		{1, "258", "FML999S", false},
		{1, "258", "FML999SPR", true},
		{1, "$258-", "FM$999MI", true},
		{1, "$2,5,8.36-", "$9,9,9.99", true},
		{1, "$2,5,8.36-", "$9,99.99", true},
		{1, "$2,5,8.36-", "$999.99", true},
		{1, "$2,5,8.36-", "$99,9.99", true},
		{1, "$2,58.36-", "$99,9.99", true},
		{1, "$25,8.36", "$9,99.99", true},
		{1, "$2,5,8.36", "$9,9,9.99", true},
		{1, "$258-", "FM$999PRMI", true},
		{1, "$258-", "FMS$999PRMI", true},
		{1, "$258-", "S$999PRMI", true},
		{1, "¥258-", "SL999PRMI", true},
		{1, "¥258-", "FMSL999PRMI", true},
		{1, "¥258-", "FMSL999MIPR", true},
		{1, "¥258-", "FMSL999MI", true},
		{1, "¥258-", "FMSL999PR", true},
		{1, "258-", "FMSL999PR", true},
		{1, "-$258", "FMS$999", true},
		{1, "$2,5,8.36-", "FM$9,9,9.99MI", true},
		{1, "$2,5,8.36-", "$9,9,9.99MI", true},
		{1, "$2,5,8.36-", "$9,9,9.99MI", true},
		{1, "$2,5,8.36", "$9,9,9.99", true},
		{1, "$25,8.36", "$9099.99", true},
		{1, "$2,5,8.36", "$9,99.99", true}, // 与预期有差异
		{1, "$2,5,8.36", "$9,9,9.99", true},
		{1, "1", "9999999EEEE", false},
		{1, "1e2", "9999999EEEE", false},
		{1, "-1e2", "9999999EEEE", false},
		{1, "-1e+2", "9999999EEEE", false},
		{1, "-1e-2", "9999999EEEE", false},
		{1, "+123.45e2", "9999999EEEE", true},
		{1, "+123.456", "9.9EEEE", true},
		{1, "+123.456e2", "99999999.99EEEE", true},
		{1, "+123.456", "9.999EEEE", true},
		{1, "+123.456", "9.9999EEEE", true},
		{1, "+123.456", "9.99999EEEE", true},
		{1, "+1E+123", "9.9EEEE", true},
		{1, "+123.456", "FM9.9EEEE", true},
		{1, "1", "9V9", false},
		{1, "1", "99V99", false},
		{1, "1", "9,9V9", false},
		{1, "1", "99V9", false},
		{1, "$12", "$99V9", true},
		{1, "17", "XXXX", false},
		{1, "017", "XXXX", false},
		{1, "17", "0XXXX", false},
		{1, "0017", "FMXXXX", false},
		{1, "17", "FMXXXX", false},
		{1, "00017", "XXXX", false},
		{1, "12", "LXXXX", true},
		{1, "$12", "S$XXXX", true},
		{1, "1", "FMLBUSXXXX", true},
		{1, "1", "RN", false},
		{1, "11", "RN", false},
		{1, "12", "LRN", true},
		{1, "$12", "S$RN", true},
		{1, "1", "FMLBUSRN", true},
		{1, "1234", "TM", false},
		{1, "1234", "TM9", false},
		{1, "12", "LTM", true},
		{1, "$12", "S$TM", true},
		{1, "1", "FMLBUSTM", true},
		{1, "1234", "tme", false},
		{1, "12", "STME", false},
		{1, "12", "FMSTME", false},
		{1, "12", "TMES", true},
		{1, "1234", "tmetme", true},
		{1, "12", "LTME", true},
		{1, "$12", "S$TME", true},
		{1, "1", "FMLBUSTME", true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			tm, err := ToCharByStr(test.ch, test.format)
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

func TestSuiteToCharByNum(t *testing.T) {
	tests := []struct {
		i         int
		f         float64
		format    string
		exception bool
	}{
		{1, 3450, "999,99", false},
		{1, 3450, "999G99", false},
		{1, 3450, "99G9G99", false},
		{1, 3450, "9,9,9,99", false},
		{1, 3450, "9,9G9G99", true},
		{1, 3450, "99999", true},
		{1, -124548, "S999999", true},
		{1, 3450, "999,99", true},
		{1, 3450, "99", true},
		{1, 3450, "9G9", true},
		{1, 124548, "99G9G999", true},
		{1, 1245.48, "99G9G9D99", true},
		{1, 23.54, "99.99", false},
		{1, 23.54, "99D99", false},
		{1, 23.54, "99D99", false},
		{1, .54, "D99", false},
		{1, 23.54, "99D99", true},
		{1, 23.54, "99D9D9", true},
		{1, 2.54, "D999", true},
		{1, 1, "0", true},
		{1, 1322526, "0099000", false},
		{1, 1322526, "0099000", false},
		{1, 1322526, "9999099", false},
		{1, 1322526, "99990999999", false},
		{1, 1322526, "999090909999999009", true},
		{1, 1322526, "99909090999009", true},
		{1, 1322526, "000000000000000", true},
		{1, 1322526, "90009000009", true},
		{1, 1322526, "900000000009", true},
		{1, 1322526, "00000000009", true},
		{1, 1322526, "99990999999", false},
		{1, 1322526, "9999", true},
		{1, 123.45, "B999.99", false},
		{1, 12345.678, "$999999.999", false},
		{1, 123.45, "L999.99", false},
		{1, 123.45, "U999.99", false},
		{1, 123.45, "U99999.990", false},
		{1, 356, "C999", false},
		{1, 1258, "9999C", false},
		{1, 1258, "999999C99", false},
		{1, 1258, "9999U", false},
		{1, 1258, "9999B", false},
		{1, 1258, "9999L", false},
		{1, 1258, "999B9L", false},
		{1, 1258, "99B99", false},
		{1, 1258.345, "9999L99", false},
		{1, 1258.235, "9999U99", false},
		{1, 1258.235, "9999C9999", false},
		{1, 1258.235, "9999B9999", false},
		{1, 1258.235, "9999$9999", false},
		{1, 1258, "99U99U", true},
		{1, 1258, "99B99B", true},
		{1, 1258, "99B9B9L", true},
		{1, 1258, "9999CL", true},
		{1, 1258, "9999C99C99", true},
		{1, 485, "FM999MI", false},
		{1, 485, "999MI", true},
		{1, -485, "999MI", true},
		{1, -485, "999MI", true},
		{1, 485, "999MI", true},
		{1, 485, "999PR", false},
		{1, -485, "999PR", false},
		{1, 485, "PR999", true},
		{1, -485, "PR999", true},
		{1, -258, "S999", false},
		{1, 258, "S999", false},
		{1, -258, "999S", false},
		{1, 258, "999S", false},
		{1, -258, "9S99", true},
		{1, -258, "9SS99", true},
		{1, -258, "S9SS99", true},
		{1, -258, "S999", false},
		{1, +258, "S999", true},
		{1, -258, "999PRS", true},
		{1, 1e+2, "9999999EEEE", false},
		{1, -1e-2, "9999999EEEE", false},
		{1, -1e-2, "9999999EEEE", true},
		{1, 12, "99V9", false},
		{1, 12, "99V999", false},
		{1, 12.45, "99V9", false},
		{1, 1, "9,9V9", false},
		{1, 1, "9V9", false},
		{1, 1, "99V99", false},
		{1, 12, "FM99V9", false},
		{1, 12, "FM99V9MI", false},
		{1, 12, "FM99V9S", false},
		{1, 12, "FM99V9PR", false},
		{1, 12, "99V9PR", false},
		{1, 12, "S99V9", false},
		{1, 12, "99V9MI", false},
		{1, 12, "L99V9", false},
		{1, 12, "S99V9MI", true},
		{1, 1, "9.9V9", true},
		{1, 17, "XXXX", false},
		{1, 017, "XXXX", false},
		{1, 17, "0XXXX", false},
		{1, 0017, "FMXXXX", false},
		{1, 17, "FMXXXX", false},
		{1, 00017, "XXXX", false},
		{1, 7, "XXXX9", true},
		{1, 6, "9XXXX", true},
		{1, 6, "X9", true},
		{1, 6, "9X", true},
		{1, 6, "LX", true},
		{1, 6, "XS", true},
		{1, 12, "FMX", true},
		{1, 12, "FMXXXXMI", true},
		{1, 12, "FMXXXXS", true},
		{1, 12, "XXXXPR", true},
		{1, 12, "SXXXX", true},
		{1, 12, "XXXXMI", true},
		{1, 1, "SX", true},
		{1, 1, "FMLBUSXXXXMIPR", true},
		{1, 11, "RN", false},
		{1, 1, "RN", false},
		{1, 14825, "RN", false},
		{1, 1485, "rn", false},
		{1, 12, "FMRN", false},
		{1, 12, "FMrn", false},
		{1, 12, "SRN", false},
		{1, 1, "SRN", false},
		{1, -12, "rn", true},
		{1, 1485, "99999RN", true},
		{1, 1485, "LRN", true},
		{1, 7, "RN9", true},
		{1, 6, "9RN", true},
		{1, 6, "RN9", true},
		{1, 6, "9RN", true},
		{1, 6, "LRN", true},
		{1, 6, "rnS", true},
		{1, 12, "FMRNS", true},
		{1, 12, "RNPR", true},
		{1, 12, "RNMI", true},
		{1, 1, "RNPR", true},
		{1, 1, "RNMI", true},
		{1, -1, "RNMI", true},
		{1, 1, "FMLBUSRNMIPR", true},
		{1, 123, "TM", false},
		{1, 1234, "TM9", false},
		{1, 12, "STM", false},
		{1, 12, "BTM", true},
		{1, 1234, "TM8", true},
		{1, 1234, "TM99", true},
		{1, 1234, "9TM9", true},
		{1, 1234, "LTM9", true},
		{1, 1234, "0TM9", true},
		{1, 1234, "$TM9", true},
		{1, 1234, "UTM9", true},
		{1, 12, "FMTM", true},
		{1, 12, "FMtm", true},
		{1, 12, "FMTM9", true},
		{1, 12, "FMtM9", true},
		{1, 12, "FMTMS", true},
		{1, 12, "TMPR", true},
		{1, 12, "TMMI", true},
		{1, 1, "STM", true},
		{1, 1, "TMTM", true},
		{1, 1, "FMLBUSTMMIPR", true},
		{1, 1234, "TME", false},
		{1, 1234, "tmE", false},
		{1, 1234, "TMe", false},
		{1, 1234, "tme", false},
		{1, 12, "STME", false},
		{1, 1, "STME", false},
		{1, 1234, "TMETME", true},
		{1, 1234, "TMETM", true},
		{1, 1234, "TM9e", true},
		{1, 1234, "tm9e", true},
		{1, 1234, "TM8", true},
		{1, 1234, "TM99", true},
		{1, 1234, "9TM9", true},
		{1, 1234, "LTM9", true},
		{1, 1234, "0TM9", true},
		{1, 1234, "$TM9", true},
		{1, 1234, "UTM9", true},
		{1, 12, "FMTME", true},
		{1, 12, "FMtme", true},
		{1, 12, "FMTM9", true},
		{1, 12, "FMTME", true},
		{1, 12, "FMTMES", true},
		{1, 12, "TMEPR", true},
		{1, 1, "FMLBUSTMEMIPR", true},
		{1, 12, "TMEMI", true},
		{1, 1, "STMETME", true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			tm, err := ToCharByNum(test.f, test.format)
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

func TestSuiteToCharByTime(t *testing.T) {
	tests := []struct {
		i         int
		dt        string
		format    string
		exception bool
	}{
		{1, "2023-10-29 01:30:56.321654789", "YYYY", false},
		{1, "2023-10-29 01:30:56.321654789", "DD", false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			dt, err := parseTime(test.dt)
			if err != nil {
				panic(err)
			}
			tm, err := ToCharByDatetime(dt, test.format)
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

var timeFormat = "2006-01-02 15:04:05.999999999"

func parseTime(timeStr string) (time.Time, error) {
	fmt.Println("timeStr:", timeStr)
	t, err := time.Parse(timeFormat, timeStr)
	return t, err
}
