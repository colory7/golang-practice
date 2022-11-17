package string_demo

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestStringConcat(t *testing.T) {
	s1 := "hello"
	s2 := "word"
	s3 := s1 + s2
	fmt.Print(s3) //s3 = "helloword"
}

func TestStringConcat2(t *testing.T) {
	s1 := "hello"
	s2 := "word"
	s3 := fmt.Sprintf("%s%s", s1, s2) //s3 = "helloword"
	fmt.Print(s3)                     //s3 = "helloword"
}

func TestStringConcat3(t *testing.T) {
	s1 := "hello"
	s2 := "word"
	var str []string = []string{s1, s2}
	s3 := strings.Join(str, "")
	fmt.Print(s3)
}

func TestStringConcat4(t *testing.T) {
	s1 := "hello"
	s2 := "word"
	var bt bytes.Buffer
	bt.WriteString(s1)
	bt.WriteString(s2)
	s3 := bt.String()
	fmt.Println(s3)

}

func TestStringConcat5(t *testing.T) {
	s1 := "hello"
	s2 := "word"
	var build strings.Builder
	build.WriteString(s1)
	build.WriteString(s2)
	s3 := build.String()
	fmt.Println(s3)

}
