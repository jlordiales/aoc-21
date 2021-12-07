package aoc_2021

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func day7FirstHalf(input []string) int {
	return minFuelCost(input, cost1)
}

func day7SecondHalf(input []string) int {
	return minFuelCost(input, cost2)
}

func minFuelCost(input []string, costFn func(int, int) int) int {
	positions := parsePositions(input[0])
	sort.Ints(positions)

	minPoint := positions[0]
	maxPoint := positions[len(positions)-1]

	minCost := math.MaxInt32

	// just calculate the total cost for each point in the range
	// and keep the min. Not too fancy but it works ok for the size
	// of the input
	for i := minPoint; i <= maxPoint; i++ {
		if cost := totalFuelCostForPoint(i, positions, costFn); cost < minCost {
			minCost = cost
		}
	}

	return minCost
}

func parsePositions(s string) []int {
	numbers := strings.Split(s, ",")

	var res []int
	for _, n := range numbers {
		parsed, _ := strconv.Atoi(n)
		res = append(res, parsed)
	}

	return res
}

func totalFuelCostForPoint(point int, positions []int, costFn func(int, int) int) int {
	cost := 0
	for _, p := range positions {
		cost += costFn(point, p)
	}
	return cost
}

func cost2(p1, p2 int) int {
	diff := int(math.Abs(float64(p1 - p2)))

	res := 0
	for i := 1; i <= diff; i++ {
		res += i
	}
	return res
}

func cost1(p1, p2 int) int {
	return int(math.Abs(float64(p1 - p2)))
}
