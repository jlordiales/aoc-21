package aoc_2021

import (
	"testing"
)

func TestDay1_FirstHalf(t *testing.T) {
	tests := []struct {
		inputFile string
		want      int
	}{
		{"testdata/day_1_example", 7},
		{"testdata/day_1_input", 1451},
	}
	for _, tt := range tests {
		t.Run(tt.inputFile, func(t *testing.T) {
			lines := ReadLines(tt.inputFile)
			if got := countIncreasesWithSlidingWindow(lines, 1); got != tt.want {
				t.Errorf("expected result to be %d but it was %d", tt.want, got)
			}
		})
	}
}

func TestDay1_SecondHalf(t *testing.T) {
	tests := []struct {
		inputFile string
		want      int
	}{
		{"testdata/day_1_example", 5},
		{"testdata/day_1_input", 1395},
	}
	for _, tt := range tests {
		t.Run(tt.inputFile, func(t *testing.T) {
			lines := ReadLines(tt.inputFile)
			if got := countIncreasesWithSlidingWindow(lines, 3); got != tt.want {
				t.Errorf("expected result to be %d but it was %d", tt.want, got)
			}
		})
	}
}
