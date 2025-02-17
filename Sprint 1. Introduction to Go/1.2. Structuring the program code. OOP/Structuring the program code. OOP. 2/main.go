package main

type User struct {
	Name   string
	Age    int
	Active bool
}

func NewUser(name string, age ...int) *User {
	userAge := 18
	if len(age) > 0 {
		userAge = age[0]
	}
	return &User{Name: name, Age: userAge, Active: true}
}
