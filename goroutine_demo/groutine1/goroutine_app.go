package main

import (
	"fmt"
	"time"
)

func printGoroutine() {
	// 协程执行函数
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Printf("Hello goroutine %v \n", i)
	}

}

func main() {
	// 主线程
	// 开启一个协程
	go printGoroutine()

	// 继续执行主协程代码
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second / 2)
		fmt.Printf("Hello main %v \n", i)
	}

}
