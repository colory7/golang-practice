package main

import (
	"fmt"
	"sync"
)

var num int
var wg sync.WaitGroup

func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		num -= 1
	}
}

func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		num += 1
	}
}

func main() {
	wg.Add(2)

	// 启动两个协程，分别是加1和减1两个函数
	go add()
	go sub()
	// 此时等待两个协程执行完毕
	wg.Wait()
	// 接着执行主协程的逻辑
	fmt.Println(num)

}
