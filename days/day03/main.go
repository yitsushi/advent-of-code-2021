package day03

import (
	"fmt"

	"github.com/yitsushi/go-aoc/perf"
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	gammaRate, epsilonRate := findMostAndLeastCommonBits(d.input, d.maskLength, d.mask)

	return fmt.Sprintf("%d", gammaRate*epsilonRate), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	o2gr := reduceFromGamma(d.input, d.maskLength, d.mask)
	co2sr := reduceFromEpsilon(d.input, d.maskLength, d.mask)

	return fmt.Sprintf("%d", o2gr*co2sr), nil
}
