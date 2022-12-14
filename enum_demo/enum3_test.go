package enum_demo

import (
	"fmt"
	"testing"
)

type FormatType int

const (
	FTByte FormatType = iota
	FTArray
	FTDefine
)

func (ft FormatType) String() string {
	switch ft {
	case FTByte:
		return "byte"
	case FTArray:
		return "array"
	case FTDefine:
		return "define"
	}
	return ""
}

func TestEnum3(t *testing.T) {
	fmt.Println(int(FTArray))
	fmt.Println(FTArray.String())
}
