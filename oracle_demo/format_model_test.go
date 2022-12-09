package oracle_demo

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
	"unicode/utf8"
)

const skipCharSize = 32

// 格式部分不匹配，报错
const dch_fmt_part_err = "Number Format error, some formats do not match near "
const num_fmt_part_err = "Datetime Format error, some formats do not match near "

// 非法字符,超出格式关键词范围
const out_keyword_range_err = "Illegal character, not in the range of Format Model keyword"

// 非法字符,超出ASCII[32-126]字符范围
const out_ascii_range_err = "Illegal character, not in ASCII [32-126] character range"

const invalid_num_err = "invalid number"

// ASCII 32-126
var formatChar = [utf8.RuneSelf]byte{
	' ',
	'!',
	'"',
	'#',
	'$',
	'%',
	'&',
	'\'',
	'(',
	')',
	'*',
	'+',
	',',
	'-',
	'.',
	'/',
	'0',
	'1',
	'2',
	'3',
	'4',
	'5',
	'6',
	'7',
	'8',
	'9',
	':',
	';',
	'<',
	'=',
	'>',
	'?',
	'@',
	'A',
	'B',
	'C',
	'D',
	'E',
	'F',
	'G',
	'H',
	'I',
	'J',
	'K',
	'L',
	'M',
	'N',
	'O',
	'P',
	'Q',
	'R',
	'S',
	'T',
	'U',
	'V',
	'W',
	'X',
	'Y',
	'Z',
	'[',
	'\\',
	']',
	'^',
	'_',
	'`',
	'a',
	'b',
	'c',
	'd',
	'e',
	'f',
	'g',
	'h',
	'i',
	'j',
	'k',
	'l',
	'm',
	'n',
	'o',
	'p',
	'q',
	'r',
	's',
	't',
	'u',
	'v',
	'w',
	'x',
	'y',
	'z',
	'{',
	'|',
	'}',
	'~',
}

const (
	one    = '1'
	two    = '2'
	plus   = '+'
	minus  = '-'
	e      = 'e'
	E      = 'E'
	dollar = '$'
)

const (
	// Number Format Model Keyword
	NUM_COMMA  = ","
	NUM_DEC    = "."
	NUM_DOLLAR = "$"
	NUM_0      = "0"
	NUM_9      = "9"
	NUM_B      = "B"
	NUM_C      = "C"
	NUM_D      = "D"
	NUM_E      = "EEEE"
	NUM_FM     = "FM"
	NUM_G      = "G"
	NUM_L      = "L"
	NUM_MI     = "MI"
	NUM_PR     = "PR"
	NUM_RN     = "RN"
	NUM_S      = "S"
	NUM_TM     = "TM"
	NUM_TM9    = "TM9"
	NUM_TME    = "TME"
	NUM_U      = "U"
	NUM_V      = "V"
	NUM_X      = "X"

	// Datetime Format Model Keyword
	DCH_A_D   = "A.D."
	DCH_A_M   = "A.M."
	DCH_AD    = "AD"
	DCH_AM    = "AM"
	DCH_B_C   = "B.C."
	DCH_BC    = "BC"
	DCH_CC    = "CC"
	DCH_DAY   = "DAY"
	DCH_DDD   = "DDD"
	DCH_DD    = "DD"
	DCH_DY    = "DY"
	DCH_D     = "D"
	DCH_FF1   = "FF1"
	DCH_FF2   = "FF2"
	DCH_FF3   = "FF3"
	DCH_FF4   = "FF4"
	DCH_FF5   = "FF5"
	DCH_FF6   = "FF6"
	DCH_FF7   = "FF7"
	DCH_FF8   = "FF8"
	DCH_FF9   = "FF9"
	DCH_FX    = "FX"
	DCH_HH24  = "HH24"
	DCH_HH12  = "HH12"
	DCH_HH    = "HH"
	DCH_IDDD  = "IDDD"
	DCH_ID    = "ID"
	DCH_IW    = "IW"
	DCH_IYYY  = "IYYY"
	DCH_IYY   = "IYY"
	DCH_IY    = "IY"
	DCH_I     = "I"
	DCH_J     = "J"
	DCH_MI    = "MI"
	DCH_MM    = "MM"
	DCH_MONTH = "MONTH"
	DCH_MON   = "MON"
	DCH_MS    = "MS"
	DCH_OF    = "OF"
	DCH_P_M   = "P.M."
	DCH_PM    = "PM"
	DCH_Q     = "Q"
	DCH_RM    = "RM"
	DCH_SSSSS = "SSSSS"
	DCH_SSSS  = "SSSS"
	DCH_SS    = "SS"
	DCH_TZH   = "TZH"
	DCH_TZM   = "TZM"
	DCH_TZ    = "TZ"
	DCH_US    = "US"
	DCH_WW    = "WW"
	DCH_W     = "W"
	DCH_Y_YYY = "YYYY"
	DCH_YYYY  = "YYYY"
	DCH_YYY   = "YYY"
	DCH_YY    = "YY"
	DCH_Y     = "Y"
)

//type keyword struct {
//	name string
//	len uint8
//	id int
//}

var numKeywords map[string]uint8
var dchKeywords map[string]uint8

type keyword struct {
	key string
	len uint8
}

func init() {
	numKeywords = map[string]uint8{
		",":    1,
		".":    1,
		"0":    1,
		"9":    1,
		"B":    1,
		"C":    1,
		"D":    1,
		"EEEE": 4,
		"FM":   2,
		"G":    1,
		"L":    1,
		"MI":   2,
		"PL":   2,
		"PR":   2,
		"RN":   2,
		"SG":   2,
		"SP":   2,
		"S":    1,
		"TH":   2,
		"V":    1,
		"b":    1,
		"c":    1,
		"d":    1,
		"eeee": 4,
		"fm":   1,
		"g":    1,
		"l":    1,
		"mi":   2,
		"pl":   2,
		"pr":   2,
		"rn":   2,
		"sg":   2,
		"sp":   2,
		"s":    1,
		"th":   2,
		"v":    1,
	}

	dchKeywords = map[string]uint8{
		"DCH_A_D":   4,
		"DCH_A_M":   4,
		"DCH_AD":    2,
		"DCH_AM":    2,
		"DCH_B_C":   4,
		"DCH_BC":    2,
		"DCH_CC":    2,
		"DCH_DAY":   3,
		"DCH_DDD":   3,
		"DCH_DD":    2,
		"DCH_DY":    2,
		"DCH_D":     1,
		"DCH_FF1":   3,
		"DCH_FF2":   3,
		"DCH_FF3":   3,
		"DCH_FF4":   3,
		"DCH_FF5":   3,
		"DCH_FF6":   3,
		"DCH_FF7":   3,
		"DCH_FF8":   3,
		"DCH_FF9":   3,
		"DCH_FX":    2,
		"DCH_HH24":  4,
		"DCH_HH12":  4,
		"DCH_HH":    2,
		"DCH_IDDD":  4,
		"DCH_ID":    2,
		"DCH_IW":    2,
		"DCH_IYYY":  4,
		"DCH_IYY":   3,
		"DCH_IY":    2,
		"DCH_I":     1,
		"DCH_J":     1,
		"DCH_MI":    2,
		"DCH_MM":    2,
		"DCH_MONTH": 5,
		"DCH_MON":   3,
		"DCH_MS":    2,
		"DCH_OF":    2,
		"DCH_P_M":   4,
		"DCH_PM":    2,
		"DCH_Q":     1,
		"DCH_RM":    2,
		"DCH_SSSSS": 5,
		"DCH_SSSS":  4,
		"DCH_SS":    2,
		"DCH_TZH":   3,
		"DCH_TZM":   3,
		"DCH_TZ":    2,
		"DCH_US":    2,
		"DCH_WW":    2,
		"DCH_W":     1,
		"DCH_Y_YYY": 5,
		"DCH_YYYY":  4,
		"DCH_YYY":   3,
		"DCH_YY":    2,
		"DCH_Y":     1,
	}
}

func TestAscii(t *testing.T) {
	theme := "狙击 start"
	for i := 0; i < len(theme); i++ {
		fmt.Printf("ascii:%c %d\n", theme[i], theme[i])
	}
}

func TestAscii2(t *testing.T) {
	theme := "狙击 start"
	for i := 0; i < len(theme); i++ {
		f := theme[i]
		if f >= 32 && f <= 127 {
			fmt.Println((string)(f))
		}
	}
}

//NumberFormatModel
//DatetimeFormatModel

func TestMatchNumberFormatModel(t *testing.T) {
	// 1.找到模式,最大匹配
	// 2.根据模式替换字符串

	//param:="111"
	format := "999"
	parseNumFormat(format)
}

// 解析数值格式
func parseNumFormat(format string) []string {
	var keywordGroup = make([]string, 4)

	//is_decimal := false
	//is_ldecimal := false
	//is_zero := false
	//is_blank := false
	//is_bracket := false
	//is_minus := false
	//is_lsign := false
	//is_plus := false
	//is_multi := false
	is_fillmode := false
	is_roman := false
	is_eeee := false

	// 格式字节长度
	formatLen := len(format)

	// 半角逗号下标
	commaIndex := -1
	// 半角句号的下标
	decIndex := -1
	BIndex := -1
	CIndex := -1
	DIndex := -1
	GIndex := -1
	LIndex := -1
	MiIndex := -1
	PrIndex := -1
	SIndex := -1
	UIndex := -1
	VIndex := -1
	TMIndex := -1

	// 最后一个关键词的索引
	lastFormatIndex := formatLen - 1

	isRerun := false
	var c byte
	for i := 0; i < formatLen; {
		// 截取一个字符

		if !isRerun {
			c = format[i]
		} else {
			isRerun = false
		}

		if c >= 32 && c <= 127 {
			fmt.Println((string)(c))

			// 这里设置为不区分大小写。NB: Oracle和Postgresql中为区分大小写
			ToUpper(&c)

			// 匹配关键词并存储
			switch c {
			case ',':
				if commaIndex == -1 {
					// 逗号不能出现在数字最右边
					if i == 0 {
						panic(errors.New("不能以逗号开头"))
					} else if i == lastFormatIndex {
						panic(errors.New("逗号不能出现在数字最右边"))
					}
					// 逗号不能出现在点号的右边
					if decIndex != -1 && decIndex < i {
						panic(errors.New("逗号不能出现在点号的右边"))
					}

					keywordGroup = append(keywordGroup, NUM_COMMA)
					commaIndex = i
					break
				} else {
					panic(errors.New("格式错误，存在多个格式符号 ,"))
				}
			case '.':
				// 只能有1个点号
				if decIndex == -1 {
					decIndex = i
				} else {
					panic(errors.New("只能有1个 ."))
				}
				keywordGroup = append(keywordGroup, NUM_DEC)
				break
			case '0':
				keywordGroup = append(keywordGroup, NUM_0)
				break
			case '9':
				keywordGroup = append(keywordGroup, NUM_9)
				break
			case 'B':
				if BIndex == -1 {
					keywordGroup = append(keywordGroup, NUM_B)
					BIndex = i
					break
				} else {
					panic(errors.New("只能有1个 B"))
				}
			case 'C':
				if CIndex != -1 {
					panic(errors.New("只能有1个 C"))
				} else if lastFormatIndex != i && 0 != i {
					panic(errors.New("C 只能在开头或者结尾"))
				}
				keywordGroup = append(keywordGroup, NUM_C)
				CIndex = i
				break
			case 'D':
				// oracle中是变量可以设置为其他值 FIXME
				if DIndex == -1 {
					keywordGroup = append(keywordGroup, NUM_D)
					DIndex = i
					c = NLS_NUMERIC_CHARACTERS[0]
					isRerun = true
					continue
				} else {
					panic(errors.New("只能有1个 D"))
				}
			case 'E':
				if !is_eeee {
					j := i + 4
					followingChars := format[i:j]
					if "EEE" == followingChars {
						keywordGroup = append(keywordGroup, NUM_E)
						i = j
						is_eeee = true
						break
					} else {
						panic(errors.New(num_fmt_part_err + "E"))
					}
				} else {
					panic(errors.New("格式不正确,只能有1组 EEEE"))
				}
			case 'F':
				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'M':
					if is_fillmode {
						panic(errors.New("只能有1组 FM"))
					}
					if 1 == i {
						keywordGroup = append(keywordGroup, NUM_FM)
						is_fillmode = true
						break
					} else {
						panic(errors.New("FM 必须在开头"))
					}

				default:
					panic(errors.New(num_fmt_part_err + "F"))
				}
				break
			case 'G':
				// oracle中是变量可以设置为其他值 FIXME
				if GIndex == -1 {
					keywordGroup = append(keywordGroup, NUM_G)
					GIndex = i
					c = NLS_NUMERIC_CHARACTERS[1]
					isRerun = true
					continue
				} else {
					panic(errors.New("只能有1个 G"))
				}
			case 'L':
				if LIndex != -1 {
					panic(errors.New("只能有1个 C"))
				} else if lastFormatIndex != i && 0 != i {
					panic(errors.New("C 只能在开头或者结尾"))
				}
				keywordGroup = append(keywordGroup, NUM_L)
				LIndex = i
				break
			case 'M':
				if MiIndex != -1 {
					panic(errors.New("只能有1个 MI"))
				} else if i == (lastFormatIndex - 1) {
					panic(errors.New("MI 只能在结尾"))
				}

				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'I':
					keywordGroup = append(keywordGroup, NUM_MI)
					MiIndex = i
					break
				default:
					panic(errors.New(num_fmt_part_err + "M"))
				}
				break
			case 'P':
				if PrIndex != -1 {
					panic(errors.New("只能有1个 PR"))
				} else if i == (lastFormatIndex - 1) {
					panic(errors.New("PR 只能在结尾"))
				}

				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'R':
					keywordGroup = append(keywordGroup, NUM_PR)
					PrIndex = i
					break
				default:
					panic(errors.New(num_fmt_part_err + "P"))
				}

			case 'R':
				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'N':
					// 判断独占 FIXME

					if is_roman {
						panic(errors.New("只能有1个 RN"))
					} else if is_fillmode && i == lastFormatIndex && i != 3 {
						panic(errors.New("包含RN的格式,除了 FM 和 RN 不能有其他格式字符"))
					} else if !is_fillmode && i == lastFormatIndex && i != 1 {
						panic(errors.New("包含RN的格式,除了 FM 和 RN 不能有其他格式字符"))
					}

					keywordGroup = append(keywordGroup, NUM_RN)
					is_roman = true
					break
				default:
					panic(errors.New(num_fmt_part_err + "R"))
				}
			case 'S':
				if SIndex != -1 {
					panic(errors.New("只能有1个 S"))
				} else if i != lastFormatIndex && i != 0 {
					panic(errors.New("S 只能在开头或者结尾"))
				}

				keywordGroup = append(keywordGroup, NUM_S)
				SIndex = i
				break
			case 'T':
				if TMIndex == -1 {
					followingTwoChars := format[i : i+3]
					if "M9" == followingTwoChars {
						keywordGroup = append(keywordGroup, NUM_TM9)
						i = i + 3
						TMIndex = i
						break
					} else if "ME" == followingTwoChars {
						keywordGroup = append(keywordGroup, NUM_TME)
						i = i + 3
						TMIndex = i
						break
					} else if 'M' == followingTwoChars[0] {
						keywordGroup = append(keywordGroup, NUM_TM)
						i = i + 2
						TMIndex = i
						break
					} else {
						panic(errors.New(num_fmt_part_err + "T"))
					}

				} else {
					panic(errors.New("只能有1个 S"))
				}
			case 'U':
				if UIndex != -1 {
					panic(errors.New("只能有1个 U"))
				} else if lastFormatIndex != i && 0 != i {
					panic(errors.New("U 只能在开头或者结尾"))
				}
				keywordGroup = append(keywordGroup, NUM_U)
				UIndex = i
				break
			case 'V':
				if VIndex != -1 {
					panic(errors.New("只能有1个 V"))
				} else if 0 != i {
					panic(errors.New("U 不能在开头"))
				}

				keywordGroup = append(keywordGroup, NUM_V)
				VIndex = i
				break
			case 'X':
				keywordGroup = append(keywordGroup, NUM_X)
				break
			default:
				panic(errors.New(out_keyword_range_err))
			}

		} else {
			panic(errors.New(out_ascii_range_err))
		}

		i++
	}

	return keywordGroup
}

// 解析日期格式
// 可以适当考虑使用字典树实现
func parseDchFormat(format string) []string {
	var keywordGroup = make([]string, 4)

	l := len(format)
	for i := 0; i < l; {
		// 截取一个字符
		c := format[i]
		if c >= 32 && c <= 127 {
			fmt.Println((string)(c))

			// 不区分大小写
			ToUpper(&c)

			// 匹配关键词并存储
			switch c {
			case 'A':
				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case '.':
					j := i + 4
					followingChars := format[i:j]
					if "A.D." == followingChars {
						keywordGroup = append(keywordGroup, DCH_A_D)
						i = j
						break
					} else if "A.M." == followingChars {
						keywordGroup = append(keywordGroup, DCH_A_M)
						i = j
						break
					} else {
						panic(errors.New(dch_fmt_part_err + "A."))
					}
				case 'D':
					keywordGroup = append(keywordGroup, DCH_AD)
					break
				case 'M':
					keywordGroup = append(keywordGroup, DCH_AM)
					break
				default:
					panic(errors.New(dch_fmt_part_err + "A"))
				}
				break
			case 'B':
				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'C':
					keywordGroup = append(keywordGroup, DCH_BC)
					break
				case '.':
					j := i + 4
					followingChars := format[i:j]
					if ".C." == followingChars {
						keywordGroup = append(keywordGroup, DCH_B_C)
					}
					i = j
					break
				default:
					panic(errors.New(dch_fmt_part_err + "B"))
				}
				break
			case 'C':
				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'C':
					keywordGroup = append(keywordGroup, DCH_CC)
					break
				default:
					panic(errors.New(dch_fmt_part_err + "C"))
				}
				break
			case 'D':
				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'D':
					keywordGroup = append(keywordGroup, DCH_DD)
					thirdChar := format[i+1]
					if thirdChar == 'D' {
						keywordGroup = append(keywordGroup, DCH_DDD)
						break
					}
					break
				case 'Y':
					keywordGroup = append(keywordGroup, DCH_DY)
					break
				default:
					followingTwoChars := format[i : i+3]
					i = i + 3
					if followingTwoChars == "AY" {
						keywordGroup = append(keywordGroup, DCH_DAY)
						break
					} else {
						panic(errors.New(dch_fmt_part_err + "D"))
					}
				}
				break
			case 'F':
				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'X':
					keywordGroup = append(keywordGroup, DCH_FX)
					break
				default:
					followingTwoChars := format[i : i+3]
					i = i + 3
					switch followingTwoChars {
					case "F1":
						keywordGroup = append(keywordGroup, DCH_FF1)
						break
					case "F2":
						keywordGroup = append(keywordGroup, DCH_FF2)
						break
					case "F3":
						keywordGroup = append(keywordGroup, DCH_FF3)
						break
					case "F4":
						keywordGroup = append(keywordGroup, DCH_FF4)
						break
					case "F5":
						keywordGroup = append(keywordGroup, DCH_FF5)
						break
					case "F6":
						keywordGroup = append(keywordGroup, DCH_FF6)
						break
					case "F7":
						keywordGroup = append(keywordGroup, DCH_FF7)
						break
					case "F8":
						keywordGroup = append(keywordGroup, DCH_FF8)
						break
					case "F9":
						keywordGroup = append(keywordGroup, DCH_FF9)
						break
					default:
						panic(errors.New(dch_fmt_part_err + "F"))
					}
				}
				break
			case 'H':
				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'H':
					keywordGroup = append(keywordGroup, DCH_HH)
					break
				default:
					followingThreeChars := format[i : i+4]
					i = i + 4
					switch followingThreeChars {
					case "H24":
						keywordGroup = append(keywordGroup, DCH_HH24)
						break
					case "H12":
						keywordGroup = append(keywordGroup, DCH_HH12)
						break
					default:
						panic(errors.New(dch_fmt_part_err + "H"))
					}
				}
				break
			case 'I':
				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'D':
					followingThreeChars := format[i : i+4]
					i = i + 4
					if "DDD" == followingThreeChars {
						keywordGroup = append(keywordGroup, DCH_IDDD)
						break
					}
					// 匹配单个字符
					keywordGroup = append(keywordGroup, DCH_D)
					break
				case 'W':
					keywordGroup = append(keywordGroup, DCH_IW)
					break
				case 'Y':
					followingTwoChars := format[i : i+3]
					if "YY" == followingTwoChars {
						keywordGroup = append(keywordGroup, DCH_IYYY)
						i = i + 3
						break
					}

					followingOneChar = followingTwoChars[0]
					if 'Y' == followingOneChar {
						keywordGroup = append(keywordGroup, DCH_IYY)
						i = i + 2
						break
					}

					keywordGroup = append(keywordGroup, DCH_IY)
					break
				default:
					panic(errors.New(dch_fmt_part_err + "I"))
				}
				// 匹配单个字符
				keywordGroup = append(keywordGroup, DCH_I)
				break
			case 'J':
				keywordGroup = append(keywordGroup, DCH_J)
				break
			case 'M':
				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'M':
					break
				case 'I':
					break
				case 'S':
					break
				default:
					start := i
					i += 2
					followingTwoChars := format[start:i]
					if "ON" == followingTwoChars {
						keywordGroup = append(keywordGroup, DCH_MON)
						start = i
						i += 2
						followingTwoChars = format[start:i]
						if "TH" == followingTwoChars {
							keywordGroup = append(keywordGroup, DCH_MONTH)
						} else {
							panic(errors.New(dch_fmt_part_err + "MON"))
						}
					} else {
						panic(errors.New(dch_fmt_part_err + "M"))
					}
				}
				break
			case 'O':
				i++
				if 'F' == format[i] {
					keywordGroup = append(keywordGroup, DCH_MONTH)
					break
				}
				panic(errors.New(dch_fmt_part_err + "O"))
			case 'P':
				i++
				if 'M' == format[i] {
					keywordGroup = append(keywordGroup, DCH_PM)
					break
				}
				start := i
				i += 3
				followingThreeChars := format[start:i]
				if ".M." == followingThreeChars {
					keywordGroup = append(keywordGroup, DCH_P_M)
					break
				}

				panic(errors.New(dch_fmt_part_err + "P"))
			case 'Q':
				keywordGroup = append(keywordGroup, DCH_Q)
				break
			case 'R':
				i++
				if 'M' == format[i] {
					keywordGroup = append(keywordGroup, DCH_RM)
					break
				}
				panic(errors.New(dch_fmt_part_err + "P"))
			case 'S':
				start := i
				i += 4
				followingFourChars := format[start:i]
				if "SSSS" == followingFourChars {
					keywordGroup = append(keywordGroup, DCH_SSSSS)
					break
				} else if "SSS" == followingFourChars[0:3] {
					keywordGroup = append(keywordGroup, DCH_SSSS)
					break
				} else if "S" == followingFourChars[0:3] {
					keywordGroup = append(keywordGroup, DCH_SS)
					break
				}
				panic(errors.New(dch_fmt_part_err + "S"))
				break
			case 'T':
				start := i
				i += 2
				followingTwoChars := format[start:i]
				if "ZH" == followingTwoChars {
					keywordGroup = append(keywordGroup, DCH_TZH)
					break
				} else if "ZM" == followingTwoChars {
					keywordGroup = append(keywordGroup, DCH_TZM)
					break
				} else if 'Z' == followingTwoChars[0] {
					keywordGroup = append(keywordGroup, DCH_TZM)
					break
				} else {
					panic(errors.New(dch_fmt_part_err + "T"))
				}
			case 'U':
				i++
				followingOneChar := format[i]
				if 'S' == followingOneChar {
					keywordGroup = append(keywordGroup, DCH_US)
					break
				}
				panic(errors.New(dch_fmt_part_err + "U"))
			case 'W':
				i++
				followingOneChar := format[i]
				if 'W' == followingOneChar {
					keywordGroup = append(keywordGroup, DCH_WW)
					break
				}
				keywordGroup = append(keywordGroup, DCH_W)
				break
			case 'Y':
				start := i
				i += 4
				followingFourChars := format[start:i]
				if ",YYYY" == followingFourChars {
					keywordGroup = append(keywordGroup, DCH_Y_YYY)
				} else if "YYY" == followingFourChars[0:3] {
					keywordGroup = append(keywordGroup, DCH_YYY)
				} else if "YY" == followingFourChars[0:2] {
					keywordGroup = append(keywordGroup, DCH_YY)
				}
				keywordGroup = append(keywordGroup, DCH_Y)
				break
			default:
				panic(errors.New(out_keyword_range_err))
			}
		} else {
			panic(errors.New(out_ascii_range_err))
		}
	}

	return keywordGroup
}

// 根据格式解析参数
func parseParam(param string, keywordGroup []string) {
	// 新的 数字或者 字符串或者 日期

	// 结果的最后1个字符的索引
	//reaultEnd := len(param) - 1
	// 获取参数的字符的下标
	paramIndex := 0
	// 参数字节长度
	paramLen := len(param)

	// 倒序存储
	inverseResult := make([]byte, paramLen)

	// 循环模式
	// 获取模式，去字符或字符串，校验，处理
	for _, keywrod := range keywordGroup {
		// 获取长度，截取字符串 TODO
		switch keywrod {
		//,(comma) 半角逗号 装饰作用
		case NUM_COMMA:
			paramIndex++
			break
		// .(period) 半角句号,点号，小数点
		case NUM_DEC:
			inverseResult = append(inverseResult, '.')
			paramIndex++
			break

		// $ 美元符号 返回的值以$符号开头 TODO
		// to_number和to_char的输出表现不一样
		case NUM_DOLLAR:
			break
		case NUM_0:
			break
		//9 替换数字，开头的0替换为空格，0除外
		case NUM_9:
			// 校验参数是否合法
			if '0' <= param[paramIndex] && param[paramIndex] <= '9' {
				inverseResult = append(inverseResult, param[paramIndex])
			} else {
				panic(errors.New(invalid_num_err))
			}

			paramIndex++
			break
		case NUM_B:
			break
		case NUM_C:
			break
		case NUM_D:
			break
		case NUM_E:
			break
		case NUM_FM:
			break
		case NUM_G:
			break
		case NUM_L:
			break
		case NUM_MI:
			break
		case NUM_PR:
			break
		case NUM_RN:
			break
		case NUM_S:
			break
		case NUM_V:
			break
		default:
			panic(errors.New("不应该到达此处.不应该出现的格式"))
		}

		// 检查 格式长度是否超过参数长度
		// 超过部分是0 还是空格?
		if paramIndex == paramLen {
			fmt.Println("参数长度比格式个数少")
			break
		}
	}

	if paramIndex != paramLen {
		panic(errors.New("参数长度比格式个数多"))
	}

	result := bytes.Buffer{}
	for i := len(inverseResult) - 1; i >= 0; i-- {
		result.WriteByte(inverseResult[i])
	}

	str := result.String()
	fmt.Println(str)
}

func ToUpper(c *byte) {
	if 'a' <= *c && *c <= 'z' {
		*c -= 32
	}
}
