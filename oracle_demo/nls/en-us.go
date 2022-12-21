package ntw

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

var englishMegas = []string{"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion", "sextillion", "septillion", "octillion", "nonillion", "decillion", "undecillion", "duodecillion", "tredecillion", "quattuordecillion"}
var englishUnits = []string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var englishTens = []string{"", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
var englishTeens = []string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

func NumToOrdinalWord(input int) string {
	return integerToEnUs(input, ORDINAL)
}

func NumToCardinalWord(input int) string {
	return integerToEnUs(input, CARDINAL)
}

func NumToWithOrdinalSuffix(input int) string {
	if input == 0 {
		return "0th"
	}

	words := bytes.Buffer{}
	words.WriteString(strconv.Itoa(input))

	remainder := input % 10
	switch remainder {
	case 1:
		words.WriteString("st")
	case 2:
		words.WriteString("nd")
	case 3:
		words.WriteString("rd")
	default:
		words.WriteString("th")
	}
	return words.String()
}

// integerToEnUs converts an integer to American English words
func integerToEnUs(input int, f flag) string {

	//log.Printf("Input: %d\n", input)
	words := []string{}

	if input < 0 {
		words = append(words, "minus")
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
