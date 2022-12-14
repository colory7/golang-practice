package oracle_demo

import (
	"bytes"
	"errors"
	"fmt"
	"math"
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

const (
	// 辅助前缀 没有冲突
	NUM_FMT_AUX_PREFIX_EMPTY = 0
	NUM_FMT_AUX_PREFIX_FM    = 1

	// 前缀 前缀互斥
	NUM_FMT_PREFIX_EMPTY  = 0
	NUM_FMT_PREFIX_DOLLAR = '$'
	NUM_FMT_PREFIX_B      = 'B'
	NUM_FMT_PREFIX_C      = 'C'
	NUM_FMT_PREFIX_L      = 'L'
	NUM_FMT_PREFIX_U      = 'U'

	// 后缀 后缀互斥 后缀决定了输出模式
	NUM_FMT_SUFFIX_EMPTY = 0
	NUM_FMT_SUFFIX_EEEE  = 1
	NUM_FMT_SUFFIX_V     = 2
	NUM_FMT_SUFFIX_RN    = 3
	NUM_FMT_SUFFIX_X     = 4
	NUM_FMT_SUFFIX_TM    = 5
	NUM_FMT_SUFFIX_TM9   = 6
	NUM_FMT_SUFFIX_TME   = 7
	NUM_FMT_SUFFIX_TMe   = 8

	// 辅助后缀 与MI PR冲突
	NUM_FMT_AUX_SUFFIX_EMPTY = 0
	NUM_FMT_AUX_SUFFIX_MI    = 1
	NUM_FMT_AUX_SUFFIX_PR    = 2

	// S
	NUM_FMT_S_EMPTY = 0
	NUM_FMT_S_START = 1
	NUM_FMT_S_END   = 2
)

type NumFmtDesc struct {
	// 辅助前缀 前缀 十进制前半部分 逗号 小数点 十进制后半部分 后缀 辅助后缀
	// 辅助前缀
	auxPrefix int
	// 互斥前缀
	prefix int
	// 前半部分 9或0的个数
	// 如果是V 模式 忽略9或0的个数，不用格式做截取，直接用参数做乘积计算
	preDec int
	// 前半部分 是否是0开头
	isLeadingZero bool
	// 0开头的个数 FIXME
	zeroCount int
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
	// 后缀
	suffix int
	// 辅助后缀
	auxSuffix int
}

const (
	empty = ""
	plus  = "+"
	minus = "-"
	dec   = "."
)

type NumParamDesc struct {
	sign      string
	preDec    string
	postDec   string
	eSign     string
	eExponent int
	// 可选
	existDec bool
	isEEEE   bool
}

func (numParam *NumParamDesc) string() string {
	var result bytes.Buffer
	if plus == numParam.sign {
		result.WriteString(plus)
	} else if minus == numParam.sign {
		result.WriteString(minus)
	} else if empty == numParam.sign {
	} else {
		panic("sign属性格式错误")
	}

	if empty != numParam.preDec {
		result.WriteString(numParam.preDec)
	} else {
		panic("格式错误,整数部分是空的")
	}

	if numParam.postDec != empty {
		result.WriteByte('.')
		result.WriteString(string(numParam.postDec))
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

	// 格式字节长度
	formatLen := len(format)

	// 最后一个关键词的索引
	lastFormatIndex := formatLen - 1

	isRerun := false
	var c byte

	readDec := false
	//var preDec = bytes.Buffer{}
	//var postDec = bytes.Buffer{}
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
				if fmtDesc.commaIndex == 0 {
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
				if !readDec {
					fmtDesc.decIndex = i
					readDec = true
				} else {
					panic(errors.New("只能有1个 ."))
				}
			case '0':
				if readDec {
					fmtDesc.postDec++
					//postDec.WriteByte('0')
				} else {
					fmtDesc.preDec++
					//preDec.WriteByte('0')
				}
			case '9':
				if readDec {
					fmtDesc.postDec++
					//postDec.WriteByte('9')
				} else {
					fmtDesc.preDec++
					//preDec.WriteByte('9')
				}
			case '$':
				if fmtDesc.prefix == NUM_FMT_PREFIX_EMPTY && i == 0 {
					fmtDesc.prefix = NUM_FMT_PREFIX_DOLLAR
				} else {
					panic(errors.New("格式前缀冲突 " + "$"))
				}
			case 'B':
				if fmtDesc.prefix == NUM_FMT_PREFIX_EMPTY {
					fmtDesc.prefix = NUM_FMT_PREFIX_B
				} else {
					panic(errors.New("格式前缀冲突 " + "B"))
				}
			case 'C':
				if fmtDesc.prefix == NUM_FMT_PREFIX_EMPTY {
					panic(errors.New("格式前缀冲突 " + "C"))
				} else if lastFormatIndex != i && 0 != i {
					panic(errors.New("C 只能在开头或者结尾"))
				}
				fmtDesc.prefix = NUM_FMT_PREFIX_C
			case 'L':
				if fmtDesc.prefix == NUM_FMT_PREFIX_EMPTY {
					fmtDesc.prefix = NUM_FMT_PREFIX_L
				} else {
					panic(errors.New("格式前缀冲突 " + "L"))
				}
			case 'U':
				if fmtDesc.prefix == NUM_FMT_PREFIX_EMPTY {
					panic(errors.New("格式前缀冲突 " + "U"))
				} else if lastFormatIndex != i && 0 != i {
					panic(errors.New("U 只能在开头或者结尾"))
				}
				fmtDesc.prefix = NUM_FMT_PREFIX_U
			case NLS_NUMERIC_CHARACTERS[0]:
			case NLS_NUMERIC_CHARACTERS[1]:
			case 'E':
				if fmtDesc.suffix == NUM_FMT_SUFFIX_EMPTY {
					j := i + 4
					followingChars := format[i+1 : j]
					if "EEE" == followingChars {
						i = j
						fmtDesc.suffix = NUM_FMT_SUFFIX_EEEE
					} else {
						panic(errors.New(num_fmt_part_err + "E"))
					}
				} else {
					panic(errors.New("conflict with E"))
				}
			case 'F':
				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'M':
					if fmtDesc.auxPrefix == NUM_FMT_AUX_PREFIX_EMPTY {
						panic(errors.New("只能有1组 FM"))
					}
					if 1 == i {
						fmtDesc.auxPrefix = NUM_FMT_AUX_PREFIX_FM
					} else {
						panic(errors.New("FM 必须在开头"))
					}
				default:
					panic(errors.New(num_fmt_part_err + "F"))
				}
			case 'M':
				if fmtDesc.auxSuffix == NUM_FMT_AUX_SUFFIX_EMPTY {
					panic(errors.New("辅助后缀冲突" + "MI"))
				} else if i == (lastFormatIndex - 1) {
					panic(errors.New("MI 只能在结尾"))
				}

				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'I':
					fmtDesc.auxSuffix = NUM_FMT_AUX_SUFFIX_MI
				default:
					panic(errors.New(num_fmt_part_err + "M"))
				}
			case 'P':
				if fmtDesc.auxSuffix == NUM_FMT_AUX_SUFFIX_EMPTY {
					panic(errors.New("辅助后缀冲突" + "PR"))
				} else if i == (lastFormatIndex - 1) {
					panic(errors.New("PR 只能在结尾"))
				}

				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'R':
					fmtDesc.auxSuffix = NUM_FMT_AUX_SUFFIX_PR
				default:
					panic(errors.New(num_fmt_part_err + "P"))
				}
			case 'R':
				// 判断独占 长度 FIXME
				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'N':
					if fmtDesc.suffix == NUM_FMT_SUFFIX_EMPTY {
						panic(errors.New("只能有1个 RN"))
					} else if fmtDesc.auxPrefix == NUM_FMT_AUX_PREFIX_FM && formatLen == 4 {
						panic(errors.New("包含RN的格式,除了 FM 和 RN 不能有其他格式字符"))
					} else if fmtDesc.auxPrefix == NUM_FMT_AUX_PREFIX_EMPTY && formatLen == 2 {
						panic(errors.New("包含RN的格式,除了 FM 和 RN 不能有其他格式字符"))
					}

					fmtDesc.suffix = NUM_FMT_SUFFIX_RN
				default:
					panic(errors.New(num_fmt_part_err + "R"))
				}
			case 'S':
				if fmtDesc.s == NUM_FMT_S_EMPTY {
					panic(errors.New("只能有1个 S"))
				} else if i == lastFormatIndex && i != 0 {
					panic(errors.New("S 只能在开头或者结尾"))
				}

				if i == 0 {
					fmtDesc.s = NUM_FMT_S_START
				} else {
					fmtDesc.s = NUM_FMT_S_END
				}
			case 'T':
				if fmtDesc.suffix == NUM_FMT_SUFFIX_EMPTY {
					i++
					followingOneChar := format[i]
					if 'M' == followingOneChar {
						if i == lastFormatIndex {
							fmtDesc.suffix = NUM_FMT_SUFFIX_TM
						} else {
							i++
							followingOneChar = format[i]
							if 'E' == followingOneChar {
								fmtDesc.suffix = NUM_FMT_SUFFIX_TME
							} else if 'e' == followingOneChar { // FIXME 已经转换为了大写
								fmtDesc.suffix = NUM_FMT_SUFFIX_TMe
							} else if '9' == followingOneChar {
								fmtDesc.suffix = NUM_FMT_SUFFIX_TM9
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
				if fmtDesc.suffix == NUM_FMT_SUFFIX_EMPTY {
					panic(errors.New("只能有1个 V"))
				} else if 0 != i {
					panic(errors.New("V 不能在开头"))
				}
				fmtDesc.suffix = NUM_FMT_SUFFIX_V
			case 'X':
				if fmtDesc.suffix == NUM_FMT_SUFFIX_EMPTY {
					panic(errors.New("conflict with X"))
				} else if 0 != i {
					panic(errors.New("V 不能在开头"))
				}

				fmtDesc.suffix = NUM_FMT_SUFFIX_X
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
func parseNumParam(num string) NumParamDesc {
	var paramDesc NumParamDesc

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
				paramDesc.existDec = true
			} else {
				panic("多个符号 " + ".")
			}
		case 'e', 'E':
			i++

			paramDesc.isEEEE = true
			var exponent = bytes.Buffer{}

			if num[i] == '+' {
				paramDesc.eSign = plus
			} else if num[i] == '-' {
				paramDesc.eSign = minus
			} else if num[i] <= '9' && num[i] >= '0' {
				paramDesc.eSign = empty
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
			paramDesc.eExponent = exponentNum
			fmt.Println(exponent.String())
		case '-':
			if i == 0 {
				paramDesc.sign = minus
			} else {
				panic("符号位置不对 " + "-")
			}
		case '+':
			if i == 0 {
				paramDesc.sign = plus
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

	paramDesc.preDec = preBuf.String()
	paramDesc.postDec = postBuf.String()
	return paramDesc
}

func parseNum(f string, num string) string {
	numFmtDesc := parseNumFormat(f)
	numFmtDescStr := fmt.Sprintf("%#v\n", numFmtDesc)
	fmt.Println(numFmtDescStr)

	numParamDesc := parseNumParam(num)
	numParamStr := fmt.Sprintf("%#v\n", numParamDesc)
	fmt.Println(numParamStr)

	result := bytes.Buffer{}

	// FM 模式，去除空格
	if numFmtDesc.auxPrefix == NUM_FMT_AUX_PREFIX_FM {

	}

	switch numFmtDesc.suffix {
	// 十进制
	case NUM_FMT_SUFFIX_EMPTY:
		// 转换科学计数
		if numParamDesc.isEEEE {
			d, err := strconv.ParseFloat(numParamDesc.preDec+"."+numParamDesc.postDec, 64)
			if err != nil {
				panic(err)
			}

			d = d * math.Pow10(numParamDesc.eExponent)
		}

		paramLen := len(numParamDesc.preDec)
		if numFmtDesc.preDec >= paramLen {
			for i := 0; i < numFmtDesc.preDec; i++ {
				result.WriteByte(numParamDesc.preDec[i])
			}

			if numFmtDesc.postDec > 0 {
				result.WriteByte('.')
				for i := 0; i < numFmtDesc.postDec; i++ {
					result.WriteByte(numParamDesc.postDec[i])
				}
			}
		} else {
			panic("格式的整数部分的长度不能比参数的整数部分的长度小")
		}
	// 科学计数
	case NUM_FMT_SUFFIX_EEEE:
		// 转换科学计数
		if numParamDesc.isEEEE {

		} else {

		}
	// 乘积
	case NUM_FMT_SUFFIX_V:
		if numParamDesc.isEEEE {
			panic("格式V 不支持科学计数参数")
		}
	// 罗马数字
	case NUM_FMT_SUFFIX_RN:
		if numParamDesc.isEEEE {
			panic("格式RN 不支持科学计数参数")
		}
	// 十六进制
	case NUM_FMT_SUFFIX_X:
		if numParamDesc.isEEEE {
			panic("格式X 不支持科学计数参数")
		}
	// 十进制 少写几个9 最小文本匹配 没有修饰符逗号，货币符号等
	case NUM_FMT_SUFFIX_TM, NUM_FMT_SUFFIX_TM9:
		if numParamDesc.isEEEE {

		}
	// 科学计数 少写几个9 最小文本匹配 没有修饰符逗号，货币符号等
	case NUM_FMT_SUFFIX_TME, NUM_FMT_SUFFIX_TMe:
		if numParamDesc.isEEEE {

		}
	default:
		panic("Theoretically unreachable")
	}

	return result.String()
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
