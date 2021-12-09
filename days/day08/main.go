package day08

import (
	"fmt"

	"github.com/yitsushi/go-aoc/perf"
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	counter := 0

	for _, entry := range d.input {
		for _, digit := range entry.OutputValue {
			if Any(digit, entry.IsOne, entry.IsFour, entry.IsSeven, entry.IsEight) {
				counter++
			}
		}
	}

	return fmt.Sprintf("%d", counter), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	var result int64

	for _, entry := range d.input {
		entry.Parse()
		entry.TranslateOutput()

		result += entry.NumericValue()
	}

	return fmt.Sprintf("%d", result), nil
}
