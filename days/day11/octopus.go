package day11

import "github.com/yitsushi/go-aoc/math"

// Octopus is one octopus.
type Octopus struct {
	brightness int
	coordinate math.Vector2DInt
}

const maxBrigthness = 9

// Brighten the octopus.
func (o *Octopus) Brighten() bool {
	o.brightness++

	return o.brightness == maxBrigthness+1
}

// Adjacents positions.
func (o Octopus) Adjacents() []math.Vector2DInt {
	result := []math.Vector2DInt{}

	targets := []math.Vector2DInt{
		{X: -1, Y: -1},
		{X: 0, Y: -1},
		{X: 1, Y: -1},
		{X: -1, Y: 0},
		{X: 1, Y: 0},
		{X: -1, Y: 1},
		{X: 0, Y: 1},
		{X: 1, Y: 1},
	}

	for _, target := range targets {
		result = append(result, o.coordinate.Add(target))
	}

	return result
}

// Clean up after flashes.
func (o *Octopus) Clean() {
	if o.brightness > maxBrigthness {
		o.brightness = 0
	}
}
