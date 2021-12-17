package day12

import (
	"strings"
)

type Node struct {
	name        string
	multi       bool
	visited     int
	connections map[string]*Node
}

func newNode(name string) *Node {
	multi := false
	for _, ch := range name {
		if ch >= 'A' && ch <= 'Z' {
			multi = true

			break
		}
	}

	return &Node{
		name:        name,
		visited:     0,
		multi:       multi,
		connections: map[string]*Node{},
	}
}

func (n Node) Name() string {
	return n.name
}

func (n Node) IsLarge() bool {
	return n.multi
}

func (n *Node) Link(node *Node) {
	if _, found := n.connections[node.Name()]; !found {
		n.connections[node.Name()] = node
	}
}

func (n Node) Options(chain []*Node, haveMoreTime bool) []*Node {
	options := []*Node{}

	for _, node := range n.connections {
		if node.Name() == "start" {
			continue
		}

		onChain := isOnChain(node, chain)
		canVisitMore := haveMoreTime && !chainHasSmallDuplicates(chain)

		if node.IsLarge() || canVisitMore || !onChain {
			options = append(options, node)
		}
	}

	return options
}

func isOnChain(node *Node, chain []*Node) bool {
	for _, visited := range chain {
		if visited == node {
			return true
		}
	}

	return false
}

func chainToString(chain []*Node) string {
	items := []string{}

	for _, node := range chain {
		items = append(items, node.Name())
	}

	return strings.Join(items, "->")
}

func chainHasSmallDuplicates(chain []*Node) bool {
	visited := map[string]bool{}

	for _, node := range chain {
		if node.IsLarge() {
			continue
		}

		if _, found := visited[node.Name()]; found {
			return true
		}

		visited[node.Name()] = true
	}

	return false
}
