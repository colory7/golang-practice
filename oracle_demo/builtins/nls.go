package builtins

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type flag int

const (
	CARDINAL flag = 0
	ORDINAL  flag = 1
)

var englishMegas = []string{"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion", "sextillion", "septillion", "octillion", "nonillion", "decillion", "undecillion", "duodecillion", "tredecillion", "quattuordecillion"}
var englishUnits = []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var englishTens = []string{"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
var englishTeens = []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

var ordinalNums = map[string]string{
	"one":           "first",
	"two":           "second",
	"three":         "third",
	"five":          "fifth",
	"eight":         "eighth",
	"nine":          "ninth",
	"twelve":        "twelfth",
	"twenty":        "twentieth",
	"twenty-one":    "twenty-first",
	"twenty-two":    "twenty-second",
	"twenty-three":  "twenty-third",
	"twenty-five":   "twenty-fifth",
	"twenty-eight":  "twenty-eighth",
	"twenty-nine":   "twenty-ninth",
	"thirty":        "thirtieth",
	"thirty-one":    "thirty-first",
	"thirty-two":    "thirty-second",
	"thirty-three":  "thirty-third",
	"thirty-five":   "thirty-fifth",
	"thirty-eight":  "thirty-eighth",
	"thirty-nine":   "thirty-ninth",
	"forty":         "fortieth",
	"forty-one":     "forty-first",
	"forty-two":     "forty-second",
	"forty-three":   "forty-third",
	"forty-five":    "forty-fifth",
	"forty-eight":   "forty-eighth",
	"forty-nine":    "forty-ninth",
	"fifty":         "fiftieth",
	"fifty-one":     "fifty-first",
	"fifty-two":     "fifty-second",
	"fifty-three":   "fifty-third",
	"fifty-five":    "fifty-fifth",
	"fifty-eight":   "fifty-eighth",
	"fifty-nine":    "fifty-ninth",
	"sixty":         "sixtieth",
	"sixty-one":     "sixty-first",
	"sixty-two":     "sixty-second",
	"sixty-three":   "sixty-third",
	"sixty-five":    "sixty-fifth",
	"sixty-eight":   "sixty-eighth",
	"sixty-nine":    "sixty-ninth",
	"seventy":       "seventieth",
	"seventy-one":   "seventy-first",
	"seventy-two":   "seventy-second",
	"seventy-three": "seventy-third",
	"seventy-five":  "seventy-fifth",
	"seventy-eight": "seventy-eighth",
	"seventy-nine":  "seventy-ninth",
	"eighty":        "eightieth",
	"eighty-one":    "eighty-first",
	"eighty-two":    "eighty-second",
	"eighty-three":  "eighty-third",
	"eighty-five":   "eighty-fifth",
	"eighty-eight":  "eighty-eighth",
	"eighty-nine":   "eighty-ninth",
	"ninety":        "ninetieth",
	"ninety-one":    "ninety-first",
	"ninety-two":    "ninety-second",
	"ninety-three":  "ninety-third",
	"ninety-five":   "ninety-fifth",
	"ninety-eight":  "ninety-eighth",
	"ninety-nine":   "ninety-ninth",
}

func NumToOrdinalWord(input int) string {
	return integerToEnUs(input, ORDINAL)
}

func NumToCardinalWord(input int) string {
	return integerToEnUs(input, CARDINAL)
}

func NumToWithOrdinalSuffix(input int) string {
	words := bytes.Buffer{}
	words.WriteString(strconv.Itoa(input))
	words.WriteString(GetOrdinalSuffix(input))
	return words.String()
}

func GetOrdinalSuffix(input int) string {
	remainder := input % 10
	var suffix string
	switch remainder {
	case 1:
		suffix = "st"
	case 2:
		suffix = "nd"
	case 3:
		suffix = "rd"
	default:
		suffix = "th"
	}
	return suffix
}

// integerToEnUs converts an integer to American English words
func integerToEnUs(input int, f flag) string {

	//log.Printf("Input: %d\n", input)
	words := []string{}

	if input < 0 {
		words = append(words, "signMinus")
		input *= -1
	}

	// split integer in triplets
	triplets := integerToTriplets(input)
	//log.Printf("Triplets: %v\n", triplets)

	// zero is a special case
	if len(triplets) == 0 {
		if f == ORDINAL {
			return "zeroth"
		} else {
			return "zero"
		}
	}

	// iterate over triplets
	for idx := len(triplets) - 1; idx >= 0; idx-- {
		triplet := triplets[idx]
		//log.Printf("Triplet: %d (idx=%d)\n", triplet, idx)

		// nothing todo for empty triplet
		if triplet == 0 {
			continue
		}

		// three-digits
		hundreds := triplet / 100 % 10
		tens := triplet / 10 % 10
		units := triplet % 10
		//log.Printf("Hundreds:%d, Tens:%d, Units:%d\n", hundreds, tens, units)
		if hundreds > 0 {
			words = append(words, englishUnits[hundreds], "hundred")
		}

		if tens == 0 && units == 0 {
			goto tripletEnd
		}

		switch tens {
		case 0:
			words = append(words, englishUnits[units])
		case 1:
			words = append(words, englishTeens[units])
		default:
			if units > 0 {
				word := fmt.Sprintf("%s-%s", englishTens[tens], englishUnits[units])
				words = append(words, word)
			} else {
				words = append(words, englishTens[tens])
			}
			break
		}

	tripletEnd:
		// mega
		if mega := englishMegas[idx]; mega != "" {
			words = append(words, mega)
		}
	}

	if f == ORDINAL {
		li := len(words) - 1
		lastWord := words[li]
		ordinalWord := ordinalNums[lastWord]
		if ordinalWord != "" {
			words[li] = ordinalWord
		} else {
			words[li] = lastWord + "th"
		}
	}

	//log.Printf("Words length: %d\n", len(words))
	return strings.Join(words, " ")
}

func integerToTriplets(number int) []int {
	triplets := []int{}

	for number > 0 {
		triplets = append(triplets, number%1000)
		number = number / 1000
	}

	return triplets
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

func ToRRRR(thisYear int, RR int) int {
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
