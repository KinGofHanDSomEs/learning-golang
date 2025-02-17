package main

func Receive(ch chan int) int {
	num := <-ch
	return num
}
