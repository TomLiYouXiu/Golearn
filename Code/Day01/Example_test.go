package main

import "testing"

func TestlengthOfNonRepeatingSubStr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"abcabcbb", 0},
		{"bbbbbb", 0},
		{"pwwkew", 0},
	}
	for _, tt := range tests {
		if actual := lengthOfNonRepeatingSubStr(tt.s); actual != tt.ans {
			t.Errorf("Got %d for input %s;"+
				"expected %d", actual, tt.s, tt.ans)
		}
	}
}
