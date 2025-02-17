package main

import (
	"io"
)

func ReadString(r io.Reader) (string, error) {
	res, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
