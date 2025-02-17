package main

import (
	"strings"
)

type UpperWriter struct {
	UpperString string
}

type Writer interface {
	Write(p []byte) (int, error)
}

func (uw *UpperWriter) Write(p []byte) (int, error) {
	uw.UpperString += strings.ToUpper(string(p))
	return len(p), nil
}
