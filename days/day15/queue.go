package day15

// Queue is a simple queue.
type Queue []*Node

// Len is the length of the queue.
func (q Queue) Len() int {
	return len(q)
}

// Less implements Less for the heap interface.
func (q Queue) Less(i, j int) bool {
	return q[i].Score < q[j].Score
}

// Swap implements Swap for the heap interface.
func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

// Push implements Push for the heap interface.
func (q *Queue) Push(node interface{}) {
	*q = append(*q, node.(*Node))
}

// Pop implements Pop for the heap interface.
func (q *Queue) Pop() interface{} {
	original, length := *q, q.Len()
	node := original[length-1]
	*q = original[:length-1]

	return node
}
