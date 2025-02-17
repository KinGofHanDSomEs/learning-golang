package main

import (
	"bytes"
	"fmt"
	"io"
)

func Contains(r io.Reader, seq []byte) (bool, error) {
	if len(seq) == 0 {
		return false, fmt.Errorf("sequence cannot be empty")
	}
	buffer := make([]byte, len(seq))
	i := 0
	for {
		n, err := r.Read(buffer[i : i+1])
		if err != nil {
			if err == io.EOF {
				return false, nil
			}
			return false, err
		}
		i = (i + n) % len(seq)
		if bytes.Equal(buffer[i:], seq[:len(buffer)-i]) && bytes.Equal(buffer[:i], seq[len(buffer)-i:]) {
			return true, nil
		}
	}
}