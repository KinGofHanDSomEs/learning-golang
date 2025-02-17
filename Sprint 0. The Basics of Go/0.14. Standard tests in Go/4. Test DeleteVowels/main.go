package main

import (
	"testing"
)

func TestDeleteVowels(t *testing.T) {
	cases := []struct {
		name  string
		value string
		want  string
	}{
		{
			name:  "test1",
			value: "hello",
			want:  "hll",
		},
		{
			name:  "test2",
			value: "HellO",
			want:  "Hll",
		},
		{
			name:  "test3",
			value: "hll",
			want:  "hll",
		},
		{
			name:  "test4",
			value: "AIeUo",
			want:  "",
		},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			got := DeleteVowels(test.value)
			if got != test.want {
				t.Errorf("DeleteVowels(%s) = %s; want %s", test.value, got, test.want)
			}
		})
	}
}
