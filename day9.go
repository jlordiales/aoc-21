package aoc_2021

import (
	"sort"
	"strconv"
)

type cavePoint struct {
	height  int
	x, y    int
	visited bool
}

type basin []cavePoint

func day9FirstHalf(input []string) int {
	matrix := parseInput(input)

	lowPoints := lowPoints(matrix)

	res := 0
	for _, p := range lowPoints {
		res += 1 + p.height
	}

	return res
}

func day9SecondHalf(input []string) int {
	matrix := parseInput(input)

	lowPoints := lowPoints(matrix)

	var basins []basin
	for _, lowPoint := range lowPoints {
		basins = append(basins, basinFromPoint(matrix, lowPoint.x, lowPoint.y))
	}

	sort.Slice(basins, func(i, j int) bool {
		return len(basins[i]) > len(basins[j])
	})

	return len(basins[0]) * len(basins[1]) * len(basins[2])

}

func basinFromPoint(m [][]cavePoint, x, y int) basin {
	point := m[x][y]
	if point.visited || point.height == 9 {
		return nil
	}

	m[point.x][point.y].visited = true
	res := []cavePoint{point}

	for _, p := range adjacentPoints(m, x, y) {
		adjBasing := basinFromPoint(m, p.x, p.y)
		res = append(res, adjBasing...)
	}

	return res
}

func lowPoints(m [][]cavePoint) []cavePoint {
	var res []cavePoint
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			n := m[i][j]
			adjacent := adjacentPoints(m, i, j)

			if isLowPoint(n.height, adjacent) {
				res = append(res, n)
			}
		}
	}

	return res
}

func isLowPoint(n int, adjacent []cavePoint) bool {
	for _, adj := range adjacent {
		if n >= adj.height {
			return false
		}
	}

	return true
}

func adjacentPoints(m [][]cavePoint, i, j int) []cavePoint {
	var res []cavePoint

	if i > 0 {
		// there's a row above us
		res = append(
			res,
			m[i-1][j],
		)
	}

	if i < len(m)-1 {
		res = append(
			res,
			m[i+1][j],
		)
	}

	if j > 0 {
		res = append(
			res,
			m[i][j-1],
		)
	}

	if j < len(m[i])-1 {
		res = append(
			res,
			m[i][j+1],
		)
	}

	return res

}

func parseInput(input []string) [][]cavePoint {
	var res [][]cavePoint

	for rowN, line := range input {
		var row []cavePoint
		for columnN, n := range line {
			height, _ := strconv.Atoi(string(n))
			row = append(row, cavePoint{height: height, x: rowN, y: columnN})
		}

		res = append(res, row)
	}

	return res
}
