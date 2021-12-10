package day09

import (
	"fmt"
	"sort"

	"github.com/yitsushi/go-aoc/math"
	"github.com/yitsushi/go-aoc/perf"
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	result := 0

	for _, location := range d.input.FindLowPoints() {
		if value, err := d.input.Position(location); err == nil {
			result += value + 1
		}
	}

	return fmt.Sprintf("%d", result), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	basinSizes := []int{}

	for _, sink := range d.input.FindLowPoints() {
		basin := append(d.input.WhoCanFlowHere(sink), sink)
		basinSizes = append(basinSizes, len(removeDuplicateValues(basin)))
	}

	sort.Slice(basinSizes, func(i, j int) bool {
		return basinSizes[i] > basinSizes[j]
	})

	result := basinSizes[0] * basinSizes[1] * basinSizes[2]

	return fmt.Sprintf("%d", result), nil
}

func removeDuplicateValues(input []math.Vector2DInt) []math.Vector2DInt {
	keys := make(map[interface{}]bool)
	list := []math.Vector2DInt{}

	for _, entry := range input {
		hash := entry.Hash()
		if _, value := keys[hash]; !value {
			keys[hash] = true

			list = append(list, entry)
		}
	}

	return list
}
