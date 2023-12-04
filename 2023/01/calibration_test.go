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
		t.Run(tc, func(t *testing.T) {
			got := calibration_value(tc)
			if got != want {
				t.Errorf("calibration_value(%q): got %d ; want %d\n", tc, got, want)
			}
		})
	}
}

func TestTextCalibration(t *testing.T) {
	tests := map[string]int{
		"two1nine":         29,
		"eightwothree":     83,
		"abcone2threexyz":  13,
		"xtwone3four":      24,
		"4nineeightseven2": 42,
		"zoneight234":      14,
		"7pqrstsixteen":    76,
	}

	for tc, want := range tests {
		t.Run(tc, func(t *testing.T) {
			got := calibration_value(tc)
			if got != want {
				t.Errorf("calibration_value(%q): got %d ; want %d\n", tc, got, want)
			}
		})
	}
}
