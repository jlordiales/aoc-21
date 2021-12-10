package aoc_2021

import (
	"testing"
)

func TestDay10_FirstHalf(t *testing.T) {
	tests := []struct {
		inputFile string
		want      int
	}{
		{"testdata/day_10_example", 26397},
		{"testdata/day_10_input", 345441},
	}
	for _, tt := range tests {
		t.Run(tt.inputFile, func(t *testing.T) {
			input := ReadLines(tt.inputFile)
			if got := day10FirstHalf(input); got != tt.want {
				t.Errorf("expected result to be %d but it was %d", tt.want, got)
			}
		})
	}
}

func TestDay10_SecondHalf(t *testing.T) {
	tests := []struct {
		inputFile string
		want      int
	}{
		{"testdata/day_10_example", 288957},
		{"testdata/day_10_input", 3235371166},
	}
	for _, tt := range tests {
		t.Run(tt.inputFile, func(t *testing.T) {
			input := ReadLines(tt.inputFile)
			if got := day10SecondHalf(input); got != tt.want {
				t.Errorf("expected result to be %d but it was %d", tt.want, got)
			}
		})
	}
}
