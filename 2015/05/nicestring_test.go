package main

import "testing"

func TestNice1(t *testing.T) {
	tests := map[string]bool{
    "ugknbfddgicrmopn": true,
    "aaa": true,
    "jchzalrnumimnmhp": false,
    "haegwjzuvuyypxyu": false,
    "dvszwmarrgswjxmb": false,
	}

	for tc, want := range tests {
		t.Run(tc, func (t *testing.T) {
			got := IsNice1(tc)
			if got != want {
				t.Errorf("Nice1(%q) -> got: %v, want %v", tc, got, want)
			}
		})
	}
}

func TestNice2(t *testing.T) {
	tests := map[string]bool{
    "qjhvhtzxzqqjkmpb": true,
    "xxyxx": true,
    "uurcxstgmygtbstg": false,
    "ieodomkazucvgmuy": false,
	}

	for tc, want := range tests {
		t.Run(tc, func (t *testing.T) {
			got := IsNice2(tc)
			if got != want {
				t.Errorf("Nice1(%q) -> got: %v, want %v", tc, got, want)
			}
		})
	}
}
