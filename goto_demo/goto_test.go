package goto_demo

import "testing"

func TestGoto(t *testing.T) {
	var i int
	for {
		println(i)
		i++
		if i > 2 {
			goto BREAK

		}
	}

BREAK:
	println("break")
}

func TestGoto2(t *testing.T) {
BREAK:
	println("break")

	var i int
	for {
		println(i)
		i++
		if i > 2 {
			goto BREAK

		}
	}
}
