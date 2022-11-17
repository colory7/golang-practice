package interface_demo

import "testing"

func TestInterface(t *testing.T) {
	c := Computer{}
	m := Mobile{}

	c.read()
	c.write()
	m.read()
	m.write()
}
