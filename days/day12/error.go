package day12

import "fmt"

type NoMoreRouteError struct {
	NodeName string
}

func (e NoMoreRouteError) Error() string {
	return fmt.Sprintf("no more routes from %s", e.NodeName)
}
