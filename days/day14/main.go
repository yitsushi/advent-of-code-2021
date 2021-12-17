package day14

import (
	"fmt"

	"github.com/yitsushi/go-aoc/perf"
)

const (
	part1Cycles = 10
	part2Cycles = 40
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	result := d.justDoIt(d.polymer, part1Cycles)

	return fmt.Sprintf("%d", result), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	result := d.justDoIt(d.polymer, part2Cycles)

	return fmt.Sprintf("%d", result), nil
}

func (d *Solver) justDoIt(polymer polymerMap, steps int) int {
	for step := 0; step < steps; step++ {
		polymer = resolve(polymer, d.transformations)
	}

	counter := countElements(polymer)
	// +1 for the last element of the polymer if it's the most of the least common
	// element.
	counter[d.lastElement]++
	min, _, max, _ := minMax(counter)

	return max - min
}
