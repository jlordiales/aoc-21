package aoc_2021

import (
	"sort"
)

var (
	closingMap = map[string]string{
		"[": "]",
		"(": ")",
		"{": "}",
		"<": ">",
	}

	scoreCorrupted = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	scoreIncomplete = map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
)

type stack []string

func (s *stack) Push(v string) {
	*s = append(*s, v)
}

func (s *stack) Pop() string {
	if len(*s) == 0 {
		return ""
	}

	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func day10FirstHalf(input []string) int {
	res := 0
	for _, line := range input {
		if status, s := status(line); status == corrupted {
			res += scoreCorrupted[s]
		}
	}
	return res
}

func day10SecondHalf(input []string) int {
	var scores []int
	for _, line := range input {
		if status, s := status(line); status == incomplete {
			score := 0
			for _, c := range s {
				charScore := scoreIncomplete[string(c)]
				score = (score * 5) + charScore
			}
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)
	return scores[(len(scores) / 2)]
}

type lineStatus int

const (
	corrupted lineStatus = iota
	incomplete
	complete
)

func status(line string) (lineStatus, string) {
	var stack stack
	for _, s := range line {
		symbol := string(s)
		switch symbol {
		case "[", "(", "{", "<":
			stack.Push(symbol)
		case "]", ")", "}", ">":
			openingSymbol := stack.Pop()
			closer := closingMap[openingSymbol]

			// return the first "corrupted" symbol
			if symbol != closer {
				return corrupted, symbol
			}
		}
	}

	res := ""
	for len(stack) > 0 {
		openingSymbol := stack.Pop()
		res += closingMap[openingSymbol]
	}

	if len(res) > 0 {
		return incomplete, res
	}
	return complete, ""
}
