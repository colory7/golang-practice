package switch_demo

import (
	"fmt"
	"testing"
)

func TestSwitch2(t *testing.T) {
	c := 'e'
	switch c {
	case '0':
	case '1':
	case '2':
	case '3':
	case '4':
	case '5':
	case '6':
	case '7':
	case '8':
	case '9':
		break
	case '.':
		break
	case 'e':
	case 'E':
		fmt.Println("222")
		break
	case ' ':
		break
	}
}

func TestSwitchFallthrough(t *testing.T) {
	a := 100
	switch a {
	case 100:
		fmt.Println("100")
		fallthrough
	case 200:
		fmt.Println("200")
	case 300:
		fmt.Println("300")
	default:
		fmt.Println("other")
	}
}

func TestSwitchFallthrough2(t *testing.T) {
	i := 0
	switch i {
	case 0:
		fmt.Println("0")
		fallthrough
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	case 3:
		fmt.Println("3")
	case 4, 5, 6:
		fmt.Println("4, 5, 6")
	default:
		fmt.Println("Default")
	}

}
