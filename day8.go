package aoc_2021

import (
	"sort"
	"strconv"
	"strings"
)

func day8FirstHalf(input []string) int {
	outputs := decodeAllOutputs(input)

	res := 0
	for _, output := range outputs {
		s := strconv.Itoa(output)
		res += strings.Count(s, "1")
		res += strings.Count(s, "4")
		res += strings.Count(s, "7")
		res += strings.Count(s, "8")
	}

	return res
}

func day8SecondHalf(input []string) int {
	count := 0
	outputs := decodeAllOutputs(input)

	for _, n := range outputs {
		count += n
	}

	return count
}

func decodeAllOutputs(input []string) []int {
	var res []int

	for _, s := range input {
		parts := strings.Split(s, "|")

		signalsPatterns := strings.TrimSpace(parts[0])
		output := strings.TrimSpace(parts[1])

		mapping := decodeEntry(signalsPatterns)
		value := outputValue(output, mapping)
		res = append(res, value)
	}

	return res
}

func outputValue(output string, inputMapping map[string]string) int {
	resNumber := ""
	fields := strings.Fields(output)

	for _, s := range fields {
		sorted := sortString(s)
		if n, ok := inputMapping[sorted]; ok {
			resNumber += n
		}
	}

	res, _ := strconv.Atoi(resNumber)
	return res
}

// given the set of 10 patterns observed for a given entry
// returns the mapping between each pattern and its corresponding
// value (as a string)
func decodeEntry(signalPatterns string) map[string]string {
	// count how many times each segment (a-g) is found
	// in the input. This + the length of each pattern
	// will determine the mapping result
	segmentsCount := make(map[string]int)
	patterns := strings.Fields(signalPatterns)

	for _, pattern := range patterns {
		for _, s := range pattern {
			segmentsCount[string(s)]++
		}
	}

	resultMapping := make(map[string]string)
	for _, pattern := range patterns {
		sorted := sortString(pattern)
		resultMapping[sorted] = valueForPattern(sorted, segmentsCount)
	}

	return resultMapping
}

/**
  0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg
*/

// ^^^ the actual segment letters in the input
// will be scrambled but this helps reason about the
// different patterns
func valueForPattern(pattern string, segmentsCount map[string]int) string {
	switch len(pattern) {
	case 2:
		// only "1" has 2 segments
		return "1"
	case 3:
		// only "7" has 3 segments
		return "7"
	case 4:
		// only "4" has 4 segments
		return "4"
	case 7:
		// only "8" has 7 segments
		return "8"
	case 5:
		// "2","3" and "5" have 5 segments
		// from those, only "2" contains segment "e"
		// and that segment is the only present in 4/10
		// numbers (0,2,6,8). So if we find that segment we know this is "2".
		// Similarly, "5" is the only one that contains segment "b" and that segment
		// is only present in 6/10 numbers (0,4,5,6,8,9)
		for _, d := range pattern {
			if segmentsCount[string(d)] == 4 {
				return "2"
			}

			if segmentsCount[string(d)] == 6 {
				return "5"
			}
		}
		return "3"
	case 6:
		// "0", "6" and "9" have 6 segments
		// from those "9" is the only one missing segment
		// "e" (which we know is the only one that appears in 4 numbers)
		isNine := true
		for _, d := range pattern {
			if segmentsCount[string(d)] == 4 {
				isNine = false
			}
		}

		if isNine {
			return "9"
		}

		// either 0 or 6 then
		// 0 is missing segment d -> present in 7/10 numbers (2,3,4,5,6,8,9)
		// segment g is also present in 7/10 numbers (0,2,3,5,6,8,9)
		// So if we find 2 segments where count == 7 that means this is a "6" (has 'd' and 'g')
		n := 0
		for _, d := range pattern {
			if segmentsCount[string(d)] == 7 {
				n++
			}
		}

		if n == 2 {
			return "6"
		}
		return "0"
	}

	return ""
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
