package time_demo

import (
	"bytes"
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

func TestFormat3(t *testing.T) {
	dch := "  公元2,023://sss    2008         -;:        05-20.333"

	f := []string{"AD", "Y,YYY", "quote", "YYYY", "MM", "DD", "FF3"}
	result := []string{}
	tmp := bytes.Buffer{}

	di := 0
	fi := 0
	for ; fi < len(f); fi++ {
		switch f[fi] {
		case "AD":
			for ; di < len(dch); di++ {
				if dch[di] == ' ' ||
					dch[di] == '-' ||
					dch[di] == ':' ||
					dch[di] == ',' ||
					dch[di] == '.' ||
					dch[di] == '/' ||
					dch[di] == ';' {
				} else {
					println("debug")
					for j := 0; j < len("公元"); j++ {
						tmp.WriteByte(dch[di])
						di++
					}
					result = append(result, tmp.String())
					tmp.Reset()
					break
				}
			}
		case "Y,YYY":
			for ; di < len(dch); di++ {
				if dch[di] == ' ' ||
					dch[di] == '-' ||
					dch[di] == ':' ||
					dch[di] == ',' ||
					dch[di] == '.' ||
					dch[di] == '/' ||
					dch[di] == ';' {
				} else {
					for j := 0; j < 5; j++ {
						tmp.WriteByte(dch[di])
						di++
					}
					result = append(result, tmp.String())
					tmp.Reset()
					break
				}
			}

		case "YYYY":
			for ; di < len(dch); di++ {
				if dch[di] == ' ' ||
					dch[di] == '-' ||
					dch[di] == ':' ||
					dch[di] == ',' ||
					dch[di] == '.' ||
					dch[di] == '/' ||
					dch[di] == ';' {
				} else {
					for j := 0; j < 4; j++ {
						tmp.WriteByte(dch[di])
						di++
					}
					result = append(result, tmp.String())
					tmp.Reset()
					break
				}
			}
		case "MM":
			for ; di < len(dch); di++ {
				if dch[di] == ' ' ||
					dch[di] == '-' ||
					dch[di] == ':' ||
					dch[di] == ',' ||
					dch[di] == '.' ||
					dch[di] == '/' ||
					dch[di] == ';' {
				} else {
					for j := 0; j < 2; j++ {
						tmp.WriteByte(dch[di])
						di++
					}
					result = append(result, tmp.String())
					tmp.Reset()
					break
				}
			}
		case "DD":
			for ; di < len(dch); di++ {
				if dch[di] == ' ' ||
					dch[di] == '-' ||
					dch[di] == ':' ||
					dch[di] == ',' ||
					dch[di] == '.' ||
					dch[di] == '/' ||
					dch[di] == ';' {
				} else {
					for j := 0; j < 2; j++ {
						tmp.WriteByte(dch[di])
						di++
					}
					result = append(result, tmp.String())
					tmp.Reset()
					break
				}
			}
		case "FF3":
			for ; di < len(dch); di++ {
				if dch[di] == ' ' ||
					dch[di] == '-' ||
					dch[di] == ':' ||
					dch[di] == ',' ||
					dch[di] == '.' ||
					dch[di] == '/' ||
					dch[di] == ';' {
				} else {
					for j := 0; j < 3; j++ {
						tmp.WriteByte(dch[di])
						di++
					}

					result = append(result, tmp.String())
					tmp.Reset()
					break
				}
			}
		case "quote":
			for ; di < len(dch); di++ {
				if dch[di] == ' ' ||
					dch[di] == '-' ||
					dch[di] == ':' ||
					dch[di] == ',' ||
					dch[di] == '.' ||
					dch[di] == '/' ||
					dch[di] == ';' {
				} else {
					for j := 0; j < 3; j++ {
						tmp.WriteByte(dch[di])
						di++
					}
					result = append(result, tmp.String())
					tmp.Reset()
					break
				}
			}
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
