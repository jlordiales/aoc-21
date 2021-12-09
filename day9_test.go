package aoc_2021

import (
	"testing"
)

func TestDay9_FirstHalf(t *testing.T) {
	tests := []struct {
		inputFile string
		want      int
	}{
		{"testdata/day_9_example", 15},
		{"testdata/day_9_input", 516},
	}
	for _, tt := range tests {
		t.Run(tt.inputFile, func(t *testing.T) {
			input := ReadLines(tt.inputFile)
			if got := day9FirstHalf(input); got != tt.want {
				t.Errorf("expected result to be %d but it was %d", tt.want, got)
			}
		})
	}
}

func TestDay9_SecondHalf(t *testing.T) {
	tests := []struct {
		inputFile string
		want      int
	}{
		{"testdata/day_9_example", 1134},
		{"testdata/day_9_input", 1023660},
	}
	for _, tt := range tests {
		t.Run(tt.inputFile, func(t *testing.T) {
			input := ReadLines(tt.inputFile)
			if got := day9SecondHalf(input); got != tt.want {
				t.Errorf("expected result to be %d but it was %d", tt.want, got)
			}
		})
	}
}
