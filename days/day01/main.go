package day01

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/go-aoc/perf"
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	var (
		previous int64 = -1
		counter  int64
	)

	for _, value := range d.input {
		if previous != -1 && previous < value {
			logrus.Infof("%d -> %d", previous, value)
			counter++
		}

		previous = value
	}

	return fmt.Sprintf("%d", counter), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	var (
		previous int64 = -1
		counter  int64
	)

	for idx := range d.input[:len(d.input)-2] {
		value := d.input[idx] + d.input[idx+1] + d.input[idx+2]

		if previous != -1 && previous < value {
			logrus.Infof("%d -> %d", previous, value)
			counter++
		}

		previous = value
	}

	return fmt.Sprintf("%d", counter), nil
}
