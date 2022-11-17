package interface_demo

import "fmt"

type Mobile struct {
}

func (c Mobile) read() {
	fmt.Println("mobile read...")
}

func (c Mobile) write() {
	fmt.Println("mobile write...")
}
