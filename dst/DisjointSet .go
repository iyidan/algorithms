package dst

// DisjointSetTreeNode 并查集树节点
type DisjointSetTreeNode struct {
	value  string
	parent *DisjointSetTreeNode
	rank   int
}

// Equal check two node equal
func (node *DisjointSetTreeNode) Equal(node1 *DisjointSetTreeNode) bool {
	if node.value == node1.value {
		return true
	}
	return false
}

// DisjointSetFind find the given node's root node
func DisjointSetFind(node *DisjointSetTreeNode) *DisjointSetTreeNode {
	if node.parent == nil || node.Equal(node.parent) {
		return node
	}
	if !node.parent.Equal(node) {
		node.parent = DisjointSetFind(node.parent)
	}
	return node.parent
}

// DisjointSetUnion union node1's tree and node2's tree into one tree
func DisjointSetUnion(node1 *DisjointSetTreeNode, node2 *DisjointSetTreeNode) {
	root1 := DisjointSetFind(node1)
	root2 := DisjointSetFind(node2)
	if root1.Equal(root2) {
		return
	}
	if root1.rank < root2.rank {
		root1.parent = root2
	} else if root1.rank > root2.rank {
		root2.parent = root1
	} else {
		root2.parent = root1
		root1.rank++
	}
}
