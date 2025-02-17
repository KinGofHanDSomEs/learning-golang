package main

import "sync"

func Write(num int) {
	mutex.Lock()
	Buf = append(Buf, num)
	mutex.Unlock()
}

func Consume() int {
	mutex.Lock()
	res := Buf[0]
	Buf = Buf[1:]
	mutex.Unlock()
	return res
}

var (
	Buf   = []int{}
	mutex sync.Mutex
)