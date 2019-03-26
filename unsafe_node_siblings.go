package tree

func (n *Node) unsafeAddSiblings(list []*Node) {
	if n.parent != nil {
		n.parent.unsafeAddChildren(list)
	} else if n.tree != nil {
		n.tree.unsafeAddRoots(list)
	}
}

func (n *Node) unsafeRemoveSiblings(list []*Node) {
	if n.parent != nil {
		n.parent.unsafeRemoveChildren(list)
	} else if n.tree != nil {
		n.tree.unsafeRemoveRoots(list)
	}
}
