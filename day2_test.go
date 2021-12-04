package aoc_2021

import (
	"testing"
)

func TestDay2_FirstHalf(t *testing.T) {
	tests := []struct {
		inputFile string
		want      int
	}{
		{"testdata/day_2_example", 150},
		{"testdata/day_2_input", 2027977},
	}
	for _, tt := range tests {
		t.Run(tt.inputFile, func(t *testing.T) {
			lines := ReadLines(tt.inputFile)
			if got := move(lines); got != tt.want {
				t.Errorf("expected result to be %d but it was %d", tt.want, got)
			}
		})
	}
}

func TestDay2_SecondHalf(t *testing.T) {
	tests := []struct {
		inputFile string
		want      int
	}{
		{"testdata/day_2_example", 900},
		{"testdata/day_2_input", 1903644897},
	}
	for _, tt := range tests {
		t.Run(tt.inputFile, func(t *testing.T) {
			lines := ReadLines(tt.inputFile)
			if got := moveWithAim(lines); got != tt.want {
				t.Errorf("expected result to be %d but it was %d", tt.want, got)
			}
		})
	}
}
