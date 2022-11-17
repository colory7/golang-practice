package time_demo

import (
	"fmt"
	"github.com/cockroachdb/apd/v3"
	"testing"
)

func TestTimeFormat(t *testing.T) {
	d1 := apd.Decimal{}
	d1.Coeff.SetString("20210101160700333666999", 10)

	fmt.Println(d1.Text('f'))
	fmt.Println(d1.String())
}
