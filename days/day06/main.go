package day06

import (
	"fmt"

	"github.com/yitsushi/go-aoc/math"
	"github.com/yitsushi/go-aoc/perf"
)

const (
	part1Days    = 80
	part2Days    = 256
	ageGroupSize = 9
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	ageGroups := make([]int64, ageGroupSize)

	for _, fish := range d.input {
		ageGroups[fish]++
	}

	for day := 0; day < part1Days; day++ {
		newAgeGroups := make([]int64, ageGroupSize)
		gaveBirth := ageGroups[0]

		for idx, value := range ageGroups[1:] {
			newAgeGroups[idx] = value
		}

		newAgeGroups[8] = gaveBirth
		newAgeGroups[6] += gaveBirth

		ageGroups = newAgeGroups
	}

	return fmt.Sprintf("%d", math.SumInt64(ageGroups...)), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	ageGroups := make([]int64, ageGroupSize)

	for _, fish := range d.input {
		ageGroups[fish]++
	}

	for day := 0; day < part2Days; day++ {
		newAgeGroups := make([]int64, ageGroupSize)
		gaveBirth := ageGroups[0]

		for idx, value := range ageGroups[1:] {
			newAgeGroups[idx] = value
		}

		newAgeGroups[8] = gaveBirth
		newAgeGroups[6] += gaveBirth

		ageGroups = newAgeGroups
	}

	return fmt.Sprintf("%d", math.SumInt64(ageGroups...)), nil
}
