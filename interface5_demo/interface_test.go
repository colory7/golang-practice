package interface5_demo

import "testing"

func TestInterface(t *testing.T) {
	var ff FlyFish
	ff = Fish{}
	ff.fly()
	ff.swim()

}
