package main

import "fmt"

type hello interface{
}

type hello1 struct{
	age int
}

func main() {
	var a hello = hello1{age: 1}
	fmt.Print(a)
}