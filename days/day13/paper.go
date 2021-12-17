package day13

import (
	"github.com/yitsushi/go-aoc/math"
)

const (
	filledCharacter = rune(0x2588)
)

type foldablePaper map[math.Vector2DInt]bool

func (f foldablePaper) Max() (int, int) {
	var maxX, maxY int

	for pos := range f {
		if pos.X > maxX {
			maxX = pos.X
		}

		if pos.Y > maxY {
			maxY = pos.Y
		}
	}

	return maxX, maxY
}

func fold(paper foldablePaper, inst instruction) foldablePaper {
	newPaper := foldablePaper{}

	for position := range paper {
		if inst.Axis == axisx {
			if position.X == inst.Value {
				continue
			}

			if position.X < inst.Value {
				newPaper[position] = true
			} else {
				shift := inst.Value - (position.X - inst.Value)
				newPaper[math.Vector2DInt{X: shift, Y: position.Y}] = true
			}
		}

		if inst.Axis == axisy {
			if position.Y == inst.Value {
				continue
			}

			if position.Y < inst.Value {
				newPaper[position] = true
			} else {
				shift := inst.Value - (position.Y - inst.Value)
				newPaper[math.Vector2DInt{X: position.X, Y: shift}] = true
			}
		}
	}

	return newPaper
}

func show(paper foldablePaper) []string {
	maxX, maxY := paper.Max()

	lines := []string{}

	for idy := 0; idy <= maxY; idy++ {
		line := ""

		for idx := 0; idx <= maxX; idx++ {
			if _, found := paper[math.Vector2DInt{X: idx, Y: idy}]; found {
				line += string(filledCharacter)
			} else {
				line += " "
			}
		}

		lines = append(lines, line)
	}

	return lines
}
