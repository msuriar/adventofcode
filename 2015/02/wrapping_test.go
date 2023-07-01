package main

import "testing"

func TestPaper(t *testing.T) {
	tests := map[string]int {
		"2x3x4": 52,
		"1x1x10": 42,
	}

	for tc, want := range tests {
		got := paper(tc)
		if got != want {
			t.Errorf("paper(%q) : got %d ; want %d\n", tc, got, want)
		}
	}
}
