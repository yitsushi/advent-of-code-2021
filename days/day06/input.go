package day06

import (
	"bufio"
	"io"

	"github.com/yitsushi/go-aoc/parsing"
)

// Solver is the main solver.
type Solver struct {
	input []int64
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		result, err := parsing.Int64Slice(scanner.Text(), ",")
		if err != nil {
			return err
		}

		for _, fish := range result {
			d.input = append(d.input, fish)
		}
	}

	return nil
}
