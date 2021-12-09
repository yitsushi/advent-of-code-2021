package day08

import (
	"bufio"
	"io"
	"strings"
)

const (
	inputParts = 2
)

// Solver is the main solver.
type Solver struct {
	input []Entry
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " | ", inputParts)

		entry := NewEntry()
		entry.UniquePatterns = strings.Split(parts[0], " ")
		entry.OutputValue = strings.Split(parts[1], " ")

		d.input = append(d.input, entry)
	}

	return nil
}
