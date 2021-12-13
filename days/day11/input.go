package day11

import (
	"bufio"
	"io"

	"github.com/yitsushi/go-aoc/math"
)

// Solver is the main solver.
type Solver struct {
	input Cave
}

const (
	asciiShift = 48
)

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)
	d.input = newCave()
	idy := 0

	for scanner.Scan() {
		text := scanner.Text()

		for idx, character := range text {
			d.input.AddOctopus(
				int(character-asciiShift),
				math.Vector2DInt{X: idx, Y: idy},
			)
		}

		idy++
	}

	return nil
}
