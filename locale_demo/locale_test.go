package locale_demo

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"testing"
)

func TestLocale(t *testing.T) {
	p := message.NewPrinter(language.French)
	p.Printf("There are %v flowers in our garden.\n", 15346.2335)
	p.Printf("There are %1.1f flowers in our garden.\n", 15346.2335)

	language.Make("el")
	language.Parse("en-UK")

	ja, _ := language.ParseBase("ja")
	jp, _ := language.ParseRegion("JP")
	jpLngTag, _ := language.Compose(ja, jp)
	fmt.Println(jpLngTag) // prints ja-JP

	fmt.Println(language.Compose(language.ParseRegion("AL"))) // prints Und-AL

}
