package aoc_2021

import (
	"strconv"
)

type energyLevel struct {
	value   int
	x, y    int
	flashed bool
}

func day11FirstHalf(input []string) int {
	octopuses := parseEnergyLevels(input)
	res := 0
	for i := 0; i < 100; i++ {
		res += step(octopuses)
	}
	return res
}

func day11SecondHalf(input []string) int {
	octopuses := parseEnergyLevels(input)

	i := 0
	for {
		i++
		step(octopuses)
		if syncronized(octopuses) {
			return i
		}
	}
}

func syncronized(m [][]*energyLevel) bool {
	for _, row := range m {
		for _, cell := range row {
			if cell.value > 0 {
				return false
			}
		}
	}

	return true
}

func step(m [][]*energyLevel) int {
	res := 0
	for _, row := range m {
		for _, cell := range row {
			cell.value++
		}
	}

	for _, row := range m {
		for _, cell := range row {
			res += flash(m, cell)
		}
	}

	// reset the flashed flag at the end of each step
	for _, row := range m {
		for _, cell := range row {
			cell.flashed = false
		}
	}

	return res
}

func flash(m [][]*energyLevel, cell *energyLevel) int {
	if cell.value <= 9 || cell.flashed {
		return 0
	}

	cell.value = 0
	cell.flashed = true
	flashesCount := 1

	for _, adj := range adjacentOctopuses(m, cell.x, cell.y) {
		if adj.flashed {
			continue
		}

		adj.value++
		flashesCount += flash(m, adj)
	}

	return flashesCount
}

func adjacentOctopuses(m [][]*energyLevel, i, j int) []*energyLevel {
	var res []*energyLevel

	// we are not on the first row
	if i > 0 {
		// above us
		res = append(res, m[i-1][j])

		if j > 0 {
			// up and left diagonal
			res = append(res, m[i-1][j-1])
		}

		if j < len(m[i])-1 {
			// up and right diagonal
			res = append(res, m[i-1][j+1])
		}

	}

	// we are not on the last row
	if i < len(m)-1 {
		// below
		res = append(res, m[i+1][j])

		if j > 0 {
			// down and left diagonal
			res = append(res, m[i+1][j-1])
		}

		if j < len(m[i])-1 {
			// down and right diagonal
			res = append(res, m[i+1][j+1])
		}
	}

	// we are not on the first column
	if j > 0 {
		// left cell
		res = append(res, m[i][j-1])
	}

	// we are not on the last column
	if j < len(m[i])-1 {
		// right cell
		res = append(res, m[i][j+1])
	}

	return res
}

func parseEnergyLevels(input []string) [][]*energyLevel {
	var res [][]*energyLevel

	for rowN, line := range input {
		var row []*energyLevel
		for columnN, n := range line {
			height, _ := strconv.Atoi(string(n))
			row = append(row, &energyLevel{value: height, x: rowN, y: columnN})
		}

		res = append(res, row)
	}

	return res
}
