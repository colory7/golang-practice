package oracle_demo

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
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

const ()

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
	DCH_AD    = "AD"
	DCH_A_D   = "A.D."
	DCH_AM    = "AM"
	DCH_A_M   = "A.M."
	DCH_BC    = "BC"
	DCH_B_C   = "B.C."
	DCH_CC    = "CC"
	DCH_SCC   = "SCC"
	DCH_DAY   = "DAY"
	DCH_DDD   = "DDD"
	DCH_DD    = "DD"
	DCH_DY    = "DY"
	DCH_DL    = "DL"
	DCH_DS    = "DS"
	DCH_D     = "D"
	DCH_E     = "E"
	DCH_EE    = "EE"
	DCH_FF1   = "FF1"
	DCH_FF2   = "FF2"
	DCH_FF3   = "FF3"
	DCH_FF4   = "FF4"
	DCH_FF5   = "FF5"
	DCH_FF6   = "FF6"
	DCH_FF7   = "FF7"
	DCH_FF8   = "FF8"
	DCH_FF9   = "FF9"
	DCH_FM    = "FM"
	DCH_FX    = "FX"
	DCH_HH24  = "HH24"
	DCH_HH12  = "HH12"
	DCH_HH    = "HH"
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
	DCH_P_M   = "P.M."
	DCH_PM    = "PM"
	DCH_Q     = "Q"
	DCH_RM    = "RM"
	DCH_RR    = "RR"
	DCH_SSSSS = "SSSSS"
	DCH_SSSS  = "SSSS"
	DCH_SS    = "SS"
	DCH_TZH   = "TZH"
	DCH_TZM   = "TZM"
	DCH_TZD   = "TZD"
	DCH_TZR   = "TZR"
	DCH_TS    = "TS"
	DCH_WW    = "WW"
	DCH_W     = "W"
	DCH_X     = "X"
	DCH_Y_YYY = "Y,YYY"
	DCH_YEAR  = "YEAR"
	DCH_SYEAR = "SYEAR"
	DCH_YYYY  = "YYYY"
	DCH_SYYYY = "SYYYY"
	DCH_YYY   = "YYY"
	DCH_YY    = "YY"
	DCH_Y     = "Y"
)

var dchSkip = [...]uint8{'-', '/', ',', '.', ';', ':', '"'}

//type keyword struct {
//	name string
//	len uint8
//	id int
//}

var numKeywords map[string]uint8
var dchKeywords map[string]uint8

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

type NumFmtDesc struct {
	// 前半部分 9或0的个数
	// 如果是V 模式 忽略9或0的个数，不用格式做截取，直接用参数做乘积计算
	preDec int
	// 前半部分 是否是0开头
	isLeadingZero bool
	// 后半部分 9或0的个数
	postDec int
	// 逗号位置 FIXME 同G
	commaIndex int
	// 点号位置 FIXME 同D
	decIndex int
	// S 位置 只能是开头或结尾
	s uint8
	// X 的个数 输出区分大小写，输出前先对参数做四舍五入，转换为正整数
	xCount int
	// 辅助前缀
	fm bool
	// 互斥前缀
	prefix int
	// 辅助后缀
	auxSuffix int
	// 辅助前缀+ 十进制+ EEEE
	isEEEE bool
	// 辅助前缀+ 独占后缀
	isRn bool
	tm   int
	isX  bool
	isV  bool
}

const (
	// 辅助前缀
	NUM_F_FILLMODE = 1 << 1

	// 互斥前缀
	NUM_F_DOLLAR = 1 << 2
	NUM_F_B      = 1 << 3
	NUM_F_C      = 1 << 4
	NUM_F_L      = 1 << 5
	NUM_F_U      = 1 << 6

	// 辅助后缀
	NUM_F_MI = 1 << 7
	NUM_F_PR = 1 << 8

	// 辅助前缀+ 十进制+ EEEE
	NUM_F_EEEE = 1 << 9

	// 辅助前缀+ 独占后缀
	NUM_F_RN  = 1 << 10
	NUM_F_TM  = 1
	NUM_F_TME = 2
	NUM_F_TM9 = 3
	NUM_F_X   = 1 << 12

	// 中缀 或后缀
	NUM_F_V = 1 << 13
	NUM_F_G = 1 << 14

	// 前缀 中缀 后缀
	NUM_F_D = 1 << 15

	//前缀 后缀
	NUM_F_START_S = 1
	NUM_F_END_S   = 1 << 1
)

const (
	empty = ""
	plus  = "+"
	minus = "-"
	dec   = "."
)

type NumParam struct {
	sign      string
	pre       string
	post      string
	eSign     string
	eExponent int
	// 可选
	isFloat bool
	isEEEE  bool
}

func (numParam *NumParam) string() string {
	var result bytes.Buffer
	if plus == numParam.sign {
		result.WriteString(plus)
	} else if minus == numParam.sign {
		result.WriteString(minus)
	} else if empty == numParam.sign {
	} else {
		panic("sign属性格式错误")
	}

	if empty != numParam.pre {
		result.WriteString(numParam.pre)
	} else {
		panic("格式错误,整数部分是空的")
	}

	if numParam.post != empty {
		result.WriteByte('.')
		result.WriteString(string(numParam.post))
	}

	if numParam.eExponent != 0 {
		result.WriteByte('e')

		if plus == numParam.sign {
			result.WriteString(plus)
		} else if minus == numParam.sign {
			result.WriteString(minus)
		} else if empty == numParam.sign {
		} else {
			panic("eSign属性格式错误")
		}
		result.WriteString(fmt.Sprint(numParam.eExponent))
	}

	return result.String()
}

// 解析数值格式
func parseNumFormat(format string) NumFmtDesc {
	var fmtDesc NumFmtDesc
	fmtDesc.commaIndex = -1
	fmtDesc.decIndex = -1

	// 格式字节长度
	formatLen := len(format)

	// 最后一个关键词的索引
	lastFormatIndex := formatLen - 1

	isRerun := false
	var c byte

	var preDec = bytes.Buffer{}
	var postDec = bytes.Buffer{}
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
				if fmtDesc.commaIndex == -1 {
					if i == 0 {
						panic(errors.New("不能以逗号开头"))
					} else if i == lastFormatIndex {
						panic(errors.New("逗号不能出现在数字最右边"))
					} else if fmtDesc.decIndex != -1 {
						panic(errors.New("逗号不能出现在点号的右边"))
					}

					fmtDesc.commaIndex = i
				} else {
					panic(errors.New("格式错误，存在多个格式符号 ,"))
				}

			case '.':
				if fmtDesc.decIndex == -1 {
					fmtDesc.decIndex = i
				} else {
					panic(errors.New("只能有1个 ."))
				}
			case '0':
				if fmtDesc.decIndex == -1 {
					preDec.WriteByte('0')
				} else {
					postDec.WriteByte('0')
				}
			case '9':
				if fmtDesc.decIndex == -1 {
					preDec.WriteByte('9')
				} else {
					postDec.WriteByte('9')
				}
			case '$':
				if fmtDesc.prefix != 0 && i == 0 {
					fmtDesc.prefix |= NUM_F_DOLLAR
				} else {
					panic(errors.New("格式前缀冲突 " + "$"))
				}
			case 'B':
				if fmtDesc.prefix != 0 {
					fmtDesc.prefix |= NUM_F_B
				} else {
					panic(errors.New("格式前缀冲突 " + "B"))
				}
			case 'C':
				if fmtDesc.prefix != 0 {
					panic(errors.New("格式前缀冲突 " + "C"))
				} else if lastFormatIndex != i && 0 != i {
					panic(errors.New("C 只能在开头或者结尾"))
				}
				fmtDesc.prefix |= NUM_F_C
			case 'L':
				if fmtDesc.prefix != 0 {
					panic(errors.New("格式前缀冲突 " + "L"))
				}
			case 'U':
				if fmtDesc.prefix != 0 {
					panic(errors.New("格式前缀冲突 " + "U"))
				} else if lastFormatIndex != i && 0 != i {
					panic(errors.New("U 只能在开头或者结尾"))
				}
				fmtDesc.prefix |= NUM_F_U
			case NLS_NUMERIC_CHARACTERS[0]:
				// 默认是 . oracle中是变量可以设置为其他值 FIXME
				if fmtDesc.decIndex != -1 {
					isRerun = true
					fmtDesc.decIndex = i
					continue
				} else {
					panic(errors.New("只能有1个 " + string(NLS_NUMERIC_CHARACTERS[0])))
				}
			case NLS_NUMERIC_CHARACTERS[1]:
				// oracle中是变量可以设置为其他值 FIXME
				if fmtDesc.commaIndex != -1 {
					isRerun = true
					continue
				} else {
					panic(errors.New("只能有1个 " + string(NLS_NUMERIC_CHARACTERS[1])))
				}
			case 'E':
				if !fmtDesc.isEEEE {
					j := i + 4
					followingChars := format[i+1 : j]
					if "EEE" == followingChars {
						i = j
						fmtDesc.isEEEE = true
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
					if !fmtDesc.fm {
						panic(errors.New("只能有1组 FM"))
					}
					if 1 == i {
						fmtDesc.fm = true
					} else {
						panic(errors.New("FM 必须在开头"))
					}
				default:
					panic(errors.New(num_fmt_part_err + "F"))
				}
			case 'M':
				if fmtDesc.auxSuffix != 0 {
					panic(errors.New("辅助后缀冲突" + "MI"))
				} else if i == (lastFormatIndex - 1) {
					panic(errors.New("MI 只能在结尾"))
				}

				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'I':
					fmtDesc.auxSuffix |= NUM_F_MI
				default:
					panic(errors.New(num_fmt_part_err + "M"))
				}
			case 'P':
				if fmtDesc.auxSuffix != 0 {
					panic(errors.New("辅助后缀冲突" + "PR"))
				} else if i == (lastFormatIndex - 1) {
					panic(errors.New("PR 只能在结尾"))
				}

				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'R':
					fmtDesc.auxSuffix |= NUM_F_PR
				default:
					panic(errors.New(num_fmt_part_err + "P"))
				}
			case 'R':
				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'N':
					// 判断独占 长度 FIXME
					if !fmtDesc.isRn {
						panic(errors.New("只能有1个 RN"))
					} else if (fmtDesc.fm) && (i == lastFormatIndex && i != 3) {
						panic(errors.New("包含RN的格式,除了 FM 和 RN 不能有其他格式字符"))
					} else if !(fmtDesc.fm) && i == lastFormatIndex && i != 1 {
						panic(errors.New("包含RN的格式,除了 FM 和 RN 不能有其他格式字符"))
					}

					fmtDesc.isRn = true
				default:
					panic(errors.New(num_fmt_part_err + "R"))
				}
			case 'S':
				if fmtDesc.s != 0 {
					panic(errors.New("只能有1个 S"))
				} else if i != lastFormatIndex && i != 0 {
					panic(errors.New("S 只能在开头或者结尾"))
				}

				if i == 0 {
					fmtDesc.s = NUM_F_START_S
				} else {
					fmtDesc.s = NUM_F_END_S
				}
			case 'T':
				if fmtDesc.tm != 0 {
					i++
					followingOneChar := format[i]
					if 'M' == followingOneChar {
						if i == lastFormatIndex {
							fmtDesc.tm = NUM_F_TM
						} else {
							i++
							followingOneChar = format[i]
							if 'E' == followingOneChar {
								fmtDesc.tm = NUM_F_TME
							} else if '9' == followingOneChar {
								fmtDesc.tm = NUM_F_TM9
							} else {
								panic(errors.New("格式错误在 TM 附近"))
							}
						}
					} else {
						panic(errors.New("格式错误在 T 附近"))
					}
				} else {
					panic(errors.New("只能有1组 TM"))
				}
			case 'V':
				if !fmtDesc.isV {
					panic(errors.New("只能有1个 V"))
				} else if 0 != i {
					panic(errors.New("V 不能在开头"))
				}
				fmtDesc.isV = true
			case 'X':
				fmtDesc.isX = true
				fmtDesc.xCount++
			default:
				panic(errors.New(out_keyword_range_err))
			}

		} else {
			panic(errors.New(out_ascii_range_err))
		}

		i++
	}

	return fmtDesc

}

// 解析数字参数
func preParseNumParam(num string) NumParam {
	var numParam NumParam

	readDec := false

	var preBuf = bytes.Buffer{}
	var postBuf = bytes.Buffer{}
	for i := 0; i < len(num); i++ {
		c := num[i]
		fmt.Println(string(c))
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if readDec {
				postBuf.WriteByte(c)
			} else {
				preBuf.WriteByte(c)
			}
		case '.':
			if readDec == false {
				readDec = true
				numParam.isFloat = true
			} else {
				panic("多个符号 " + ".")
			}
		case 'e', 'E':
			i++

			numParam.isEEEE = true
			var exponent = bytes.Buffer{}

			if num[i] == '+' {
				numParam.eSign = plus
			} else if num[i] == '-' {
				numParam.eSign = minus
			} else if num[i] <= '9' && num[i] >= '0' {
				numParam.eSign = empty
				exponent.WriteByte(num[i])
			}

			for i++; i < len(num); i++ {
				if num[i] <= '9' && num[i] >= '0' {
					exponent.WriteByte(num[i])
				} else {
					panic(errors.New("科学计数的指数使用了非法字符 " + string(num[i])))
				}
			}

			exponentNum, err := strconv.Atoi(exponent.String())
			if err != nil {
				panic(err)
			}
			numParam.eExponent = exponentNum
			fmt.Println(exponent.String())
		case '-':
			if i == 0 {
				numParam.sign = minus
			} else {
				panic("符号位置不对 " + "-")
			}
		case '+':
			if i == 0 {
				numParam.sign = plus
			} else {
				panic("符号位置不对 " + "+")
			}
		case ',':
			panic("暂时不支持 " + ",")
		default:
			panic(errors.New("不支持的数字符号"))
		}
	}

	// EEEE
	// 十进制 符号可选
	// 十进制 符号可选 逗号分组

	return numParam
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
					} else if "A.M." == followingChars {
						keywordGroup = append(keywordGroup, DCH_A_M)
						i = j
					} else {
						panic(errors.New(dch_fmt_part_err + "A."))
					}
				case 'D':
					keywordGroup = append(keywordGroup, DCH_AD)
				case 'M':
					keywordGroup = append(keywordGroup, DCH_AM)
				default:
					panic(errors.New(dch_fmt_part_err + "A"))
				}
			case 'B':
				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'C':
					keywordGroup = append(keywordGroup, DCH_BC)
				case '.':
					j := i + 4
					followingChars := format[i:j]
					if ".C." == followingChars {
						keywordGroup = append(keywordGroup, DCH_B_C)
					}
					i = j
				default:
					panic(errors.New(dch_fmt_part_err + "B"))
				}
			case 'C':
				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'C':
					keywordGroup = append(keywordGroup, DCH_CC)
				default:
					panic(errors.New(dch_fmt_part_err + "C"))
				}
			case 'D':
				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'D':
					keywordGroup = append(keywordGroup, DCH_DD)
					thirdChar := format[i+1]
					if thirdChar == 'D' {
						keywordGroup = append(keywordGroup, DCH_DDD)
					}
				case 'Y':
					keywordGroup = append(keywordGroup, DCH_DY)
				default:
					followingTwoChars := format[i : i+3]
					i = i + 3
					if followingTwoChars == "AY" {
						keywordGroup = append(keywordGroup, DCH_DAY)
					} else {
						panic(errors.New(dch_fmt_part_err + "D"))
					}
				}
			case 'F':
				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'X':
					keywordGroup = append(keywordGroup, DCH_FX)
				default:
					followingTwoChars := format[i : i+3]
					i = i + 3
					switch followingTwoChars {
					case "F1":
						keywordGroup = append(keywordGroup, DCH_FF1)
					case "F2":
						keywordGroup = append(keywordGroup, DCH_FF2)
					case "F3":
						keywordGroup = append(keywordGroup, DCH_FF3)
					case "F4":
						keywordGroup = append(keywordGroup, DCH_FF4)
					case "F5":
						keywordGroup = append(keywordGroup, DCH_FF5)
					case "F6":
						keywordGroup = append(keywordGroup, DCH_FF6)
					case "F7":
						keywordGroup = append(keywordGroup, DCH_FF7)
					case "F8":
						keywordGroup = append(keywordGroup, DCH_FF8)
					case "F9":
						keywordGroup = append(keywordGroup, DCH_FF9)
					default:
						panic(errors.New(dch_fmt_part_err + "F"))
					}
				}
			case 'H':
				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'H':
					keywordGroup = append(keywordGroup, DCH_HH)
				default:
					followingThreeChars := format[i : i+4]
					i = i + 4

					switch followingThreeChars {
					case "H24":
						keywordGroup = append(keywordGroup, DCH_HH24)
					case "H12":
						keywordGroup = append(keywordGroup, DCH_HH12)
					default:
						panic(errors.New(dch_fmt_part_err + "H"))
					}
				}
			case 'I':
				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'W':
					keywordGroup = append(keywordGroup, DCH_IW)
				case 'Y':
					followingTwoChars := format[i : i+3]
					if "YY" == followingTwoChars {
						keywordGroup = append(keywordGroup, DCH_IYYY)
						i = i + 3
					}

					followingOneChar = followingTwoChars[0]
					if 'Y' == followingOneChar {
						keywordGroup = append(keywordGroup, DCH_IYY)
						i = i + 2
					}
					keywordGroup = append(keywordGroup, DCH_IY)
				default:
					panic(errors.New(dch_fmt_part_err + "I"))
				}
				// 匹配单个字符
				keywordGroup = append(keywordGroup, DCH_I)
			case 'J':
				keywordGroup = append(keywordGroup, DCH_J)
			case 'M':
				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'M':
					keywordGroup = append(keywordGroup, DCH_MM)
				case 'I':
					keywordGroup = append(keywordGroup, DCH_MI)
				default:
					start := i
					i += 2
					if i < l {
						followingTwoChars := format[start:i]
						if "ON" == followingTwoChars {
							keywordGroup = append(keywordGroup, DCH_MON)
							start = i
							i += 2
							if i < l {
								followingTwoChars = format[start:i]
								if "TH" == followingTwoChars {
									keywordGroup = append(keywordGroup, DCH_MONTH)
								} else {
									panic(errors.New(dch_fmt_part_err + "MON"))
								}
							} else {
								panic(errors.New(dch_fmt_part_err + "MON"))
							}
						} else {
							panic(errors.New(dch_fmt_part_err + "M"))
						}
					} else {
						panic(errors.New(dch_fmt_part_err + "M"))
					}
				}
			case 'O':
				i++
				if 'F' == format[i] {
					keywordGroup = append(keywordGroup, DCH_MONTH)
				}
				panic(errors.New(dch_fmt_part_err + "O"))
			case 'P':
				i++
				if i < l {
					if 'M' == format[i] {
						keywordGroup = append(keywordGroup, DCH_PM)
					} else {
						start := i
						i += 3
						if i < l {
							followingThreeChars := format[start:i]
							if ".M." == followingThreeChars {
								keywordGroup = append(keywordGroup, DCH_P_M)
							} else {
								panic(errors.New(dch_fmt_part_err + "P"))
							}
						} else {
							panic(errors.New(dch_fmt_part_err + "P"))
						}
					}
				} else {
					panic(errors.New(dch_fmt_part_err + "P"))
				}
			case 'Q':
				keywordGroup = append(keywordGroup, DCH_Q)
			case 'R':
				i++
				if 'M' == format[i] {
					keywordGroup = append(keywordGroup, DCH_RM)
				} else {
					panic(errors.New(dch_fmt_part_err + "R"))
				}
			case 'S':
				start := i
				i += 4
				followingFourChars := format[start:i]
				if "SSSS" == followingFourChars {
					keywordGroup = append(keywordGroup, DCH_SSSSS)
				} else if "SSS" == followingFourChars[0:3] {
					keywordGroup = append(keywordGroup, DCH_SSSS)
				} else if "S" == followingFourChars[0:3] {
					keywordGroup = append(keywordGroup, DCH_SS)
				} else {
					panic(errors.New(dch_fmt_part_err + "S"))
				}
			case 'T':
				start := i
				i += 2
				followingTwoChars := format[start:i]
				if "ZH" == followingTwoChars {
					keywordGroup = append(keywordGroup, DCH_TZH)
				} else if "ZM" == followingTwoChars {
					keywordGroup = append(keywordGroup, DCH_TZM)
				} else if 'Z' == followingTwoChars[0] {
					keywordGroup = append(keywordGroup, DCH_TZM)
				} else {
					panic(errors.New(dch_fmt_part_err + "T"))
				}
			case 'W':
				i++
				followingOneChar := format[i]
				if 'W' == followingOneChar {
					keywordGroup = append(keywordGroup, DCH_WW)
				} else {
					keywordGroup = append(keywordGroup, DCH_W)
				}
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
				} else {
					keywordGroup = append(keywordGroup, DCH_Y)
				}
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
func parseNumParam(param string, keywordGroup []string, numDesc NumFmtDesc) string {
	// 新的 数字或者 字符串或者 日期

	// 结果的最后1个字符的索引
	//reaultEnd := len(param) - 1
	// 获取参数的字符的下标
	paramIndex := 0
	// 参数字节长度
	paramLen := len(param)

	// 倒序存储
	inverseResult := make([]byte, paramLen)

	result := bytes.Buffer{}

	// 循环模式
	// 获取模式，去字符或字符串，校验，处理
	for _, keyword := range keywordGroup {
		// 获取长度，截取字符串 TODO
		switch keyword {
		//,(comma) 半角逗号 装饰作用
		case NUM_COMMA:
			paramIndex++
		// .(period) 半角句号,点号，小数点
		case NUM_DEC:
			inverseResult = append(inverseResult, '.')
			paramIndex++
		// $ 美元符号 返回的值以$符号开头 TODO
		// to_number和to_char的输出表现不一样
		case NUM_DOLLAR:
		case NUM_0:
		//9 替换数字，开头的0替换为空格，0除外
		case NUM_9:
			// 校验参数是否合法
			if '0' <= param[paramIndex] && param[paramIndex] <= '9' {
				inverseResult = append(inverseResult, param[paramIndex])
			} else {
				panic(errors.New(invalid_num_err))
			}

			paramIndex++
		case NUM_B:
		case NUM_C:
		case NUM_D:
		case NUM_E:
			scienceFmt := "%" + fmt.Sprint(numDesc.postDec) + ".E"
			//
			result.WriteString(fmt.Sprintf(scienceFmt, param))
		case NUM_FM:
		case NUM_G:
		case NUM_L:
		case NUM_MI:
		case NUM_PR:
			// 如果开头的字符是负号，则用尖括号包裹值
			// FIXME
			result.WriteByte('<')
			result.WriteByte('>')
		case NUM_RN:
			d, err := strconv.Atoi(param)
			if err != nil {
				panic(err)
			}
			result = intToRoman(d)
		case NUM_S:
		case NUM_V:
		case NUM_X:
		default:
			panic(errors.New("不应该到达此处.不应该出现的格式"))
		}

		// 检查 格式长度是否超过参数长度
		// 超过部分是0 还是空格?
		//if paramIndex == paramLen {
		//	break
		//}
	}

	//if paramIndex != paramLen {
	//	panic(errors.New("参数长度比格式个数多"))
	//}

	//result := bytes.Buffer{}
	//for i := len(inverseResult) - 1; i >= 0; i-- {
	//	result.WriteByte(inverseResult[i])
	//}

	//if len(result) == 0 {
	//
	//}

	return result.String()
}

func ToUpper(c *byte) {
	if 'a' <= *c && *c <= 'z' {
		*c -= 32
	}
}

func intToRoman(num int) bytes.Buffer {
	romes := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	numbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	rm := bytes.Buffer{}
	for i := 0; i < len(numbers); i++ {
		for num >= numbers[i] {
			num -= numbers[i]
			rm.WriteString(romes[i])
		}
	}
	return rm
}

func TestParseNumFmt(t *testing.T) {
	f := "99EEEE"
	numFmtDesc := parseNumFormat(f)

	str := fmt.Sprintf("%#v\n", numFmtDesc)
	fmt.Println(str)
}

func TestParseNumParam(t *testing.T) {
	num := "-36.25e+97"
	numProc := preParseNumParam(num)
	str := fmt.Sprintf("%#v\n", numProc)
	fmt.Println(str)
}

func TestToString(t *testing.T) {
	var numParam NumParam
	numParam.sign = plus
	numParam.pre = "36"
	numParam.post = "25"
	numParam.eSign = empty
	numParam.eExponent = 12

	fmt.Println(numParam.string())
}
