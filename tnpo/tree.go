package tnpo

type Tree struct {
	root    *Node
	nodeSet map[int]*Node
}

func NewTree() *Tree {
	return &Tree{nil, map[int]*Node{}}
}

func (tree *Tree) GetNode(n int) *Node {
	if node, exist := tree.nodeSet[n]; exist {
		return node
	}
	return nil
}

func (tree *Tree) Contains(n int) bool {
	_, exits := tree.nodeSet[n]
	return exits
}

func (tree *Tree) CreateNode(n int, parent *Node) *Node {
	node := NewNode(n)
	if tree.AddNode(node, parent) {
		// fmt.Println("CreateNode", n)
		return node
	}
	return nil
}

func (tree *Tree) AddNode(node *Node, parent *Node) bool {
	// fmt.Println("AddNode", node, parent, tree.root)
	if tree.root == nil && parent == nil {
		tree.root = node
		tree.nodeSet = map[int]*Node{node.data: node}
		return true
	}
	if tree.root != nil && parent == nil {
		return false
	}
	if _, exits := tree.nodeSet[node.data]; exits {
		return false
	}
	if _, exits := tree.nodeSet[parent.data]; !exits {
		return false
	}
	parent.AddSuccessor(node)
	tree.nodeSet[node.data] = node
	return true
}

func (tree *Tree) Clear() {
	tree.root = nil
	tree.nodeSet = map[int]*Node{}
}
