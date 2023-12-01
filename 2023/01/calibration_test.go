package main

import "testing"

func TestCalibrationValues(t *testing.T) {
	tests := map[string]int{
		"1abc2":       12,
		"pqr3stu8vwx": 38,
		"a1b2c3d4e5f": 15,
		"treb7uchet":  77,
	}

	for tc, want := range tests {
		t.Run(tc, func (t *testing.T) {
			got := calibration_value(tc)
			if got != want {
				t.Errorf("calibration_value(%q): got %d ; want %d\n", tc, got, want)
			}
		})
	}
}
