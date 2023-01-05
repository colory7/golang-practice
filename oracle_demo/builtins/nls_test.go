package builtins

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	fmt.Println(integerToEnUs(4236, CARDINAL))
	fmt.Println(integerToEnUs(1111111, CARDINAL))
	fmt.Println(integerToEnUs(111111120, CARDINAL))
	fmt.Println(integerToEnUs(111111121, CARDINAL))
	fmt.Println(integerToEnUs(111111130, CARDINAL))
	fmt.Println(integerToEnUs(111111131, CARDINAL))
	fmt.Println(integerToEnUs(111111100, CARDINAL))
	fmt.Println(integerToEnUs(111000000, CARDINAL))
	fmt.Println("====")
	fmt.Println(NumToOrdinalWord(4236))
	fmt.Println(NumToOrdinalWord(1111111))
	fmt.Println(NumToOrdinalWord(111111120))
	fmt.Println(NumToOrdinalWord(111111121))
	fmt.Println(NumToOrdinalWord(111111130))
	fmt.Println(NumToOrdinalWord(111111131))
	fmt.Println(NumToOrdinalWord(111111100))
	fmt.Println(NumToOrdinalWord(111000000))
	fmt.Println("====")
	fmt.Println(NumToCardinalWord(4236))
	fmt.Println(NumToCardinalWord(1111111))
	fmt.Println(NumToCardinalWord(111111120))
	fmt.Println(NumToCardinalWord(111111121))
	fmt.Println(NumToCardinalWord(111111130))
	fmt.Println(NumToCardinalWord(111111131))
	fmt.Println(NumToCardinalWord(111111100))
	fmt.Println(NumToCardinalWord(111000000))
}

func ExampleIntegerToEnUs() {
	fmt.Println(integerToEnUs(42, CARDINAL))
	// Output: forty-two
}

func Test232(t *testing.T) {
	n := NumToCardinalWord(1)
	fmt.Println(n)
	m := NumToWithOrdinalSuffix(1)
	fmt.Println(m)

	n = NumToCardinalWord(23)
	fmt.Println(n)
	m = NumToWithOrdinalSuffix(23)
	fmt.Println(m)

	n = NumToCardinalWord(123)
	fmt.Println(n)
	m = NumToWithOrdinalSuffix(123)
	fmt.Println(m)

	n = NumToCardinalWord(1123)
	fmt.Println(n)
	m = NumToWithOrdinalSuffix(1123)
	fmt.Println(m)

	n = NumToCardinalWord(11123)
	fmt.Println(n)
	m = NumToWithOrdinalSuffix(11123)
	fmt.Println(m)

	n = NumToCardinalWord(111123)
	fmt.Println(n)
	m = NumToWithOrdinalSuffix(111123)
	fmt.Println(m)

	n = NumToCardinalWord(1111123)
	fmt.Println(n)
	m = NumToWithOrdinalSuffix(1111123)
	fmt.Println(m)

	n = NumToCardinalWord(11111123)
	fmt.Println(n)
	m = NumToWithOrdinalSuffix(11111123)
	fmt.Println(m)

	n = NumToCardinalWord(11111111111111123)
	fmt.Println(n)
	m = NumToWithOrdinalSuffix(11111111111111123)
	fmt.Println(m)

	n = NumToCardinalWord(111111111111111123)
	fmt.Println(n)
	m = NumToWithOrdinalSuffix(111111111111111123)
	fmt.Println(m)

	n = NumToCardinalWord(1111111111111111123)
	fmt.Println(n)
	m = NumToWithOrdinalSuffix(1111111111111111123)
	fmt.Println(m)
}

func TestIntegerToEnUs(t *testing.T) {
	t.Parallel()

	tests := map[int]string{
		-1:            "minus one",
		0:             "zero",
		1:             "one",
		9:             "nine",
		10:            "ten",
		11:            "eleven",
		19:            "nineteen",
		20:            "twenty",
		21:            "twenty-one",
		80:            "eighty",
		90:            "ninety",
		99:            "ninety-nine",
		100:           "one hundred",
		101:           "one hundred one",
		111:           "one hundred eleven",
		120:           "one hundred twenty",
		121:           "one hundred twenty-one",
		900:           "nine hundred",
		909:           "nine hundred nine",
		919:           "nine hundred nineteen",
		990:           "nine hundred ninety",
		999:           "nine hundred ninety-nine",
		1000:          "one thousand",
		2000:          "two thousand",
		4000:          "four thousand",
		5000:          "five thousand",
		11000:         "eleven thousand",
		21000:         "twenty-one thousand",
		999000:        "nine hundred ninety-nine thousand",
		999999:        "nine hundred ninety-nine thousand nine hundred ninety-nine",
		1000000:       "one million",
		2000000:       "two million",
		4000000:       "four million",
		5000000:       "five million",
		100100100:     "one hundred million one hundred thousand one hundred",
		500500500:     "five hundred million five hundred thousand five hundred",
		606606606:     "six hundred six million six hundred six thousand six hundred six",
		999000000:     "nine hundred ninety-nine million",
		999000999:     "nine hundred ninety-nine million nine hundred ninety-nine",
		999999000:     "nine hundred ninety-nine million nine hundred ninety-nine thousand",
		999999999:     "nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
		1174315110:    "one billion one hundred seventy-four million three hundred fifteen thousand one hundred ten",
		1174315119:    "one billion one hundred seventy-four million three hundred fifteen thousand one hundred nineteen",
		15174315119:   "fifteen billion one hundred seventy-four million three hundred fifteen thousand one hundred nineteen",
		35174315119:   "thirty-five billion one hundred seventy-four million three hundred fifteen thousand one hundred nineteen",
		935174315119:  "nine hundred thirty-five billion one hundred seventy-four million three hundred fifteen thousand one hundred nineteen",
		2935174315119: "two trillion nine hundred thirty-five billion one hundred seventy-four million three hundred fifteen thousand one hundred nineteen",
	}

	for input, expectedOutput := range tests {
		name := fmt.Sprintf("%d", input)
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, expectedOutput, integerToEnUs(input, CARDINAL))
		})
	}
}
