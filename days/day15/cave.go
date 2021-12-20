package day15

import (
	"container/heap"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/go-aoc/math"
)

const (
	riskModulo = 9
)

// Cave is the map of the cave with a shit ton of chitons.
type Cave struct {
	Chitons map[math.Vector2DInt]*Node
}

// At returns with a chiton at given coordinate.
func (c *Cave) At(position math.Vector2DInt) *Node {
	return c.Chitons[position]
}

// MaxValues returns with the max X and max Y coordinates.
func (c Cave) MaxValues() (int, int) {
	maxx, maxy := 0, 0

	for location := range c.Chitons {
		if location.X > maxx {
			maxx = location.X
		}

		if location.Y > maxy {
			maxy = location.Y
		}
	}

	return maxx, maxy
}

// AddChiton to the cave.
func (c *Cave) AddChiton(value int, coordinate math.Vector2DInt) {
	node := NewNode(value, coordinate)

	c.Chitons[coordinate] = node

	for _, coord := range coordinate.Neighbours() {
		if coord.X != node.Coordinate.X && coord.Y != node.Coordinate.Y {
			continue
		}

		if target, found := c.Chitons[coord]; found {
			target.AddConnection(node)
			node.AddConnection(target)
		}
	}
}

// Route calculates the safest route.
func (c *Cave) Route(fromNode, toNode *Node) NodeGroup {
	queue := Queue{}

	heap.Init(&queue)
	heap.Push(&queue, fromNode)

	history := map[math.Vector2DInt]*Node{}
	scoreMap := map[math.Vector2DInt]int{fromNode.Coordinate: 0}

	for queue.Len() > 0 {
		current, _ := heap.Pop(&queue).(*Node)

		if current.Coordinate == toNode.Coordinate {
			path := []*Node{toNode}
			previousNode := c.At(history[toNode.Coordinate].Coordinate)

			for previousNode != fromNode {
				path = append(path, previousNode)
				previousNode = c.At(history[previousNode.Coordinate].Coordinate)
			}

			return newNodeGroup(path)
		}

		for _, neighbour := range current.Connections {
			score := scoreMap[current.Coordinate]
			localScore := score + neighbour.Value
			nScore := scoreMap[neighbour.Coordinate]

			if nScore == 0 || localScore < nScore {
				history[neighbour.Coordinate] = current
				scoreMap[neighbour.Coordinate] = localScore

				xDistance := toNode.Coordinate.X - neighbour.Coordinate.X
				yDistance := toNode.Coordinate.Y - neighbour.Coordinate.Y

				newNode := &Node{
					Coordinate:  neighbour.Coordinate,
					Value:       neighbour.Value,
					Score:       localScore + xDistance + yDistance,
					Connections: neighbour.Connections,
				}

				heap.Push(&queue, newNode)
			}
		}
	}

	return newNodeGroup([]*Node{})
}

// Show the cave with the path.
func (c Cave) Show(path NodeGroup) {
	maxx, maxy := c.MaxValues()

	for idy := 0; idy <= maxy; idy++ {
		var line string

		for idx := 0; idx <= maxx; idx++ {
			node := c.At(math.Vector2DInt{X: idx, Y: idy})
			format := "%d"

			if path.Contains(node) {
				format = "\x1b[31m%d\033[m"
			}

			line += fmt.Sprintf(format, node.Value)
		}

		logrus.Info(line)
	}
}

// ExtendCave based on part2.
func ExtendCave(cave Cave, times int) Cave {
	newCave := Cave{Chitons: map[math.Vector2DInt]*Node{}}
	maxx, maxy := cave.MaxValues()

	for idy := 0; idy < times; idy++ {
		for idx := 0; idx < times; idx++ {
			offset := idy + idx

			for _, chiton := range cave.Chitons {
				newCoordinates := math.Vector2DInt{
					X: chiton.Coordinate.X + ((maxx + 1) * idx),
					Y: chiton.Coordinate.Y + ((maxy + 1) * idy),
				}

				logrus.Info(newCoordinates)

				newValue := (chiton.Value + offset) % riskModulo
				if newValue == 0 {
					newValue = 9
				}

				newCave.AddChiton(newValue, newCoordinates)
			}
		}
	}

	return newCave
}
