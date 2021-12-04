package aoc_2021

import (
	"strconv"
	"strings"
)

func move(lines []string) int {
	depth := 0
	horizontalPosition := 0

	for _, l := range lines {
		parts := strings.Split(l, " ")
		direction := parts[0]
		value, _ := strconv.Atoi(parts[1])
		switch direction {
		case "forward":
			horizontalPosition += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}

	return depth * horizontalPosition
}

func moveWithAim(lines []string) int {
	depth := 0
	horizontalPosition := 0
	aim := 0

	for _, l := range lines {
		parts := strings.Split(l, " ")
		direction := parts[0]
		value, _ := strconv.Atoi(parts[1])

		switch direction {
		case "forward":
			depth += value * aim
			horizontalPosition += value
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}

	return depth * horizontalPosition
}
