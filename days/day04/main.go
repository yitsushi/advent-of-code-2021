package day04

import (
	"fmt"

	"github.com/yitsushi/go-aoc/perf"
	"github.com/yitsushi/go-aoc/puzzle"
)

func sum(values []int64) int64 {
	var result int64

	for _, value := range values {
		result += value
	}

	return result
}

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	for _, draw := range d.drawOrder {
		for _, board := range d.boards {
			if board.Check(draw) {
				return fmt.Sprintf("%d", draw*sum(board.UnmarkedValues())), nil
			}
		}
	}

	return "", puzzle.NoSolutionError{}
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	wins := 0

	for _, draw := range d.drawOrder {
		for _, board := range d.boards {
			if board.DidWin() {
				continue
			}

			if board.Check(draw) {
				wins++

				if wins == len(d.boards) {
					return fmt.Sprintf("%d", draw*sum(board.UnmarkedValues())), nil
				}
			}
		}
	}

	return "", puzzle.NoSolutionError{}
}
