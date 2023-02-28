package mysql_func_demo

import (
	"fmt"
	"testing"
	"time"
)

func TestToSeconds(tx *testing.T) {
	t1, _ := time.Parse(time.RFC3339, "2023-02-21T00:00:00+00:00")
	fmt.Println(t1.Unix())
	fmt.Println(t1.Unix() + 62167219200)
}
