package day15

import (
	"github.com/yitsushi/go-aoc/slice"
)

// NodeGroup is a group of nodes, used to show path.
type NodeGroup struct {
	Chain []*Node
	Sum   int
}

func newNodeGroup(group []*Node) NodeGroup {
	sum := 0

	for _, node := range group {
		sum += node.Value
	}

	return NodeGroup{
		Chain: group,
		Sum:   sum,
	}
}

// Path is the string representation of a path.
func (g *NodeGroup) Path() string {
	path := []int{}

	for _, node := range g.Chain {
		path = append(path, node.Value)
	}

	return slice.JoinInt(path, "-")
}

// Value is the sum of all the risk values.
func (g NodeGroup) Value() int {
	return g.Sum
}

// Contains checks if a node is in the node group.
func (g *NodeGroup) Contains(node *Node) bool {
	for _, target := range g.Chain {
		if target == node {
			return true
		}
	}

	return false
}
