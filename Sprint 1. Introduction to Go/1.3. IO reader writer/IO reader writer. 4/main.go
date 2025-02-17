package main

import (
	"io"
)

func Copy(r io.Reader, w io.Writer, n uint) error {
	data := make([]byte, n)
	bytesRead, err := r.Read(data)
	if err != nil {
		return err
	}
	_, err = w.Write(data[:bytesRead])
	if err != nil {
		return err
	}
	return nil
}
