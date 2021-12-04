package aoc_2021

import (
	"strconv"
	"strings"
)

type number struct {
	n      int
	marked bool
}

type board [][]number

func (b board) markNumber(n int) {
	for rowIdx, row := range b {
		for colIdx, cell := range row {
			if cell.n == n {
				b[rowIdx][colIdx].marked = true
				return
			}
		}
	}
}

func (b board) column(colNumber int) []number {
	var result []number
	for _, row := range b {
		result = append(result, row[colNumber])
	}

	return result
}

func (b board) isComplete() bool {
	if len(b) == 0 {
		return false
	}

	for _, row := range b {
		if allMarked(row) {
			return true
		}
	}

	numberOfColumns := len(b[0])
	for i := 0; i < numberOfColumns; i++ {
		if allMarked(b.column(i)) {
			return true
		}
	}

	return false
}

func (b board) sumUnmarked() int {
	var res int
	for _, r := range b {
		for _, number := range r {
			if !number.marked {
				res += number.n
			}
		}
	}
	return res
}

func allMarked(row []number) bool {
	for _, r := range row {
		if !r.marked {
			return false
		}
	}

	return true
}

func buildBoard(input []string) board {
	var res board
	for _, r := range input {
		var row []number
		cells := strings.Split(strings.TrimSpace(r), " ")
		for _, cell := range cells {
			if len(cell) == 0 {
				continue
			}

			value, _ := strconv.Atoi(cell)
			row = append(row, number{n: value})
		}

		res = append(res, row)
	}

	return res
}

type winResult struct {
	board           board
	lastNumberDrawn int
}

func (w winResult) score() int {
	return w.board.sumUnmarked() * w.lastNumberDrawn
}

func allBoardsCompleted(boards []*board) bool {
	for _, board := range boards {
		if !board.isComplete() {
			return false
		}
	}
	return true
}

func playUntil(input []string, keepPlayingFn func([]winResult) bool) []winResult {
	var res []winResult

	var drawnNumbers []int
	for _, n := range strings.Split(input[0], ",") {
		drawnNumber, _ := strconv.Atoi(n)
		drawnNumbers = append(drawnNumbers, drawnNumber)
	}

	var boards []*board
	for i := 2; i <= len(input)-5; i = i + 6 {
		b := buildBoard(input[i : i+5])
		boards = append(boards, &b)
	}

	for _, drawnNumber := range drawnNumbers {
		// no need to keep drawing numbers if all boards
		// are already completed
		if allBoardsCompleted(boards) {
			return res
		}

		for i := range boards {
			// skip the board if it's already complete
			if boards[i].isComplete() {
				continue
			}

			boards[i].markNumber(drawnNumber)
			if boards[i].isComplete() {
				res = append(res, winResult{
					board:           *boards[i],
					lastNumberDrawn: drawnNumber,
				})
			}

			if !keepPlayingFn(res) {
				return res
			}
		}
	}

	return res
}

func bingoToWin(input []string) int {
	results := playUntil(input, func(results []winResult) bool {
		// play until we get the first winner
		return len(results) < 1
	})

	return results[0].score()
}

func bingoToLoose(input []string) int {
	results := playUntil(input, func(results []winResult) bool {
		// play until all boards are complete (or we run out of numbers)
		return true
	})

	return results[len(results)-1].score()
}
