package day09

import "fmt"

// OutOfBoundError happens when we want to reach a position ouside of the map.
type OutOfBoundError struct {
	Index int
}

func (e OutOfBoundError) Error() string {
	return fmt.Sprintf("out of bound error: %d", e.Index)
}
