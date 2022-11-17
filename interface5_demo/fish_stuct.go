package interface5_demo

import "fmt"

type Fish struct {
}

func (fish Fish) fly() {
	fmt.Println("fly...")
}

func (fish Fish) swim() {
	fmt.Println("swim...")
}
