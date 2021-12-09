package day07

import (
	"fmt"
	mmath "math"
	"sort"

	"github.com/yitsushi/go-aoc/math"
	"github.com/yitsushi/go-aoc/perf"
)

const (
	half = 2
)

func abs(value int64) int64 {
	if value < 0 {
		return -value
	}

	return value
}

func median(values ...int64) int64 {
	valuesLength := len(values)

	if valuesLength%2 == 1 {
		return values[valuesLength/2]
	}

	return values[valuesLength/2-1]
}

func mean(values ...int64) int64 {
	sum := float64(math.SumInt64(values...))
	valuesLength := float64(len(values))

	return int64(mmath.Floor(sum / valuesLength))
}

func triangular(value int64) int64 {
	return value * (value + 1) / half
}

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	sort.Slice(d.input, func(i, j int) bool { return d.input[i] < d.input[j] })

	var cost int64

	target := median(d.input...)

	for _, position := range d.input {
		cost += abs(position - target)
	}

	return fmt.Sprintf("%d", cost), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	var cost int64

	target := mean(d.input...)

	for _, position := range d.input {
		cost += triangular(abs(position - target))
	}

	return fmt.Sprintf("%d", cost), nil
}
