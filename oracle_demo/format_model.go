package oracle_demo

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

// ASCII Table
const (
	ASCII_SPACE                = ' '
	ASCII_EXCLAMATION          = '!'
	ASCII_QUOTATION            = '"'
	ASCII_NUMBER_SIGN          = '#'
	ASCII_DOLLAR               = '$'
	ASCII_PERCENT              = '%'
	ASCII_AMPERSAND            = '&'
	ASCII_APOSTROPHE           = '\''
	ASCII_LEFT_PARENTHESIS     = '('
	ASCII_RIGHT_PARENTHESIS    = ')'
	ASCII_ASTERISK             = '*'
	ASCII_PLUS                 = '+'
	ASCII_COMMA                = ','
	ASCII_HYPHEN               = '-'
	ASCII_PERIOD               = '.'
	ASCII_SLASH                = '/'
	ASCII_0                    = '0'
	ASCII_1                    = '1'
	ASCII_2                    = '2'
	ASCII_3                    = '3'
	ASCII_4                    = '4'
	ASCII_5                    = '5'
	ASCII_6                    = '6'
	ASCII_7                    = '7'
	ASCII_8                    = '8'
	ASCII_9                    = '9'
	ASCII_COLON                = ':'
	ASCII_SEMICOLON            = ';'
	ASCII_LT                   = '<'
	ASCII_EQ                   = '='
	ASCII_GT                   = '>'
	ASCII_QUESTION             = '?'
	ASCII_AT                   = '@'
	ASCII_A                    = 'A'
	ASCII_B                    = 'B'
	ASCII_C                    = 'C'
	ASCII_D                    = 'D'
	ASCII_E                    = 'E'
	ASCII_F                    = 'F'
	ASCII_G                    = 'G'
	ASCII_H                    = 'H'
	ASCII_I                    = 'I'
	ASCII_J                    = 'J'
	ASCII_K                    = 'K'
	ASCII_L                    = 'L'
	ASCII_M                    = 'M'
	ASCII_N                    = 'N'
	ASCII_O                    = 'O'
	ASCII_P                    = 'P'
	ASCII_Q                    = 'Q'
	ASCII_R                    = 'R'
	ASCII_S                    = 'S'
	ASCII_T                    = 'T'
	ASCII_U                    = 'U'
	ASCII_V                    = 'V'
	ASCII_W                    = 'W'
	ASCII_X                    = 'X'
	ASCII_Y                    = 'Y'
	ASCII_Z                    = 'Z'
	ASCII_LEFT_SQUARE_BRACKET  = '['
	ASCII_BACKSLASH            = '\\'
	ASCII_RIGHT_SQUARE_BRACKET = ']'
	ASCII_CARET                = '^'
	ASCII_UNDERSCORE           = '_'
	ASCII_GRAVE_ACCENT         = '`'
	ASCII_a                    = 'a'
	ASCII_b                    = 'b'
	ASCII_c                    = 'c'
	ASCII_d                    = 'd'
	ASCII_e                    = 'e'
	ASCII_f                    = 'f'
	ASCII_g                    = 'g'
	ASCII_h                    = 'h'
	ASCII_i                    = 'i'
	ASCII_j                    = 'j'
	ASCII_k                    = 'k'
	ASCII_l                    = 'l'
	ASCII_m                    = 'm'
	ASCII_n                    = 'n'
	ASCII_o                    = 'o'
	ASCII_p                    = 'p'
	ASCII_q                    = 'q'
	ASCII_r                    = 'r'
	ASCII_s                    = 's'
	ASCII_t                    = 't'
	ASCII_u                    = 'u'
	ASCII_v                    = 'v'
	ASCII_w                    = 'w'
	ASCII_x                    = 'x'
	ASCII_y                    = 'y'
	ASCII_z                    = 'z'
	ASCII_LEFT_CURLY_BRACE     = '{'
	ASCII_VERTICAL_BAR         = '|'
	ASCII_RIGHT_CURLY_BRACE    = '}'
	ASCII_TILDE                = '~'
)

// Oracle 内置变量
const (
	// for number format 'L'
	NLS_CURRENCY = "￥"
	// for number format 'U'
	NLS_DUAL_CURRENCY = "￥"
)

// 异常
const (
	// 格式部分不匹配，报错
	dch_fmt_mismatch_err      = "Date Format error, some formats do not match near "
	dch_fmt_length_err        = "Date Format error, incorrect format length near "
	num_fmt_part_err          = "Datetime Format error, some formats do not match near "
	not_support_err           = "not support"
	format_conflict_err       = "format conflict with "
	format_err                = "format err "
	unreachable_err           = "unreachable code"
	format_length_smaller_err = "Format length is smaller than parameter length"
	// 非法字符,超出格式关键词范围
	out_keyword_range_err = "Illegal character, not in the range of Format Model keyword"
	// 非法字符,超出ASCII[32-126]字符范围
	out_ascii_range_err = "Illegal character, not in ASCII [32-126] character range"
	invalid_num_err     = "invalid number"
)

// 模式
const (
	mode_flag_fm = 1
	mode_flag_fx = 1 << 1
	mode_flag_th = 1 << 2
	mode_flag_sp = 1 << 3
)

// 数值类型相关常量
const (
	matchModeEmpty matchMode = 0
	matchModeFm    matchMode = 1

	signEmpty sign = sign(0)
	signSpace sign = ' '
	signPlus  sign = '+'
	signMinus sign = '-'
	signGt    sign = '>'
	signLt    sign = '<'

	signModeEmpty  signMode = 0
	signModePR     signMode = 1
	signModeMI     signMode = 2
	signModeSStart signMode = 3
	signModeSEnd   signMode = 4

	currencySymbolEmpty  currencySymbol = ""
	currencySymbolDollar currencySymbol = "$"
	currencySymbolB      currencySymbol = " "
	currencySymbolC      currencySymbol = "cny"
	currencySymbolL      currencySymbol = NLS_CURRENCY
	currencySymbolU      currencySymbol = NLS_DUAL_CURRENCY

	// 后缀 后缀互斥 后缀决定了输出模式
	outputModeEmpty outputMode = 0
	outputModeEEEE  outputMode = 1
	outputModeV     outputMode = 2
	outputModeRN    outputMode = 3
	outputModeX     outputMode = 4
	outputModeTM    outputMode = 5
	outputModeTME   outputMode = 6
)

// 数值相关类型
type matchMode int
type sign byte
type currencySymbol string
type outputMode int
type signMode int

// 日期相关类型
type dchKeyword int
type dtType int

var dchKeywords map[int]int
var NLS_WEEKS = map[time.Weekday]string{
	time.Sunday:    "星期日",
	time.Monday:    "星期一",
	time.Tuesday:   "星期二",
	time.Wednesday: "星期三",
	time.Thursday:  "星期四",
	time.Friday:    "星期五",
	time.Saturday:  "星期六",
}

var NLS_MONTHS = map[time.Month]string{
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
var NLS_MONTHS_REVERSE = map[string]time.Month{
	"1月":  time.January,
	"2月":  time.February,
	"3月":  time.March,
	"4月":  time.April,
	"5月":  time.May,
	"6月":  time.June,
	"7月":  time.July,
	"8月":  time.August,
	"9月":  time.September,
	"10月": time.October,
	"11月": time.November,
	"12月": time.December,
}

// 日期类型相关变量
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

	empty_str   = ""
	empty_int   = 0
	empty_float = 0.0
	tsFormat    = "15:04:05"
	//dateFormat = "YYYY-MM-DD HH24:MI:SS"
	dateLayout = "2006-01-02 15:04:05"
)

// 日期模型相关常量
const (
	dt_type_date         dtType = 1
	dt_type_timestamp    dtType = 2
	dt_type_timestamp_tz dtType = 3

	dt_flag_year   = 1
	dt_flag_month  = 1 << 1
	dt_flag_day    = 1 << 2
	dt_flag_hour   = 1 << 3
	dt_flag_minute = 1 << 4
	dt_flag_second = 1 << 5
	dt_flag_nansec = 1 << 6
	dt_flag_tzr    = 1 << 7
	dt_flag_tzh    = 1 << 8
	dt_flag_tzm    = 1 << 9
	dt_flag_adbc   = 1 << 10
	dt_flag_ampm   = 1 << 11
)

// 日期类型模型关键词
const (
	DCH_EMPTY        dchKeyword = 0
	DCH_MINUS        dchKeyword = 1
	DCH_SLASH        dchKeyword = 2
	DCH_COMMA        dchKeyword = 3
	DCH_DEC          dchKeyword = 4
	DCH_SEMICOLON    dchKeyword = 5
	DCH_COLON        dchKeyword = 6
	DCH_SPACE        dchKeyword = 7
	DCH_DOUBLE_QUOTE dchKeyword = 8
	DCH_AD           dchKeyword = 9
	DCH_A_D_         dchKeyword = 10
	DCH_AM           dchKeyword = 11
	DCH_A_M_         dchKeyword = 12
	DCH_BC           dchKeyword = 13
	DCH_B_C_         dchKeyword = 14
	DCH_CC           dchKeyword = 15
	DCH_SCC          dchKeyword = 16
	DCH_DAY          dchKeyword = 17
	DCH_DDD          dchKeyword = 18
	DCH_DD           dchKeyword = 19
	DCH_DL           dchKeyword = 20
	DCH_DS           dchKeyword = 21
	DCH_DY           dchKeyword = 22
	DCH_D            dchKeyword = 23
	DCH_E            dchKeyword = 24
	DCH_EE           dchKeyword = 25
	DCH_FF1          dchKeyword = 26
	DCH_FF2          dchKeyword = 27
	DCH_FF3          dchKeyword = 28
	DCH_FF4          dchKeyword = 29
	DCH_FF5          dchKeyword = 30
	DCH_FF6          dchKeyword = 31
	DCH_FF7          dchKeyword = 32
	DCH_FF8          dchKeyword = 33
	DCH_FF9          dchKeyword = 34
	DCH_FF           dchKeyword = 35
	DCH_HH24         dchKeyword = 36
	DCH_HH12         dchKeyword = 37
	DCH_HH           dchKeyword = 38
	DCH_IW           dchKeyword = 39
	DCH_IYYY         dchKeyword = 40
	DCH_IYY          dchKeyword = 41
	DCH_IY           dchKeyword = 42
	DCH_I            dchKeyword = 43
	DCH_J            dchKeyword = 44
	DCH_MI           dchKeyword = 45
	DCH_MM           dchKeyword = 46
	DCH_MONTH        dchKeyword = 47
	DCH_MON          dchKeyword = 48
	DCH_P_M_         dchKeyword = 49
	DCH_PM           dchKeyword = 50
	DCH_Q            dchKeyword = 51
	DCH_RM           dchKeyword = 52
	DCH_RR           dchKeyword = 53
	DCH_RRRR         dchKeyword = 54
	DCH_SP           dchKeyword = 55
	DCH_SSSSS        dchKeyword = 56
	DCH_SS           dchKeyword = 57
	DCH_TZH          dchKeyword = 58
	DCH_TZM          dchKeyword = 59
	DCH_TZD          dchKeyword = 60
	DCH_TZR          dchKeyword = 61
	DCH_TS           dchKeyword = 62
	DCH_WW           dchKeyword = 63
	DCH_W            dchKeyword = 64
	DCH_X            dchKeyword = 65
	DCH_Y_YYY        dchKeyword = 66
	DCH_YEAR         dchKeyword = 67
	DCH_SYEAR        dchKeyword = 68
	DCH_YYYY         dchKeyword = 69
	DCH_SYYYY        dchKeyword = 70
	DCH_YYY          dchKeyword = 71
	DCH_YY           dchKeyword = 72
	DCH_Y            dchKeyword = 73
)

// 数值的格式描述
type NumFmtDesc struct {
	// 匹配模式
	matchMode matchMode
	// 左符号 + - < 空
	//leftSign sign
	// 货币符号
	currencySymbol currencySymbol
	// 数值模型 前半部分 9 0 ,
	preSep string
	// 除去逗号的有效长度
	preSepValidLen int
	// 数值模型 后半部分 9 0
	postSep string
	// 分隔的位置 用. 或者V 分割
	//sepIndex int
	// 输出模式
	outputMode outputMode
	// 右符号 + - > 空
	//rightSign sign
	// 符号模式
	signMode signMode
}

// 数值的参数描述
type NumParamDesc struct {
	nSign     sign
	preDec    string
	postDec   string
	eSign     sign
	hasE      bool
	eExponent int
}

func (numParam *NumParamDesc) decimal() (float64, error) {
	base := bytes.Buffer{}
	base.WriteByte(byte(numParam.nSign))
	base.WriteString(numParam.preDec)
	base.WriteString(numParam.postDec)

	f, err := strconv.ParseFloat(base.String(), 64)
	if err != nil {
		return 0, err
	}

	if numParam.hasE {
		if numParam.eSign == signMinus {
			f *= math.Pow10(-numParam.eExponent)
		} else {
			f *= math.Pow10(numParam.eExponent)
		}
	}
	return f, nil
}

func (numParam *NumParamDesc) string() (string, error) {
	var result bytes.Buffer
	if signPlus == numParam.nSign {
		result.WriteByte(byte(signPlus))
	} else if numParam.nSign == signMinus {
		result.WriteByte(byte(signMinus))
	} else if numParam.nSign == signEmpty {
	} else {
		return empty_str, errors.New("sign属性格式错误")
	}

	if empty_str != numParam.preDec {
		result.WriteString(numParam.preDec)
	} else {
		return empty_str, errors.New("格式错误,整数部分是空的")
	}

	if numParam.postDec != empty_str {
		result.WriteByte('.')
		result.WriteString(numParam.postDec)
	}

	if numParam.eExponent != 0 {
		result.WriteByte('e')

		if numParam.nSign == signPlus {
			result.WriteByte(byte(signPlus))
		} else if numParam.nSign == signMinus {
			result.WriteByte(byte(signMinus))
		} else if numParam.nSign == signEmpty {
		} else {
			return empty_str, errors.New("eSign属性格式错误")
		}
		result.WriteString(fmt.Sprint(numParam.eExponent))
	}

	return result.String(), nil
}

// 解析数值格式
func parseNumFormat(format string) (*NumFmtDesc, error) {
	fmtDesc := &NumFmtDesc{}

	// 格式字节长度
	flen := len(format)
	li := flen - 1

	var c byte

	readDec := false
	readV := false

	preSep := bytes.Buffer{}
	postSep := bytes.Buffer{}
	for fi := 0; fi < flen; {
		// 截取一个字符
		c = format[fi]
		if c >= 32 && c <= 127 {
			// 匹配关键词并存储
			switch c {
			case 'F', 'f':
				fi++
				if format[fi] == 'M' || format[fi] == 'm' {
					if fmtDesc.matchMode != matchModeEmpty {
						return fmtDesc, errors.New(format_conflict_err + "FM")
					}
					if fi == 1 {
						fmtDesc.matchMode = matchModeFm
					} else {
						return nil, errors.New("FM can only be at the beginning")
					}
				} else {
					return nil, errors.New(num_fmt_part_err + "F")
				}
				fi++
			case 'R', 'r':
				// 判断独占 长度 FIXME
				fi++
				if format[fi] == 'N' || format[fi] == 'n' {
					if fmtDesc.outputMode != outputModeEmpty {
						return fmtDesc, errors.New(format_conflict_err + "RN")
					} else if fmtDesc.matchMode == matchModeFm && flen == 4 {
						return nil, errors.New(format_conflict_err + "RN")
					} else if fmtDesc.matchMode == matchModeEmpty && flen == 2 {
						return nil, errors.New(format_conflict_err + "RN")
					}
					fmtDesc.outputMode = outputModeRN
				} else {
					return nil, errors.New(num_fmt_part_err + "R")
				}
				if fi != li {
					return nil, errors.New(format_err + "RN")
				} else {
					break
				}
			case 'T', 't':
				if fmtDesc.outputMode == outputModeEmpty {
					fi++
					if format[fi] == 'M' || format[fi] == 'm' {
						if fi == li {
							fmtDesc.outputMode = outputModeTM
						} else {
							fi++
							if format[fi] == 'E' || format[fi] == 'e' {
								fmtDesc.outputMode = outputModeTME
							} else if format[fi] == '9' {
								fmtDesc.outputMode = outputModeTM
							} else {
								return nil, errors.New(format_err + "TM")
							}
						}
					} else {
						return nil, errors.New(format_err + string(c))
					}
				} else {
					return nil, errors.New(format_conflict_err + "TM")
				}

				if fi != li {
					return nil, errors.New(format_err + "TM")
				} else {
					break
				}
			case 'X', 'x':
				if fmtDesc.outputMode != outputModeX {
					return nil, errors.New(format_conflict_err + string(c))
				}
				fmtDesc.outputMode = outputModeX
				for ; fi < flen; fi++ {
					if format[fi] == 'X' || format[fi] == 'x' {
						fmtDesc.preSepValidLen++
					} else {
						return nil, errors.New("can only have X or x in hexadecimal format")
					}
				}
				break
			default:
				signAffixSetup := false

				for fi < flen {
					// 截取一个字符
					c = format[fi]
					if c >= 32 && c <= 127 {
						switch c {
						case '9':
							if readDec || readV {
								postSep.WriteByte('9')
							} else {
								preSep.WriteByte('9')
								fmtDesc.preSepValidLen++
							}
							fi++
						case '.':
							if !readDec {
								readDec = true
							} else {
								return fmtDesc, errors.New("there can only be 1 period")
							}
							fi++
						case '0':
							if readDec || readV {
								postSep.WriteByte('0')
							} else {
								preSep.WriteByte('0')
								fmtDesc.preSepValidLen++
							}
							fi++
						case ',':
							if fi == 0 {
								return nil, errors.New("cannot begin with a comma")
							} else if fi == li {
								return fmtDesc, errors.New("comma cannot appear on the right most side of a number")
							} else if readDec {
								return fmtDesc, errors.New("the comma cannot appear on the right side of the period")
							}
							preSep.WriteByte(',')
							fi++
						case '$':
							if fmtDesc.currencySymbol == currencySymbolEmpty && fi == 0 {
								fmtDesc.currencySymbol = currencySymbolDollar
							} else {
								return fmtDesc, errors.New(format_conflict_err + string(c))
							}
							fi++
						case 'B', 'b':
							if fmtDesc.currencySymbol == currencySymbolEmpty {
								fmtDesc.currencySymbol = currencySymbolB
							} else {
								return fmtDesc, errors.New(format_conflict_err + string(c))
							}
							fi++
						case 'C', 'c':
							if fmtDesc.currencySymbol == currencySymbolEmpty {
								return nil, errors.New(format_conflict_err + string(c))
							} else if li != fi && 0 != fi {
								return fmtDesc, errors.New("C can only be at the beginning or end")
							}
							fmtDesc.currencySymbol = currencySymbolC
							fi++
						case 'L', 'l':
							if fmtDesc.currencySymbol == currencySymbolEmpty {
								fmtDesc.currencySymbol = currencySymbolL
							} else {
								return nil, errors.New(format_conflict_err + string(c))
							}
							fi++
						case 'U', 'u':
							if fmtDesc.currencySymbol == currencySymbolEmpty {
								return nil, errors.New(format_conflict_err + "U")
							} else if li != fi && 0 != fi {
								return fmtDesc, errors.New("U can only be at the beginning or end")
							}
							fmtDesc.currencySymbol = currencySymbolU
							fi++
						case 'M', 'm':
							if signAffixSetup {
								return nil, errors.New(format_conflict_err + "M")
							} else if fi == (li - 1) {
								return nil, errors.New("MI can only be at the end")
							}

							fi++
							if format[fi] == 'I' || format[fi] == 'i' {
								fmtDesc.signMode = signModeMI
							} else {
								return nil, errors.New(num_fmt_part_err + "M")
							}
							fi++
						case 'P', 'p':
							if signAffixSetup {
								return nil, errors.New(format_conflict_err + "PR")
							} else if fi != (li - 1) {
								return fmtDesc, errors.New("PR can only be at the end")
							}

							fi++
							if format[fi] == 'R' || format[fi] == 'r' {
								fmtDesc.signMode = signModePR
							} else {
								return nil, errors.New(num_fmt_part_err + "P")
							}
							fi++
						case 'S', 's':
							if signAffixSetup {
								return nil, errors.New(format_conflict_err + string(c))
							} else if fi == li {
								fmtDesc.signMode = signModeSEnd
							} else if fi != 0 {
								fmtDesc.signMode = signModeSStart
							} else {
								return fmtDesc, errors.New("S can only be at the beginning or end")
							}

							fi++
						case 'E', 'e':
							if fmtDesc.outputMode == outputModeEmpty {
								start := fi + 1
								fi = start + 3
								if "EEE" == strings.ToUpper(format[start:fi]) {
									fmtDesc.outputMode = outputModeEEEE
								} else {
									return nil, errors.New(num_fmt_part_err + string(c))
								}
							} else {
								return nil, errors.New(format_conflict_err + "E")
							}
						case 'V', 'v':
							if readDec {
								return nil, errors.New(format_conflict_err + ".")
							}
							if fmtDesc.outputMode != outputModeEmpty {
								return nil, errors.New(format_conflict_err + string(c))
							} else if 0 == fi {
								return nil, errors.New("can not start with " + string(c))
							}
							readV = true
							fmtDesc.outputMode = outputModeV
							fi++
						default:
							return nil, errors.New(out_keyword_range_err)
						}
					} else {
						return nil, errors.New(out_ascii_range_err)
					}
				}
			}
		} else {
			return nil, errors.New(out_ascii_range_err)
		}
	}

	fmtDesc.preSep = preSep.String()
	fmtDesc.postSep = postSep.String()
	return fmtDesc, nil

}

// 解析数字参数
func parseNumParam(num string) (*NumParamDesc, error) {
	paramDesc := &NumParamDesc{}

	// 读取到小数点
	readDec := false

	var preBuf = bytes.Buffer{}
	var postBuf = bytes.Buffer{}
	for i := 0; i < len(num); i++ {
		c := num[i]
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
			} else {
				return paramDesc, errors.New("多个符号 " + ".")
			}
		case 'e', 'E':
			i++
			paramDesc.hasE = true
			var exponent = bytes.Buffer{}

			if num[i] == '+' {
				paramDesc.eSign = signPlus
			} else if num[i] == '-' {
				paramDesc.eSign = signMinus
			} else if num[i] <= '9' && num[i] >= '0' {
				paramDesc.eSign = signEmpty
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
		case '-':
			if i == 0 {
				paramDesc.nSign = signMinus
			} else {
				return paramDesc, errors.New("符号位置不对 " + "-")
			}
		case '+':
			if i == 0 {
				paramDesc.nSign = signPlus
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
	if paramDesc.hasE {
		if preBuf.Len() > 0 {
			paramDesc.preDec = preBuf.String()
		} else {
			paramDesc.preDec = "0"
		}

		if postBuf.Len() > 0 {
			paramDesc.postDec = postBuf.String()
		}
		ff := paramDesc.preDec + "." + paramDesc.postDec
		d, err := strconv.ParseFloat(ff, 64)
		if err != nil {
			return paramDesc, err
		}
		fff := "%" + paramDesc.preDec + "." + paramDesc.postDec + "f"
		v := fmt.Sprintf(fff, d)
		//FIXME
		fmt.Println(v)

	} else {
		paramDesc.preDec = preBuf.String()
		paramDesc.postDec = postBuf.String()
	}

	return paramDesc, nil
}

// 字符串类型转换成数字
func ToNumber(num string, format string) (float64, error) {
	numFmtDesc, err := parseNumFormat(format)
	if err != nil {
		return empty_float, err
	}
	//log.Printf("%#v\n", numFmtDesc)

	numParamDesc, err := parseNumParam(num)
	if err != nil {
		return empty_float, err
	}
	//log.Printf("%#v\n", numParamDesc)

	switch numFmtDesc.outputMode {
	// 十进制
	case outputModeEmpty:
		if numFmtDesc.preSepValidLen < len(numParamDesc.preDec) {
			return empty_float, errors.New(format_length_smaller_err)
		}
		f, err := strconv.ParseFloat(num, 64)
		if err != nil {
			return empty_float, err
		}
		return f, nil
	// 十六进制
	case outputModeX:
		d, err := strconv.ParseInt(num, 16, 64)
		if err != nil {
			return empty_float, err
		}
		if len(numParamDesc.preDec) > numFmtDesc.preSepValidLen {
			return empty_float, errors.New(format_length_smaller_err)
		}
		return float64(d), nil
	default:
		return empty_float, errors.New(not_support_err)
	}

}

// 字符串 转 时间戳
func ToTimestamp(dch string, format string) (*time.Time, error) {
	return toDatetime(dch, format, dt_type_timestamp)
}

// 字符串 转 带时区的时间戳
func ToTimestampTimeZone(dch string, format string) (*time.Time, error) {
	return toDatetime(dch, format, dt_type_timestamp_tz)
}

// 字符串 转 日期
func ToDate(dch string, format string) (*time.Time, error) {
	return toDatetime(dch, format, dt_type_date)
}

func toDatetime(dch string, format string, tp dtType) (*time.Time, error) {
	dchKeywords, quoted, aux_flag, err := parseFmt(format)
	if err != nil {
		return nil, nil
	}

	year, month, day := 0, time.Month(0), 0
	hour, min, sec, nsec := 0, 0, 0, 0
	tzr := time.Local

	now := time.Now()

	var parseDch func(*string, *int, *int, int) (string, error)
	if aux_flag&mode_flag_fx == 0 {
		parseDch = parseDchNotFX
	} else {
		parseDch = parseDchFX
	}

	qi := 0
	di := 0
	dlen := len(dch)
	dt_flag := 0

	for ki := 0; ki < len(dchKeywords); ki++ {
		switch dchKeywords[ki] {
		case DCH_DOUBLE_QUOTE:
			field, err := parseDch(&dch, &dlen, &di, len(quoted[qi]))
			if err != nil {
				return nil, err
			}
			if field != quoted[qi] {
				return nil, errors.New("引号内容不匹配")
			}
			qi++
		case DCH_SPACE:
			if aux_flag&mode_flag_fx == 1 {
				if dch[di] != ' ' {
					return nil, errors.New("严格模式下` `不匹配")
				}
				di++
			}
		case DCH_MINUS:
			if aux_flag&mode_flag_fx == 1 {
				if dch[di] != '-' {
					return nil, errors.New("严格模式下`-`不匹配")
				}
				di++
			}
		case DCH_SLASH:
			if aux_flag&mode_flag_fx == 1 {
				if dch[di] != '-' {
					return nil, errors.New("严格模式下`/`不匹配")
				}
				di++
			}
		case DCH_COMMA:
			if aux_flag&mode_flag_fx == 1 {
				if dch[di] != '-' {
					return nil, errors.New("严格模式下`,`不匹配")
				}
				di++
			}
		case DCH_DEC:
			if aux_flag&mode_flag_fx == 1 {
				if dch[di] != '-' {
					return nil, errors.New("严格模式下`.`不匹配")
				}
				di++
			}
		case DCH_COLON:
			if aux_flag&mode_flag_fx == 1 {
				if dch[di] != '-' {
					return nil, errors.New("严格模式下`:`不匹配")
				}
				di++
			}
		case DCH_SEMICOLON:
			if aux_flag&mode_flag_fx == 1 {
				if dch[di] != '-' {
					return nil, errors.New("严格模式下`;`不匹配")
				}
				di++
			}
		case DCH_DD:
			if dt_flag&dt_flag_day == 0 {
				field, err := parseDch(&dch, &dlen, &di, 2)
				if err != nil {
					return nil, err
				}
				day, err = strconv.Atoi(field)
				if err != nil {
					return nil, err
				}
				dt_flag |= dt_flag_day
			} else {
				return nil, errors.New("格式 日 已经重复")
			}
		case DCH_HH24, DCH_HH12, DCH_HH:
			if dt_flag&dt_flag_hour == 0 {
				field, err := parseDch(&dch, &dlen, &di, 2)
				if err != nil {
					return nil, err
				}
				hour, err = strconv.Atoi(field)
				if err != nil {
					return nil, err
				}
				dt_flag |= dt_flag_hour
			} else {
				return nil, errors.New("格式 小时 已经重复")
			}
		case DCH_MI:
			if dt_flag&dt_flag_minute == 0 {
				field, err := parseDch(&dch, &dlen, &di, 2)
				if err != nil {
					return nil, err
				}
				min, err = strconv.Atoi(field)
				if err != nil {
					return nil, err
				}
				dt_flag |= dt_flag_minute
			} else {
				return nil, errors.New("格式 分钟 已经重复")
			}
		case DCH_MM:
			if dt_flag&dt_flag_month == 0 {
				field, err := parseDch(&dch, &dlen, &di, 2)
				if err != nil {
					return nil, err
				}
				mon, err := strconv.Atoi(field)
				month = time.Month(mon)
				if err != nil {
					return nil, err
				}
				dt_flag |= dt_flag_month
			} else {
				return nil, errors.New("格式 月 已经重复")
			}
		case DCH_MONTH, DCH_MON: //FIXME
			if dt_flag&dt_flag_month == 0 {
				field, err := parseDch(&dch, &dlen, &di, 2)
				if err != nil {
					return nil, err
				}
				month = NLS_MONTHS_REVERSE[field]
				if err != nil {
					return nil, err
				}
				dt_flag |= dt_flag_month
			} else {
				return nil, errors.New("格式 月 已经重复")
			}
		case DCH_RR:
			if dt_flag&dt_flag_year == 0 {
				field, err := parseDch(&dch, &dlen, &di, 2)
				if err != nil {
					return nil, err
				}
				RR, err := strconv.Atoi(field)
				if err != nil {
					return nil, err
				}
				year = toRRRR(now.Year(), RR)
				dt_flag |= dt_flag_year
			} else {
				return nil, errors.New("格式 年 已经重复")
			}
		case DCH_RRRR:
			if dt_flag&dt_flag_year == 0 {
				field, err := parseDch(&dch, &dlen, &di, 4)
				if err != nil {
					return nil, err
				}
				year, err = strconv.Atoi(field)
				if err != nil {
					return nil, err
				}
				dt_flag |= dt_flag_year
			} else {
				return nil, errors.New("格式 年 已经重复")
			}
		case DCH_SS:
			if dt_flag&dt_flag_second == 0 {
				field, err := parseDch(&dch, &dlen, &di, 2)
				if err != nil {
					return nil, err
				}
				sec, err = strconv.Atoi(field)
				if err != nil {
					return nil, err
				}
				dt_flag |= dt_flag_second
			} else {
				return nil, errors.New("格式 秒 已经重复")
			}
		case DCH_Y_YYY:
			if dt_flag&dt_flag_year == 0 {

				field, err := parseDch(&dch, &dlen, &di, 5)
				if err != nil {
					return nil, err
				}
				year, err = strconv.Atoi(field[0:1] + field[2:5])
				if err != nil {
					return nil, err
				}
				dt_flag |= dt_flag_year
			} else {
				return nil, errors.New("格式 年 已经重复")
			}
		case DCH_YYYY:
			if dt_flag&dt_flag_year == 0 {
				field, err := parseDch(&dch, &dlen, &di, 4)
				if err != nil {
					return nil, err
				}
				year, err = strconv.Atoi(field)
				if err != nil {
					return nil, err
				}
				dt_flag |= dt_flag_year
			} else {
				return nil, errors.New("格式 年 已经重复")
			}
		case DCH_YYY:
			if dt_flag&dt_flag_year == 0 {
				field, err := parseDch(&dch, &dlen, &di, 3)
				if err != nil {
					return nil, err
				}
				year, err = strconv.Atoi(strconv.Itoa(now.Year())[0:1] + field)
				if err != nil {
					return nil, err
				}
				dt_flag |= dt_flag_year
			} else {
				return nil, errors.New("格式 年 已经重复")
			}
		case DCH_YY:
			if dt_flag&dt_flag_year == 0 {
				field, err := parseDch(&dch, &dlen, &di, 2)
				if err != nil {
					return nil, err
				}
				year, err = strconv.Atoi(strconv.Itoa(now.Year())[0:2] + field)
				if err != nil {
					return nil, err
				}
				dt_flag |= dt_flag_year
			} else {
				return nil, errors.New("格式 年 已经重复")
			}
		case DCH_Y:
			if dt_flag&dt_flag_year == 0 {
				field, err := parseDch(&dch, &dlen, &di, 1)
				if err != nil {
					return nil, err
				}
				year, err = strconv.Atoi(strconv.Itoa(now.Year())[0:3] + field)
				if err != nil {
					return nil, err
				}
				dt_flag |= dt_flag_year
			} else {
				return nil, errors.New("格式 年 已经重复")
			}
		case DCH_TZH:
			if tp == dt_type_timestamp_tz {
				if dt_flag&dt_flag_tzr == 0 && dt_flag&dt_flag_tzh == 0 {
					// TODO
					dt_flag |= dt_flag_tzh
				} else {
					return nil, errors.New("格式 时区的小时 已经重复")
				}
			} else {
				return nil, errors.New("只有带时区的时间戳类型支持时区")
			}
		case DCH_TZM:
			if tp == dt_type_timestamp_tz {
				if dt_flag&dt_flag_tzr == 0 && dt_flag&dt_flag_tzh == 0 {
					// TODO
					dt_flag |= dt_flag_tzm
				} else {
					return nil, errors.New("格式 时区的分钟 已经重复")
				}
			} else {
				return nil, errors.New("只有带时区的时间戳类型支持时区")
			}
		case DCH_TZR:
			if tp == dt_type_timestamp_tz {
				if dt_flag&dt_flag_tzr == 0 && dt_flag&dt_flag_tzh == 0 && dt_flag&dt_flag_tzm == 0 {
					// TODO
					dt_flag |= dt_flag_tzr
				} else {
					return nil, errors.New("格式 时区 已经重复")
				}
			} else {
				return nil, errors.New("只有带时区的时间戳类型支持时区")
			}
		case DCH_FF1:
			if tp != dt_type_date {
				if dt_flag&dt_flag_nansec == 0 {
					// TODO
					dt_flag |= dt_flag_nansec
				} else {
					return nil, errors.New("格式 纳秒 已经重复")
				}
			} else {
				return nil, errors.New("日期类型不支持小数秒")
			}
		case DCH_FF2:
			if tp != dt_type_date {
				if dt_flag&dt_flag_nansec == 0 {
					// TODO
					dt_flag |= dt_flag_nansec
				} else {
					return nil, errors.New("格式 纳秒 已经重复")
				}
			} else {
				return nil, errors.New("日期类型不支持小数秒")
			}
		case DCH_FF3:
			if tp != dt_type_date {
				if dt_flag&dt_flag_nansec == 0 {
					// TODO
					dt_flag |= dt_flag_nansec
				} else {
					return nil, errors.New("格式 纳秒 已经重复")
				}
			} else {
				return nil, errors.New("日期类型不支持小数秒")
			}
		case DCH_FF4:
			if tp != dt_type_date {
				if dt_flag&dt_flag_nansec == 0 {
					// TODO
					dt_flag |= dt_flag_nansec
				} else {
					return nil, errors.New("格式 纳秒 已经重复")
				}
			} else {
				return nil, errors.New("日期类型不支持小数秒")
			}
		case DCH_FF5:
			if tp != dt_type_date {
				if dt_flag&dt_flag_nansec == 0 {
					// TODO
					dt_flag |= dt_flag_nansec
				} else {
					return nil, errors.New("格式 纳秒 已经重复")
				}
			} else {
				return nil, errors.New("日期类型不支持小数秒")
			}
		case DCH_FF6:
			if tp != dt_type_date {
				if dt_flag&dt_flag_nansec == 0 {
					// TODO
					dt_flag |= dt_flag_nansec
				} else {
					return nil, errors.New("格式 纳秒 已经重复")
				}
			} else {
				return nil, errors.New("日期类型不支持小数秒")
			}
		case DCH_FF7:
			if tp != dt_type_date {
				if dt_flag&dt_flag_nansec == 0 {
					// TODO
					dt_flag |= dt_flag_nansec
				} else {
					return nil, errors.New("格式 纳秒 已经重复")
				}
			} else {
				return nil, errors.New("日期类型不支持小数秒")
			}
		case DCH_FF8:
			if tp != dt_type_date {
				if dt_flag&dt_flag_nansec == 0 {
					// TODO
					dt_flag |= dt_flag_nansec
				} else {
					return nil, errors.New("格式 纳秒 已经重复")
				}
			} else {
				return nil, errors.New("日期类型不支持小数秒")
			}
		case DCH_FF9:
			if tp != dt_type_date {
				if dt_flag&dt_flag_nansec == 0 {
					// TODO
					dt_flag |= dt_flag_nansec
				} else {
					return nil, errors.New("格式 纳秒 已经重复")
				}
			} else {
				return nil, errors.New("日期类型不支持小数秒")
			}
		case DCH_FF:
			if tp != dt_type_date {
				if dt_flag&dt_flag_nansec == 0 {
					// TODO
					dt_flag |= dt_flag_nansec
				} else {
					return nil, errors.New("格式 纳秒 已经重复")
				}
			} else {
				return nil, errors.New("日期类型不支持小数秒")
			}
		// FIXME 暂时不支持
		//case DCH_AD:
		//	if dItems[ki] != NLS_AD {
		//		return nil, errors.New("格式字符不匹配")
		//	}
		//case DCH_A_D_:
		//	if dItems[ki] != NLS_A_D_ {
		//		return nil, errors.New("格式字符不匹配")
		//	}
		//case DCH_AM:
		//	if dItems[ki] != NLS_AM {
		//		return nil, errors.New("格式字符不匹配")
		//	}
		//case DCH_A_M_:
		//	if dItems[ki] != NLS_A_M_ {
		//		return nil, errors.New("格式字符不匹配")
		//	}
		//case DCH_PM:
		//case DCH_P_M_:
		//case DCH_BC:
		//	if dItems[ki] != NLS_BC {
		//		return nil, errors.New("格式字符不匹配")
		//	}
		//case DCH_B_C_:
		//	if dItems[ki] != NLS_B_C_ {
		//		return nil, errors.New("格式字符不匹配")
		//	}
		default:
			return nil, errors.New(not_support_err)
		}
	}

	if qi != len(quoted) {
		return nil, errors.New("引号内容未遍历完，不匹配")
	}

	if year == 0 {
		year = now.Year()
	}
	if month == 0 {
		month = now.Month()
	}
	if day == 0 {
		day = 1
	}

	t := time.Date(year, month, day, hour, min, sec, nsec, tzr)
	return &t, nil
}

func parseDchFX(dch *string, dlen *int, di *int, size int) (string, error) {
	start := *di
	*di += size
	if *di > *dlen {
		return "", errors.New("格式长度不匹配")
	}
	return (*dch)[start:*di], nil
}

func parseDchNotFX(dch *string, dlen *int, di *int, size int) (string, error) {
	tmp := bytes.Buffer{}
	for ; *di < *dlen; *di++ {
		if (*dch)[*di] == ' ' ||
			(*dch)[*di] == '-' ||
			(*dch)[*di] == ':' ||
			(*dch)[*di] == ',' ||
			(*dch)[*di] == '.' ||
			(*dch)[*di] == '/' ||
			(*dch)[*di] == ';' {
		} else {
			for j := 0; j < size; j++ {
				// FIXME 05 DD 和 5 DD
				tmp.WriteByte((*dch)[*di])
				*di++
			}
			return tmp.String(), nil
		}
	}

	return empty_str, errors.New("未找到格式对应的匹配项")
}

// 数字类型 转 格式化字符串
func ToCharByNum(numFloat float64, format string) (string, error) {
	numStr := strconv.FormatFloat(numFloat, 'f', -1, 64)
	return ToChar(numStr, numFloat, format)
}

// 字符串类型 转 格式化字符串
func ToCharByStr(numStr string, format string) (string, error) {
	numFloat, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return empty_str, err
	}
	return ToChar(numStr, numFloat, format)
}

func ToChar(numStr string, numFloat float64, format string) (string, error) {
	negative := numFloat < 0
	numFloat = math.Abs(numFloat)
	numFmtDesc, err := parseNumFormat(format)
	if err != nil {
		return empty_str, err
	}
	log.Printf("%#v\n", numFmtDesc)

	numParamDesc, err := parseNumParam(numStr)
	if err != nil {
		return empty_str, err
	}
	log.Printf("%#v\n", numParamDesc)

	result := bytes.Buffer{}

	if numFmtDesc.preSepValidLen < len(numParamDesc.preDec) {
		return empty_str, errors.New("格式的整数部分的长度不能比参数的整数部分的长度小")
	}

	// 左符号
	leftSign, rightSign := decorateSign(negative, numFmtDesc.signMode)
	if leftSign != signEmpty {
		if numFmtDesc.matchMode != matchModeFm {
			result.WriteByte(byte(leftSign))
		}
	}
	// 货币符号
	result.WriteString(string(numFmtDesc.currencySymbol))

	switch numFmtDesc.outputMode {
	// 十进制
	case outputModeEmpty:
		// 左符号
		// 货币符号
		// 分隔符前半部分: 0 9 逗号
		// 分隔符: 小数点
		// 分隔符后半部分: 0 9
		// 右符号
		outputDecimal(numParamDesc, numFmtDesc, negative, &result)
		// 右符号
		if rightSign != signEmpty {
			if numFmtDesc.matchMode != matchModeFm {
				result.WriteByte(byte(rightSign))
			}
		}
	// 科学计数
	case outputModeEEEE:
		// 左符号
		// 货币符号
		// 分隔符前半部分: 0 9
		// 分隔符: 小数点
		// 分隔符后半部分: 0 9
		// 右符号
		// TODO NB: Oracle中,会对最后一位进行四舍五入
		result.WriteString(strconv.FormatFloat(numFloat, 'E', len(numFmtDesc.postSep), 64))
		// 右符号
		if rightSign != signEmpty {
			if numFmtDesc.matchMode != matchModeFm {
				result.WriteByte(byte(rightSign))
			}
		}
	// 乘积 V 9 独占
	case outputModeV:
		// 左符号
		// 货币符号
		// 分隔符前半部分: 0 9 逗号
		// 分隔符: V
		// 分隔符后半部分: 0 9
		// 右符号
		// TODO NB: Oracle中,会对最后一位进行四舍五入
		result.WriteString(strconv.FormatInt(int64(numFloat*math.Pow10(len(numFmtDesc.postSep))), 10))
		// 右符号
		if rightSign != signEmpty {
			if numFmtDesc.matchMode != matchModeFm {
				result.WriteByte(byte(rightSign))
			}
		}
	// 十六进制 X 独占
	case outputModeX:
		result.WriteString(strconv.FormatFloat(numFloat, 'f', -1, 64))
	// 罗马计数 RN 独占
	case outputModeRN:
		result.WriteString(toRoman(int(numFloat)).String())
	// 最小文本 TM 独占
	case outputModeTM:
		result.WriteString(strconv.FormatFloat(numFloat, 'f', -1, 64))
	// 最小文本 TME 独占
	case outputModeTME:
		result.WriteString(strconv.FormatFloat(numFloat, 'E', -1, 64))
	default:
		return empty_str, errors.New(not_support_err)
	}
	return result.String(), nil
}

func decorateSign(negative bool, signMode signMode) (sign, sign) {
	leftSign := signEmpty
	rightSign := signEmpty

	switch signMode {
	case signModePR:
		if negative {
			leftSign = signLt
			rightSign = signGt
		}
	case signModeMI:
		if negative {
			leftSign = signSpace
			rightSign = signMinus
		}
	case signModeSStart:
		if negative {
			leftSign = signMinus
		} else {
			leftSign = signPlus
		}
	case signModeSEnd:
		if negative {
			rightSign = signMinus
		} else {
			rightSign = signPlus
		}
	}
	if leftSign == signEmpty && negative {
		leftSign = signMinus
	}
	return leftSign, rightSign
}

func outputDecimal(numParamDesc *NumParamDesc, numFmtDesc *NumFmtDesc, negative bool, result *bytes.Buffer) {
	// 分隔符前半部分: 0 9 逗号
	pPreLen := len(numParamDesc.preDec)
	fPreLen := len(numFmtDesc.preSep)

	commaCnt := fPreLen - numFmtDesc.preSepValidLen
	pRightLen := commaCnt + pPreLen
	pLeftLen := fPreLen - pRightLen

	firstZeroFound := false
	for i := 0; i < pLeftLen; i++ {
		if !firstZeroFound && numFmtDesc.preSep[i] == '0' {
			firstZeroFound = true
		}
		if firstZeroFound {
			result.WriteByte('0')
		}
	}

	j := 0
	for i := pLeftLen; i < fPreLen; i++ {
		if numFmtDesc.preSep[i] == ',' {
			if j != 0 {
				result.WriteByte(',')
			}
		} else {
			result.WriteByte(numParamDesc.preDec[j])
			j++
		}
	}

	// 分隔符: 小数点
	if numFmtDesc.postSep != empty_str {
		result.WriteByte('.')
		// TODO NB: Oracle中,会对最后一位进行四舍五入
		// 分隔符后半部分: 0 9
		for i := 0; i < len(numFmtDesc.postSep) && i < len(numParamDesc.postDec); i++ {
			result.WriteByte(numParamDesc.postDec[i])
		}
	}
}

func ToCharByDatetime(t time.Time, format string) (string, error) {
	fmKeywords, quoted, aux_flag, err := parseFmt(format)
	if err != nil {
		return empty_str, nil
	}

	//fixme
	println(aux_flag)

	result := bytes.Buffer{}

	qi := 0
	for i := 0; i < len(fmKeywords); i++ {
		switch fmKeywords[i] {
		case DCH_DOUBLE_QUOTE:
			result.WriteString(quoted[qi])
			qi++
		case DCH_SPACE:
			result.WriteByte(' ')
		case DCH_MINUS:
			result.WriteByte('-')
		case DCH_SLASH:
			result.WriteByte('/')
		case DCH_COMMA:
			result.WriteByte(',')
		case DCH_DEC:
			result.WriteByte('.')
		case DCH_SEMICOLON:
			result.WriteByte(';')
		case DCH_COLON:
			result.WriteByte(':')
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
			return empty_str, errors.New(not_support_err)
		case DCH_DAY:
			result.WriteString(NLS_WEEKS[t.Weekday()])
		case DCH_DDD:
			result.WriteString(strconv.Itoa(t.YearDay()))
		case DCH_DD:
			result.WriteString(strconv.Itoa(t.Day()))
		case DCH_DL:
			tmp, err := ToCharByDatetime(t, NLS_DL)
			if err != nil {
				return empty_str, nil
			}
			result.WriteString(tmp)
		case DCH_DS:
			tmp, err := ToCharByDatetime(t, NLS_DS)
			if err != nil {
				return empty_str, nil
			}
			result.WriteString(tmp)
		case DCH_DY:
			result.WriteString(NLS_WEEKS[t.Weekday()])
		case DCH_D:
			result.WriteString(strconv.Itoa(int(t.Weekday())))
		case DCH_E:
			return empty_str, errors.New(not_support_err)
		case DCH_EE:
			return empty_str, errors.New(not_support_err)
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
			result.WriteString(strconv.Itoa(toJulian(t.Year(), int(t.Month()), t.Day())))
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
			result.WriteString(toRoman(int(t.Month())).String())
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
			result.WriteString(NumToCardinalWord(t.Year()))
		case DCH_SYEAR:
			result.WriteString(NumToCardinalWord(t.Year() / 100))
			result.WriteString(SPACE)
			result.WriteString(NumToCardinalWord(t.Year() % 100))
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

// 解析日期格式
func parseFmt(format string) ([]dchKeyword, []string, int, error) {
	dchKeywords := []dchKeyword{}

	flen := len(format)

	quoted := []string{}
	aux_flag := 0

	var keyword dchKeyword
	var err error
	for fi := 0; fi < flen; {
		// 截取一个字符
		c := format[fi]
		if c >= 32 && c <= 127 {
			switch c {
			// DCH reproduced
			case '-':
				keyword = DCH_MINUS
				fi++
			case '/':
				keyword = DCH_SLASH
				fi++
			case ',':
				keyword = DCH_COMMA
				fi++
			case '.':
				keyword = DCH_DEC
				fi++
			case ';':
				keyword = DCH_SEMICOLON
				fi++
			case ':':
				keyword = DCH_COLON
				fi++
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
			case 'A', 'a':
				keyword, err = parsePrefixA(&fi, flen, format)
			case 'B', 'b':
				keyword, err = parsePrefixB(&fi, flen, format)
			case 'C', 'c':
				keyword, err = parsePrefixC(&fi, flen, format)
			case 'D', 'd':
				keyword, err = parsePrefixD(&fi, flen, format)
			case 'E', 'e':
				// TODO EE E
				return nil, nil, empty_int, errors.New(not_support_err)
			case 'F', 'f':
				keyword, err = parsePrefixF(&fi, flen, format, &aux_flag)
			case 'H', 'h':
				keyword, err = parsePrefixH(&fi, flen, format)
			case 'I', 'i':
				keyword, err = parsePrefixI(&fi, flen, format)
			case 'J', 'j':
				keyword = DCH_J
				fi++
			case 'M', 'm':
				keyword, err = parsePrefixM(&fi, flen, format)
			case 'P', 'p':
				keyword, err = parsePrefixP(&fi, flen, format)
			case 'Q', 'q':
				keyword = DCH_Q
				fi++
			case 'R', 'r':
				keyword, err = parsePrefixR(&fi, flen, format)
			case 'S', 's':
				keyword, err = parsePrefixS(&fi, flen, format, &aux_flag)
			case 'T', 't':
				keyword, err = parsePrefixT(&fi, flen, format, &aux_flag)
			case 'W', 'w':
				fi++
				if fi < flen && format[fi] == 'W' {
					// DCH WW
					keyword = DCH_WW
					fi++
				} else {
					// DCH W
					keyword = DCH_W
				}
			case 'X', 'x':
				keyword = DCH_X
				fi++
			case 'Y', 'y':
				keyword, err = parsePrefixY(&fi, flen, format)
			default:
				return nil, nil, aux_flag, errors.New(out_keyword_range_err)
			}

			if err != nil {
				return nil, nil, empty_int, err
			}
			dchKeywords = append(dchKeywords, keyword)
		} else {
			return nil, nil, empty_int, errors.New(out_ascii_range_err + string(c))
		}
	}

	return dchKeywords, quoted, aux_flag, nil
}

func parsePrefixA(fi *int, flen int, format string) (dchKeyword, error) {
	var keyword dchKeyword
	*fi++
	if *fi < flen {
		switch format[*fi] {
		case '.':
			*fi++
			start := *fi
			*fi += 2
			if *fi <= flen {
				if format[start:*fi] == "D." || format[start:*fi] == "d." {
					// DCH A.D.
					keyword = DCH_A_D_
				} else if format[start:*fi] == "M." || format[start:*fi] == "m." {
					// DCH A.M.
					keyword = DCH_A_M_
				} else {
					return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "A.")
				}
			} else {
				return DCH_EMPTY, errors.New(dch_fmt_length_err + "A.")
			}
		case 'D', 'd':
			// DCH AD
			keyword = DCH_AD
			*fi++
		case 'M', 'm':
			// DCH AM
			keyword = DCH_AM
			*fi++
		default:
			return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "A")
		}
	} else {
		return DCH_EMPTY, errors.New(dch_fmt_length_err + "A")
	}
	return keyword, nil
}

func parsePrefixB(fi *int, flen int, format string) (dchKeyword, error) {
	var keyword dchKeyword
	*fi++
	if *fi < flen {
		switch format[*fi] {
		case 'C', 'c':
			// DCH BC
			keyword = DCH_BC
			*fi++
		case '.':
			*fi++
			start := *fi
			*fi += 2
			if *fi <= flen && (format[start:*fi] == "C." || format[start:*fi] == "c.") {
				// DCH B.C.
				keyword = DCH_B_C_
			} else {
				return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "B.")
			}
		default:
			return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "B")
		}
	} else {
		return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "B")
	}

	return keyword, nil
}

func parsePrefixC(fi *int, flen int, format string) (dchKeyword, error) {
	var keyword dchKeyword
	*fi++
	if *fi < flen {
		switch format[*fi] {
		case 'C', 'c':
			// DCH CC
			keyword = DCH_CC
			*fi++
		default:
			return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "C")
		}
	} else {
		return DCH_EMPTY, errors.New(dch_fmt_length_err + "C")
	}

	return keyword, nil
}

func parsePrefixD(fi *int, flen int, format string) (dchKeyword, error) {
	var keyword dchKeyword
	*fi++
	if *fi < flen {
		if format[*fi] == 'A' || format[*fi] == 'a' {
			*fi++
			if *fi < flen && format[*fi] == 'Y' {
				// DCH DAY 同 DY
				keyword = DCH_DAY
				*fi++
			} else {
				return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "DA")
			}
		} else if format[*fi] == 'D' || format[*fi] == 'd' {
			*fi++
			if *fi < flen && (format[*fi] == 'D' || format[*fi] == 'd') {
				// DCH DDD
				keyword = DCH_DDD
				*fi++
			} else {
				// DCH DD
				keyword = DCH_DD
			}
		} else if format[*fi] == 'L' || format[*fi] == 'l' {
			keyword = DCH_DL
			*fi++
		} else if format[*fi] == 'S' || format[*fi] == 's' {
			keyword = DCH_DS
			*fi++
		} else if format[*fi] == 'Y' || format[*fi] == 'y' {
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

func parsePrefixF(fi *int, flen int, format string, flag *int) (dchKeyword, error) {
	var keyword dchKeyword
	*fi++
	if *fi < flen {
		switch format[*fi] {
		case 'X', 'x':
			// TODO 最后处理
			*flag |= mode_flag_fx
		case 'M', 'm':
			// TODO 最后处理
			*flag |= mode_flag_fm
		case 'F', 'f':
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
			return DCH_EMPTY, errors.New(dch_fmt_length_err + "F")
		}
	} else {
		return DCH_EMPTY, errors.New(dch_fmt_length_err + "F")
	}
	return keyword, nil
}

func parsePrefixH(fi *int, flen int, format string) (dchKeyword, error) {
	var keyword dchKeyword
	*fi++
	if *fi < flen {
		if format[*fi] == 'H' || format[*fi] == 'h' {
			*fi++
			if format[*fi] == '2' {
				*fi++
				if *fi < flen {
					// DCH HH24
					if format[*fi] == '4' {
						keyword = DCH_HH24
					}
				}
			} else if format[*fi] == '1' {
				*fi++
				if *fi < flen {
					// DCH HH12
					if format[*fi] == '2' {
						keyword = DCH_HH12
					}
				}
			}

			if keyword == DCH_EMPTY {
				// DCH HH 同 HH12
				keyword = DCH_HH
			}
		} else {
			return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "H")
		}
	} else {
		return DCH_EMPTY, errors.New(dch_fmt_length_err + "H")
	}
	*fi++
	return keyword, nil
}

func parsePrefixI(fi *int, flen int, format string) (dchKeyword, error) {
	var keyword dchKeyword
	*fi++
	if *fi < flen {
		switch format[*fi] {
		case 'W', 'w':
			// DCH IW
			keyword = DCH_IW
			*fi++
		case 'Y', 'y':
			*fi++
			if *fi < flen && (format[*fi] == 'Y' || format[*fi] == 'y') {
				*fi++
				if *fi < flen && (format[*fi] == 'Y' || format[*fi] == 'y') {
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

func parsePrefixM(fi *int, flen int, format string) (dchKeyword, error) {
	var keyword dchKeyword
	*fi++
	if *fi < flen && (format[*fi] == 'I' || format[*fi] == 'i') {
		// DCH MI
		keyword = DCH_MI
		*fi++
	} else if *fi < flen && (format[*fi] == 'M' || format[*fi] == 'm') {
		// DCH MM
		keyword = DCH_MM
		*fi++
	} else if *fi < flen && (format[*fi] == 'O' || format[*fi] == 'o') {
		*fi++
		if *fi < flen && (format[*fi] == 'N' || format[*fi] == 'n') {
			*fi++
			start := *fi
			*fi += 2
			if *fi <= flen && strings.ToUpper(format[start:*fi]) == "TH" {
				// DCH MONTH
				keyword = DCH_MONTH
			} else {
				// DCH MON
				keyword = DCH_MON
				*fi -= 2
			}
		} else {
			return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "MO")
		}
	} else {
		return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "M")
	}
	return keyword, nil
}

func parsePrefixP(fi *int, flen int, format string) (dchKeyword, error) {
	var keyword dchKeyword
	*fi++
	if *fi < flen {
		if format[*fi] == 'M' || format[*fi] == 'm' {
			keyword = DCH_PM
			*fi++
		} else if format[*fi] == '.' {
			*fi++
			start := *fi
			*fi += 2
			if *fi <= flen {
				if "M." == format[start:*fi] {
					keyword = DCH_P_M_
				} else {
					return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "P")
				}
			} else {
				return DCH_EMPTY, errors.New(dch_fmt_length_err + "P")
			}
		} else {
			return DCH_EMPTY, errors.New(dch_fmt_length_err + "P")
		}
	} else {
		return DCH_EMPTY, errors.New(dch_fmt_length_err + "P")
	}
	return keyword, nil
}

func parsePrefixR(fi *int, flen int, format string) (dchKeyword, error) {
	var keyword dchKeyword
	*fi++
	if *fi < flen {
		if format[*fi] == 'M' || format[*fi] == 'm' {
			keyword = DCH_RM
			*fi++
		} else if format[*fi] == 'R' || format[*fi] == 'r' {
			*fi++
			start := *fi
			*fi += 2
			if *fi <= flen && strings.ToUpper(format[start:*fi]) == "RR" {
				// DCH RRRR
				keyword = DCH_RRRR
			} else {
				// DCH RR
				keyword = DCH_RR
				*fi -= 2
			}
		} else {
			return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "R")
		}
	} else {
		return DCH_EMPTY, errors.New(dch_fmt_length_err + "R")
	}
	return keyword, nil
}

func parsePrefixS(fi *int, flen int, format string, flag *int) (dchKeyword, error) {
	var keyword dchKeyword
	*fi++
	if *fi < flen {
		switch format[*fi] {
		case 'P', 'p':
			// DCH SP TODO 最后处理
			*flag |= mode_flag_sp
			*fi++
		case 'S', 's':
			*fi++
			start := *fi
			*fi += 3
			if *fi <= flen && strings.ToUpper(format[start:*fi]) == "SSS" {
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
				if strings.ToUpper(format[start:*fi]) == "YYY" {
					// TODO golang 好像不支持公元前
					// DCH SYYYY 正负号+数字
					//if 公元前 {result.WriteByte('-')}
					keyword = DCH_SYYYY
				} else if strings.ToUpper(format[start:*fi]) == "EAR" {
					// FIXME oracle中将4位的年分成了 2个2位数
					// DCH SYEAR 正负号+基数词
					//if 公元前 {result.WriteByte('-')}
					keyword = DCH_SYEAR
				} else {
					return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "SY")
				}
			} else {
				return DCH_EMPTY, errors.New(dch_fmt_length_err + "S")
			}
		default:
			return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "S")
		}
	} else {
		return DCH_EMPTY, errors.New(dch_fmt_length_err + "S")
	}
	return keyword, nil
}

func parsePrefixT(fi *int, flen int, format string, flag *int) (dchKeyword, error) {
	var keyword dchKeyword
	// TODO 更换类型后更改时区
	*fi++
	if *fi < flen {
		if format[*fi] == 'S' || format[*fi] == 's' {
			// DCH TS 下午 9:30:00
			keyword = DCH_TS
		} else if format[*fi] == 'Z' || format[*fi] == 'z' {
			*fi++
			if *fi < flen && format[*fi] == 'D' {
				// DCH TZD PDT 时区
				keyword = DCH_TZD
			} else if *fi < flen && (format[*fi] == 'H' || format[*fi] == 'h') {
				// DCH TZH -07 时区小时
				keyword = DCH_TZH
			} else if *fi < flen && (format[*fi] == 'M' || format[*fi] == 'm') {
				// DCH TZM 00 时区分
				keyword = DCH_TZM
			} else if *fi < flen && (format[*fi] == 'R' || format[*fi] == 'r') {
				// DCH TZR US/PACIFIC 时区区域
				keyword = DCH_TZR
			} else {
				return DCH_EMPTY, errors.New("格式错误")
			}
		} else if format[*fi] == 'H' || format[*fi] == 'h' {
			// DCH TH TODO 最后处理
			//keyword = DCH_TH
			*flag |= mode_flag_th
		} else {
			return DCH_EMPTY, errors.New("格式错误")
		}
	} else {
		return DCH_EMPTY, errors.New("格式错误")
	}
	*fi++
	return keyword, nil
}
func parsePrefixY(fi *int, flen int, format string) (dchKeyword, error) {
	var keyword dchKeyword
	*fi++
	if *fi < flen {
		if format[*fi] == ',' {
			*fi++
			start := *fi
			*fi += 3
			if *fi <= flen {
				if strings.ToUpper(format[start:*fi]) == "YYY" {
					// DCH Y,YYY
					keyword = DCH_Y_YYY
				} else {
					return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "Y,")
				}
			} else {
				return DCH_EMPTY, errors.New(dch_fmt_length_err + "Y,")
			}
		} else if format[*fi] == 'Y' || format[*fi] == 'y' {
			*fi++

			if *fi < flen && (format[*fi] == 'Y' || format[*fi] == 'y') {
				*fi++
				if *fi < flen && (format[*fi] == 'Y' || format[*fi] == 'y') {
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
		} else if format[*fi] == 'E' || format[*fi] == 'e' {
			*fi++
			start := *fi
			*fi += 2
			if *fi <= flen && strings.ToUpper(format[start:*fi]) == "AR" {
				// DCH YEAR 基数词
				keyword = DCH_YEAR
			} else {
				return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "YE")
			}
		}
	} else {
		// DCH Y
		keyword = DCH_Y
	}
	return keyword, nil
}
