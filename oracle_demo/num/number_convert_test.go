package num

import (
	"fmt"
	"testing"
)

// https://baike.baidu.com/item/%E5%9F%BA%E6%95%B0%E8%AF%8D
var ordinal_num map[int]string

// https://baike.baidu.com/item/%E5%BA%8F%E6%95%B0%E8%AF%8D/3811128
var cardinal_num map[int]string

func init() {

	cardinal_num = map[int]string{
		1:             "one",
		2:             "two",
		3:             "three",
		4:             "four",
		5:             "five",
		6:             "six",
		7:             "seven",
		8:             "eight",
		9:             "nine",
		10:            "ten",
		11:            "eleven",
		12:            "twelve",
		13:            "thirteen",
		14:            "fourteen",
		15:            "fifteen",
		16:            "sixteen",
		17:            "seventeen",
		18:            "eighteen",
		19:            "nineteen",
		20:            "twenty",
		30:            "thirty",
		40:            "forty",
		50:            "fifty",
		60:            "sixty",
		70:            "seventy",
		80:            "eighty",
		90:            "ninety",
		100:           "one hundred",
		101:           "a hundred and one",
		110:           "a hundred and ten",
		120:           "a hundred and twenty",
		200:           "two hundred",
		1000:          "one thousand",
		1001:          "a thousand and one",
		1010:          "a thousand and ten",
		100000:        "one hundred thousand",
		1000000:       "one million",
		2000000:       "two million",
		1000000000:    "billion",
		1000000000000: "trillion",
	}

}

func TestCardinalNum2OrdinalNum(t *testing.T) {
	fmt.Println(ordinalNum2Word(342578))
	fmt.Println(ordinalNum2Word(27))
}

func ordinalNum2Word(num int) string {
	tens := num / 10
	ones := num % 10

	return cardinal_num[tens*10] + "-" + ordinal_num[ones]
}

func TestOrdinalNum2CardinalNum(t *testing.T) {

}
