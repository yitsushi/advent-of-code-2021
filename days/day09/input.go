package day09

import (
	"bufio"
	"io"
)

// Solver is the main solver.
type Solver struct {
	input CaveMap
}

const (
	asciiShift = 48
)

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)
	d.input = newCaveMap()

	for scanner.Scan() {
		text := scanner.Text()
		line := []int{}

		for _, character := range text {
			line = append(line, int(character-asciiShift))
		}

		d.input.AddLayer(line)
	}

	return nil
}
