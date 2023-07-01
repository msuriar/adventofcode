package main

import "testing"

func TestFloors(t *testing.T) {
	tests := map[string]int{
		"(())": 0,
		"(((": 3,
	}

	for tc, want := range tests {
		t.Run(tc, func (t *testing.T) {
			got := floor(tc)
			if got != want {
				t.Errorf("floor(%q) : got %d ; want %d\n", tc, got, want)
			}
		})
	}
}
