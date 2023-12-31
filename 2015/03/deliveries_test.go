package main

import (
	"testing"
)

func TestSoloDeliveries(t *testing.T) {
	tests := map[string]int{
		">":          2,
		"^>v<":       4,
		"^v^v^v^v^v": 2,
	}

	for tc, want := range tests {
		t.Run(tc, func(t *testing.T) {
			c := NewCourse()
			c.Navigate(tc)
			got := c.UniqueHouses()

			if got != want {
				t.Errorf("UniqueHouses(%q) : got %d ; want %d\n", tc, got, want)
			}
		})
	}
}

func TestPartnerDeliveries(t *testing.T) {
	tests := map[string]int{
    "^v": 3,
    "^>v<": 3,
    "^v^v^v^v^v": 11,
	}

	for tc, want := range tests {
		t.Run(tc, func(t *testing.T) {
			c := NewCourse()
			c.PartnerNavigate(tc)
			got := c.UniqueHouses()

			if got != want {
				t.Errorf("UniqueHouses(%q) : got %d ; want %d\n", tc, got, want)
			}
		})
	}
}
