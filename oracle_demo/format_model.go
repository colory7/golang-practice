package oracle_demo

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// 格式部分不匹配，报错
const dch_fmt_mismatch_err = "Date Format error, some formats do not match near "
const dch_fmt_length_err = "Date Format error, incorrect format length near "
const num_fmt_part_err = "Datetime Format error, some formats do not match near "
const not_support_err = "not support"
const unreachable_err = "unreachable"

// 非法字符,超出格式关键词范围
const out_keyword_range_err = "Illegal character, not in the range of Format Model keyword"

// 非法字符,超出ASCII[32-126]字符范围
const out_ascii_range_err = "Illegal character, not in ASCII [32-126] character range"

const invalid_num_err = "invalid number"

type FMKeyword int

type DateFmtDesc struct {
}

const (
	mode_flag_fm = 1
	mode_flag_fx = 1 << 1
	mode_flag_th = 1 << 2
	mode_flag_sp = 1 << 3
)

type dtType int

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
	DCH_EMPTY = iota
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

var dchKeywords map[int]int

var NLS_WEEKS = map[time.Weekday]string{}
var NLS_MONTHS = map[time.Month]string{}
var NLS_MONTHS_REVERSE = map[string]time.Month{}

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
	empty_str   = ""
	empty_int   = 0
	empty_float = 0.0
	tsFormat    = "15:04:05"
	//dateFormat = "YYYY-MM-DD HH24:MI:SS"
	dateLayout = "2006-01-02 15:04:05"
)

const (
	empty = ""
	plus  = "+"
	minus = "-"
	dec   = "."
)

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
	NLS_MONTHS_REVERSE = map[string]time.Month{
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

}

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
	// 后缀 输出模式
	suffix int
	// 辅助后缀
	auxSuffix int
}

type NumParamDesc struct {
	sign      string
	preDec    string
	postDec   string
	eSign     string
	eExponent int
	// 可选
	hasDec bool
	isEEEE bool
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
	li := flen - 1

	var c byte

	readDec := false
	//var preDec = bytes.Buffer{}
	//var postDec = bytes.Buffer{}
	for fi := 0; fi < flen; {
		// 截取一个字符
		c = format[fi]
		if c >= 32 && c <= 127 {
			// 这里设置为不区分大小写。NB: Oracle和Postgresql中为区分大小写
			toUpper(&c)

			// 匹配关键词并存储
			switch c {
			case ',':
				if fmtDesc.commaIndex == 0 {
					if fi == 0 {
						return fmtDesc, errors.New("不能以逗号开头")
					} else if fi == li {
						return fmtDesc, errors.New("逗号不能出现在数字最右边")
					} else if fmtDesc.decIndex != -1 {
						return fmtDesc, errors.New("逗号不能出现在点号的右边")
					}

					fmtDesc.commaIndex = fi
				} else {
					return fmtDesc, errors.New("格式错误，存在多个格式符号 ,")
				}

				fi++
			case '.':
				if !readDec {
					fmtDesc.decIndex = fi
					readDec = true
				} else {
					return fmtDesc, errors.New("只能有1个 .")
				}
				fi++
			case '0':
				if readDec {
					fmtDesc.postDec++
					//postDec.WriteByte('0')
				} else {
					fmtDesc.preDec++
					//preDec.WriteByte('0')
				}
				fi++
			case '9':
				if readDec {
					fmtDesc.postDec++
					//postDec.WriteByte('9')
				} else {
					fmtDesc.preDec++
					//preDec.WriteByte('9')
				}
				fi++
			case '$':
				if fmtDesc.prefix == NUM_FMT_PREFIX_EMPTY && fi == 0 {
					fmtDesc.prefix = NUM_FMT_PREFIX_DOLLAR
				} else {
					return fmtDesc, errors.New("格式前缀冲突 " + "$")
				}
				fi++
			case 'B':
				if fmtDesc.prefix == NUM_FMT_PREFIX_EMPTY {
					fmtDesc.prefix = NUM_FMT_PREFIX_B
				} else {
					return fmtDesc, errors.New("格式前缀冲突 " + "B")
				}
				fi++
			case 'C':
				if fmtDesc.prefix == NUM_FMT_PREFIX_EMPTY {
					return fmtDesc, errors.New("格式前缀冲突 " + "C")
				} else if li != fi && 0 != fi {
					return fmtDesc, errors.New("C 只能在开头或者结尾")
				}
				fmtDesc.prefix = NUM_FMT_PREFIX_C
				fi++
			case 'L':
				if fmtDesc.prefix == NUM_FMT_PREFIX_EMPTY {
					fmtDesc.prefix = NUM_FMT_PREFIX_L
				} else {
					return fmtDesc, errors.New("格式前缀冲突 " + "L")
				}
				fi++
			case 'U':
				if fmtDesc.prefix == NUM_FMT_PREFIX_EMPTY {
					return fmtDesc, errors.New("格式前缀冲突 " + "U")
				} else if li != fi && 0 != fi {
					return fmtDesc, errors.New("U 只能在开头或者结尾")
				}
				fmtDesc.prefix = NUM_FMT_PREFIX_U
				fi++
			case 'E':
				if fmtDesc.suffix == NUM_FMT_SUFFIX_EMPTY {
					start := fi + 1
					fi += 4
					if "EEE" == format[start:fi] {
						fmtDesc.suffix = NUM_FMT_SUFFIX_EEEE
					} else {
						return fmtDesc, errors.New(num_fmt_part_err + "E")
					}
				} else {
					return fmtDesc, errors.New("conflict with E")
				}
			case 'F':
				fi++

				if format[fi] == 'M' {
					if fmtDesc.auxPrefix == NUM_FMT_AUX_PREFIX_EMPTY {
						return fmtDesc, errors.New("只能有1组 FM")
					}
					if fi == 1 {
						fmtDesc.auxPrefix = NUM_FMT_AUX_PREFIX_FM
					} else {
						return fmtDesc, errors.New("FM 必须在开头")
					}
				} else {
					return fmtDesc, errors.New(num_fmt_part_err + "F")
				}
				fi++
			case 'M':
				if fmtDesc.auxSuffix == NUM_FMT_AUX_SUFFIX_EMPTY {
					return fmtDesc, errors.New("辅助后缀冲突" + "MI")
				} else if fi == (li - 1) {
					return fmtDesc, errors.New("MI 只能在结尾")
				}

				fi++
				if format[fi] == 'I' {
					fmtDesc.auxSuffix = NUM_FMT_AUX_SUFFIX_MI
				} else {
					return fmtDesc, errors.New(num_fmt_part_err + "M")
				}
				fi++
			case 'P':
				if fmtDesc.auxSuffix == NUM_FMT_AUX_SUFFIX_EMPTY {
					return fmtDesc, errors.New("辅助后缀冲突" + "PR")
				} else if fi == (li - 1) {
					return fmtDesc, errors.New("PR 只能在结尾")
				}

				fi++
				if format[fi] == 'R' {
					fmtDesc.auxSuffix = NUM_FMT_AUX_SUFFIX_PR
				} else {
					return fmtDesc, errors.New(num_fmt_part_err + "P")
				}
				fi++
			case 'R':
				// 判断独占 长度 FIXME
				fi++
				if format[fi] == 'N' {
					if fmtDesc.suffix == NUM_FMT_SUFFIX_EMPTY {
						return fmtDesc, errors.New("只能有1个 RN")
					} else if fmtDesc.auxPrefix == NUM_FMT_AUX_PREFIX_FM && flen == 4 {
						return fmtDesc, errors.New("包含RN的格式,除了 FM 和 RN 不能有其他格式字符")
					} else if fmtDesc.auxPrefix == NUM_FMT_AUX_PREFIX_EMPTY && flen == 2 {
						return fmtDesc, errors.New("包含RN的格式,除了 FM 和 RN 不能有其他格式字符")
					}

					fmtDesc.suffix = NUM_FMT_SUFFIX_RN
				} else {
					return fmtDesc, errors.New(num_fmt_part_err + "R")
				}
				fi++
			case 'S':
				if fmtDesc.s == NUM_FMT_S_EMPTY {
					return fmtDesc, errors.New("只能有1个 S")
				} else if fi == li && fi != 0 {
					return fmtDesc, errors.New("S 只能在开头或者结尾")
				}

				if fi == 0 {
					fmtDesc.s = NUM_FMT_S_START
				} else {
					fmtDesc.s = NUM_FMT_S_END
				}
				fi++
			case 'T':
				if fmtDesc.suffix == NUM_FMT_SUFFIX_EMPTY {
					fi++
					if 'M' == format[fi] {
						if fi == li {
							fmtDesc.suffix = NUM_FMT_SUFFIX_TM
						} else {
							fi++
							if 'E' == format[fi] {
								fmtDesc.suffix = NUM_FMT_SUFFIX_TME
							} else if 'e' == format[fi] { // FIXME 已经转换为了大写
								fmtDesc.suffix = NUM_FMT_SUFFIX_TMe
							} else if '9' == format[fi] {
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
				fi++
			case 'V':
				if fmtDesc.suffix == NUM_FMT_SUFFIX_EMPTY {
					return fmtDesc, errors.New("只能有1个 V")
				} else if 0 != fi {
					return fmtDesc, errors.New("V 不能在开头")
				}
				fmtDesc.suffix = NUM_FMT_SUFFIX_V
				fi++
			case 'X':
				if fmtDesc.suffix == NUM_FMT_SUFFIX_EMPTY {
					return fmtDesc, errors.New("conflict with X")
				} else if 0 != fi {
					return fmtDesc, errors.New("V 不能在开头")
				}

				fmtDesc.suffix = NUM_FMT_SUFFIX_X
				fmtDesc.xCount++
				fi++
			default:
				return fmtDesc, errors.New(out_keyword_range_err)
			}
		} else {
			return fmtDesc, errors.New(out_ascii_range_err)
		}

	}

	return fmtDesc, nil

}

// 解析数字参数
func parseNumParam(num string) (NumParamDesc, error) {
	var paramDesc NumParamDesc

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
				paramDesc.hasDec = true
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
func ToNumberByStr(num string, format string) (float64, error) {
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

	flen := len(format)
	switch numFmtDesc.suffix {
	// 十进制
	case NUM_FMT_SUFFIX_EMPTY:
		if numFmtDesc.preDec < len(numParamDesc.preDec) {
			return empty_float, errors.New("格式的整数部分的长度不能比参数的整数部分的长度小")
		}
		f, err := strconv.ParseFloat(num, 64)
		if err != nil {
			return empty_float, err
		}
		return f, nil
	// 十六进制
	case NUM_FMT_SUFFIX_X:
		d, err := strconv.ParseInt(num, 16, 64)
		if err != nil {
			return empty_float, err
		}
		if len(numParamDesc.preDec) > flen {
			return empty_float, errors.New("格式长度比数值长度小")
		}
		return float64(d), nil
	default:
		return empty_float, errors.New(not_support_err)
	}

	return empty_float, nil
}

func ToTimestamp(dch string, format string) (*time.Time, error) {
	return toDT(dch, format, dt_type_timestamp)
}

func ToTimestampTimeZone(dch string, format string) (*time.Time, error) {
	return toDT(dch, format, dt_type_timestamp_tz)
}

func ToDate(dch string, format string) (*time.Time, error) {
	return toDT(dch, format, dt_type_date)
}

func toDT(dch string, format string, tp dtType) (*time.Time, error) {
	fmKeywords, quoted, aux_flag, err := parseFmt(format)
	if err != nil {
		return nil, nil
	}

	year, month, day := 0, time.Month(0), 0
	hour, min, sec, nsec := 0, 0, 0, 0
	tzr := time.Local

	//dItems := ParseDchByTime(dch, flag)

	//if len(fmKeywords) != len(dItems) {
	//	return nil, errors.New("格式长度与参数长度不匹配")
	//}

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

	for ki := 0; ki < len(fmKeywords); ki++ {
		switch fmKeywords[ki] {
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
				day, err = strconv.Atoi(field)
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

func ToDatetimeChar(t time.Time, format string) (string, error) {
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
			tmp, err := ToDatetimeChar(t, NLS_DL)
			if err != nil {
				return empty_str, nil
			}
			result.WriteString(tmp)
		case DCH_DS:
			tmp, err := ToDatetimeChar(t, NLS_DS)
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
func parseFmt(format string) ([]int, []string, int, error) {
	fmKeywords := []int{}

	flen := len(format)

	quoted := []string{}
	aux_flag := 0

	var keyword FMKeyword
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
				return nil, nil, empty_int, errors.New(not_support_err)
			case 'F':
				keyword, err = parsePrefixF(&fi, flen, format, &aux_flag)
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
				keyword, err = parsePrefixS(&fi, flen, format, &aux_flag)
			case 'T':
				keyword, err = parsePrefixT(&fi, flen, format, &aux_flag)
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
				return nil, nil, aux_flag, errors.New(out_keyword_range_err)
			}

			if err != nil {
				return nil, nil, empty_int, err
			}
			fmKeywords = append(fmKeywords, int(keyword))
		} else {
			return nil, nil, empty_int, errors.New(out_ascii_range_err + string(c))
		}
	}

	return fmKeywords, quoted, aux_flag, nil
}

func parsePrefixA(fi *int, flen int, format string) (FMKeyword, error) {
	var keyword FMKeyword
	*fi++
	if *fi < flen {
		switch format[*fi] {
		case '.':
			*fi++
			start := *fi
			*fi += 2
			if *fi <= flen {
				if "D." == format[start:*fi] {
					// DCH A.D.
					keyword = DCH_A_D_
				} else if "M." == format[start:*fi] {
					// DCH A.M.
					keyword = DCH_A_M_
				} else {
					return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "A.")
				}
			} else {
				return DCH_EMPTY, errors.New(dch_fmt_length_err + "A.")
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
			return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "A")
		}
	} else {
		return DCH_EMPTY, errors.New(dch_fmt_length_err + "A")
	}
	return keyword, nil
}

func parsePrefixB(fi *int, flen int, format string) (FMKeyword, error) {
	var keyword FMKeyword
	*fi++
	if *fi < flen {
		switch format[*fi] {
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

func parsePrefixC(fi *int, flen int, format string) (FMKeyword, error) {
	var keyword FMKeyword
	*fi++
	if *fi < flen {
		switch format[*fi] {
		case 'C':
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
				return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "DA")
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

func parsePrefixF(fi *int, flen int, format string, flag *int) (FMKeyword, error) {
	var keyword FMKeyword
	*fi++
	if *fi < flen {
		switch format[*fi] {
		case 'X':
			// TODO 最后处理
			*flag |= mode_flag_fx
		case 'M':
			// TODO 最后处理
			*flag |= mode_flag_fm
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
			return DCH_EMPTY, errors.New(dch_fmt_length_err + "F")
		}
	} else {
		return DCH_EMPTY, errors.New(dch_fmt_length_err + "F")
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
					return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "H2")
				}
			} else {
				return DCH_EMPTY, errors.New(dch_fmt_length_err + "H2")
			}
		case '1':
			*fi++
			if *fi < flen {
				// DCH HH12
				if format[*fi] == '2' {
					keyword = DCH_HH12
				} else {
					return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "H2")
				}
			} else {
				return DCH_EMPTY, errors.New(dch_fmt_length_err + "H1")
			}
		}
	} else {
		return DCH_EMPTY, errors.New(dch_fmt_length_err + "H")
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
			return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "MO")
		}
	} else {
		return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "M")
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

func parsePrefixR(fi *int, flen int, format string) (FMKeyword, error) {
	var keyword FMKeyword
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
			return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "R")
		}
	} else {
		return DCH_EMPTY, errors.New(dch_fmt_length_err + "R")
	}
	return keyword, nil
}

func parsePrefixS(fi *int, flen int, format string, flag *int) (FMKeyword, error) {
	var keyword FMKeyword
	*fi++
	if *fi < flen {
		switch format[*fi] {
		case 'P':
			// DCH SP TODO 最后处理
			*flag |= mode_flag_sp
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

func parsePrefixT(fi *int, flen int, format string, flag *int) (FMKeyword, error) {
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
				return DCH_EMPTY, errors.New("格式错误")
			}
		} else if format[*fi] == 'H' {
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
					return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "Y,")
				}
			} else {
				return DCH_EMPTY, errors.New(dch_fmt_length_err + "Y,")
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
				return DCH_EMPTY, errors.New(dch_fmt_mismatch_err + "YE")
			}
		}
	} else {
		// DCH Y
		keyword = DCH_Y
	}
	return keyword, nil
}
func toUpper(c *byte) {
	if 'a' <= *c && *c <= 'z' {
		*c -= 32
	}
}

func toRoman(num int) *bytes.Buffer {
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

func toJulian(year int, month int, day int) int {
	adj := (14 - month) / 12
	y := year + 4800 - adj
	m := month + 12*adj - 3
	return day + (153*m+2)/5 + y*365 + y/4 - y/100 + y/400 - 32045
}

func toRRRR(thisYear int, RR int) int {
	year := 0
	firstTwo := thisYear / 100
	lastTwo := thisYear % 100

	// 0-49
	if lastTwo >= 0 && lastTwo <= 49 {
		if RR >= 0 && lastTwo <= 49 {
			year = firstTwo*100 + RR
		} else {
			year = (firstTwo-1)*100 + RR
		}
	} else {
		// 50-99
		if RR >= 50 && lastTwo <= 99 {
			year = firstTwo*100 + RR
		} else {
			year = (firstTwo+1)*100 + RR
		}
	}
	return year
}
