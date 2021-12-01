package day01

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

const (
	base10     = 10
	intBitSize = 64
)

// Solver is the main solver.
type Solver struct {
	input []int64
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		value, err := strconv.ParseInt(scanner.Text(), base10, intBitSize)
		if err != nil {
			return fmt.Errorf("failed to parse input: %w", err)
		}

		d.input = append(d.input, value)
	}

	if scanner.Err() != nil {
		return fmt.Errorf("failed to parse input: %w", scanner.Err())
	}

	return nil
}
