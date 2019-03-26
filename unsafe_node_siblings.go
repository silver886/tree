package tree

func (n *Node) unsafeAddSiblings(nodes []*Node) {
	if n.parent != nil {
		n.parent.unsafeAddChildren(nodes)
	} else if n.tree != nil {
		n.tree.unsafeAddRoots(nodes)
	}
}

func (n *Node) unsafeRemoveSiblings(nodes []*Node) {
	if n.parent != nil {
		n.parent.unsafeRemoveChildren(nodes)
	} else if n.tree != nil {
		n.tree.unsafeRemoveRoots(nodes)
	}
}
