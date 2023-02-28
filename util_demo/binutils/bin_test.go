package binutils

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"testing"
)

func BinarySum(a, b string) string {
	index_a := len(a) - 1
	index_b := len(b) - 1
	jinwei := 0
	result := ""
	for index_a >= 0 && index_b >= 0 {
		ia := a[index_a] - '0' // -'0' 可以得到对应数字
		ib := b[index_b] - '0'
		sum := int(ia) + int(ib) + jinwei

		if sum >= 2 {
			jinwei = 1
		} else {
			jinwei = 0
		}
		res := sum%2 + '0'
		result = fmt.Sprintf("%c%s", res, result)
		index_a--
		index_b--
	}

	for index_a >= 0 {
		ia := a[index_a] - '0'
		sum := int(ia) + jinwei
		if sum >= 2 {
			jinwei = 1
		} else {
			jinwei = 0
		}
		res := sum%2 + '0'
		result = fmt.Sprintf("%c%s", res, result)
		index_a--
	}
	for index_b >= 0 {
		ib := b[index_b] - '0'
		sum := int(ib) + jinwei
		if sum >= 2 {
			jinwei = 1
		} else {
			jinwei = 0
		}
		res := sum%2 + '0'
		result = fmt.Sprintf("%c%s", res, result)
		index_b--
	}
	if jinwei == 1 {
		result = fmt.Sprintf("1%s", result)
	}
	return result
}

func BinaryAdd(a, b string) string {
	ai, _ := new(big.Int).SetString(a, 2)
	bo, _ := new(big.Int).SetString(b, 2)
	return ai.Add(ai, bo).Text(2)
}

func TestBinaryAdd(t *testing.T) {
	fmt.Println(BinarySum("1", "0"))
	fmt.Println(BinaryAdd("1", "0"))

	fmt.Println(BinarySum("1", "-1"))
	fmt.Println(BinaryAdd("1", "-1"))

	fmt.Println(BinarySum("12", "-1"))
	fmt.Println(BinaryAdd("12", "-1"))
}

func TestBinaryAdd2(t *testing.T) {
	fmt.Println(BinarySum("12", "-1"))
	fmt.Println(BinaryAdd("1", "-1"))
}

func TestBinary3(t *testing.T) {
	fmt.Println(Complement(3))
	fmt.Println(Complement(-3))
	fmt.Println(Complement(-9876448))
	// 1111111111111111111111111111111111111111011010010100110000100000
}

func ReverseBits(num uint32) uint32 {
	var rev uint32 = 0
	for i := 0; i < 32; i++ {
		rev <<= 1
		if num%2 == 1 {
			rev += 1
		}
		num >>= 1
	}
	return rev
}

func Test(t *testing.T) {
	initial := new(big.Int)
	initial.SetString("-42", 10)

	value, _ := new(big.Int).SetString("-42", 10)

	var result [2]int64

	result[0] = value.Int64()
	result[1] = value.Rsh(value, 64).Int64()

	leRepresentation := make([]byte, 16)

	binary.LittleEndian.PutUint64(leRepresentation[:8], uint64(result[0]))
	binary.LittleEndian.PutUint64(leRepresentation[8:], uint64(result[1]))

	fmt.Println(leRepresentation)

	fmt.Println(result)

	reverse := big.NewInt(result[1])
	reverse.Lsh(reverse, 64)
	reverse.Add(reverse, big.NewInt(result[0]))

	fmt.Println(reverse.String())

	fmt.Println(initial.String() == reverse.String())
}
