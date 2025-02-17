package main

import "fmt"

type Logger interface {
	Log(message string)
}

type LogLevel string

const (
	Error LogLevel = "Error"
	Info  LogLevel = "Info"
)

type Log struct {
	Level LogLevel
}

func (l Log) Log(message string) {
	switch l.Level {
	case Error:
		fmt.Printf("ERROR: %s", message)
	case Info:
		fmt.Printf("INFO: %s", message)
	}
}
