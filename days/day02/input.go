package day02

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/yitsushi/go-aoc/puzzle"
)

// Solver is the main solver.
type Solver struct {
	input []Instruction
}

const (
	base10           = 10
	intBitSize       = 64
	instructionParts = 2
)

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), " ", instructionParts)
		inst := NewInstruction()

		if len(parts) != instructionParts {
			return puzzle.InputParseError{
				Message: fmt.Sprintf("invalid instruction: %s", parts[0]),
			}
		}

		value, err := strconv.ParseInt(parts[1], base10, intBitSize)
		if err != nil {
			return puzzle.InputParseError{Message: err.Error()}
		}

		inst.Value = value

		switch parts[0] {
		case "forward":
			inst.Command = ForwardCommand
		case "down":
			inst.Command = DownCommand
		case "up":
			inst.Command = UpCommand
		default:
			return puzzle.InputParseError{
				Message: fmt.Sprintf("invalid instruction: %s", parts[0]),
			}
		}

		d.input = append(d.input, inst)
	}

	if scanner.Err() != nil {
		return puzzle.InputParseError{Message: scanner.Err().Error()}
	}

	return nil
}
