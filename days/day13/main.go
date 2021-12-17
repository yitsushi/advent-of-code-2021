package day13

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/go-aoc/perf"
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	// show(d.paper)

	paper := fold(d.paper, d.instructions[0])

	for _, line := range show(paper) {
		logrus.Info(line)
	}

	return fmt.Sprintf("%d", len(paper)), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	paper := d.paper

	for _, inst := range d.instructions {
		paper = fold(paper, inst)
	}

	return strings.Join(show(paper), "\n"), nil // puzzle.NoSolutionError{}
}
