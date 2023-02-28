package string_demo

import (
	"fmt"
	"strings"
	"testing"
)

type MyLocale struct {
	name         string
	description  string
	decimalPoint byte
	thousandSep  byte
}

var mysqlLocales = map[string]MyLocale{
	"en_US": {name: "en_US", description: "English - United States", decimalPoint: '.', thousandSep: ','},
	"en_GB": {name: "en_GB", description: "English - United Kingdom", decimalPoint: '.', thousandSep: ','},
	"ja_JP": {name: "ja_JP", description: "Japanese - Japan", decimalPoint: '.', thousandSep: ','},
	"sv_SE": {name: "sv_SE", description: "Swedish - Sweden", decimalPoint: ',', thousandSep: ' '},
	"de_DE": {name: "de_DE", description: "German - Germany", decimalPoint: ',', thousandSep: '.'},
	"fr_FR": {name: "fr_FR", description: "French - France", decimalPoint: ',', thousandSep: byte(0)},
	"ar_AE": {name: "ar_AE", description: "Arabic - United Arab Emirates", decimalPoint: '.', thousandSep: ','},
	"ar_BH": {name: "ar_BH", description: "Arabic - Bahrain", decimalPoint: '.', thousandSep: ','},
	"ar_JO": {name: "ar_JO", description: "Arabic - Jordan", decimalPoint: '.', thousandSep: ','},
	"ar_SA": {name: "ar_SA", description: "Arabic - Saudi Arabia", decimalPoint: '.', thousandSep: byte(0)},
	"ar_SY": {name: "ar_SY", description: "Arabic - Syria", decimalPoint: '.', thousandSep: ','},
	"be_BY": {name: "be_BY", description: "Belarusian - Belarus", decimalPoint: ',', thousandSep: '.'},
	"bg_BG": {name: "bg_BG", description: "Bulgarian - Bulgaria", decimalPoint: ',', thousandSep: byte(0)},
	"ca_ES": {name: "ca_ES", description: "Catalan - Catalan", decimalPoint: ',', thousandSep: byte(0)},
	"cs_CZ": {name: "cs_CZ", description: "Czech - Czech Republic", decimalPoint: ',', thousandSep: ' '},
	"da_DK": {name: "da_DK", description: "Danish - Denmark", decimalPoint: ',', thousandSep: '.'},
	"de_AT": {name: "de_AT", description: "German - Austria", decimalPoint: ',', thousandSep: byte(0)},
	"es_ES": {name: "es_ES", description: "Spanish - Spain", decimalPoint: ',', thousandSep: '.'},
	"et_EE": {name: "et_EE", description: "Estonian - Estonia", decimalPoint: ',', thousandSep: ' '},
	"eu_ES": {name: "eu_ES", description: "Basque - Basque", decimalPoint: ',', thousandSep: byte(0)},
	"fi_FI": {name: "fi_FI", description: "Finnish - Finland", decimalPoint: ',', thousandSep: ' '},
	"fo_FO": {name: "fo_FO", description: "Faroese - Faroe Islands", decimalPoint: ',', thousandSep: '.'},
	"gl_ES": {name: "gl_ES", description: "Galician - Galician", decimalPoint: ',', thousandSep: byte(0)},
	"gu_IN": {name: "gu_IN", description: "Gujarati - India", decimalPoint: '.', thousandSep: ','},
	"he_IL": {name: "he_IL", description: "Hebrew - Israel", decimalPoint: '.', thousandSep: ','},
	"hi_IN": {name: "hi_IN", description: "Hindi - India", decimalPoint: '.', thousandSep: ','},
	"hr_HR": {name: "hr_HR", description: "Croatian - Croatia", decimalPoint: ',', thousandSep: byte(0)},
	"hu_HU": {name: "hu_HU", description: "Hungarian - Hungary", decimalPoint: ',', thousandSep: '.'},
	"id_ID": {name: "id_ID", description: "Indonesian - Indonesia", decimalPoint: ',', thousandSep: '.'},
	"is_IS": {name: "is_IS", description: "Icelandic - Iceland", decimalPoint: ',', thousandSep: '.'},
	"it_CH": {name: "it_CH", description: "Italian - Switzerland", decimalPoint: ',', thousandSep: '\''},
	"ko_KR": {name: "ko_KR", description: "Korean - Korea", decimalPoint: '.', thousandSep: ','},
	"lt_LT": {name: "lt_LT", description: "Lithuanian - Lithuania", decimalPoint: ',', thousandSep: '.'},
	"lv_LV": {name: "lv_LV", description: "Latvian - Latvia", decimalPoint: ',', thousandSep: ' '},
	"mk_MK": {name: "mk_MK", description: "Macedonian - FYROM", decimalPoint: ',', thousandSep: ' '},
	"mn_MN": {name: "mn_MN", description: "Mongolia - Mongolian", decimalPoint: ',', thousandSep: '.'},
	"ms_MY": {name: "ms_MY", description: "Malay - Malaysia", decimalPoint: '.', thousandSep: ','},
	"nb_NO": {name: "nb_NO", description: "Norwegian(Bokml) - Norway", decimalPoint: ',', thousandSep: '.'},
	"nl_NL": {name: "nl_NL", description: "Dutch - The Netherlands", decimalPoint: ',', thousandSep: byte(0)},
	"pl_PL": {name: "pl_PL", description: "Polish - Poland", decimalPoint: ',', thousandSep: byte(0)},
	"pt_BR": {name: "pt_BR", description: "Portugese - Brazil", decimalPoint: ',', thousandSep: byte(0)},
	"pt_PT": {name: "pt_PT", description: "Portugese - Portugal", decimalPoint: ',', thousandSep: byte(0)},
	"ro_RO": {name: "ro_RO", description: "Romanian - Romania", decimalPoint: ',', thousandSep: '.'},
	"ru_RU": {name: "ru_RU", description: "Russian - Russia", decimalPoint: ',', thousandSep: ' '},
	"ru_UA": {name: "ru_UA", description: "Russian - Ukraine", decimalPoint: ',', thousandSep: '.'},
	"sk_SK": {name: "sk_SK", description: "sk_SK", decimalPoint: ',', thousandSep: ' '},
	"sl_SI": {name: "sl_SI", description: "Slovenian - Slovenia", decimalPoint: ',', thousandSep: ' '},
	"sq_AL": {name: "sq_AL", description: "Albanian - Albania", decimalPoint: ',', thousandSep: '.'},
	"sr_RS": {name: "sr_RS", description: "Serbian - Serbia", decimalPoint: '.', thousandSep: byte(0)},
	"ta_IN": {name: "ta_IN", description: "Tamil - India", decimalPoint: '.', thousandSep: ','},
	"te_IN": {name: "te_IN", description: "Telugu - India", decimalPoint: '.', thousandSep: ','},
	"th_TH": {name: "th_TH", description: "Thai - Thailand", decimalPoint: '.', thousandSep: ','},
	"tr_TR": {name: "tr_TR", description: "Turkish - Turkey", decimalPoint: ',', thousandSep: '.'},
	"uk_UA": {name: "uk_UA", description: "Ukrainian - Ukraine", decimalPoint: ',', thousandSep: '.'},
	"ur_PK": {name: "ur_PK", description: "Urdu - Pakistan", decimalPoint: '.', thousandSep: ','},
	"vi_VN": {name: "vi_VN", description: "Vietnamese - Vietnam", decimalPoint: ',', thousandSep: '.'},
	"zh_CN": {name: "zh_CN", description: "Chinese - Peoples Republic of China", decimalPoint: '.', thousandSep: ','},
	"zh_TW": {name: "zh_TW", description: "Chinese - Taiwan", decimalPoint: '.', thousandSep: ','},
	"ar_DZ": {name: "ar_DZ", description: "Arabic - Algeria", decimalPoint: '.', thousandSep: ','},
	"ar_EG": {name: "ar_EG", description: "Arabic - Egypt", decimalPoint: '.', thousandSep: ','},
	"ar_IN": {name: "ar_IN", description: "Arabic - Iran", decimalPoint: '.', thousandSep: ','},
	"ar_IQ": {name: "ar_IQ", description: "Arabic - Iraq", decimalPoint: '.', thousandSep: ','},
	"ar_KW": {name: "ar_KW", description: "Arabic - Kuwait", decimalPoint: '.', thousandSep: ','},
	"ar_LB": {name: "ar_LB", description: "Arabic - Lebanon", decimalPoint: '.', thousandSep: ','},
	"ar_LY": {name: "ar_LY", description: "Arabic - Libya", decimalPoint: '.', thousandSep: ','},
	"ar_MA": {name: "ar_MA", description: "Arabic - Morocco", decimalPoint: '.', thousandSep: ','},
	"ar_OM": {name: "ar_OM", description: "Arabic - Oman", decimalPoint: '.', thousandSep: ','},
	"ar_QA": {name: "ar_QA", description: "Arabic - Qatar", decimalPoint: '.', thousandSep: ','},
	"ar_SD": {name: "ar_SD", description: "Arabic - Sudan", decimalPoint: '.', thousandSep: ','},
	"ar_TN": {name: "ar_TN", description: "Arabic - Tunisia", decimalPoint: '.', thousandSep: ','},
	"ar_YE": {name: "ar_YE", description: "Arabic - Yemen", decimalPoint: '.', thousandSep: ','},
	"de_BE": {name: "de_BE", description: "German - Belgium", decimalPoint: ',', thousandSep: '.'},
	"de_CH": {name: "de_CH", description: "German - Switzerland", decimalPoint: '.', thousandSep: '\''},
	"de_LU": {name: "de_LU", description: "German - Luxembourg", decimalPoint: ',', thousandSep: '.'},
	"en_AU": {name: "en_AU", description: "English - Australia", decimalPoint: '.', thousandSep: ','},
	"en_CA": {name: "en_CA", description: "English - Canada", decimalPoint: '.', thousandSep: ','},
	"en_IN": {name: "en_IN", description: "English - India", decimalPoint: '.', thousandSep: ','},
	"en_NZ": {name: "en_NZ", description: "English - New Zealand", decimalPoint: '.', thousandSep: ','},
	"en_PH": {name: "en_PH", description: "English - Philippines", decimalPoint: '.', thousandSep: ','},
	"en_ZA": {name: "en_ZA", description: "English - South Africa", decimalPoint: '.', thousandSep: ','},
	"en_ZW": {name: "en_ZW", description: "English - Zimbabwe", decimalPoint: '.', thousandSep: ','},
	"es_AR": {name: "es_AR", description: "Spanish - Argentina", decimalPoint: ',', thousandSep: '.'},
	"es_BO": {name: "es_BO", description: "Spanish - Bolivia", decimalPoint: ',', thousandSep: '.'},
	"es_CL": {name: "es_CL", description: "Spanish - Chile", decimalPoint: ',', thousandSep: '.'},
	"es_CO": {name: "es_CO", description: "Spanish - Columbia", decimalPoint: ',', thousandSep: '.'},
	"es_CR": {name: "es_CR", description: "Spanish - Costa Rica", decimalPoint: ',', thousandSep: ' '},
	"es_DO": {name: "es_DO", description: "Spanish - Dominican Republic", decimalPoint: '.', thousandSep: ','},
	"es_EC": {name: "es_EC", description: "Spanish - Ecuador", decimalPoint: ',', thousandSep: '.'},
	"es_GT": {name: "es_GT", description: "Spanish - Guatemala", decimalPoint: '.', thousandSep: ','},
	"es_HN": {name: "es_HN", description: "Spanish - Honduras", decimalPoint: '.', thousandSep: ','},
	"es_MX": {name: "es_MX", description: "Spanish - Mexico", decimalPoint: '.', thousandSep: ','},
	"es_NI": {name: "es_NI", description: "Spanish - Nicaragua", decimalPoint: '.', thousandSep: ','},
	"es_PA": {name: "es_PA", description: "Spanish - Panama", decimalPoint: '.', thousandSep: ','},
	"es_PE": {name: "es_PE", description: "Spanish - Peru", decimalPoint: '.', thousandSep: ','},
	"es_PR": {name: "es_PR", description: "Spanish - Puerto Rico", decimalPoint: '.', thousandSep: ','},
	"es_PY": {name: "es_PY", description: "Spanish - Paraguay", decimalPoint: ',', thousandSep: '.'},
	"es_SV": {name: "es_SV", description: "Spanish - El Salvador", decimalPoint: '.', thousandSep: ','},
	"es_US": {name: "es_US", description: "Spanish - United States", decimalPoint: '.', thousandSep: ','},
	"es_UY": {name: "es_UY", description: "Spanish - Uruguay", decimalPoint: ',', thousandSep: '.'},
	"es_VE": {name: "es_VE", description: "Spanish - Venezuela", decimalPoint: ',', thousandSep: '.'},
	"fr_BE": {name: "fr_BE", description: "French - Belgium", decimalPoint: ',', thousandSep: '.'},
	"fr_CA": {name: "fr_CA", description: "French - Canada", decimalPoint: ',', thousandSep: ' '},
	"fr_CH": {name: "fr_CH", description: "French - Switzerland", decimalPoint: ',', thousandSep: byte(0)},
	"fr_LU": {name: "fr_LU", description: "French - Luxembourg", decimalPoint: ',', thousandSep: byte(0)},
	"it_IT": {name: "it_IT", description: "Italian - Italy", decimalPoint: ',', thousandSep: byte(0)},
	"nl_BE": {name: "nl_BE", description: "Dutch - Belgium", decimalPoint: ',', thousandSep: '.'},
	"no_NO": {name: "no_NO", description: "Norwegian - Norway", decimalPoint: ',', thousandSep: '.'},
	"sv_FI": {name: "sv_FI", description: "Swedish - Finland", decimalPoint: ',', thousandSep: ' '},
	"zh_HK": {name: "zh_HK", description: "Chinese - Hong Kong SAR", decimalPoint: '.', thousandSep: ','},
	"el_GR": {name: "el_GR", description: "Greek - Greece", decimalPoint: ',', thousandSep: '.'},
	"rm_CH": {name: "rm_CH", description: "Romansh - Switzerland", decimalPoint: ',', thousandSep: '\''},
}

func TestFormatFloat(t *testing.T) {
	fmt.Println(numFormat("3232378998456.97355433256377889"))

}

func numFormat(str string) string {
	numStr := strings.Split(str, ".")[0] //如果有小数获取整数部分
	length := len(numStr)
	if length < 4 {
		return str
	}
	count := (length - 1) / 3
	for i := 0; i < count; i++ {
		numStr = numStr[:length-(i+1)*3] + "," + numStr[length-(i+1)*3:]
	}
	return numStr
}

func TestFormat2(t *testing.T) {
}

//func Format(str string, frac int, locale string) string {
//	// locale
//	format :=
//	integ := strings.Split(str, ".")[0]
//	length := len(integ)
//	if length < 4 {
//		return str
//	}
//	ret := []string{}
//	length := len(integ)
//	for i := 0; i < length; i++ {
//		ret = append(ret, string(integ[i]))
//	}
//	strings.Join(ret, ",")
//}
