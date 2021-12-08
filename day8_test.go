package aoc_2021

import (
	"testing"
)

func TestDay8_FirstHalf(t *testing.T) {
	tests := []struct {
		inputFile string
		want      int
	}{
		{"testdata/day_8_example", 26},
		{"testdata/day_8_input", 476},
	}
	for _, tt := range tests {
		t.Run(tt.inputFile, func(t *testing.T) {
			input := ReadLines(tt.inputFile)
			if got := day8FirstHalf(input); got != tt.want {
				t.Errorf("expected result to be %d but it was %d", tt.want, got)
			}
		})
	}
}

func TestDay8_SecondHalf(t *testing.T) {
	tests := []struct {
		inputFile string
		want      int
	}{
		{"testdata/day_8_example", 61229},
		{"testdata/day_8_input", 1011823},
	}
	for _, tt := range tests {
		t.Run(tt.inputFile, func(t *testing.T) {
			input := ReadLines(tt.inputFile)
			if got := day8SecondHalf(input); got != tt.want {
				t.Errorf("expected result to be %d but it was %d", tt.want, got)
			}
		})
	}
}
