package main

import "fmt"

func PrettyArrayOutput(array [9]string) {
	for i := 0; i < len(array)-2; i++ {
        fmt.Println(i+1, "я уже сделал:", array[i])
	}
	for i := 7; i < len(array); i++ {
        fmt.Println(i+1, "не успел сделать:", array[i])
	}
}