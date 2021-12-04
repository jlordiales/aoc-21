package aoc_2021

import (
	"strconv"
	"strings"
)

func powerConsumption(input []string) int64 {
	if len(input) == 0 {
		return 0
	}

	// assume all lines in the input have the same length
	numberOfBits := len(input[0])

	bitsByPosition := make([]string, numberOfBits)
	for _, line := range input {
		for i, c := range line {
			bitsByPosition[i] += string(c)
		}
	}

	gammaRateBits := ""
	epsilonRateBits := ""
	for _, s := range bitsByPosition {
		oneBits := strings.Count(s, "1")
		zeroBits := strings.Count(s, "0")

		if oneBits > zeroBits {
			gammaRateBits += "1"
			epsilonRateBits += "0"
		} else {
			gammaRateBits += "0"
			epsilonRateBits += "1"
		}
	}

	gammaRate, _ := strconv.ParseInt(gammaRateBits, 2, 64)
	epsilonRate, _ := strconv.ParseInt(epsilonRateBits, 2, 64)

	return gammaRate * epsilonRate
}

func lifeSupportRating(input []string) int64 {
	if len(input) == 0 {
		return 0
	}

	oxygenGenRating := ratingWithCriteria(input, mostCommonBit, 0)
	scrubberRating := ratingWithCriteria(input, leastCommonBit, 0)
	return oxygenGenRating * scrubberRating

}

func ratingWithCriteria(input []string, criteria ratingCriteria, currentBit int) int64 {
	if len(input) == 0 {
		return 0
	}

	if len(input) == 1 {
		result, _ := strconv.ParseInt(input[0], 2, 64)
		return result
	}

	bitsByPosition := bitsByPosition(input)

	commonBit := criteria(bitsByPosition[currentBit])

	var filtered []string
	for _, s := range input {
		if string(s[currentBit]) == commonBit {
			filtered = append(filtered, s)
		}
	}

	return ratingWithCriteria(filtered, criteria, currentBit+1)
}

func bitsByPosition(input []string) []string {
	// assume all lines in the input have the same length
	numberOfBits := len(input[0])
	result := make([]string, numberOfBits)
	for _, line := range input {
		for i, c := range line {
			result[i] += string(c)
		}
	}

	return result
}

type ratingCriteria func(string) string

func mostCommonBit(s string) string {
	oneBits := strings.Count(s, "1")
	zeroBits := strings.Count(s, "0")

	if oneBits >= zeroBits {
		return "1"
	}
	return "0"
}

func leastCommonBit(s string) string {
	if mostCommonBit(s) == "1" {
		return "0"
	}
	return "1"

}
