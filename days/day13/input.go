package day13

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/yitsushi/go-aoc/math"
)

// Solver is the main solver.
type Solver struct {
	paper        foldablePaper
	instructions []instruction
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)
	d.paper = foldablePaper{}
	d.instructions = []instruction{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "fold") {
			inst := instruction{Axis: 0, Value: 0}

			fmt.Sscanf(line, "fold along %c=%d", &inst.Axis, &inst.Value)

			d.instructions = append(d.instructions, inst)

			continue
		}

		position := math.Vector2DInt{X: 0, Y: 0}

		fmt.Sscanf(line, "%d,%d", &position.X, &position.Y)

		d.paper[position] = true
	}

	return nil
}
