package main

import "testing"

func TestLights(t *testing.T) {
	tests := map[string]int {
		"turn on 0,0 through 999,999": 1000000,
		"toggle 0,0 through 999,0": 1000000-1000,
		"turn off 499,499 through 500,500": 1000000-1000-4,
	}

	g := NewGrid(1000, 1000)

	for tc, want := range tests {
		t.Run(tc, func (t *testing.T) {
			g.Process(tc)
			got := g.NumOn()
			if got != want {
				t.Errorf("after %q -> got: %v, want %v", tc, got, want)
			}
		})
	}
}
