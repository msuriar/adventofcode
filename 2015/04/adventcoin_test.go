package main

import "testing"

func TestAdventCoin(t *testing.T) {
	tests := map[string]int{
		"abcdef": 609043,
		"pqrstuv": 1048970,
	}

	for tc, want := range tests {
		t.Run(tc, func (t *testing.T) {
			got := FindHashSuffix(tc, 5)
			if got != want {
				t.Errorf("FindHashSuffix(%s, 5) => got %d ; want %d\n", tc, got, want)
			}
		})
	}
}
