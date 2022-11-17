package interface_demo

import "fmt"

type Computer struct {
}

func (c Computer) read() {
	fmt.Println("computer read...")
}

func (c Computer) write() {
	fmt.Println("computer write...")
}
