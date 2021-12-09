package day05

import (
	"bufio"
	"fmt"
	"io"

	"github.com/yitsushi/go-aoc/math"
)

// Solver is the main solver.
type Solver struct {
	input []Range
}

type Range struct {
	From math.Vector2D
	To   math.Vector2D
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		from := math.Vector2D{}
		to := math.Vector2D{}

		fmt.Sscanf(line, "%f,%f -> %f,%f", &from.X, &from.Y, &to.X, &to.Y)

		d.input = append(d.input, Range{From: from, To: to})
	}

	return nil
}
