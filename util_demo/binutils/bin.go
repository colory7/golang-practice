package binutils

import (
	"bytes"
	"errors"
	"strconv"
)

func Complement(num int64) (string, error) {
	if num > 0 {
		return "", errors.New("not support")
	}

	bStr := strconv.FormatInt(-num, 2)
	//b1, _ := new(big.Int).SetString(bStr, 2)
	//b2, _ := new(big.Int).SetString("-1", 2)
	//br := b1.Add(b1, b2).Text(2)

	br := AddBinary(bStr, "1111111111111111111111111111111111111111111111111111111111111111")

	bf := bytes.Buffer{}
	l := len(br)
	for i := 64; i >= 0; i-- {
		if i == l {
			break
		} else {
			bf.WriteByte('1')
		}
	}
	for i := l - 1; i >= 0; i-- {
		if br[i] == '0' {
			bf.WriteByte('1')
		} else {
			bf.WriteByte('0')
		}
	}
	return bf.String(), nil
}

func AddBinary(a string, b string) string {
	ans := ""
	carry := 0
	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)

	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
