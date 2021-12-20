package day15

import "github.com/yitsushi/go-aoc/math"

// Node is a Chiton.
type Node struct {
	Value       int
	Coordinate  math.Vector2DInt
	Connections []*Node
	Score       int
}

// NewNode creates a new Node.
func NewNode(value int, coordinate math.Vector2DInt) *Node {
	return &Node{
		Value:       value,
		Connections: []*Node{},
		Coordinate:  coordinate,
		Score:       0,
	}
}

// AddConnection to the Node.
func (n *Node) AddConnection(node *Node) {
	n.Connections = append(n.Connections, node)
}
