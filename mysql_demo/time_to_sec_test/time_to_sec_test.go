package mysql_func_demo

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeToSeconds(tx *testing.T) {
	t1, _ := time.Parse(time.RFC3339, "2023-02-21T19:30:10+00:00")
	fmt.Println(t1.Hour()*3600 + t1.Minute()*60 + t1.Second())

}
