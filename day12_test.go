package aoc_2021

import (
	"testing"
)

func TestDay12_FirstHalf(t *testing.T) {
	tests := []struct {
		inputFile string
		want      int
	}{
		{"testdata/day_12_example", 10},
		{"testdata/day_12_input", 4773},
	}
	for _, tt := range tests {
		t.Run(tt.inputFile, func(t *testing.T) {
			input := ReadLines(tt.inputFile)
			if got := day12FirstHalf(input); got != tt.want {
				t.Errorf("expected result to be %d but it was %d", tt.want, got)
			}
		})
	}
}

func TestDay12_SecondHalf(t *testing.T) {
	tests := []struct {
		inputFile string
		want      int
	}{
		{"testdata/day_12_example", 36},
		{"testdata/day_12_input", 116985},
	}
	for _, tt := range tests {
		t.Run(tt.inputFile, func(t *testing.T) {
			input := ReadLines(tt.inputFile)
			if got := day12SecondHalf(input); got != tt.want {
				t.Errorf("expected result to be %d but it was %d", tt.want, got)
			}
		})
	}
}
