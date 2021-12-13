package day11

import (
	"fmt"

	"github.com/yitsushi/go-aoc/perf"
	"github.com/yitsushi/go-aoc/puzzle"
)

const (
	part1Steps = 100
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	d.input.ResetFlashCount()

	for i := 1; i <= part1Steps; i++ {
		d.input.Cycle()
	}

	return fmt.Sprintf("%d", d.input.FlashCount()), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	for i := 1; true; i++ {
		d.input.ResetFlashCount()
		d.input.Cycle()

		if d.input.FlashCount() == d.input.Size() {
			return fmt.Sprintf("%d", i), nil
		}
	}

	return "", puzzle.NoSolutionError{}
}
