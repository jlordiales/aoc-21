package aoc_2021

import (
	"strconv"
	"strings"
)

func countFishAfter(input []string, days int) int {
	var fish []int
	for _, n := range strings.Split(input[0], ",") {
		days, _ := strconv.Atoi(n)
		fish = append(fish, days)
	}

	for i := 0; i < days; i++ {
		var newOnes []int
		for j := 0; j < len(fish); j++ {
			if fish[j] == 0 {
				fish[j] = 6
				newOnes = append(newOnes, 8)
			} else {
				fish[j]--
			}
		}

		fish = append(fish, newOnes...)
	}

	return len(fish)
}
