package day15

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/go-aoc/math"
	"github.com/yitsushi/go-aoc/perf"
	"github.com/yitsushi/go-aoc/puzzle"
)

const (
	part2Multiplier = 5
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	path := findPath(d.cave)

	if len(path.Chain) == 0 {
		return "", puzzle.NoSolutionError{}
	}

	return fmt.Sprintf("%d", path.Value()), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	cave := ExtendCave(d.cave, part2Multiplier)

	path := findPath(cave)

	if len(path.Chain) == 0 {
		return "", puzzle.NoSolutionError{}
	}

	return fmt.Sprintf("%d", path.Value()), nil
}

func findPath(cave Cave) NodeGroup {
	maxx, maxy := cave.MaxValues()
	start := cave.At(math.Vector2DInt{X: 0, Y: 0})
	end := cave.At(math.Vector2DInt{X: maxx, Y: maxy})

	logrus.Info(start, end)

	path := cave.Route(start, end)

	logrus.Info(path.Path())
	logrus.Info(path.Value())

	cave.Show(path)

	return path
}
