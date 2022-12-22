package main

import (
	"fmt"
	"sync"
)

var num int
var wg sync.WaitGroup
var lock sync.Mutex // 声明一个互斥锁变量

func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock() // 加锁，当前只有该协程执行
		num -= 1
		lock.Unlock() // 解锁
	}
}

func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		lock.Lock() // 读数据时加锁，因为此时可能还会出现资源竞争
		num += 1
		lock.Unlock()
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
