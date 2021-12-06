package aoc_2021

import (
	"testing"
)

func TestDay6_FirstHalf(t *testing.T) {
	tests := []struct {
		inputFile string
		want      int
	}{
		{"testdata/day_6_example", 5934},
		{"testdata/day_6_input", 383160},
	}
	for _, tt := range tests {
		t.Run(tt.inputFile, func(t *testing.T) {
			lines := ReadLines(tt.inputFile)
			if got := countFishAfter(lines, 80); got != tt.want {
				t.Errorf("expected result to be %d but it was %d", tt.want, got)
			}
		})
	}
}

func TestDay6_SecondHalf(t *testing.T) {
	tests := []struct {
		inputFile string
		want      int
	}{
		{"testdata/day_6_example", 26984457539},
		{"testdata/day_6_input", 1721148811504},
	}
	for _, tt := range tests {
		t.Run(tt.inputFile, func(t *testing.T) {
			lines := ReadLines(tt.inputFile)
			if got := countFishAfter(lines, 256); got != tt.want {
				t.Errorf("expected result to be %d but it was %d", tt.want, got)
			}
		})
	}
}
