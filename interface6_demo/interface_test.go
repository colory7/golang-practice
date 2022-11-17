package interface6_demo

import "testing"

func TestInterface(t *testing.T) {
	dog := Dog{}
	cat := Cat{}
	per := Person{}

	per.care(dog)
	per.care(cat)

}
