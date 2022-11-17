package interface4_demo

import "fmt"

type Mobile struct {
}

func (m Mobile) playMusic() {
	fmt.Println("播放音乐")
}

func (m Mobile) playVideo() {
	fmt.Println("播放视频")
}
