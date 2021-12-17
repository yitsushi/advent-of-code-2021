package day12

import "fmt"

// NoMoreRouteError when we hit a dead end.
type NoMoreRouteError struct {
	NodeName string
}

func (e NoMoreRouteError) Error() string {
	return fmt.Sprintf("no more routes from %s", e.NodeName)
}
