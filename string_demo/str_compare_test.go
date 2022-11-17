package string_demo

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func TestStr(t *testing.T) {

	fmt.Println("go" == "go")
	fmt.Println("GO" == "go")

	fmt.Println(strings.Compare("GO", "go"))
	fmt.Println(strings.Compare("go", "go"))

	fmt.Println(strings.EqualFold("GO", "go"))
}

// operator compare
func compareOperatorsTxt(a string) bool {
	file, err := os.Open("names.txt")
	result := false
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if strings.ToLower(a) == strings.ToLower(scanner.Text()) {
			result = true
		} else {
			result = false
		}
	}
	file.Close()
	return result
}

// strings compare
func compareStringTxt(a string) bool {
	file, err := os.Open("names.txt")
	result := false
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if strings.Compare(strings.ToLower(a), strings.ToLower(scanner.Text())) == 0 {
			result = true
		} else {
			result = false
		}
	}
	file.Close()
	return result
}

// EqualFold compare
func compareEFTxt(a string) bool {
	file, err := os.Open("names.txt")
	result := false
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if strings.EqualFold(a, scanner.Text()) {
			result = true
		} else {
			result = false
		}
	}
	file.Close()
	return result
}

//字符串长度相同而字符不同
func compareByCountTxt(a string) bool {
	file, err := os.Open("names.txt")
	result := false
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if len(a) == len(scanner.Text()) && strings.EqualFold(a, scanner.Text()) {
			result = true
		} else {
			result = false
		}
	}
	file.Close()
	return result
}
