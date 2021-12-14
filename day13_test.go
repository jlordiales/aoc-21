package aoc_2021

import (
	"io/ioutil"
	"testing"
)

func TestDay13_FirstHalf(t *testing.T) {
	tests := []struct {
		inputFile string
		want      int
	}{
		{"testdata/day_13_example", 17},
		{"testdata/day_13_input", 621},
	}
	for _, tt := range tests {
		t.Run(tt.inputFile, func(t *testing.T) {
			input := ReadLines(tt.inputFile)
			if got := day13FirstHalf(input); got != tt.want {
				t.Errorf("expected result to be %d but it was %d", tt.want, got)
			}
		})
	}
}

func TestDay13_SecondHalf(t *testing.T) {
	tests := []struct {
		inputFile string
	}{
		{inputFile: "testdata/day_13_example"},
		{"testdata/day_13_input"},
	}
	for _, tt := range tests {
		t.Run(tt.inputFile, func(t *testing.T) {
			input := ReadLines(tt.inputFile)

			want, _ := ioutil.ReadFile(tt.inputFile + "_output")
			got := day13SecondHalf(input)
			wantStr := string(want)

			if got != wantStr {
				t.Errorf("expected result to be \n%s but it was \n%s", wantStr, got)
			}
		})
	}
}
