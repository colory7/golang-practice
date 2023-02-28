package locale_demo

import (
	"golang.org/x/text/message"
	"testing"
)

func TestLocale2(t *testing.T) {
	p := message.NewPrinter(message.MatchLanguage("en"))
	p.Println(2323123456.7328) // Prints 123,456.78

	p.Println("%d", 2323123456.7328) // Prints 4,331 ducks in a row

	p = message.NewPrinter(message.MatchLanguage("nl"))
	p.Println("%.1f", 234956.7328) // Prints Hoogte: 1,244.9 meter

}
