package day01

import (
	"bufio"
	"io"
	"strconv"

	"github.com/yitsushi/go-aoc/puzzle"
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
			return puzzle.InputParseError{Message: err.Error()}
		}

		d.input = append(d.input, value)
	}

	if scanner.Err() != nil {
		return puzzle.InputParseError{Message: scanner.Err().Error()}
	}

	return nil
}
