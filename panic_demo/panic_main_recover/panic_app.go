package main

import (
	"errors"
	"log"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("err:", err)
		}
	}()
	go test1()
	time.Sleep(time.Second * 3)
	panic(errors.New("stop test1"))
	log.Println("123")
	select {}
}

func test1() {
	for {
		tm := time.NewTicker(time.Second)
		select {
		case <-tm.C:
			log.Println("test1")
		}
	}
}
