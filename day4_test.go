package aoc_2021

import (
	"testing"
)

func TestDay4_FirstHalf(t *testing.T) {
	tests := []struct {
		inputFile string
		want      int
	}{
		{"testdata/day_4_example", 4512},
		{"testdata/day_4_input", 38913},
	}
	for _, tt := range tests {
		t.Run(tt.inputFile, func(t *testing.T) {
			lines := ReadLines(tt.inputFile)
			if got := bingoToWin(lines); got != tt.want {
				t.Errorf("expected result to be %d but it was %d", tt.want, got)
			}
		})
	}
}

func TestDay4_SecondHalf(t *testing.T) {
	tests := []struct {
		inputFile string
		want      int
	}{
		{"testdata/day_4_example", 1924},
		{"testdata/day_4_input", 16836},
	}
	for _, tt := range tests {
		t.Run(tt.inputFile, func(t *testing.T) {
			lines := ReadLines(tt.inputFile)
			if got := bingoToLoose(lines); got != tt.want {
				t.Errorf("expected result to be %d but it was %d", tt.want, got)
			}
		})
	}
}
