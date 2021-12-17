package day12

import (
	"bufio"
	"io"
	"strings"
)

const linkParts = 2

// Solver is the main solver.
type Solver struct {
	input Map
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)
	d.input = newMap()

	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), "-", linkParts)

		d.input.AddConnection(parts[0], parts[1])
	}

	return nil
}
