package interface4_demo

import "testing"

func TestInterface(t *testing.T) {
	m := Mobile{}
	m.playMusic()
	m.playVideo()
}
