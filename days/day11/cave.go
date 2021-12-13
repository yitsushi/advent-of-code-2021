package day11

import (
	"github.com/yitsushi/go-aoc/math"
)

// Cave is the cave with octopuses.
type Cave struct {
	octopuses  map[interface{}]*Octopus
	flashCount int64
	maxY       int
	maxX       int
}

func newCave() Cave {
	return Cave{
		octopuses:  map[interface{}]*Octopus{},
		flashCount: 0,
		maxY:       0,
		maxX:       0,
	}
}

// AddOctopus into the cave.
func (c *Cave) AddOctopus(brightness int, coord math.Vector2DInt) {
	octopus := Octopus{
		brightness: brightness,
		coordinate: coord,
	}

	if c.maxY < coord.Y {
		c.maxY = coord.Y
	}

	if c.maxX < coord.X {
		c.maxX = coord.X
	}

	c.octopuses[coord.Hash()] = &octopus
}

// Get octopus at position.
func (c Cave) Get(coord math.Vector2DInt) *Octopus {
	if octopus, found := c.octopuses[coord.Hash()]; found {
		return octopus
	}

	return nil
}

// Brighten an octopus and its neighbours if flashed.
func (c *Cave) Brighten(octopus *Octopus) {
	if octopus == nil {
		return
	}

	if octopus.Brighten() {
		c.flashCount++

		for _, target := range octopus.Adjacents() {
			c.Brighten(c.Get(target))
		}
	}
}

// Cycle is one cycle int their precious life of blinking.
func (c *Cave) Cycle() {
	for _, octopus := range c.octopuses {
		c.Brighten(octopus)
	}

	// Cleanup
	for _, octopus := range c.octopuses {
		octopus.Clean()
	}
}

// FlashCount returns the number of flashes.
func (c Cave) FlashCount() int64 {
	return c.flashCount
}

// ResetFlashCount resets the flash counter.
func (c *Cave) ResetFlashCount() {
	c.flashCount = 0
}

// Size of the cave, how many octopuses are in there.
func (c Cave) Size() int64 {
	return int64(len(c.octopuses))
}
