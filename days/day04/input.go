package day04

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/yitsushi/go-aoc/parsing"
	"github.com/yitsushi/go-aoc/puzzle"
)

// Solver is the main solver.
type Solver struct {
	boards    []*Board
	drawOrder []int64
}

// SetInput receives the input and parses its content.
func (d *Solver) SetInput(input io.Reader) error {
	scanner := bufio.NewScanner(input)
	d.boards = []*Board{}

	var board *Board

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			newBoard := Board{}
			board = &newBoard

			d.boards = append(d.boards, &newBoard)

			continue
		}

		if strings.Contains(line, ",") {
			var err error

			d.drawOrder, err = parsing.Int64Slice(line, ",")
			if err != nil {
				return puzzle.InputParseError{
					Message: fmt.Sprintf("unable to parse draw order: %s", line),
				}
			}

			continue
		}

		line = strings.ReplaceAll(line, "  ", " ")
		line = strings.TrimPrefix(line, " ")

		numbers, err := parsing.Int64Slice(line, " ")
		if err != nil {
			return puzzle.InputParseError{
				Message: fmt.Sprintf("unable to parse board line: %s", line),
			}
		}

		for _, n := range numbers {
			board.Add(n)
		}
	}

	return nil
}
