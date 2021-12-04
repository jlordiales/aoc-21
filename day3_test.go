package aoc_2021

import (
	"testing"
)

func TestDay3_FirstHalf(t *testing.T) {
	tests := []struct {
		inputFile string
		want      int64
	}{
		{"testdata/day_3_example", 198},
		{"testdata/day_3_input", 4160394},
	}
	for _, tt := range tests {
		t.Run(tt.inputFile, func(t *testing.T) {
			lines := ReadLines(tt.inputFile)
			if got := powerConsumption(lines); got != tt.want {
				t.Errorf("expected result to be %d but it was %d", tt.want, got)
			}
		})
	}
}

func TestDay3_SecondHalf(t *testing.T) {
	tests := []struct {
		inputFile string
		want      int64
	}{
		{"testdata/day_3_example", 230},
		{"testdata/day_3_input", 4125600},
	}
	for _, tt := range tests {
		t.Run(tt.inputFile, func(t *testing.T) {
			lines := ReadLines(tt.inputFile)
			if got := lifeSupportRating(lines); got != tt.want {
				t.Errorf("expected result to be %d but it was %d", tt.want, got)
			}
		})
	}
}
