package aoc_2021

import (
	"strconv"
	"strings"
)

func day13FirstHalf(input []string) int {
	paper, foldInst := parsePaper(input)

	paper.foldPaper(foldInst[0])

	return len(paper.points)
}

func day13SecondHalf(input []string) string {
	paper, foldInst := parsePaper(input)

	for _, f := range foldInst {
		paper.foldPaper(f)
	}

	return paper.String()
}

type paper struct {
	points map[paperPoint]struct{}
}

type paperPoint struct {
	x, y int
}

func (paper paper) String() string {
	var maxX, maxY int
	for p := range paper.points {
		if p.x > maxX {
			maxX = p.x
		}

		if p.y > maxY {
			maxY = p.y
		}
	}

	res := ""
	for i := 0; i <= maxY; i++ {
		for j := 0; j <= maxX; j++ {
			if _, ok := paper.points[paperPoint{x: j, y: i}]; ok {
				res += "#"
			} else {
				res += " "
			}
		}
		res += "\n"
	}

	return res
}

func (paper *paper) foldPaper(foldInstruction paperPoint) {
	if foldInstruction.y > 0 {
		paper.foldHorizontal(foldInstruction)
		return
	}
	paper.foldVertical(foldInstruction)
}

func (paper *paper) foldVertical(foldInstruction paperPoint) {
	foldCol := foldInstruction.x

	for point := range paper.points {
		if point.x > foldCol {
			colDiff := point.x - foldCol
			newPoint := paperPoint{
				y: point.y,
				x: foldCol - colDiff,
			}
			paper.points[newPoint] = struct{}{}
			delete(paper.points, point)
		}
	}
}

func (paper *paper) foldHorizontal(foldInstruction paperPoint) {
	foldRow := foldInstruction.y

	for point := range paper.points {
		if point.y > foldRow {
			rowDiff := point.y - foldRow
			newPoint := paperPoint{
				x: point.x,
				y: foldRow - rowDiff,
			}
			paper.points[newPoint] = struct{}{}
			delete(paper.points, point)
		}
	}
}

func parsePaper(input []string) (paper, []paperPoint) {
	p := paper{
		points: make(map[paperPoint]struct{}),
	}

	var foldInst []paperPoint
	for _, l := range input {
		if strings.HasPrefix(l, "fold") {
			foldInst = append(foldInst, parseFoldInstruction(l))
			continue
		}

		if len(l) > 0 {
			split := strings.Split(l, ",")
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])
			p.points[paperPoint{x: x, y: y}] = struct{}{}
		}
	}

	return p, foldInst
}

func parseFoldInstruction(l string) paperPoint {
	split := strings.Split(l, "=")
	value, _ := strconv.Atoi(split[1])

	if strings.Contains(l, "x") {
		return paperPoint{x: value}
	}
	return paperPoint{y: value}
}
