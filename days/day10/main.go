package day10

import (
	"fmt"
	"sort"

	"github.com/yitsushi/go-aoc/perf"
)

const (
	missingScoreMultiplier = 5
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	finalScore := 0

	for _, sequence := range d.input {
		valid, item, _ := parse(sequence)
		if valid {
			continue
		}

		finalScore += scoreError(item)
	}

	return fmt.Sprintf("%d", finalScore), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	scoreList := []int{}

	for _, sequence := range d.input {
		score := 0

		valid, _, stack := parse(sequence)
		if !valid {
			continue
		}

		for !stack.IsEmpty() {
			item, ok := stack.Pop().(Item)
			if !ok {
				continue
			}

			score = score*missingScoreMultiplier + scoreMissing(item)
		}

		scoreList = append(scoreList, score)
	}

	sort.Ints(scoreList)

	return fmt.Sprintf("%d", scoreList[len(scoreList)/2]), nil
}
