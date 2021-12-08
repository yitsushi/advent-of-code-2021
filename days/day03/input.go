package day03

import (
	"bufio"
	"io"
	"strconv"

	"github.com/yitsushi/go-aoc/puzzle"
)

const (
	base2   = 2
	intSize = 64
)

// Solver is the main solver.
type Solver struct {
	maskLength int
	mask       int64
	input      []int64
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		text := scanner.Text()

		if d.mask == 0 {
			d.maskLength = len(text)

			for i := 0; i < d.maskLength; i++ {
				d.mask = (d.mask << 1) + 1
			}
		}

		v, err := strconv.ParseInt(text, base2, intSize)
		if err != nil {
			return puzzle.InputParseError{Message: text}
		}

		d.input = append(d.input, v)
	}

	return nil
}
