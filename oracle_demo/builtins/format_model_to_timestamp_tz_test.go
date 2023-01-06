package builtins

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestToTimestampTZPositive(t *testing.T) {
	dch := "2023-10-29 01:30:56"
	format := "YYYY-MM-DD HH:MI:SS"
	tm, err := ToTimestampTZ(dch, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(tm)
}

func TestToTimestampTZNegative(t *testing.T) {
	dch := "2023-10-32"
	format := "YYYY-MM-DD HH:MI:SS"
	tm, err := ToTimestampTZ(dch, format)
	assert.Error(t, err)
	fmt.Println(tm)
}

func TestSuiteToTimestampTZ(t *testing.T) {
	tests := []struct {
		i         int
		dch       string
		format    string
		exception bool
	}{
		{1, "2023-10-31", "YYYY-MM-DD HH:MI:SS", true},
		{2, "2023-10-29 01:30:56", "YYYY-MM-DD HH:MI:SS", false},
		{3, "2023-10-29", "YYYY-MM-DD", false},
		{4, "2023////10//29", "YYYY////MM//DD", false},
		{5, "2023,,,,,10//29", "YYYY,,,,,MM//DD", false},
		{6, "2023....10..29", "YYYY....MM..DD", false},
		{7, "2023....10..29", "YYYY;;;;MM..DD", false},
		{8, "2023....10..29", "YYYY::::MM..DD", false},
		{9, "2023abcd10..29", "YYYY\"abcd\"MM,,DD", false},
		{10, "2023abcd10..29", "YYYY\"abcdMM,,DD", true},
		{11, "2023abcdef10..29", "YYYY\"abcd\"\"ef\"MM,,DD", false},
		{12, "2023abcdef10..29", "YYYY\"abcd\"\"efMM,,DD", true},
		{13, "2008\"xxxx\"05,,20", "YYYY\"abcd\"MM,,DD", true}, // 内容不匹配
		{14, "2008 05 20", "YYYY\"abcd\"\"efMM,,DD", true},
		{15, "2008 05 20", "YYYY MM DD", false},
		{16, "2008-05-20", "YYYY-MM-DD", false},
		{17, "09:26:50", "HH:MI:SS", false},
		{18, "23--26-50", "HH24--MI-SS", false},
		{19, "2023-01-04 09:26:50", "YYYY-MM-DD HH:MI:SS", false},
		{20, "2023-1月-04 09:26:50", "YYYY-MON-DD HH:MI:SS", false},
		{21, "2023-1月-04 09:26:50", "FMYYYY-MON-DD HH:MI:SS", false},
		{22, "2023-1月-04 09:26:50", "FXYYYY-MON-DD HH:MI:SS", false},
		{23, "2023-1月-04 09:26:50", "FXYYYY-MON-DD/HH:MI:SS", true},
		{24, "2023-1月-04 09:26:50", "YYYY-MON-DD/HH:MI:SS", false},

		{25, "2023-01-04 09:26:50.231456897", "YYYY-MM-DD HH:MI:SS.FF9", false},
		{26, "2023-01-04 09:26:50.2314568", "YYYY-MM-DD HH:MI:SS.FF7", false},
		{27, "2023-01-04 09:26:50.23145", "YYYY-MM-DD HH:MI:SS.FF5", false},
		{28, "2023-01-04 09:26:50.231", "YYYY-MM-DD HH:MI:SS.FF5", false},

		{29, "13:26:50", "HH:MI:SS", true},
		{30, "25:26:50", "HH24:MI:SS", true},
		{31, "2023-01-04 09:26:50.231456897", "YYYY-MM-DD HH:MI:SS.FF7", true},
		{32, "2023-01-04 09:26:50.231456897", "YYYY-MM-DD HH:MI:SS.FF5", true},
		{33, "2023-01-04 09:26:50. 231", "YYYY-MM-DD HH:MI:SS.FF5", true},
		{34, "2023-13-04 09:26:50.231456897", "YYYY-MM-DD HH:MI:SS.FF9", true},
		{35, "2023-01-32 09:26:50.231456897", "YYYY-MM-DD HH:MI:SS.FF9", true},
		{36, "2023-01-04 26:26:50.231456897", "YYYY-MM-DD HH:MI:SS.FF9", true},
		{37, "2023-01-04 09:68:50.231456897", "YYYY-MM-DD HH:MI:SS.FF9", true},
		{38, "2023-01-04 09:26:73.231456897", "YYYY-MM-DD HH:MI:SS.FF9", true},
		{39, "2023-01-04 09:26:50.2314568979865", "YYYY-MM-DD HH:MI:SS.FF9", true},
		{39, "20223-01-04 09:26:50.2314568979", "YYYY-MM-DD HH:MI:SS.FF9", true},

		{40, "2023-01-04 09:26:50.231456897 09", "YYYY-MM-DD HH:MI:SS.FF9 TZH", false},
		{41, "2023-01-04 09:26:50.231456897 09:02", "YYYY-MM-DD HH:MI:SS.FF9 TZH:TZM", false},
		{42, "2023-01-04 09:26:50.231456897 +09:02", "YYYY-MM-DD HH:MI:SS.FF9 TZH:TZM", false},
		{43, "2023-01-04 09:26:50.231456897 -07:05", "YYYY-MM-DD HH:MI:SS.FF9 TZH:TZM", false},
		{44, "2023-01-04 09:26:50.231456897 Asia/Shanghai", "YYYY-MM-DD HH:MI:SS.FF9 TZR", false},
		{45, "2023-01-04 09:26:50.231456897 Asia/Urumqi", "YYYY-MM-DD HH:MI:SS.FF9 TZR", false},
		{46, "2023-01-04 09:26:50.231456897 -09:03", "YYYY-MM-DD HH:MI:SS.FF9 TZR", false},
		{47, "2023-01-04 09:26:50.231456897 CST", "YYYY-MM-DD HH:MI:SS.FF9 TZR", false},
		{48, "2023-01-04 09:26:50.231456897 GMT", "YYYY-MM-DD HH:MI:SS.FF9 TZR", false},
		{49, "2023-01-04 09:26:50.231456897 UTC", "YYYY-MM-DD HH:MI:SS.FF9 TZR", false},
		{50, "2023-01-04 09:26:50.231456897 PST", "YYYY-MM-DD HH:MI:SS.FF9 TZR", false},
		{51, "2023-01-04 09:26:50.231456897 EDT", "YYYY-MM-DD HH:MI:SS.FF9 TZR", true},
		{52, "2023-01-04 09:26:50.231456897 aaa", "YYYY-MM-DD HH:MI:SS.FF9 TZR", true},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.i), func(t *testing.T) {
			tm, err := ToTimestampTZ(test.dch, test.format)
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

func TestZone(t *testing.T) {
	var cstSh, _ = time.LoadLocation("Asia/Urumqi") //上海
	fmt.Println("SH : ", time.Now().In(cstSh))

	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	fmt.Println("SH : ", time.Now().In(cstZone))

	var cstZone2 = time.FixedZone("PST", 8*3600) // 东八
	fmt.Println("SH : ", time.Now().In(cstZone2))

	var cstZone3 = time.FixedZone("aaa", 8*3600) // 东八
	fmt.Println("SH : ", time.Now().In(cstZone3))

	var cstZone4 = time.FixedZone(" ", 8*3600) // 东八
	fmt.Println("SH : ", time.Now().In(cstZone4))

	fmt.Println(time.Now().Location().String())
	fmt.Println(time.Now().Zone())

	fmt.Println(time.UTC.String())

}
