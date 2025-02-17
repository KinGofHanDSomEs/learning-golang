package main

import (
	"testing"
)

func TestGetUTFLength(t *testing.T) {
	cases := []struct {
		name   string
		values []byte
		want   int
		err    error
	}{
		{
			name:   "test1",
			values: []byte("Hello"),
			want:   5,
			err:    nil,
		},
		{
			name:   "test2",
			values: []byte("Тест123"),
			want:   7,
			err:    nil,
		},
		{
			name:   "test3",
			values: []byte{0xff},
			want:   0,
			err:    ErrInvalidUTF8,
		},
	}
	for _, test := range cases {
		length, err := GetUTFLength(test.values)
		if length != test.want {
			t.Errorf("For input %v, expected %d, but got %d", test.values, test.want, length)
		}
		if err != test.err {
			t.Errorf("For input %v, expected error %v, but got %v", test.values, test.err, err)
		}
	}
}
