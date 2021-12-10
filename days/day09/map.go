package day09

import (
	"github.com/yitsushi/go-aoc/math"
)

const (
	topValue = 9
)

// CaveMap is the map of the cave.
type CaveMap struct {
	layers [][]int
}

func newCaveMap() CaveMap {
	return CaveMap{layers: [][]int{}}
}

// AddLayer to the map.
func (c *CaveMap) AddLayer(line []int) {
	c.layers = append(c.layers, line)
}

// Layers of the map.
func (c CaveMap) Layers() [][]int {
	return c.layers
}

// Layer of the map.
func (c CaveMap) Layer(n int) ([]int, error) {
	if n < 0 || n >= len(c.layers) {
		return []int{}, OutOfBoundError{Index: n}
	}

	return c.layers[n], nil
}

// Position value on the map.
func (c CaveMap) Position(coord math.Vector2DInt) (int, error) {
	targetRow, err := c.Layer(coord.Y)
	if err != nil {
		return -1, err
	}

	if coord.X < 0 || coord.X >= len(targetRow) {
		return -1, OutOfBoundError{Index: coord.X}
	}

	return targetRow[coord.X], nil
}

// Adjacents positions.
func (c CaveMap) Adjacents(coord math.Vector2DInt) []math.Vector2DInt {
	result := []math.Vector2DInt{}

	targets := []math.Vector2DInt{
		{X: 0, Y: 1},
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: -1, Y: 0},
	}

	for _, target := range targets {
		if _, err := c.Position(coord.Add(target)); err == nil {
			result = append(result, coord.Add(target))
		}
	}

	return result
}

// IsLowest checks if all adjacents are higher or not.
func (c CaveMap) IsLowest(coord math.Vector2DInt) bool {
	value, err := c.Position(coord)
	if err != nil {
		return false
	}

	for _, adjacent := range c.Adjacents(coord) {
		adjacentValue, err := c.Position(adjacent)
		if err != nil {
			continue
		}

		if adjacentValue <= value {
			return false
		}
	}

	return true
}

// FindLowPoints returns with all the positions where it's the lowest based
// on its adjacent positions.
func (c CaveMap) FindLowPoints() []math.Vector2DInt {
	result := []math.Vector2DInt{}

	for idy, line := range c.Layers() {
		for idx := range line {
			if c.IsLowest(math.Vector2DInt{Y: idy, X: idx}) {
				result = append(
					result,
					math.Vector2DInt{X: idx, Y: idy},
				)
			}
		}
	}

	return result
}

// WhoCanFlowHere walks up on the map from a point to discover basins.
func (c CaveMap) WhoCanFlowHere(coord math.Vector2DInt) []math.Vector2DInt {
	result := []math.Vector2DInt{}

	value, err := c.Position(coord)
	if err != nil {
		return result
	}

	for _, target := range c.Adjacents(coord) {
		targetValue, err := c.Position(target)
		if err != nil {
			continue
		}

		if targetValue == topValue {
			continue
		}

		if targetValue > value {
			result = append(result, target)
		}
	}

	for _, target := range result {
		result = append(result, c.WhoCanFlowHere(target)...)
	}

	return result
}
