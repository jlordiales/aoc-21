package aoc_2021

import (
	"math"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type line struct {
	start point
	end   point
}

func (l line) isHorizontal() bool {
	return l.start.y == l.end.y
}

func (l line) isVertical() bool {
	return l.start.x == l.end.x
}

func (l line) isDiagonal() bool {
	return math.Abs(float64(l.start.x)-float64(l.end.x)) == math.Abs(float64(l.start.y)-float64(l.end.y))
}

func movementDirection(start, end int) int {
	diff := end - start
	if diff == 0 {
		return 0
	}

	if diff < 0 {
		return -1
	}

	return 1
}

// return all points that make up this line
func (l line) points() []point {
	xDirection := movementDirection(l.start.x, l.end.x)
	yDirection := movementDirection(l.start.y, l.end.y)

	// make sure we haven't gone too far in either direction
	keepMoving := func(x, y int) bool {
		if xDirection > 0 && x > l.end.x {
			return false
		}
		if xDirection < 0 && x < l.end.x {
			return false
		}
		if yDirection > 0 && y > l.end.y {
			return false
		}
		if yDirection < 0 && y < l.end.y {
			return false
		}

		return true
	}

	var res []point
	for x, y := l.start.x, l.start.y; keepMoving(x, y); x, y = x+xDirection, y+yDirection {
		res = append(res, point{x: x, y: y})
	}

	return res
}

func parsePoint(s string) point {
	parts := strings.Split(s, ",")

	x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
	y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
	return point{
		x: x,
		y: y,
	}
}

func parseLines(input []string, criteria func(line) bool) []line {
	var lines []line

	for _, l := range input {
		lineParts := strings.Split(l, "->")

		firstPoint := parsePoint(lineParts[0])
		secondPoint := parsePoint(lineParts[1])

		line := line{start: firstPoint, end: secondPoint}

		if criteria(line) {
			lines = append(lines, line)
		}
	}

	return lines
}

func filterByMinOverlaps(lines []line, numberOfOverlaps int) []point {
	overlapsByPoint := make(map[point]int)

	for _, l := range lines {
		for _, point := range l.points() {
			overlapsByPoint[point] = overlapsByPoint[point] + 1
		}
	}

	var res []point
	for k, v := range overlapsByPoint {
		if v >= numberOfOverlaps {
			res = append(res, k)
		}
	}

	return res
}

func numberOfPointsStraight(input []string) int {
	lines := parseLines(input, func(l line) bool {
		return l.isHorizontal() || l.isVertical()
	})

	return len(filterByMinOverlaps(lines, 2))
}

func numberOfPointsDiagonal(input []string) int {
	lines := parseLines(input, func(l line) bool {
		return l.isHorizontal() || l.isVertical() || l.isDiagonal()
	})

	return len(filterByMinOverlaps(lines, 2))
}
