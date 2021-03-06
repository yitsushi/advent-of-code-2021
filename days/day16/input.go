package day16

import (
	"bufio"
	"io"
)

// Solver is the main solver.
type Solver struct {
	input []string
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		d.input = append(d.input, scanner.Text())
	}

	return nil
}
