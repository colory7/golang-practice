package time_demo

import (
	"bytes"
	"errors"
	"testing"
)

func TestFormat2(t *testing.T) {
	s := "  公元2,023://sss    2008         -;:        05-20"
	println(s)

	arr := []string{}

	tmp := bytes.Buffer{}
	for i := 0; i < len(s); i++ {

		if s[i] == ' ' ||
			s[i] == '-' ||
			s[i] == ':' ||
			//s[i] == ',' ||
			s[i] == '.' ||
			s[i] == '/' ||
			s[i] == ';' {
			if tmp.Len() > 0 {
				arr = append(arr, tmp.String())
				tmp.Reset()
			}
		} else {
			tmp.WriteByte(s[i])
		}
	}

	arr = append(arr, tmp.String())

	println("====")
	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}

	println(len(arr))
}

const (
	flag_year   = 1
	flag_month  = 1 << 1
	flag_day    = 1 << 2
	flag_hour   = 1 << 3
	flag_minute = 1 << 4
	flag_second = 1 << 5
	flag_nansec = 1 << 6
	flag_tz     = 1 << 7
)

func TestFXFormat(t *testing.T) {
	//dch := "  公元2,023://sss    8         -;:        05-20 2:3:4.333"
	dch := "公元2,023sss2008-05-20 12:53:64.333"
	dlen := len(dch)

	f := []string{"AD", "Y,YYY", "quote", "YYYY", "MINUS", "MM", "MINUS", "DD", "SPACE", "HH", "COLON", "MI", "COLON", "SS", "DOT", "FF3"}
	result := []string{}

	di := 0
	fi := 0
	quotedLen := len("sss")
	for ; fi < len(f); fi++ {
		switch f[fi] {
		case "AD":
			d, err := parseDchFM(&dch, dlen, &di, len("公元"))
			if err != nil {
				panic(err)
			}
			result = append(result, d)
		case "Y,YYY":
			d, err := parseDchFM(&dch, dlen, &di, 5)
			if err != nil {
				panic(err)
			}
			result = append(result, d)
		case "YYYY":
			d, err := parseDchFM(&dch, dlen, &di, 4)
			if err != nil {
				panic(err)
			}
			println(d)

			result = append(result, d)
		case "MM":
			d, err := parseDchFM(&dch, dlen, &di, 2)
			if err != nil {
				panic(err)
			}
			result = append(result, d)
		case "DD":
			d, err := parseDchFM(&dch, dlen, &di, 2)
			if err != nil {
				panic(err)
			}
			result = append(result, d)
		case "FF3":
			d, err := parseDchFM(&dch, dlen, &di, 3)
			if err != nil {
				panic(err)
			}
			result = append(result, d)
		case "HH":
			d, err := parseDchFM(&dch, dlen, &di, 2)
			if err != nil {
				panic(err)
			}
			result = append(result, d)
		case "MI":
			d, err := parseDchFX(&dch, dlen, &di, 2)
			if err != nil {
				panic(err)
			}
			result = append(result, d)
		case "SS":
			d, err := parseDchFM(&dch, dlen, &di, 2)
			if err != nil {
				panic(err)
			}
			result = append(result, d)
		case "quote":
			d, err := parseDchFM(&dch, dlen, &di, quotedLen)
			if err != nil {
				panic(err)
			}
			result = append(result, d)
		case "MINUS":
			d, err := parseDchFM(&dch, dlen, &di, 1)
			if err != nil {
				panic(err)
			}
			result = append(result, d)
		case "COLON":
			d, err := parseDchFM(&dch, dlen, &di, 1)
			if err != nil {
				panic(err)
			}
			result = append(result, d)
		case "DOT":
			d, err := parseDchFM(&dch, dlen, &di, 1)
			if err != nil {
				panic(err)
			}
			result = append(result, d)
		case "SPACE":
			d, err := parseDchFM(&dch, dlen, &di, 1)
			if err != nil {
				panic(err)
			}
			result = append(result, d)
		default:
			panic("never reach")
		}
	}

	for i := 0; i < len(result); i++ {
		println(result[i])
	}

	if len(f) != len(result) {
		panic("格式 未匹配完")
	}

	if di != len(dch) {
		panic("参数 未匹配完")
	}
}

func TestFMFormat(t *testing.T) {
	dch := "  公元2,023://sss    8         -;:        05-20 2:3:4.333"
	dlen := len(dch)

	f := []string{"AD", "Y,YYY", "quote", "YYYY", "MM", "DD", "HH", "MI", "SS", "FF3"}
	result := []string{}

	di := 0
	fi := 0
	quotedLen := len("sss")
	for ; fi < len(f); fi++ {
		switch f[fi] {
		case "AD":
			d, err := parseDchFX(&dch, dlen, &di, len("公元"))
			if err != nil {
				panic(err)
			}
			result = append(result, d)
		case "Y,YYY":
			d, err := parseDchFX(&dch, dlen, &di, 5)
			if err != nil {
				panic(err)
			}
			result = append(result, d)
		case "YYYY":
			d, err := parseDchFX(&dch, dlen, &di, 4)
			if err != nil {
				panic(err)
			}
			println(d)

			result = append(result, d)
		case "MM":
			d, err := parseDchFX(&dch, dlen, &di, 2)
			if err != nil {
				panic(err)
			}
			result = append(result, d)
		case "DD":
			d, err := parseDchFX(&dch, dlen, &di, 2)
			if err != nil {
				panic(err)
			}
			result = append(result, d)
		case "FF3":
			d, err := parseDchFX(&dch, dlen, &di, 3)
			if err != nil {
				panic(err)
			}
			result = append(result, d)
		case "HH":
			d, err := parseDchFX(&dch, dlen, &di, 2)
			if err != nil {
				panic(err)
			}
			result = append(result, d)
		case "MI":
			d, err := parseDchFX(&dch, dlen, &di, 2)
			if err != nil {
				panic(err)
			}
			result = append(result, d)
		case "SS":
			d, err := parseDchFX(&dch, dlen, &di, 2)
			if err != nil {
				panic(err)
			}
			result = append(result, d)
		case "quote":
			d, err := parseDchFX(&dch, dlen, &di, quotedLen)
			if err != nil {
				panic(err)
			}
			result = append(result, d)
		default:
			panic("never reach")
		}
	}

	for i := 0; i < len(result); i++ {
		println(result[i])
	}

	if len(f) != len(result) {
		panic("格式 未匹配完")
	}

	if di != len(dch) {
		panic("参数 未匹配完")
	}
}

func parseDchFM(dch *string, dlen int, di *int, size int) (string, error) {
	start := *di
	*di += size
	if *di > dlen {
		return "", errors.New("格式长度不匹配")
	}
	return (*dch)[start:*di], nil
}

func parseDchFX(dch *string, dlen int, di *int, size int) (string, error) {
	tmp := bytes.Buffer{}
	for ; *di < dlen; *di++ {
		if (*dch)[*di] == ' ' ||
			(*dch)[*di] == '-' ||
			(*dch)[*di] == ':' ||
			(*dch)[*di] == ',' ||
			(*dch)[*di] == '.' ||
			(*dch)[*di] == '/' ||
			(*dch)[*di] == ';' {
		} else {
			for j := 0; j < size; j++ {
				tmp.WriteByte((*dch)[*di])
				*di++
			}
			return tmp.String(), nil
		}
	}

	return "", errors.New("未找到格式对应的匹配项")
}

func TestZone(t *testing.T) {

}
