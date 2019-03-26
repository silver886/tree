package tree

func (n *Node) unsafeSetTree(tree *Tree) {
	tree.roots = append(tree.roots, n)
	n.tree = tree
}

func (n *Node) unsafeRemoveTree() {
	for i, v := range n.tree.roots {
		if v == n {
			n.tree.roots = append(n.tree.roots[:i], n.tree.roots[i+1:]...)
			continue
		}
	}
	n.tree = nil
}
