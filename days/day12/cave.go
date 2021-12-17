package day12

import "github.com/sirupsen/logrus"

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

func (m *Map) DoIHaveMoreFuckingTime(answer bool) {
	m.haveMoreTime = answer
}

func (m *Map) AddConnection(from, to string) {
	if _, found := m.caves[from]; !found {
		m.caves[from] = newNode(from)
	}

	if _, found := m.caves[to]; !found {
		m.caves[to] = newNode(to)
	}

	m.caves[from].Link(m.caves[to])
	m.caves[to].Link(m.caves[from])
}

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
