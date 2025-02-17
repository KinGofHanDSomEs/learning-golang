package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

func (p Person) Print() {
	fmt.Printf("Name: %s\nAge: %d\nAddress: %s\n", p.name, p.age, p.address)
}
