package main

import (
	"context"
	"errors"
	"io"
)

func Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error) {
	if len(seq) == 0 {
		return false, errors.New("empty sequence")
	}

	data := make([]byte, len(seq))
	for {
		select {
		case <-ctx.Done():
			return false, ctx.Err()
		default:
			buf := make([]byte, 1)
			_, err := r.Read(buf)
			if err != nil {
				if err == io.EOF {
					return false, nil
				}
				return false, err
			}
			if len(data) == len(seq) {
				data = data[1:]
			}
			data = append(data, buf[0])
			f := true
			for i := 0; i < len(data); i++ {
				if data[i] != seq[i] {
					f = false
					break
				}
			}
			if f {
				return true, nil
			}
		}
	}
	return false, nil
}
