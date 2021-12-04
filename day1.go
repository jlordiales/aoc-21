package aoc_2021

import (
	"strconv"
)

func countIncreasesWithSlidingWindow(lines []string, windowSize int) int {
	if len(lines) <= windowSize {
		return 0
	}

	var values []int
	for i := 0; i <= len(lines)-windowSize; i++ {
		window := lines[i : i+windowSize]
		values = append(values, sumValues(window))
	}

	var result int
	for i := 1; i < len(values); i++ {
		if values[i] > values[i-1] {
			result++
		}
	}

	return result
}

func sumValues(values []string) int {
	var result int
	for _, v := range values {
		intValue, _ := strconv.Atoi(v)
		result += intValue
	}
	return result
}
