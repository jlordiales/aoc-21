package aoc_2021

import (
	"strconv"
	"strings"
)

func countFishAfter(input []string, days int) int {
	fish := make(map[int]int)
	for _, n := range strings.Split(input[0], ",") {
		days, _ := strconv.Atoi(n)
		fish[days]++
	}

	for i := 0; i < days; i++ {
		for j := 0; j <= 8; j++ {
			if fish[j] > 0 {
				fish[j-1] += fish[j]
				fish[j] = 0
			}
		}

		fish[6] += fish[-1]
		fish[8] += fish[-1]
		fish[-1] = 0
	}

	var total int
	for _, v := range fish {
		total += v
	}

	return total
}
