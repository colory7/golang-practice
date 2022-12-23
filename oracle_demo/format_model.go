package oracle_demo

import (
	"bytes"
	"errors"
	"fmt"
	"golang_practice/oracle_demo/nls"
	"strconv"
	"time"
	"unicode/utf8"
)

const skipCharSize = 32

// 格式部分不匹配，报错
const dch_fmt_mismatch_err = "Date Format error, some formats do not match near "
const dch_fmt_length_err = "Date Format error, incorrect format length near "
const num_fmt_part_err = "Datetime Format error, some formats do not match near "
const not_support_err = "not support"

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
	// Format Model关键词
	// Number Format Model Keyword
	NUM_COMMA = iota
	NUM_DEC
	NUM_DOLLAR
	NUM_0
	NUM_9
	NUM_B
	NUM_C
	NUM_D
	NUM_E
	NUM_FM
	NUM_G
	NUM_L
	NUM_MI
	NUM_PR
	NUM_RN
	NUM_S
	NUM_TM
	NUM_TM9
	NUM_TME
	NUM_U
	NUM_V
	NUM_X

	// Datetime Format Model Keyword
	DCH_COPY = iota
	DCH_MINUS
	DCH_SLASH
	DCH_COMMA
	DCH_DEC
	DCH_SEMICOLON
	DCH_COLON
	DCH_SPACE
	DCH_DOUBLE_QUOTE
	DCH_AD
	DCH_A_D_
	DCH_AM
	DCH_A_M_
	DCH_BC
	DCH_B_C_
	DCH_CC
	DCH_SCC
	DCH_DAY
	DCH_DDD
	DCH_DD
	DCH_DL
	DCH_DS
	DCH_DY
	DCH_D
	DCH_E
	DCH_EE
	DCH_FF1
	DCH_FF2
	DCH_FF3
	DCH_FF4
	DCH_FF5
	DCH_FF6
	DCH_FF7
	DCH_FF8
	DCH_FF9
	DCH_FF
	DCH_FM
	DCH_FX
	DCH_HH24
	DCH_HH12
	DCH_HH
	DCH_IW
	DCH_IYYY
	DCH_IYY
	DCH_IY
	DCH_I
	DCH_J
	DCH_MI
	DCH_MM
	DCH_MONTH
	DCH_MON
	DCH_P_M_
	DCH_PM
	DCH_Q
	DCH_RM
	DCH_RR
	DCH_RRRR
	DCH_SP
	DCH_SSSSS
	DCH_SS
	DCH_TZH
	DCH_TZM
	DCH_TZD
	DCH_TZR
	DCH_TS
	DCH_TH
	DCH_WW
	DCH_W
	DCH_X
	DCH_Y_YYY
	DCH_YEAR
	DCH_SYEAR
	DCH_YYYY
	DCH_SYYYY
	DCH_YYY
	DCH_YY
	DCH_Y
)

type keyword struct {
	key        int
	replaceLen int
}

var dchKeywords map[int]int

var NLS_WEEKS = map[time.Weekday]string{}
var NLS_MONTHS = map[time.Month]string{}

func init() {

	NLS_WEEKS = map[time.Weekday]string{
		time.Sunday:    "星期日",
		time.Monday:    "星期一",
		time.Tuesday:   "星期二",
		time.Wednesday: "星期三",
		time.Thursday:  "星期四",
		time.Friday:    "星期五",
		time.Saturday:  "星期六",
	}

	NLS_MONTHS = map[time.Month]string{
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

	dchKeywords = map[int]int{
		DCH_A_D_:  len(NLS_A_D_),
		DCH_AD:    len(NLS_AD),
		DCH_AM:    len(NLS_AM),
		DCH_A_M_:  len(NLS_A_M_),
		DCH_BC:    len(NLS_BC),
		DCH_B_C_:  len(NLS_B_C_),
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
		DCH_P_M_:  4,
		DCH_PM:    2,
		DCH_Q:     1,
		DCH_RM:    2,
		DCH_SSSSS: 5,
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

func (numParam *NumParamDesc) string() (string, error) {
	var result bytes.Buffer
	if plus == numParam.sign {
		result.WriteString(plus)
	} else if minus == numParam.sign {
		result.WriteString(minus)
	} else if empty == numParam.sign {
	} else {
		return empty_str, errors.New("sign属性格式错误")
	}

	if empty != numParam.preDec {
		result.WriteString(numParam.preDec)
	} else {
		return empty_str, errors.New("格式错误,整数部分是空的")
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
			return empty_str, errors.New("eSign属性格式错误")
		}
		result.WriteString(fmt.Sprint(numParam.eExponent))
	}

	return result.String(), nil
}

// 解析数值格式
func parseNumFormat(format string) (NumFmtDesc, error) {
	var fmtDesc NumFmtDesc

	// 格式字节长度
	flen := len(format)

	// 最后一个关键词的索引
	lastFormatIndex := flen - 1

	isRerun := false
	var c byte

	readDec := false
	//var preDec = bytes.Buffer{}
	//var postDec = bytes.Buffer{}
	for i := 0; i < flen; {
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
						return fmtDesc, errors.New("不能以逗号开头")
					} else if i == lastFormatIndex {
						return fmtDesc, errors.New("逗号不能出现在数字最右边")
					} else if fmtDesc.decIndex != -1 {
						return fmtDesc, errors.New("逗号不能出现在点号的右边")
					}

					fmtDesc.commaIndex = i
				} else {
					return fmtDesc, errors.New("格式错误，存在多个格式符号 ,")
				}

			case '.':
				if !readDec {
					fmtDesc.decIndex = i
					readDec = true
				} else {
					return fmtDesc, errors.New("只能有1个 .")
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
					return fmtDesc, errors.New("格式前缀冲突 " + "$")
				}
			case 'B':
				if fmtDesc.prefix == NUM_FMT_PREFIX_EMPTY {
					fmtDesc.prefix = NUM_FMT_PREFIX_B
				} else {
					return fmtDesc, errors.New("格式前缀冲突 " + "B")
				}
			case 'C':
				if fmtDesc.prefix == NUM_FMT_PREFIX_EMPTY {
					return fmtDesc, errors.New("格式前缀冲突 " + "C")
				} else if lastFormatIndex != i && 0 != i {
					return fmtDesc, errors.New("C 只能在开头或者结尾")
				}
				fmtDesc.prefix = NUM_FMT_PREFIX_C
			case 'L':
				if fmtDesc.prefix == NUM_FMT_PREFIX_EMPTY {
					fmtDesc.prefix = NUM_FMT_PREFIX_L
				} else {
					return fmtDesc, errors.New("格式前缀冲突 " + "L")
				}
			case 'U':
				if fmtDesc.prefix == NUM_FMT_PREFIX_EMPTY {
					return fmtDesc, errors.New("格式前缀冲突 " + "U")
				} else if lastFormatIndex != i && 0 != i {
					return fmtDesc, errors.New("U 只能在开头或者结尾")
				}
				fmtDesc.prefix = NUM_FMT_PREFIX_U
			case ntw.NLS_NUMERIC_CHARACTERS[0]:
			case ntw.NLS_NUMERIC_CHARACTERS[1]:
			case 'E':
				if fmtDesc.suffix == NUM_FMT_SUFFIX_EMPTY {
					j := i + 4
					followingChars := format[i+1 : j]
					if "EEE" == followingChars {
						i = j
						fmtDesc.suffix = NUM_FMT_SUFFIX_EEEE
					} else {
						return fmtDesc, errors.New(num_fmt_part_err + "E")
					}
				} else {
					return fmtDesc, errors.New("conflict with E")
				}
			case 'F':
				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'M':
					if fmtDesc.auxPrefix == NUM_FMT_AUX_PREFIX_EMPTY {
						return fmtDesc, errors.New("只能有1组 FM")
					}
					if 1 == i {
						fmtDesc.auxPrefix = NUM_FMT_AUX_PREFIX_FM
					} else {
						return fmtDesc, errors.New("FM 必须在开头")
					}
				default:
					return fmtDesc, errors.New(num_fmt_part_err + "F")
				}
			case 'M':
				if fmtDesc.auxSuffix == NUM_FMT_AUX_SUFFIX_EMPTY {
					return fmtDesc, errors.New("辅助后缀冲突" + "MI")
				} else if i == (lastFormatIndex - 1) {
					return fmtDesc, errors.New("MI 只能在结尾")
				}

				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'I':
					fmtDesc.auxSuffix = NUM_FMT_AUX_SUFFIX_MI
				default:
					return fmtDesc, errors.New(num_fmt_part_err + "M")
				}
			case 'P':
				if fmtDesc.auxSuffix == NUM_FMT_AUX_SUFFIX_EMPTY {
					return fmtDesc, errors.New("辅助后缀冲突" + "PR")
				} else if i == (lastFormatIndex - 1) {
					return fmtDesc, errors.New("PR 只能在结尾")
				}

				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'R':
					fmtDesc.auxSuffix = NUM_FMT_AUX_SUFFIX_PR
				default:
					return fmtDesc, errors.New(num_fmt_part_err + "P")
				}
			case 'R':
				// 判断独占 长度 FIXME
				i++
				followingOneChar := format[i]
				switch followingOneChar {
				case 'N':
					if fmtDesc.suffix == NUM_FMT_SUFFIX_EMPTY {
						return fmtDesc, errors.New("只能有1个 RN")
					} else if fmtDesc.auxPrefix == NUM_FMT_AUX_PREFIX_FM && flen == 4 {
						return fmtDesc, errors.New("包含RN的格式,除了 FM 和 RN 不能有其他格式字符")
					} else if fmtDesc.auxPrefix == NUM_FMT_AUX_PREFIX_EMPTY && flen == 2 {
						return fmtDesc, errors.New("包含RN的格式,除了 FM 和 RN 不能有其他格式字符")
					}

					fmtDesc.suffix = NUM_FMT_SUFFIX_RN
				default:
					return fmtDesc, errors.New(num_fmt_part_err + "R")
				}
			case 'S':
				if fmtDesc.s == NUM_FMT_S_EMPTY {
					return fmtDesc, errors.New("只能有1个 S")
				} else if i == lastFormatIndex && i != 0 {
					return fmtDesc, errors.New("S 只能在开头或者结尾")
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
								return fmtDesc, errors.New("格式错误在 TM 附近")
							}
						}
					} else {
						return fmtDesc, errors.New("格式错误在 T 附近")
					}
				} else {
					return fmtDesc, errors.New("只能有1组 TM")
				}
			case 'V':
				if fmtDesc.suffix == NUM_FMT_SUFFIX_EMPTY {
					return fmtDesc, errors.New("只能有1个 V")
				} else if 0 != i {
					return fmtDesc, errors.New("V 不能在开头")
				}
				fmtDesc.suffix = NUM_FMT_SUFFIX_V
			case 'X':
				if fmtDesc.suffix == NUM_FMT_SUFFIX_EMPTY {
					return fmtDesc, errors.New("conflict with X")
				} else if 0 != i {
					return fmtDesc, errors.New("V 不能在开头")
				}

				fmtDesc.suffix = NUM_FMT_SUFFIX_X
				fmtDesc.xCount++
			default:
				return fmtDesc, errors.New(out_keyword_range_err)
			}

		} else {
			return fmtDesc, errors.New(out_ascii_range_err)
		}

		i++
	}

	return fmtDesc, nil

}

// 解析数字参数
func parseNumParam(num string) (NumParamDesc, error) {
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
				return paramDesc, errors.New("多个符号 " + ".")
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
					return paramDesc, errors.New("科学计数的指数使用了非法字符 " + string(num[i]))
				}
			}

			exponentNum, err := strconv.Atoi(exponent.String())
			if err != nil {
				return paramDesc, err
			}
			paramDesc.eExponent = exponentNum
			fmt.Println(exponent.String())
		case '-':
			if i == 0 {
				paramDesc.sign = minus
			} else {
				return paramDesc, errors.New("符号位置不对 " + "-")
			}
		case '+':
			if i == 0 {
				paramDesc.sign = plus
			} else {
				return paramDesc, errors.New("符号位置不对 " + "+")
			}
		case ',':
			return paramDesc, errors.New("暂时不支持 " + ",")
		default:
			return paramDesc, errors.New("不支持的数字符号")
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
			return paramDesc, err
		}
		fff := "%" + paramDesc.preDec + "." + paramDesc.postDec + "f"
		v := fmt.Sprintf(fff, d)
		fmt.Println(v)

	} else {
		paramDesc.preDec = preBuf.String()
		paramDesc.postDec = postBuf.String()
	}

	return paramDesc, nil
}

func toNumber(num string, format string) (string, error) {
	numFmtDesc, err := parseNumFormat(format)
	if err != nil {
		return empty_str, err
	}
	numFmtDescStr := fmt.Sprintf("%#v\n", numFmtDesc)
	fmt.Println(numFmtDescStr)

	numParamDesc, err := parseNumParam(num)
	if err != nil {
		return empty_str, err
	}
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
			return empty_str, errors.New("格式的整数部分的长度不能比参数的整数部分的长度小")
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
			return empty_str, errors.New("格式V 不支持科学计数参数")
		}
	// 罗马数字
	case NUM_FMT_SUFFIX_RN:
		if numParamDesc.isEEEE {
			return empty_str, errors.New("格式RN 不支持科学计数参数")
		}
	// 十六进制
	case NUM_FMT_SUFFIX_X:
		if numParamDesc.isEEEE {
			return empty_str, errors.New("格式X 不支持科学计数参数")
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
		return empty_str, errors.New("Theoretically unreachable")
	}

	result.Write(bs)
	return result.String(), nil
}

const (
	SPACE       = " "
	ASSIC_SPACE = ' '
	NLS_AD      = "公元"
	NLS_A_D_    = "公元"
	NLS_AM      = "上午"
	NLS_A_M_    = "上午"
	NLS_BC      = "公元前"
	NLS_B_C_    = "公元前"
	NLS_PM      = "下午"
	NLS_P_M_    = "下午"
	NLS_DL      = "YYYY\"年\"MM\"月\"DD\"日\" DAY"
	NLS_DS      = "YYYY-MM-DD"
	NLS_X       = "."
)

const (
	empty_str = ""
	tsFormat  = "15:04:05"
	//dateFormat = "YYYY-MM-DD HH24:MI:SS"
	dateLayout = "2006-01-02 15:04:05"
)

func toDate(dch string, format string) (*time.Time, error) {
	fmKeywords, quoted, err := ParseDchByTime(format)
	if err != nil {
		return nil, nil
	}

	year, month, day := 0, time.Month(0), 0
	hour, min, sec, nsec := 0, 0, 0, 0

	di := 0
	qi := 0
	for ki := 0; ki < len(fmKeywords); ki++ {
		switch fmKeywords[ki] {
		case DCH_DOUBLE_QUOTE:
			di += len(quoted[qi])
			qi++
		case DCH_MINUS, DCH_SLASH, DCH_COMMA, DCH_SEMICOLON, DCH_COLON:
			di++
		case DCH_AD:
			di += 2
		case DCH_A_D_:
			di += 4
		case DCH_AM:
			di += 2
		case DCH_A_M_:
			di += 4
		case DCH_BC:
			di += 2
		case DCH_B_C_:
			di += 4
		case DCH_CC:
			return nil, errors.New(not_support_err)
		case DCH_SCC:
			return nil, errors.New(not_support_err)
		case DCH_DAY:
			return nil, errors.New(not_support_err)
		case DCH_DDD:
			return nil, errors.New(not_support_err)
		case DCH_DD:
			start := di
			di += 2
			day, err = strconv.Atoi(dch[start:di])
			if err != nil {
				return nil, err
			}
		case DCH_FM:
			//TODO
		case DCH_FX:
			//TODO
		case DCH_HH24, DCH_HH12, DCH_HH:
			start := di
			di += 2
			day, err = strconv.Atoi(dch[start:di])
			if err != nil {
				return nil, err
			}
		case DCH_MI:
		case DCH_MM:
		case DCH_MONTH, DCH_MON:
		case DCH_RR:
		case DCH_RRRR:
		case DCH_TZH:
		case DCH_TZM:
		case DCH_TZD:
		case DCH_TZR:
		case DCH_TS:
		case DCH_Y_YYY:
		case DCH_YYYY:
		case DCH_YYY:
		case DCH_YY:
		case DCH_Y:
		default:
			return nil, errors.New(not_support_err)
		}
	}

	if di != len(quoted) {
		panic("引号内容未遍历完，不匹配")
	}
	if qi != len(dch) {
		panic("参数未遍历完,不匹配")
	}

	now := time.Now()

	if year == 0 {
		year = now.Year()
	}
	if month == 0 {
		month = now.Month()
	}
	if day == 0 {
		day = 1
	}
	t := time.Date(year, month, day, hour, min, sec, nsec, time.Local)
	fmt.Println(t.Format(dateLayout))

	return &t, nil
}

func toChar(t time.Time, format string) (string, error) {
	fmKeywords, quoted, err := ParseDchByTime(format)
	if err != nil {
		return empty_str, nil
	}

	result := bytes.Buffer{}

	qi := 0
	for i := 0; i < len(fmKeywords); i++ {
		switch fmKeywords[i] {
		case DCH_DOUBLE_QUOTE:
			result.WriteString(quoted[qi])
			qi++
		case DCH_MINUS, DCH_SLASH, DCH_COMMA, DCH_SEMICOLON, DCH_COLON:
			result.WriteString(string(fmKeywords[i]))
		case DCH_AD, DCH_A_D_:
			result.WriteString(NLS_AD)
		case DCH_AM, DCH_A_M_:
			result.WriteString(NLS_AM)
		case DCH_BC:
			result.WriteString(NLS_BC)
		case DCH_B_C_:
			result.WriteString(NLS_B_C_)
		case DCH_CC:
			result.WriteString(strconv.Itoa((t.Year() + 99) / 100))
		case DCH_SCC:
			//TODO 公元前 正负号
		case DCH_DAY:
			result.WriteString(NLS_WEEKS[t.Weekday()])
		case DCH_DDD:
			result.WriteString(strconv.Itoa(t.YearDay()))
		case DCH_DD:
			result.WriteString(strconv.Itoa(t.Day()))
		case DCH_DL:
			tmp, err := toChar(t, NLS_DL)
			if err != nil {
				return empty_str, nil
			}
			result.WriteString(tmp)
		case DCH_DS:
			tmp, err := toChar(t, NLS_DS)
			if err != nil {
				return empty_str, nil
			}
			result.WriteString(tmp)
		case DCH_DY:
			result.WriteString(NLS_WEEKS[t.Weekday()])
		case DCH_D:
			result.WriteString(strconv.Itoa(int(t.Weekday())))
		case DCH_E:
			return empty_str, errors.New("not support")
		case DCH_EE:
			return empty_str, errors.New("not support")
		case DCH_FF1:
			result.WriteString(strconv.Itoa(t.Nanosecond() / 1e8))
		case DCH_FF2:
			result.WriteString(strconv.Itoa(t.Nanosecond() / 1e7))
		case DCH_FF3:
			result.WriteString(strconv.Itoa(t.Nanosecond() / 1e6))
		case DCH_FF4:
			result.WriteString(strconv.Itoa(t.Nanosecond() / 1e5))
		case DCH_FF5:
			result.WriteString(strconv.Itoa(t.Nanosecond() / 1e4))
		case DCH_FF6:
			result.WriteString(strconv.Itoa(t.Nanosecond() / 1e3))
		case DCH_FF7:
			result.WriteString(strconv.Itoa(t.Nanosecond() / 1e2))
		case DCH_FF8:
			result.WriteString(strconv.Itoa(t.Nanosecond() / 1e1))
		case DCH_FF9, DCH_FF:
			result.WriteString(strconv.Itoa(t.Nanosecond()))
		case DCH_FM:
		case DCH_FX:
		case DCH_HH24:
			if t.Hour() < 10 {
				result.WriteByte('0')
			}
			result.WriteString(strconv.Itoa(t.Hour()))
		case DCH_HH12, DCH_HH:
			hour := t.Hour()
			if hour > 12 {
				hour = hour - 12
			}
			if hour < 10 {
				result.WriteByte('0')
			}
			result.WriteString(strconv.Itoa(hour))
		case DCH_IW:
			_, w := t.ISOWeek()
			result.WriteString(strconv.Itoa(w))
		case DCH_IYYY:
			y, _ := t.ISOWeek()
			result.WriteString(strconv.Itoa(y))
		case DCH_IYY:
			y, _ := t.ISOWeek()
			result.WriteString(strconv.Itoa(y)[1:])
		case DCH_IY:
			y, _ := t.ISOWeek()
			result.WriteString(strconv.Itoa(y)[2:])
		case DCH_I:
			y, _ := t.ISOWeek()
			result.WriteString(strconv.Itoa(y)[3:])
		case DCH_J:
			result.WriteString(strconv.Itoa(ToJulian(t.Year(), int(t.Month()), t.Day())))
		case DCH_MI:
			result.WriteString(strconv.Itoa(t.Minute()))
		case DCH_MM:
			if t.Month() < 10 {
				result.WriteByte('0')
			}
			result.WriteString(strconv.Itoa(int(t.Month())))
		case DCH_MONTH, DCH_MON:
			result.WriteString(NLS_MONTHS[t.Month()])
		case DCH_P_M_:
			result.WriteString(NLS_P_M_)
		case DCH_PM:
			result.WriteString(NLS_PM)
		case DCH_Q:
			result.WriteString(strconv.Itoa(int(t.Month()+2) / 3))
		case DCH_RM:
			result.WriteString(ToRoman(int(t.Month())).String())
		case DCH_RR:
			result.WriteString(strconv.Itoa(t.Year())[2:])
		case DCH_RRRR:
			result.WriteString(strconv.Itoa(t.Year()))
		case DCH_SSSSS:
			result.WriteString(strconv.Itoa((t.Hour()*60+t.Minute())*60 + t.Second()))
		case DCH_SS:
			result.WriteString(strconv.Itoa(t.Second()))
		case DCH_TZH:
			result.WriteString(t.Format("-07"))
		case DCH_TZM:
			result.WriteString(t.Format("-0700")[3:])
		case DCH_TZD:
			zone, _ := t.Local().Zone()
			result.WriteString(zone)
		case DCH_TZR:
			result.WriteString(t.Location().String())
		case DCH_TS:
			if t.Hour() > 12 {
				result.WriteString(NLS_AM)
				result.WriteByte(ASSIC_SPACE)
			} else {
				result.WriteString(NLS_AM)
				result.WriteByte(ASSIC_SPACE)
			}
			result.WriteString(t.Format(tsFormat))
		case DCH_WW:
			result.WriteString(strconv.Itoa((t.YearDay() + 6) / 7))
		case DCH_W:
			result.WriteString(strconv.Itoa((t.Day() + 6) / 7))
		case DCH_X:
			result.WriteString(NLS_X)
		case DCH_Y_YYY:
			year := strconv.Itoa(t.Year())
			result.WriteString(year[:1] + "," + year[1:])
		case DCH_YEAR:
			result.WriteString(ntw.NumToCardinalWord(t.Year()))
		case DCH_SYEAR:
			result.WriteString(ntw.NumToCardinalWord(t.Year() / 100))
			result.WriteString(SPACE)
			result.WriteString(ntw.NumToCardinalWord(t.Year() % 100))
		case DCH_YYYY:
			year := strconv.Itoa(t.Year())
			result.WriteString(year)
		case DCH_SYYYY:
			result.WriteString(strconv.Itoa(t.Year()))
		case DCH_YYY:
			year := strconv.Itoa(t.Year())
			result.WriteString(year[1:])
		case DCH_YY:
			year := strconv.Itoa(t.Year())
			result.WriteString(year[2:])
		case DCH_Y:
			result.WriteString(strconv.Itoa(t.Year())[3:])
		default:
			return empty_str, errors.New("unrechable")
		}
	}

	return result.String(), nil
}

type FMKeyword int

// 解析日期格式
func ParseDchByTime(format string) ([]int, []string, error) {
	fmKeywords := []int{}

	flen := len(format)

	//println(format)

	//aux_flag_fm := false
	//aux_flag_fx := false
	//aux_flag_sp := false
	//aux_flag_th := false

	quoted := []string{}

	var keyword FMKeyword
	var err error
	for fi := 0; fi < flen; {
		// 截取一个字符
		c := format[fi]
		if c >= 32 && c <= 127 {
			//log.Println("debug: c-> " + (string)(c))
			//log.Println(result.String())
			//frest := flen - fi
			switch c {
			// DCH reproduced
			case '-':
				keyword = DCH_MINUS
			case '/':
				keyword = DCH_SLASH
			case ',':
				keyword = DCH_COMMA
			case '.':
				keyword = DCH_DEC
			case ';':
				keyword = DCH_SEMICOLON
			case ':':
				keyword = DCH_COLON
			case ' ':
				keyword = DCH_SPACE
				fi++
			// DCH 左双引号
			case '"':
				keyword = DCH_DOUBLE_QUOTE

				tmp := bytes.Buffer{}
				fi++
				for ; fi < flen; fi++ {
					if '"' == format[fi] {
						// DCH 右双引号
						quoted = append(quoted, tmp.String())
						break
					} else {
						// DCH 双引号中的内容
						tmp.WriteByte(format[fi])
					}
				}
			case 'A':
				keyword, err = parsePrefixA(&fi, flen, format)
			case 'B':
				keyword, err = parsePrefixB(&fi, flen, format)
			case 'C':
				keyword, err = parsePrefixC(&fi, flen, format)
			case 'D':
				keyword, err = parsePrefixD(&fi, flen, format)
			case 'E':
				// TODO EE E
				return nil, nil, errors.New("not support")
			case 'F':
				keyword, err = parsePrefixF(&fi, flen, format)
			case 'H':
				keyword, err = parsePrefixH(&fi, flen, format)
			case 'I':
				keyword, err = parsePrefixI(&fi, flen, format)
			case 'J':
				keyword = DCH_J
				fi++
			case 'M':
				keyword, err = parsePrefixM(&fi, flen, format)
			case 'P':
				keyword, err = parsePrefixP(&fi, flen, format)
			case 'Q':
				keyword = DCH_Q
				fi++
			case 'R':
				keyword, err = parsePrefixR(&fi, flen, format)
			case 'S':
				keyword, err = parsePrefixS(&fi, flen, format)
			case 'T':
				keyword, err = parsePrefixT(&fi, flen, format)
			case 'W':
				fi++
				if fi < flen && format[fi] == 'W' {
					// DCH WW
					keyword = DCH_WW
					fi++
				} else {
					// DCH W
					keyword = DCH_W
				}
			case 'X':
				keyword = DCH_X
				fi++
			case 'Y':
				keyword, err = parsePrefixY(&fi, flen, format)
			default:
				return nil, nil, errors.New(out_keyword_range_err)
			}

			if err != nil {
				return nil, nil, err
			}
			fmKeywords = append(fmKeywords, keyword)
		} else {
			return nil, nil, errors.New(out_ascii_range_err + string(c))
		}
	}

	//if aux_flag_fm {
	//
	//}
	//if aux_flag_fx {
	//
	//}
	//if aux_flag_sp {
	//
	//}
	//if aux_flag_th {
	//
	//}

	return fmKeywords, quoted, nil
}

func parsePrefixA(fi *int, flen int, format string) (FMKeyword, error) {
	var keyword FMKeyword
	*fi++
	if *fi < flen {
		followingOneChar := format[*fi]
		switch followingOneChar {
		case '.':
			*fi++
			start := *fi
			*fi += 2
			if *fi <= flen {
				followingChars := format[start:*fi]
				if "D." == followingChars {
					// DCH A.D.
					keyword = DCH_A_D_
				} else if "M." == followingChars {
					// DCH A.M.
					keyword = DCH_A_M_
				} else {
					return empty_str, errors.New(dch_fmt_mismatch_err + "A.")
				}
			} else {
				return empty_str, errors.New(dch_fmt_length_err + "A.")
			}
		case 'D':
			// DCH AD
			keyword = DCH_AD
			*fi++
		case 'M':
			// DCH AM
			keyword = DCH_AM
			*fi++
		default:
			return empty_str, errors.New(dch_fmt_mismatch_err + "A")
		}
	} else {
		return empty_str, errors.New(dch_fmt_length_err + "A")
	}
	return keyword, nil
}

func parsePrefixB(fi *int, flen int, format string) (FMKeyword, error) {
	var keyword FMKeyword
	*fi++
	if *fi < flen {
		followingOneChar := format[*fi]
		switch followingOneChar {
		case 'C':
			// DCH BC
			keyword = DCH_BC
			*fi++
		case '.':
			*fi++
			start := *fi
			*fi += 2
			if *fi <= flen && "C." == format[start:*fi] {
				// DCH B.C.
				keyword = DCH_B_C_
			} else {
				return empty_str, errors.New(dch_fmt_mismatch_err + "B.")
			}
		default:
			return empty_str, errors.New(dch_fmt_mismatch_err + "B")
		}
	} else {
		return empty_str, errors.New(dch_fmt_mismatch_err + "B")
	}

	return keyword, nil
}

func parsePrefixC(fi *int, flen int, format string) (FMKeyword, error) {
	var keyword FMKeyword
	*fi++
	if *fi < flen {
		followingOneChar := format[*fi]
		switch followingOneChar {
		case 'C':
			// DCH CC
			keyword = DCH_CC
			*fi++
		default:
			return empty_str, errors.New(dch_fmt_mismatch_err + "C")
		}
	} else {
		return empty_str, errors.New(dch_fmt_length_err + "C")
	}

	return keyword, nil
}

func parsePrefixD(fi *int, flen int, format string) (FMKeyword, error) {
	var keyword FMKeyword
	*fi++
	if *fi < flen {
		if format[*fi] == 'A' {
			*fi++
			if *fi < flen && format[*fi] == 'Y' {
				// DCH DAY 同 DY
				keyword = DCH_DAY
				*fi++
			} else {
				return empty_str, errors.New(dch_fmt_mismatch_err + "DA")
			}
		} else if format[*fi] == 'D' {
			*fi++
			if *fi < flen && format[*fi] == 'D' {
				// DCH DDD
				keyword = DCH_DDD
				*fi++
			} else {
				// DCH DD
				keyword = DCH_DD
			}
		} else if format[*fi] == 'L' {
			keyword = DCH_DL
			*fi++
		} else if format[*fi] == 'S' {
			keyword = DCH_DS
			*fi++
		} else if format[*fi] == 'Y' {
			// DCH DY
			keyword = DCH_DY
			*fi++
		} else {
			// DCH D
			keyword = DCH_D
		}
	} else {
		// DCH D
		keyword = DCH_D
	}
	return keyword, nil
}

func parsePrefixF(fi *int, flen int, format string) (FMKeyword, error) {
	var keyword FMKeyword
	*fi++
	if *fi < flen {
		followingOneChar := format[*fi]
		switch followingOneChar {
		case 'X':
			// TODO 最后处理
			keyword = DCH_FX
		case 'M':
			// TODO 最后处理
			keyword = DCH_FM
		case 'F':
			*fi++
			if *fi < flen {
				switch format[*fi] {
				case '1':
					keyword = DCH_FF1
					*fi++
				case '2':
					keyword = DCH_FF2
					*fi++
				case '3':
					keyword = DCH_FF3
					*fi++
				case '4':
					keyword = DCH_FF4
					*fi++
				case '5':
					keyword = DCH_FF5
					*fi++
				case '6':
					keyword = DCH_FF6
					*fi++
				case '7':
					keyword = DCH_FF7
					*fi++
				case '8':
					keyword = DCH_FF8
					*fi++
				case '9':
					keyword = DCH_FF9
					*fi++
				default:
					keyword = DCH_FF
				}
			} else {
				keyword = DCH_FF
			}
		default:
			return empty_str, errors.New(dch_fmt_length_err + "F")
		}
	} else {
		return empty_str, errors.New(dch_fmt_length_err + "F")
	}
	return keyword, nil
}

func parsePrefixH(fi *int, flen int, format string) (FMKeyword, error) {
	var keyword FMKeyword
	*fi++
	if *fi < flen {
		switch format[*fi] {
		case 'H':
			// DCH HH 同 HH12
			keyword = DCH_HH
		case '2':
			*fi++
			if *fi < flen {
				// DCH HH24
				if format[*fi] == '4' {
					keyword = DCH_HH24
				} else {
					return empty_str, errors.New(dch_fmt_mismatch_err + "H2")
				}
			} else {
				return empty_str, errors.New(dch_fmt_length_err + "H2")
			}
		case '1':
			*fi++
			if *fi < flen {
				// DCH HH12
				if format[*fi] == '2' {
					keyword = DCH_HH12
				} else {
					return empty_str, errors.New(dch_fmt_mismatch_err + "H2")
				}
			} else {
				return empty_str, errors.New(dch_fmt_length_err + "H1")
			}
		}
	} else {
		return empty_str, errors.New(dch_fmt_length_err + "H")
	}
	*fi++
	return keyword, nil
}

func parsePrefixI(fi *int, flen int, format string) (FMKeyword, error) {
	var keyword FMKeyword
	*fi++
	if *fi < flen {
		switch format[*fi] {
		case 'W':
			// DCH IW
			keyword = DCH_IW
			*fi++
		case 'Y':
			*fi++
			if *fi < flen && format[*fi] == 'Y' {
				*fi++
				if *fi < flen && format[*fi] == 'Y' {
					// DCH IYYY
					keyword = DCH_IYYY
					*fi++
				} else {
					// DCH IYY
					keyword = DCH_IYY
				}
			} else {
				// DCH IY
				keyword = DCH_IY
			}
		}
	} else {
		// DCH I
		keyword = DCH_I
	}
	return keyword, nil
}

func parsePrefixM(fi *int, flen int, format string) (FMKeyword, error) {
	var keyword FMKeyword
	*fi++
	if *fi < flen && format[*fi] == 'I' {
		// DCH MI
		keyword = DCH_MI
		*fi++
	} else if *fi < flen && format[*fi] == 'M' {
		// DCH MM
		keyword = DCH_MM
		*fi++
	} else if *fi < flen && format[*fi] == 'O' {
		*fi++
		if *fi < flen && format[*fi] == 'N' {
			*fi++
			start := *fi
			*fi += 2
			if *fi <= flen && format[start:*fi] == "TH" {
				// DCH MONTH
				keyword = DCH_MONTH
			} else {
				// DCH MON
				keyword = DCH_MON
				*fi -= 2
			}
		} else {
			return empty_str, errors.New(dch_fmt_mismatch_err + "MO")
		}
	} else {
		return empty_str, errors.New(dch_fmt_mismatch_err + "M")
	}
	return keyword, nil
}

func parsePrefixP(fi *int, flen int, format string) (FMKeyword, error) {
	var keyword FMKeyword
	*fi++
	if *fi < flen {
		if 'M' == format[*fi] {
			keyword = DCH_PM
			*fi++
		} else if '.' == format[*fi] {
			*fi++
			start := *fi
			*fi += 2
			if *fi <= flen {
				if "M." == format[start:*fi] {
					keyword = DCH_P_M_
				} else {
					return empty_str, errors.New(dch_fmt_mismatch_err + "P")
				}
			} else {
				return empty_str, errors.New(dch_fmt_length_err + "P")
			}
		} else {
			return empty_str, errors.New(dch_fmt_length_err + "P")
		}
	} else {
		return empty_str, errors.New(dch_fmt_length_err + "P")
	}
	return keyword, nil
}

func parsePrefixR(fi *int, flen int, format string) (FMKeyword, error) {
	var keyword int
	*fi++
	if *fi < flen {
		if 'M' == format[*fi] {
			keyword = DCH_RM
			*fi++
		} else if 'R' == format[*fi] {
			*fi++
			start := *fi
			*fi += 2
			if *fi <= flen && format[start:*fi] == "RR" {
				// DCH RRRR
				keyword = DCH_RRRR
			} else {
				// DCH RR
				keyword = DCH_RR
				*fi -= 2
			}
		} else {
			return empty_str, errors.New(dch_fmt_mismatch_err + "R")
		}
	} else {
		return empty_str, errors.New(dch_fmt_length_err + "R")
	}
	return keyword, nil
}

func parsePrefixS(fi *int, flen int, format string) (FMKeyword, error) {
	var keyword FMKeyword
	*fi++
	if *fi < flen {
		switch format[*fi] {
		case 'P':
			// DCH SP TODO 最后处理
			keyword = DCH_SP
			*fi++
		case 'S':
			*fi++
			start := *fi
			*fi += 3
			if *fi <= flen && format[start:*fi] == "SSS" {
				// DCH SSSSS 午夜过后的秒
				keyword = DCH_SSSSS
			} else {
				// DCH SS
				keyword = DCH_SS
				*fi -= 3
			}
		case 'Y':
			*fi++
			start := *fi
			*fi += 3
			if *fi <= flen {
				if format[start:*fi] == "YYY" {
					// TODO golang 好像不支持公元前
					// DCH SYYYY 正负号+数字
					//if 公元前 {result.WriteByte('-')}
					keyword = DCH_SYYYY
				} else if format[start:*fi] == "EAR" {
					// FIXME oracle中将4位的年分成了 2个2位数
					// DCH SYEAR 正负号+基数词
					//if 公元前 {result.WriteByte('-')}
					keyword = DCH_SYEAR
				} else {
					return empty_str, errors.New(dch_fmt_mismatch_err + "SY")
				}
			} else {
				return empty_str, errors.New(dch_fmt_length_err + "S")
			}
		default:
			return empty_str, errors.New(dch_fmt_mismatch_err + "S")
		}
	} else {
		return empty_str, errors.New(dch_fmt_length_err + "S")
	}
	return keyword, nil
}

func parsePrefixT(fi *int, flen int, format string) (FMKeyword, error) {
	var keyword FMKeyword
	// TODO 更换类型后更改时区
	*fi++
	if *fi < flen {
		if format[*fi] == 'S' {
			// DCH TS 下午 9:30:00
			keyword = DCH_TS
		} else if format[*fi] == 'Z' {
			*fi++
			if *fi < flen && format[*fi] == 'D' {
				// DCH TZD PDT 时区
				keyword = DCH_TZD
			} else if *fi < flen && format[*fi] == 'H' {
				// DCH TZH -07 时区小时
				keyword = DCH_TZH
			} else if *fi < flen && format[*fi] == 'M' {
				// DCH TZM 00 时区分
				keyword = DCH_TZM
			} else if *fi < flen && format[*fi] == 'R' {
				// DCH TZR US/PACIFIC 时区区域
				keyword = DCH_TZR
			} else {
				return empty_str, errors.New("格式错误")
			}
		} else if format[*fi] == 'H' {
			// DCH TH TODO 最后处理
			keyword = DCH_TH
		} else {
			return empty_str, errors.New("格式错误")
		}
	} else {
		return empty_str, errors.New("格式错误")
	}
	*fi++
	return keyword, nil
}
func parsePrefixY(fi *int, flen int, format string) (FMKeyword, error) {
	var keyword FMKeyword
	*fi++
	if *fi < flen {
		if format[*fi] == ',' {
			*fi++
			start := *fi
			*fi += 3
			if *fi <= flen {
				if format[start:*fi] == "YYY" {
					// DCH Y,YYY
					keyword = DCH_Y_YYY
				} else {
					return empty_str, errors.New(dch_fmt_mismatch_err + "Y,")
				}
			} else {
				return empty_str, errors.New(dch_fmt_length_err + "Y,")
			}
		} else if format[*fi] == 'Y' {
			*fi++

			if *fi < flen && format[*fi] == 'Y' {
				*fi++
				if *fi < flen && format[*fi] == 'Y' {
					// DCH YYYY
					keyword = DCH_YYYY
					*fi++
				} else {
					// DCH YYY
					keyword = DCH_YYY
				}
			} else {
				// DCH YY
				keyword = DCH_YY
			}
		} else if format[*fi] == 'E' {
			*fi++
			start := *fi
			*fi += 2
			if *fi <= flen && format[start:*fi] == "AR" {
				// DCH YEAR 基数词
				keyword = DCH_YEAR
			} else {
				return empty_str, errors.New(dch_fmt_mismatch_err + "YE")
			}
		}
	} else {
		// DCH Y
		keyword = DCH_Y
	}
	return keyword, nil
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
