package mysql_func_demo

import (
	"fmt"
	"testing"
	"time"
)

func TestToDays(tx *testing.T) {
	t1, _ := time.Parse(time.RFC3339, "2023-02-20T22:08:41+00:00")
	fmt.Println(t1.Unix() / 86400)

	t1, _ = time.Parse(time.RFC3339, "0000-01-01T00:00:00+00:00")
	// -62167219200
	fmt.Println(t1.Unix())
	// -719528
	fmt.Println(t1.Unix() / 86400)

	t1, _ = time.Parse(time.RFC3339, "1970-01-02T00:00:00+00:00")
	// 86340
	fmt.Println(t1.Unix())
	// 1
	fmt.Println(t1.Unix() / 86400)

	t1, _ = time.Parse(time.RFC3339, "2023-02-21T00:00:00+00:00")
	fmt.Println(t1.Unix())
	fmt.Println((t1.Unix() / 86400) + 719528)
	fmt.Println((t1.Unix() + 62167219200) / 86400)

}
