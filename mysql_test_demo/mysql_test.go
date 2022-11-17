package mysql

import (
	"fmt"
	"testing"
)

func Test_findByPk(t *testing.T) {
	var aa int
	aa = 2
	fmt.Println(aa)
	num := findByPk(1)
	t.Log(num)
}
