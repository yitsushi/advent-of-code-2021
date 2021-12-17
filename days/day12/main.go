package day12

import (
	"fmt"

	"github.com/yitsushi/go-aoc/perf"
	"github.com/yitsushi/go-aoc/puzzle"
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	routes, _, _ := d.input.AllPossibleRoute(nil, []*Node{})

	return fmt.Sprintf("%d", len(routes)), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	return "", puzzle.NotImplementedError{}
}
