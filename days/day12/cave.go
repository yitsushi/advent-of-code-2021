package day12

import "github.com/sirupsen/logrus"

// Map is the map of the cave system.
type Map struct {
	startNode    *Node
	endNode      *Node
	caves        map[string]*Node
	haveMoreTime bool
}

func newMap() Map {
	start := newNode("start")
	end := newNode("end")

	return Map{
		caves: map[string]*Node{
			"start": start,
			"end":   end,
		},
		startNode:    start,
		endNode:      end,
		haveMoreTime: false,
	}
}

// DoIHaveMoreFuckingTime is basically part2.
func (m *Map) DoIHaveMoreFuckingTime(answer bool) {
	m.haveMoreTime = answer
}

// AddConnection between caves.
func (m *Map) AddConnection(fromCave, toCave string) {
	if _, found := m.caves[fromCave]; !found {
		m.caves[fromCave] = newNode(fromCave)
	}

	if _, found := m.caves[toCave]; !found {
		m.caves[toCave] = newNode(toCave)
	}

	m.caves[fromCave].Link(m.caves[toCave])
	m.caves[toCave].Link(m.caves[fromCave])
}

// AllPossibleRoute generates all possible routes.
func (m *Map) AllPossibleRoute(node *Node, chain []*Node) ([][]*Node, bool, error) {
	if node == nil {
		node = m.startNode
	}

	if node == m.endNode {
		return [][]*Node{{node}}, true, nil
	}

	chain = append(chain, node)

	routes := [][]*Node{}

	options := node.Options(chain, m.haveMoreTime)
	if len(options) == 0 {
		return routes, false, NoMoreRouteError{NodeName: node.Name()}
	}

	for _, option := range options {
		route := []*Node{node}

		all, _, err := m.AllPossibleRoute(option, chain)
		if err != nil {
			logrus.WithField("node", *option).Info(err.Error())

			continue
		}

		for _, r := range all {
			routes = append(routes, append(route, r...))
		}
	}

	return routes, false, nil
}
