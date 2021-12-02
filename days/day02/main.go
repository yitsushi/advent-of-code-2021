package day02

import (
	"fmt"

	"github.com/yitsushi/go-aoc/math"
	"github.com/yitsushi/go-aoc/perf"
)

// Part1 for this day.
func (d *Solver) Part1() (string, error) {
	defer perf.Duration(perf.Track("Part1"))

	submarine := math.Vector3D{X: 0, Y: 0, Z: 0}

	for _, inst := range d.input {
		switch inst.Command {
		case ForwardCommand:
			submarine.X += float64(inst.Value)
		case DownCommand:
			submarine.Z += float64(inst.Value)
		case UpCommand:
			submarine.Z -= float64(inst.Value)
		case NoCommand:
		}
	}

	return fmt.Sprintf("%0.0f", submarine.X*submarine.Z), nil
}

// Part2 for this day.
func (d *Solver) Part2() (string, error) {
	defer perf.Duration(perf.Track("Part2"))

	submarine := math.Vector3D{X: 0, Y: 0, Z: 0}

	for _, inst := range d.input {
		switch inst.Command {
		case ForwardCommand:
			submarine.X += float64(inst.Value)
			submarine.Z += submarine.Y * float64(inst.Value)
		case DownCommand:
			submarine.Y += float64(inst.Value)
		case UpCommand:
			submarine.Y -= float64(inst.Value)
		case NoCommand:
		}
	}

	return fmt.Sprintf("%0.0f", submarine.X*submarine.Z), nil
}
