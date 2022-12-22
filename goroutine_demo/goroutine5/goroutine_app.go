package main

import (
	"fmt"
	"sync"
	"time"
)

var mapNum = make(map[int]int, 50)
var wg sync.WaitGroup
var rwlock sync.RWMutex // 声明一个读写锁变量

func write() {
	defer wg.Done()
	rwlock.Lock()
	fmt.Printf("写入数据\n")
	time.Sleep(time.Second * 2)
	fmt.Printf("写入结束\n")
	rwlock.Unlock()
}

func read() {
	defer wg.Done()
	rwlock.RLock()
	fmt.Printf("读取数据\n")
	time.Sleep(time.Second)
	fmt.Printf("读取结束\n")
	rwlock.RUnlock()
}

func main() {
	// 启动10个读，10个写的协程
	wg.Add(22)

	for i := 0; i < 20; i++ {
		go write()
	}

	for i := 0; i < 2; i++ {
		go read()
	}

	// 此时等待20个协程执行完毕
	wg.Wait()

}
