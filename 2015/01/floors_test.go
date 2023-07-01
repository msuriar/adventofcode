package main

import "testing"

func TestFloors(t *testing.T) {
	tests := map[string]int{
		"(())": 0,
		"()()": 0,
		"(((": 3,
		"(()(()(": 3,
		"))(((((": 3,
		"())": -1,
		"))(": -1,
		")))": -3,
		")())())": -3,
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

func TestBasement(t *testing.T) {
	tests := map[string]int {
		")": 1,
		"()())": 5,
	}

	for tc, want := range tests {
		got := basement(tc)
		if got != want {
			t.Errorf("basement(%q) : got %d ; want %d\n", tc, got, want)
		}
	}
}
