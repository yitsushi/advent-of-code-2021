package day14

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Solver is the main solver.
type Solver struct {
	lastElement     byte
	polymer         polymerMap
	transformations transformationMap
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)
	d.transformations = map[string]string{}
	d.polymer = map[string]int{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		if !strings.Contains(line, "->") {
			for idx := 0; idx < len(line)-1; idx++ {
				d.polymer[line[idx:idx+2]]++
			}

			d.lastElement = line[len(line)-1]
		}

		var from, to string

		fmt.Sscanf(line, "%s -> %s", &from, &to)

		d.transformations[from] = to
	}

	return nil
}
