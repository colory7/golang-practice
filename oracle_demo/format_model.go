package oracle_demo

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"time"
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

type FMKeyword string

const (
	// Format Model关键词
	// Number Format Model Keyword
	NUM_COMMA  FMKeyword = ","
	NUM_DEC    FMKeyword = "."
	NUM_DOLLAR FMKeyword = "$"
	NUM_0      FMKeyword = "0"
	NUM_9      FMKeyword = "9"
	NUM_B      FMKeyword = "B"
	NUM_C      FMKeyword = "C"
	NUM_D      FMKeyword = "D"
	NUM_E      FMKeyword = "EEEE"
	NUM_FM     FMKeyword = "FM"
	NUM_G      FMKeyword = "G"
	NUM_L      FMKeyword = "L"
	NUM_MI     FMKeyword = "MI"
	NUM_PR     FMKeyword = "PR"
	NUM_RN     FMKeyword = "RN"
	NUM_S      FMKeyword = "S"
	NUM_TM     FMKeyword = "TM"
	NUM_TM9    FMKeyword = "TM9"
	NUM_TME    FMKeyword = "TME"
	NUM_U      FMKeyword = "U"
	NUM_V      FMKeyword = "V"
	NUM_X      FMKeyword = "X"

	// Datetime Format Model Keyword
	DCH_MINUS        FMKeyword = "-"
	DCH_SLASH        FMKeyword = "/"
	DCH_COMMA        FMKeyword = ","
	DCH_SEMICOLON    FMKeyword = ";"
	DCH_COLON        FMKeyword = ":"
	DCH_DOUBLE_QUOTE FMKeyword = "\""
	DCH_AD           FMKeyword = "AD"
	DCH_A_D          FMKeyword = "A.D."
	DCH_AM           FMKeyword = "AM"
	DCH_A_M          FMKeyword = "A.M."
	DCH_BC           FMKeyword = "BC"
	DCH_B_C          FMKeyword = "B.C."
	DCH_CC           FMKeyword = "CC"
	DCH_SCC          FMKeyword = "SCC"
	DCH_DAY          FMKeyword = "DAY"
	DCH_DDD          FMKeyword = "DDD"
	DCH_DD           FMKeyword = "DD"
	DCH_DY           FMKeyword = "DY"
	DCH_DL           FMKeyword = "DL"
	DCH_DS           FMKeyword = "DS"
	DCH_D            FMKeyword = "D"
	DCH_E            FMKeyword = "E"
	DCH_EE           FMKeyword = "EE"
	DCH_FF1          FMKeyword = "FF1"
	DCH_FF2          FMKeyword = "FF2"
	DCH_FF3          FMKeyword = "FF3"
	DCH_FF4          FMKeyword = "FF4"
	DCH_FF5          FMKeyword = "FF5"
	DCH_FF6          FMKeyword = "FF6"
	DCH_FF7          FMKeyword = "FF7"
	DCH_FF8          FMKeyword = "FF8"
	DCH_FF9          FMKeyword = "FF9"
	DCH_FM           FMKeyword = "FM"
	DCH_FX           FMKeyword = "FX"
	DCH_HH24         FMKeyword = "HH24"
	DCH_HH12         FMKeyword = "HH12"
	DCH_HH           FMKeyword = "HH"
	DCH_IW           FMKeyword = "IW"
	DCH_IYYY         FMKeyword = "IYYY"
	DCH_IYY          FMKeyword = "IYY"
	DCH_IY           FMKeyword = "IY"
	DCH_I            FMKeyword = "I"
	DCH_J            FMKeyword = "J"
	DCH_MI           FMKeyword = "MI"
	DCH_MM           FMKeyword = "MM"
	DCH_MONTH        FMKeyword = "MONTH"
	DCH_MON          FMKeyword = "MON"
	DCH_P_M          FMKeyword = "P.M."
	DCH_PM           FMKeyword = "PM"
	DCH_Q            FMKeyword = "Q"
	DCH_RM           FMKeyword = "RM"
	DCH_RR           FMKeyword = "RR"
	DCH_SSSSS        FMKeyword = "SSSSS"
	DCH_SSSS         FMKeyword = "SSSS"
	DCH_SS           FMKeyword = "SS"
	DCH_TZH          FMKeyword = "TZH"
	DCH_TZM          FMKeyword = "TZM"
	DCH_TZD          FMKeyword = "TZD"
	DCH_TZR          FMKeyword = "TZR"
	DCH_TS           FMKeyword = "TS"
	DCH_WW           FMKeyword = "WW"
	DCH_W            FMKeyword = "W"
	DCH_X            FMKeyword = "X"
	DCH_Y_YYY        FMKeyword = "Y,YYY"
	DCH_YEAR         FMKeyword = "YEAR"
	DCH_SYEAR        FMKeyword = "SYEAR"
	DCH_YYYY         FMKeyword = "YYYY"
	DCH_SYYYY        FMKeyword = "SYYYY"
	DCH_YYY          FMKeyword = "YYY"
	DCH_YY           FMKeyword = "YY"
	DCH_Y            FMKeyword = "Y"
)

var dchSkip = [...]uint8{'-', '/', ',', '.', ';', ':', '"'}

type keyword struct {
	key        FMKeyword
	replaceLen int
}

var dchKeywords map[FMKeyword]int

var weekDays = map[time.Weekday]string{}
var month = map[time.Month]string{}

func init() {

	weekDays = map[time.Weekday]string{
		time.Sunday:    "星期日",
		time.Monday:    "星期一",
		time.Tuesday:   "星期二",
		time.Wednesday: "星期三",
		time.Thursday:  "星期四",
		time.Friday:    "星期五",
		time.Saturday:  "星期六",
	}

	month = map[time.Month]string{
		time.January:   "1月",
		time.February:  "2月",
		time.March:     "3月",
		time.April:     "4月",
		time.May:       "5月",
		time.June:      "6月",
		time.July:      "7月",
		time.August:    "8月",
		time.September: "9月",
		time.October:   "10月",
		time.November:  "11月",
		time.December:  "12月",
	}

	dchKeywords = map[FMKeyword]int{
		DCH_A_D:   len(NLS_A_D_),
		DCH_AD:    len(NLS_AD),
		DCH_AM:    len(NLS_AM),
		DCH_A_M:   len(NLS_A_M_),
		DCH_BC:    len(NLS_BC),
		DCH_B_C:   len(NLS_B_C_),
		DCH_CC:    2,
		DCH_DAY:   3,
		DCH_DDD:   3,
		DCH_DD:    2,
		DCH_DY:    2,
		DCH_D:     1,
		DCH_FF1:   3,
		DCH_FF2:   3,
		DCH_FF3:   3,
		DCH_FF4:   3,
		DCH_FF5:   3,
		DCH_FF6:   3,
		DCH_FF7:   3,
		DCH_FF8:   3,
		DCH_FF9:   3,
		DCH_FX:    2,
		DCH_HH24:  4,
		DCH_HH12:  4,
		DCH_HH:    2,
		DCH_IW:    2,
		DCH_IYYY:  4,
		DCH_IYY:   3,
		DCH_IY:    2,
		DCH_I:     1,
		DCH_J:     1,
		DCH_MI:    2,
		DCH_MM:    2,
		DCH_MONTH: 5,
		DCH_MON:   3,
		DCH_P_M:   4,
		DCH_PM:    2,
		DCH_Q:     1,
		DCH_RM:    2,
		DCH_SSSSS: 5,
		DCH_SSSS:  4,
		DCH_SS:    2,
		DCH_TZH:   3,
		DCH_TZM:   3,
		DCH_WW:    2,
		DCH_W:     1,
		DCH_Y_YYY: 5,
		DCH_YYYY:  4,
		DCH_YYY:   3,
		DCH_YY:    2,
		DCH_Y:     1,
	}
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

	// 十进制 符号可选
	// 十进制 符号可选 逗号分组

	// 科学计数 转换为 十进制
	// TODO 提前转换 还是不转换 十进制
	if paramDesc.isEEEE {
		if preBuf.Len() > 0 {
			paramDesc.preDec = preBuf.String()
		} else {
			paramDesc.preDec = "0"
		}

		if postBuf.Len() > 0 {
			paramDesc.postDec = postBuf.String()
		}
		ff := paramDesc.preDec + "." + paramDesc.postDec
		fmt.Println(ff)
		d, err := strconv.ParseFloat(ff, 64)
		if err != nil {
			panic(err)
		}
		fff := "%" + paramDesc.preDec + "." + paramDesc.postDec + "f"
		v := fmt.Sprintf(fff, d)
		fmt.Println(v)

	} else {
		paramDesc.preDec = preBuf.String()
		paramDesc.postDec = postBuf.String()
	}

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

	var bs []byte

	// FM 模式，去除空格

	wroteSign := false
	switch numFmtDesc.suffix {
	// 十进制
	case NUM_FMT_SUFFIX_EMPTY:
		paramLen := len(numParamDesc.preDec)
		if numFmtDesc.preDec >= paramLen {
			for i := numFmtDesc.preDec; i >= 0; i-- {
				if i < paramLen {
					if !wroteSign {
						if numFmtDesc.auxPrefix == NUM_FMT_AUX_PREFIX_FM {
							if numParamDesc.sign == minus {
								bs = append(bs, '-')
							}
						} else {
							if numParamDesc.sign == minus {
								bs = append(bs, '-')
							} else {
								bs = append(bs, ' ')
							}
						}
						wroteSign = true
					}
					bs = append(bs, numParamDesc.preDec[paramLen-i-1])
				} else {
					bs = append(bs, ' ')
				}
			}

			if numFmtDesc.postDec > 0 {
				bs = append(bs, '.')
				for i := 0; i < numFmtDesc.postDec; i++ {
					bs = append(bs, numParamDesc.postDec[i])
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
	// 十进制
	// 少写几个9 最小文本匹配 没有修饰符逗号，货币符号等
	// 超过64个字符则转换为科学计数
	case NUM_FMT_SUFFIX_TM, NUM_FMT_SUFFIX_TM9:
		if numParamDesc.isEEEE {

		}
	// 科学计数
	// 少写几个9 最小文本匹配 没有修饰符逗号，货币符号等
	// 小数部分最多36个字符，整数部分最多1个字符
	case NUM_FMT_SUFFIX_TME, NUM_FMT_SUFFIX_TMe:
		if numParamDesc.isEEEE {

		}
	default:
		panic("Theoretically unreachable")
	}

	result.Write(bs)
	return result.String()
}

const (
	NLS_AD   = "公元"
	NLS_A_D_ = "公元"
	NLS_AM   = "上午"
	NLS_A_M_ = "上午"
	NLS_BC   = "公元前"
	NLS_B_C_ = "公元前"
	NLS_PM   = "下午"
	NLS_P_M_ = "下午"
	NLS_DL   = "YYYY年MM月DD日 DAY"
	NLS_DS   = "YYYY-MM-DD"
	NLS_X    = "."
)

// 解析日期格式
// 可以适当考虑使用字典树实现
func parseDchByStr(param string, format string) string {
	//var keywordGroup = make([]keyword, 4)

	result := bytes.Buffer{}
	flen := len(format)

	pi := 0
	for fi := 0; fi < flen; {
		// 截取一个字符
		c := format[fi]
		if c >= 32 && c <= 127 {
			fmt.Println("c: " + (string)(c))

			frest := flen - fi

			// 不区分大小写 FIXME
			ToUpper(&c)

			// 匹配关键词并存储
			switch c {
			// 跳过字符
			case '-', '/', ',', '.', ';', ':', ' ':
				fmt.Println("skip: " + string(c))
				fmt.Println(fi)
				pc := param[pi]
				switch pc {
				case '-', '/', ',', '.', ';', ':', ' ':
					fmt.Println("TODO: skip char " + string(pc))
					fi++
					pi++
				default:
					panic("不匹配的字符: " + string(pc))
				}
			case '"':
				var skipWord bytes.Buffer
				for ; fi < flen; fi++ {
					if '"' == format[fi] {
						break
					} else {
						skipWord.WriteByte(format[fi])
					}
				}
				result.Write(skipWord.Bytes())
				// 日期类型参数 输出
				// 字符串类型参数 只做匹配
			case 'A':
				fi++
				followingOneChar := format[fi]
				switch followingOneChar {
				case '.':
					fi++
					j := fi + 2
					if j <= len(format) {
						followingChars := format[fi:j]
						if "D." == followingChars {
							// DCH A.D.
							pe := pi + len(NLS_AD)
							v := param[pi:pe]
							if v == NLS_AD {
								fmt.Println("TODO: " + NLS_AD)
								result.WriteString(NLS_AD)
								fi = j
								pi = pe
							} else {
								panic("语法错误,参数与 A.D. 格式不匹配")
							}
						} else if "M." == followingChars {
							// DCH A.M.
							pe := pi + len(NLS_AD)
							v := param[pi:pe]
							if v == NLS_AM {
								fmt.Println("TODO: " + NLS_AM)
								result.WriteString(NLS_AM)
								fi = j
								pi = pe

							} else {
								panic("语法错误,参数与 A.M. 格式不匹配")
							}
						} else {
							panic(errors.New(dch_fmt_part_err + "A."))
						}
					} else {
						panic(errors.New(dch_fmt_part_err + "A."))
					}
				case 'D':
					// DCH AD
					fi++
					pe := pi + len(NLS_AD)
					v := param[pi:pe]
					if v == NLS_AD {
						fmt.Println("TODO: " + NLS_AD)
						result.WriteString(NLS_AD)
						pi = pe
					} else {
						panic("语法错误,参数与 AD 格式不匹配")
					}
				case 'M':
					// DCH AM
					fi++
					pe := pi + len(NLS_AD)
					v := param[pi:pe]
					if v == NLS_AM {
						fmt.Println("TODO: " + NLS_AM)
						result.WriteString(NLS_AM)
						pi = pe
					} else {
						panic("语法错误,参数与 AM 格式不匹配")
					}
				default:
					panic(errors.New(dch_fmt_part_err + "A"))
				}
				// 同上
			case 'B':
				fi++
				followingOneChar := format[fi]
				switch followingOneChar {
				case 'C':
					result.WriteString(NLS_BC)
				case '.':
					fe := fi + 4
					followingChars := format[fi:fe]
					if ".C." == followingChars {
						result.WriteString(NLS_BC)
					}
					fi = fe
				default:
					panic(errors.New(dch_fmt_part_err + "B"))
				}
				// 只适用于时间类型参数
			case 'C':
				fi++
				followingOneChar := format[fi]
				switch followingOneChar {
				case 'C':
					year := time.Now().Year()
					result.WriteString(strconv.Itoa((year + 99) / 100))
				default:
					panic(errors.New(dch_fmt_part_err + "C"))
				}
				// 字符串类型参数的 D 周中的日和julia 冲突
				// 时间类型参数的 D
			case 'D':
				fi++

				// DAY 同 DY
				if frest >= 2 && format[fi:fi+2] == "AY" || frest >= 1 && format[fi] == 'Y' {
					weekDay := time.Now().Weekday()
					result.WriteString(weekDays[weekDay])
				} else if frest >= 1 && format[fi] == 'D' {
					day := time.Now().Day()
					result.WriteString(strconv.Itoa(day))
				} else if frest >= 1 && format[fi] == 'L' {
					// TODO 插入格式 continue
					result.WriteString(parseDchByStr("???", NLS_DL))

				} else if frest >= 1 && format[fi] == 'S' {
					// TODO 插入格式 continue
					result.WriteString(parseDchByStr("???", NLS_DS))

				} else {
					weekDay := time.Now().Weekday()
					result.WriteString(strconv.Itoa(int(weekDay)))
				}

			case 'F':
				fi++
				followingOneChar := format[fi]
				switch followingOneChar {
				case 'X':
					//keywordGroup = append(keywordGroup, DCH_FX)
				case 'M':
					//FIXME
				default:
					followingTwoChars := format[fi : fi+3]
					fi = fi + 3
					switch followingTwoChars {
					case "F1":
						//keywordGroup = append(keywordGroup, DCH_FF1)
					case "F2":
						//keywordGroup = append(keywordGroup, DCH_FF2)
					case "F3":
						//keywordGroup = append(keywordGroup, DCH_FF3)
					case "F4":
						//keywordGroup = append(keywordGroup, DCH_FF4)
					case "F5":
						//keywordGroup = append(keywordGroup, DCH_FF5)
					case "F6":
						//keywordGroup = append(keywordGroup, DCH_FF6)
					case "F7":
						//keywordGroup = append(keywordGroup, DCH_FF7)
					case "F8":
						//keywordGroup = append(keywordGroup, DCH_FF8)
					case "F9":
						//keywordGroup = append(keywordGroup, DCH_FF9)
					default:
						panic(errors.New(dch_fmt_part_err + "F"))
					}
				}
			case 'H':
				fi++
				followingOneChar := format[fi]
				switch followingOneChar {
				case 'H':
					//keywordGroup = append(keywordGroup, DCH_HH)
				default:
					followingThreeChars := format[fi : fi+4]
					fi = fi + 4

					switch followingThreeChars {
					case "H24":
						//keywordGroup = append(keywordGroup, DCH_HH24)
					case "H12":
						//keywordGroup = append(keywordGroup, DCH_HH12)
					default:
						panic(errors.New(dch_fmt_part_err + "H"))
					}
				}
			case 'I':
				fi++
				followingOneChar := format[fi]
				switch followingOneChar {
				case 'W':
					//keywordGroup = append(keywordGroup, DCH_IW)
				case 'Y':
					followingTwoChars := format[fi : fi+3]
					if "YY" == followingTwoChars {
						//keywordGroup = append(keywordGroup, DCH_IYYY)
						fi = fi + 3
					}

					followingOneChar = followingTwoChars[0]
					if 'Y' == followingOneChar {
						//keywordGroup = append(keywordGroup, DCH_IYY)
						fi = fi + 2
					}
					//keywordGroup = append(keywordGroup, DCH_IY)
				default:
					panic(errors.New(dch_fmt_part_err + "I"))
				}
				// 匹配单个字符
				//keywordGroup = append(keywordGroup, DCH_I)
			case 'J':
				t := time.Now()
				result.WriteString(strconv.Itoa(ToJulian(t.Year(), int(t.Month()), t.Day())))
			case 'M':
				t := time.Now()
				fi++
				if fi <= flen && format[fi] == 'I' {
					// DCH MI
					result.WriteString(strconv.Itoa(t.Minute()))
				} else if fi <= flen && format[fi] == 'M' {
					// DCH MM
					if t.Month() < 10 {
						result.WriteString("0")
					}
					result.WriteString(strconv.Itoa(int(t.Month())))
				} else if fi <= flen && format[fi] == 'O' {
					fi++
					if fi <= flen && format[fi] == 'N' {
						fe := fi + 2
						if fi <= flen && format[fi:fe] == "TH" {
							// DCH MONTH
							fi = fe
						} else {
							// DCH MON
						}
						result.WriteString(month[t.Month()])
					} else {
						panic(errors.New(dch_fmt_part_err + "MO"))
					}
				} else {
					panic(errors.New(dch_fmt_part_err + "M"))
				}
			case 'P':
				fi++
				if fi < flen {
					if 'M' == format[fi] {
						//keywordGroup = append(keywordGroup, DCH_PM)
					} else {
						start := fi
						fi += 3
						if fi < flen {
							followingThreeChars := format[start:fi]
							if ".M." == followingThreeChars {
								//keywordGroup = append(keywordGroup, DCH_P_M)
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
				t := time.Now()
				result.WriteString(strconv.Itoa(int(t.Month()+2) / 3))
			case 'R':
				fi++
				if 'M' == format[fi] {
					//keywordGroup = append(keywordGroup, DCH_RM)
				} else {
					panic(errors.New(dch_fmt_part_err + "R"))
				}
			case 'S':
				start := fi
				fi += 4
				followingFourChars := format[start:fi]
				if "SSSS" == followingFourChars {
					//keywordGroup = append(keywordGroup, DCH_SSSSS)
				} else if "SSS" == followingFourChars[0:3] {
					//keywordGroup = append(keywordGroup, DCH_SSSS)
				} else if "S" == followingFourChars[0:3] {
					//keywordGroup = append(keywordGroup, DCH_SS)
				} else {
					panic(errors.New(dch_fmt_part_err + "S"))
				}
			case 'T':
				start := fi
				fi += 2
				followingTwoChars := format[start:fi]
				if "ZH" == followingTwoChars {
					//keywordGroup = append(keywordGroup, DCH_TZH)
				} else if "ZM" == followingTwoChars {
					//keywordGroup = append(keywordGroup, DCH_TZM)
				} else if 'Z' == followingTwoChars[0] {
					//keywordGroup = append(keywordGroup, DCH_TZM)
				} else {
					panic(errors.New(dch_fmt_part_err + "T"))
				}
			case 'W':
				fi++
				if format[fi] == 'W' {
					fi++

				} else {

				}
			case 'Y':
				fi++
				start := fi
				frest--
				if frest >= 4 {
					fi += 4
				} else {
					fi += frest
				}
				followingNChars := format[start:fi]
				missed := true

				if frest >= 4 && (",YYY" == followingNChars[0:4] || ",yyy" == followingNChars[0:4]) {
					// DCH Y,YYY
					pe := pi + 5
					v := param[pi:pe]
					result.WriteString(v)
					missed = false
				}
				if missed && frest >= 3 && ("YYY" == followingNChars[0:3] || "yyy" == followingNChars[0:3]) {
					// DCH YYYY
					pe := pi + 4
					v := param[pi:pe]
					result.WriteString(v)
					missed = false
				}
				if missed && frest >= 2 && ("YY" == followingNChars[0:2] || "yy" == followingNChars[0:4]) {
					// DCH YYY
					pe := pi + 3
					v := param[pi:pe]
					result.WriteString(v)
					missed = false
				}
				if missed && frest >= 1 && ("Y" == followingNChars[0:1] || "y" == followingNChars[0:4]) {
					// DCH YY
					pe := pi + 2
					v := param[pi:pe]
					result.WriteString(v)
					missed = false
				}

				if missed {
					// DCH Y
					pe := pi + 1
					v := param[pi:pe]
					result.WriteString(v)
				}

			default:
				panic(errors.New(out_keyword_range_err))
			}
		} else {
			panic(errors.New(out_ascii_range_err))
		}
	}

	return result.String()
}

// 解析日期格式
func parseDchByTime(t time.Time, format string) string {
	result := bytes.Buffer{}
	flen := len(format)

	for fi := 0; fi < flen; {
		// 截取一个字符
		c := format[fi]
		if c >= 32 && c <= 127 {
			fmt.Println("c: " + (string)(c))

			frest := flen - fi
			switch c {
			// 跳过字符
			case '-', '/', ',', '.', ';', ':', ' ':
				result.WriteByte(c)
			case '"':
				for ; fi < flen; fi++ {
					if '"' == format[fi] {
						break
					} else {
						result.WriteByte(format[fi])
					}
				}
			case 'A':
				fi++
				followingOneChar := format[fi]
				switch followingOneChar {
				case '.':
					fi++
					j := fi + 2
					if j <= len(format) {
						followingChars := format[fi:j]
						if "D." == followingChars {
							// DCH A.D.
							result.WriteString(NLS_AD)
							fi = j
						} else if "M." == followingChars {
							// DCH A.M.
							result.WriteString(NLS_AM)
							fi = j
						} else {
							panic(errors.New(dch_fmt_part_err + "A."))
						}
					} else {
						panic(errors.New(dch_fmt_part_err + "A."))
					}
				case 'D':
					// DCH AD
					fi++
					result.WriteString(NLS_AD)
				case 'M':
					// DCH AM
					fi++
					result.WriteString(NLS_AM)
				default:
					panic(errors.New(dch_fmt_part_err + "A"))
				}
			case 'B':
				fi++
				followingOneChar := format[fi]
				switch followingOneChar {
				case 'C':
					result.WriteString(NLS_BC)
				case '.':
					fe := fi + 4
					followingChars := format[fi:fe]
					if ".C." == followingChars {
						result.WriteString(NLS_BC)
					}
					fi = fe
				default:
					panic(errors.New(dch_fmt_part_err + "B"))
				}
			case 'C':
				fi++
				followingOneChar := format[fi]
				switch followingOneChar {
				case 'C':
					result.WriteString(strconv.Itoa((t.Year() + 99) / 100))
				default:
					panic(errors.New(dch_fmt_part_err + "C"))
				}
			case 'D':
				fi++

				// DAY 同 DY
				if frest >= 2 && format[fi:fi+2] == "AY" || frest >= 1 && format[fi] == 'Y' {
					weekDay := time.Now().Weekday()
					result.WriteString(weekDays[weekDay])
				} else if frest >= 1 && format[fi] == 'D' {
					day := time.Now().Day()
					result.WriteString(strconv.Itoa(day))
				} else if frest >= 1 && format[fi] == 'L' {
					// TODO 插入格式 continue
					result.WriteString(parseDchByStr("???", NLS_DL))

				} else if frest >= 1 && format[fi] == 'S' {
					// TODO 插入格式 continue
					result.WriteString(parseDchByStr("???", NLS_DS))

				} else {
					weekDay := t.Weekday()
					result.WriteString(strconv.Itoa(int(weekDay)))
				}
			case 'E':
				// TODO EE E
			case 'F':
				fi++
				followingOneChar := format[fi]
				switch followingOneChar {
				case 'X':
					// TODO
				case 'M':
					// TODO
				default:
					followingTwoChars := format[fi : fi+3]
					fi = fi + 3
					switch followingTwoChars {
					case "F1":
						result.WriteString(strconv.Itoa(t.Nanosecond() / 1e8))
					case "F2":
						result.WriteString(strconv.Itoa(t.Nanosecond() / 1e7))
					case "F3":
						result.WriteString(strconv.Itoa(t.Nanosecond() / 1e6))
					case "F4":
						result.WriteString(strconv.Itoa(t.Nanosecond() / 1e5))
					case "F5":
						result.WriteString(strconv.Itoa(t.Nanosecond() / 1e4))
					case "F6":
						result.WriteString(strconv.Itoa(t.Nanosecond() / 1e3))
					case "F7":
						result.WriteString(strconv.Itoa(t.Nanosecond() / 1e2))
					case "F8":
						result.WriteString(strconv.Itoa(t.Nanosecond() / 10))
					case "F9":
						result.WriteString(strconv.Itoa(t.Nanosecond()))
					default:
						panic(errors.New(dch_fmt_part_err + "F"))
					}
				}
			case 'H':
				fi++
				followingOneChar := format[fi]
				switch followingOneChar {
				case 'H':
					// DCH HH
					hour := t.Hour()
					if hour > 12 {
						result.WriteString(strconv.Itoa(hour - 12))
					} else {
						result.WriteString(strconv.Itoa(hour))
					}
				default:
					followingThreeChars := format[fi : fi+4]
					fi = fi + 4

					switch followingThreeChars {
					case "H24":
						// DCH HH24
						result.WriteString(strconv.Itoa(t.Hour()))
					case "H12":
						// DCH HH12
						hour := t.Hour()
						if hour > 12 {
							result.WriteString(strconv.Itoa(hour - 12))
						} else {
							result.WriteString(strconv.Itoa(hour))
						}
					default:
						panic(errors.New(dch_fmt_part_err + "H"))
					}
				}
			case 'I':
				fi++
				y, w := t.ISOWeek()

				if fi < flen {
					followingOneChar := format[fi]
					switch followingOneChar {
					case 'W':
						result.WriteString(strconv.Itoa(w))
					case 'Y':
						followingTwoChars := format[fi : fi+3]
						if "YY" == followingTwoChars {
							// DCH IYYY
							fi = fi + 3
							result.WriteString(strconv.Itoa(y))
						}

						followingOneChar = followingTwoChars[0]
						if 'Y' == followingOneChar {
							// DCH IYY
							result.WriteString(strconv.Itoa(y)[1:])
							fi = fi + 2
						} else {
							// DCH IY
							result.WriteString(strconv.Itoa(y)[2:])
						}
					}
				} else {
					// DCH I
					result.WriteString(strconv.Itoa(y)[3:])
				}
			case 'J':
				result.WriteString(strconv.Itoa(ToJulian(t.Year(), int(t.Month()), t.Day())))
			case 'M':
				fi++
				if fi <= flen && format[fi] == 'I' {
					// DCH MI
					result.WriteString(strconv.Itoa(t.Minute()))
				} else if fi <= flen && format[fi] == 'M' {
					// DCH MM
					if t.Month() < 10 {
						result.WriteString("0")
					}
					result.WriteString(strconv.Itoa(int(t.Month())))
				} else if fi <= flen && format[fi] == 'O' {
					fi++
					if fi <= flen && format[fi] == 'N' {
						fe := fi + 2
						if fi <= flen && format[fi:fe] == "TH" {
							// DCH MONTH
							fi = fe
						} else {
							// DCH MON
						}
						result.WriteString(month[t.Month()])
					} else {
						panic(errors.New(dch_fmt_part_err + "MO"))
					}
				} else {
					panic(errors.New(dch_fmt_part_err + "M"))
				}
			case 'P':
				fi++
				if fi < flen {
					if 'M' == format[fi] {
						result.WriteString(NLS_PM)
					} else {
						start := fi
						fi += 3
						if fi < flen {
							followingThreeChars := format[start:fi]
							if ".M." == followingThreeChars {
								result.WriteString(NLS_P_M_)
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
				result.WriteString(strconv.Itoa(int(t.Month()+2) / 3))
			case 'R':
				fi++
				if 'M' == format[fi] {
					result.WriteString(ToRoman(int(t.Month())).String())
				} else {
					panic(errors.New(dch_fmt_part_err + "R"))
				}

				// DCH RRRR
				result.WriteString(strconv.Itoa(t.Year()))
				// DCH RR
				result.WriteString(strconv.Itoa(t.Year())[2:])
			case 'S':
				// DCH SYEAR
				// DCH SYYYY
				// DCH SSSSS
				// DCH SP

				start := fi
				fi += 4
				followingFourChars := format[start:fi]
				if "SSSS" == followingFourChars {
					//keywordGroup = append(keywordGroup, DCH_SSSSS)
					//result.WriteString()
				} else if "S" == followingFourChars[0:3] {
					//keywordGroup = append(keywordGroup, DCH_SS)
				} else {
					panic(errors.New(dch_fmt_part_err + "S"))
				}
			case 'T':
				// DCH TS
				// DCH TZD
				// DCH TZH
				// DCH TZM
				// DCH TZR
				// DCH TH
			case 'W':
				fi++
				if format[fi] == 'W' {
					fi++
					// DCH WW
					result.WriteString(strconv.Itoa((t.YearDay() + 6) / 7))
				} else {
					// DCH W
					result.WriteString(strconv.Itoa((t.Day() + 6) / 7))
				}
			case 'X':
				result.WriteString(NLS_X)
			case 'Y':
				fi++

				year := strconv.Itoa(t.Year())

				// DCH Y,YYY
				result.WriteString(year[:1] + "," + year[1:])
				// DCH YEAR
				// DCH YYYY
				result.WriteString(year)
				// DCH YYY
				result.WriteString(year[1:])
				// DCH YY
				result.WriteString(year[2:])
				// DCH Y
				result.WriteString(year[3:])

			default:
				panic(errors.New(out_keyword_range_err))
			}
		} else {
			panic(errors.New(out_ascii_range_err))
		}
	}

	return result.String()
}

func ToUpper(c *byte) {
	if 'a' <= *c && *c <= 'z' {
		*c -= 32
	}
}

func ToRoman(num int) *bytes.Buffer {
	romes := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	numbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	rm := &bytes.Buffer{}
	for i := 0; i < len(numbers); i++ {
		for num >= numbers[i] {
			num -= numbers[i]
			rm.WriteString(romes[i])
		}
	}
	return rm
}

func ToJulian(year int, month int, day int) int {
	adj := (14 - month) / 12
	y := year + 4800 - adj
	m := month + 12*adj - 3
	return day + (153*m+2)/5 + y*365 + y/4 - y/100 + y/400 - 32045
}
