package day05

import (
	"fmt"

	"github.com/yitsushi/go-aoc/perf"
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	fields := d.fill(func(data Range) bool {
		return data.From.X != data.To.X && data.From.Y != data.To.Y
	})
	moreThanOne := 0

	for _, value := range fields {
		if value > 1 {
			moreThanOne++
		}
	}

	return fmt.Sprintf("%d", moreThanOne), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	fields := d.fill(func(_ Range) bool { return false })
	moreThanOne := 0

	for _, value := range fields {
		if value > 1 {
			moreThanOne++
		}
	}

	return fmt.Sprintf("%d", moreThanOne), nil
}
