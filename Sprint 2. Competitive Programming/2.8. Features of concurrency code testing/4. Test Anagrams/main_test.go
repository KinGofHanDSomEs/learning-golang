package main

import "testing"

func TestAreAnagrams(t *testing.T) {
	cases := []struct {
		str1 string
		str2 string
		want bool
	}{
		{"listen", "silent", true},
		{"Triangle", "Integral", true},
		{"Apple", "Pabble", false},
		{"Hello", "World", false},
		{"Dormitory", "Dirty Room", false},
		{"", "", true},
		{"Listen", "Silent", true},
		{"Abc", "abc", true},
		{"A", "a", true},
		{"123", "321", true},
		{"a1b2c", "2b1ca", true},
		{"a b c", "c b a", true},
		{"!@#$%^", "%^$#@!", true},
		{"", "a", false},
		{"ab", "", false},
	}
	for _, cs := range cases {
		t.Run("test", func(t *testing.T) {
			got := AreAnagrams(cs.str1, cs.str2)
			if cs.want != got {
				t.Fatalf("expected: %v, got: %v", cs.want, got)
			}
		})
	}
}
