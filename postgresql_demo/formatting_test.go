package postgresql_demo

//
//// var DCH_keywords map[string]KeyWord
//var DCH_keywords []KeyWord
//var NUM_keywords []NumKeyWord
//
///* ----------
// * Full months
// * ----------
// */
//
//var months_full []string
//var days_short []string
//
//var rm_months_upper []string
//var rm_months_lower []string
//var rm1 []string
//var rm10 []string
//var rm100 []string
//
//var numTH []string
//var numth []string
//
//var adbc_strings []string
//var adbc_strings_long []string
//
//var ampm_strings []string
//var ampm_strings_long []string
//
//func init() {
//	months_full = []string{
//		"January",
//		"February",
//		"March",
//		"April",
//		"May",
//		"June",
//		"July",
//		"August",
//		"September",
//		"October",
//		"November",
//		"December",
//	}
//
//	days_short = []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
//
//	rm_months_upper = []string{"XII", "XI", "X", "IX", "VIII", "VII", "VI", "V", "IV", "III", "II", "I"}
//	rm_months_lower = []string{"xii", "xi", "x", "ix", "viii", "vii", "vi", "v", "iv", "iii", "ii", "i"}
//
//	/* ----------
//	* Roman numbers
//	* ----------
//	 */
//	rm1 = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
//	rm10 = []string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
//	rm100 = []string{"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
//
//	numTH = []string{"ST", "ND", "RD", "TH"}
//	numth = []string{"st", "nd", "rd", "th"}
//
//	DCH_keywords = []KeyWord{
//		/*	name, len, id, is_digit, date_mode */
//		{"A.D.", 4, DCH_A_D, false, FROM_CHAR_DATE_NONE}, /* A */
//		{"A.M.", 4, DCH_A_M, false, FROM_CHAR_DATE_NONE},
//		{"AD", 2, DCH_AD, false, FROM_CHAR_DATE_NONE},
//		{"AM", 2, DCH_AM, false, FROM_CHAR_DATE_NONE},
//		{"B.C.", 4, DCH_B_C, false, FROM_CHAR_DATE_NONE}, /* B */
//		{"BC", 2, DCH_BC, false, FROM_CHAR_DATE_NONE},
//		{"CC", 2, DCH_CC, true, FROM_CHAR_DATE_NONE},    /* C */
//		{"DAY", 3, DCH_DAY, false, FROM_CHAR_DATE_NONE}, /* D */
//		{"DDD", 3, DCH_DDD, true, FROM_CHAR_DATE_GREGORIAN},
//		{"DD", 2, DCH_DD, true, FROM_CHAR_DATE_GREGORIAN},
//		{"DY", 2, DCH_DY, false, FROM_CHAR_DATE_NONE},
//		{"D", 1, DCH_D, true, FROM_CHAR_DATE_GREGORIAN},
//		{"FF1", 3, DCH_FF1, false, FROM_CHAR_DATE_NONE}, /* F */
//		{"FF2", 3, DCH_FF2, false, FROM_CHAR_DATE_NONE},
//		{"FF3", 3, DCH_FF3, false, FROM_CHAR_DATE_NONE},
//		{"FF4", 3, DCH_FF4, false, FROM_CHAR_DATE_NONE},
//		{"FF5", 3, DCH_FF5, false, FROM_CHAR_DATE_NONE},
//		{"FF6", 3, DCH_FF6, false, FROM_CHAR_DATE_NONE},
//		{"FX", 2, DCH_FX, false, FROM_CHAR_DATE_NONE},
//		{"HH24", 4, DCH_HH24, true, FROM_CHAR_DATE_NONE}, /* H */
//		{"HH12", 4, DCH_HH12, true, FROM_CHAR_DATE_NONE},
//		{"HH", 2, DCH_HH, true, FROM_CHAR_DATE_NONE},
//		{"IDDD", 4, DCH_IDDD, true, FROM_CHAR_DATE_ISOWEEK}, /* I */
//		{"ID", 2, DCH_ID, true, FROM_CHAR_DATE_ISOWEEK},
//		{"IW", 2, DCH_IW, true, FROM_CHAR_DATE_ISOWEEK},
//		{"IYYY", 4, DCH_IYYY, true, FROM_CHAR_DATE_ISOWEEK},
//		{"IYY", 3, DCH_IYY, true, FROM_CHAR_DATE_ISOWEEK},
//		{"IY", 2, DCH_IY, true, FROM_CHAR_DATE_ISOWEEK},
//		{"I", 1, DCH_I, true, FROM_CHAR_DATE_ISOWEEK},
//		{"J", 1, DCH_J, true, FROM_CHAR_DATE_NONE},   /* J */
//		{"MI", 2, DCH_MI, true, FROM_CHAR_DATE_NONE}, /* M */
//		{"MM", 2, DCH_MM, true, FROM_CHAR_DATE_GREGORIAN},
//		{"MONTH", 5, DCH_MONTH, false, FROM_CHAR_DATE_GREGORIAN},
//		{"MON", 3, DCH_MON, false, FROM_CHAR_DATE_GREGORIAN},
//		{"MS", 2, DCH_MS, true, FROM_CHAR_DATE_NONE},
//		{"OF", 2, DCH_OF, false, FROM_CHAR_DATE_NONE},    /* O */
//		{"P.M.", 4, DCH_P_M, false, FROM_CHAR_DATE_NONE}, /* P */
//		{"PM", 2, DCH_PM, false, FROM_CHAR_DATE_NONE},
//		{"Q", 1, DCH_Q, true, FROM_CHAR_DATE_NONE},         /* Q */
//		{"RM", 2, DCH_RM, false, FROM_CHAR_DATE_GREGORIAN}, /* R */
//		{"SSSSS", 5, DCH_SSSS, true, FROM_CHAR_DATE_NONE},  /* S */
//		{"SSSS", 4, DCH_SSSS, true, FROM_CHAR_DATE_NONE},
//		{"SS", 2, DCH_SS, true, FROM_CHAR_DATE_NONE},
//		{"TZH", 3, DCH_TZH, false, FROM_CHAR_DATE_NONE}, /* T */
//		{"TZM", 3, DCH_TZM, true, FROM_CHAR_DATE_NONE},
//		{"TZ", 2, DCH_TZ, false, FROM_CHAR_DATE_NONE},
//		{"US", 2, DCH_US, true, FROM_CHAR_DATE_NONE},      /* U */
//		{"WW", 2, DCH_WW, true, FROM_CHAR_DATE_GREGORIAN}, /* W */
//		{"W", 1, DCH_W, true, FROM_CHAR_DATE_GREGORIAN},
//		{"Y,YYY", 5, DCH_Y_YYY, true, FROM_CHAR_DATE_GREGORIAN}, /* Y */
//		{"YYYY", 4, DCH_YYYY, true, FROM_CHAR_DATE_GREGORIAN},
//		{"YYY", 3, DCH_YYY, true, FROM_CHAR_DATE_GREGORIAN},
//		{"YY", 2, DCH_YY, true, FROM_CHAR_DATE_GREGORIAN},
//		{"Y", 1, DCH_Y, true, FROM_CHAR_DATE_GREGORIAN},
//		{"a.d.", 4, DCH_a_d, false, FROM_CHAR_DATE_NONE}, /* a */
//		{"a.m.", 4, DCH_a_m, false, FROM_CHAR_DATE_NONE},
//		{"ad", 2, DCH_ad, false, FROM_CHAR_DATE_NONE},
//		{"am", 2, DCH_am, false, FROM_CHAR_DATE_NONE},
//		{"b.c.", 4, DCH_b_c, false, FROM_CHAR_DATE_NONE}, /* b */
//		{"bc", 2, DCH_bc, false, FROM_CHAR_DATE_NONE},
//		{"cc", 2, DCH_CC, true, FROM_CHAR_DATE_NONE},    /* c */
//		{"day", 3, DCH_day, false, FROM_CHAR_DATE_NONE}, /* d */
//		{"ddd", 3, DCH_DDD, true, FROM_CHAR_DATE_GREGORIAN},
//		{"dd", 2, DCH_DD, true, FROM_CHAR_DATE_GREGORIAN},
//		{"dy", 2, DCH_dy, false, FROM_CHAR_DATE_NONE},
//		{"d", 1, DCH_D, true, FROM_CHAR_DATE_GREGORIAN},
//		{"ff1", 3, DCH_FF1, false, FROM_CHAR_DATE_NONE}, /* f */
//		{"ff2", 3, DCH_FF2, false, FROM_CHAR_DATE_NONE},
//		{"ff3", 3, DCH_FF3, false, FROM_CHAR_DATE_NONE},
//		{"ff4", 3, DCH_FF4, false, FROM_CHAR_DATE_NONE},
//		{"ff5", 3, DCH_FF5, false, FROM_CHAR_DATE_NONE},
//		{"ff6", 3, DCH_FF6, false, FROM_CHAR_DATE_NONE},
//		{"fx", 2, DCH_FX, false, FROM_CHAR_DATE_NONE},
//		{"hh24", 4, DCH_HH24, true, FROM_CHAR_DATE_NONE}, /* h */
//		{"hh12", 4, DCH_HH12, true, FROM_CHAR_DATE_NONE},
//		{"hh", 2, DCH_HH, true, FROM_CHAR_DATE_NONE},
//		{"iddd", 4, DCH_IDDD, true, FROM_CHAR_DATE_ISOWEEK}, /* i */
//		{"id", 2, DCH_ID, true, FROM_CHAR_DATE_ISOWEEK},
//		{"iw", 2, DCH_IW, true, FROM_CHAR_DATE_ISOWEEK},
//		{"iyyy", 4, DCH_IYYY, true, FROM_CHAR_DATE_ISOWEEK},
//		{"iyy", 3, DCH_IYY, true, FROM_CHAR_DATE_ISOWEEK},
//		{"iy", 2, DCH_IY, true, FROM_CHAR_DATE_ISOWEEK},
//		{"i", 1, DCH_I, true, FROM_CHAR_DATE_ISOWEEK},
//		{"j", 1, DCH_J, true, FROM_CHAR_DATE_NONE},   /* j */
//		{"mi", 2, DCH_MI, true, FROM_CHAR_DATE_NONE}, /* m */
//		{"mm", 2, DCH_MM, true, FROM_CHAR_DATE_GREGORIAN},
//		{"month", 5, DCH_month, false, FROM_CHAR_DATE_GREGORIAN},
//		{"mon", 3, DCH_mon, false, FROM_CHAR_DATE_GREGORIAN},
//		{"ms", 2, DCH_MS, true, FROM_CHAR_DATE_NONE},
//		{"p.m.", 4, DCH_p_m, false, FROM_CHAR_DATE_NONE}, /* p */
//		{"pm", 2, DCH_pm, false, FROM_CHAR_DATE_NONE},
//		{"q", 1, DCH_Q, true, FROM_CHAR_DATE_NONE},         /* q */
//		{"rm", 2, DCH_rm, false, FROM_CHAR_DATE_GREGORIAN}, /* r */
//		{"sssss", 5, DCH_SSSS, true, FROM_CHAR_DATE_NONE},  /* s */
//		{"ssss", 4, DCH_SSSS, true, FROM_CHAR_DATE_NONE},
//		{"ss", 2, DCH_SS, true, FROM_CHAR_DATE_NONE},
//		{"tz", 2, DCH_tz, false, FROM_CHAR_DATE_NONE},     /* t */
//		{"us", 2, DCH_US, true, FROM_CHAR_DATE_NONE},      /* u */
//		{"ww", 2, DCH_WW, true, FROM_CHAR_DATE_GREGORIAN}, /* w */
//		{"w", 1, DCH_W, true, FROM_CHAR_DATE_GREGORIAN},
//		{"y,yyy", 5, DCH_Y_YYY, true, FROM_CHAR_DATE_GREGORIAN}, /* y */
//		{"yyyy", 4, DCH_YYYY, true, FROM_CHAR_DATE_GREGORIAN},
//		{"yyy", 3, DCH_YYY, true, FROM_CHAR_DATE_GREGORIAN},
//		{"yy", 2, DCH_YY, true, FROM_CHAR_DATE_GREGORIAN},
//		{"y", 1, DCH_Y, true, FROM_CHAR_DATE_GREGORIAN},
//	}
//
//	NUM_keywords = []NumKeyWord{
//		/*	name, len, id			is in Index */
//		{",", 1, NUM_COMMA}, /* , */
//		{".", 1, NUM_DEC},   /* . */
//		{"0", 1, NUM_0},     /* 0 */
//		{"9", 1, NUM_9},     /* 9 */
//		{"B", 1, NUM_B},     /* B */
//		{"C", 1, NUM_C},     /* C */
//		{"D", 1, NUM_D},     /* D */
//		{"EEEE", 4, NUM_E},  /* E */
//		{"FM", 2, NUM_FM},   /* F */
//		{"G", 1, NUM_G},     /* G */
//		{"L", 1, NUM_L},     /* L */
//		{"MI", 2, NUM_MI},   /* M */
//		{"PL", 2, NUM_PL},   /* P */
//		{"PR", 2, NUM_PR},
//		{"RN", 2, NUM_RN}, /* R */
//		{"SG", 2, NUM_SG}, /* S */
//		{"SP", 2, NUM_SP},
//		{"S", 1, NUM_S},
//		{"TH", 2, NUM_TH},  /* T */
//		{"V", 1, NUM_V},    /* V */
//		{"b", 1, NUM_B},    /* b */
//		{"c", 1, NUM_C},    /* c */
//		{"d", 1, NUM_D},    /* d */
//		{"eeee", 4, NUM_E}, /* e */
//		{"fm", 2, NUM_FM},  /* f */
//		{"g", 1, NUM_G},    /* g */
//		{"l", 1, NUM_L},    /* l */
//		{"mi", 2, NUM_MI},  /* m */
//		{"pl", 2, NUM_PL},  /* p */
//		{"pr", 2, NUM_PR},
//		{"rn", 2, NUM_rn}, /* r */
//		{"sg", 2, NUM_SG}, /* s */
//		{"sp", 2, NUM_SP},
//		{"s", 1, NUM_S},
//		{"th", 2, NUM_th}, /* t */
//		{"v", 1, NUM_V},   /* v */
//	}
//
//	adbc_strings = []string{ad_STR, bc_STR, AD_STR, BC_STR}
//	adbc_strings_long = []string{a_d_STR, b_c_STR, A_D_STR, B_C_STR}
//
//	ampm_strings = []string{am_STR, pm_STR, AM_STR, PM_STR}
//	ampm_strings_long = []string{a_m_STR, p_m_STR, A_M_STR, P_M_STR}
//}
