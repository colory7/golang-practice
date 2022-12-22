package main

import (
	"errors"
	"log"
	"time"
)

func main() {
	test1()
	time.Sleep(time.Second * 3)
	log.Println("123")
}
func test1() {
	log.Println("test1 start")
	test2()
	log.Println("test1 end")
}
func test2() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("test2 err:", err)
		}
	}()
	log.Println("test2 start")
	panic(errors.New("stop test2"))
	log.Println("test2 end")
}
