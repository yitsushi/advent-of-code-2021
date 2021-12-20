package day15

import (
	"bufio"
	"io"

	"github.com/yitsushi/go-aoc/math"
)

const (
	asciiShift = 48
)

// Solver is the main solver.
type Solver struct {
	cave Cave
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)
	d.cave = Cave{Chitons: map[math.Vector2DInt]*Node{}}
	idy := 0

	for scanner.Scan() {
		line := scanner.Text()

		for idx, character := range line {
			d.cave.AddChiton(
				int(character-asciiShift),
				math.Vector2DInt{X: idx, Y: idy},
			)
		}

		idy++
	}

	return nil
}
