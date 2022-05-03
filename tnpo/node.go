package tnpo

import "fmt"

type Node struct {
	data        int
	predecessor *Node
	successors  map[int]*Node
}

func NewNode(data int) *Node {
	return &Node{
		data:        data,
		predecessor: nil,
		successors:  make(map[int]*Node)}
}

func (node *Node) Successors() map[int]*Node {
	return node.successors
}

func (node *Node) Data() int {
	return node.data
}

func (node *Node) Clear() {
	node.predecessor = nil
	node.successors = map[int]*Node{}
}

func (node *Node) RemoveSuccessor(successor Node) *Node {
	val, exist := node.successors[successor.data]
	if exist {
		delete(node.successors, successor.data)
	}
	return val
}

func (node *Node) AddSuccessor(successor *Node) {
	successor.predecessor = node
	node.successors[successor.data] = successor
}

func (node *Node) CreateSuccessor(successorData int) *Node {
	successor := NewNode(successorData)
	successor.predecessor = node
	node.successors[successor.data] = successor
	return successor
}

func (node *Node) TraceUp() []*Node {
	el := node
	trace := []*Node{}
	for el != nil {
		trace = append(trace, el)
		el = el.predecessor
	}
	return trace
}

func (node *Node) DataTraceUp() []int {
	el := node
	trace := []int{}
	for el != nil {
		trace = append(trace, el.data)
		el = el.predecessor
	}
	return trace
}

func (node *Node) String() string {
	return fmt.Sprint("Node(", node.data, ")")
}
