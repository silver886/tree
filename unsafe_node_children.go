package tree

func (n *Node) unsafeAddChildren(list []*Node) {
	for _, v := range list {
		v.unsafeSetParent(n)
	}
}

func (n *Node) unsafeRemoveChildren(list []*Node) {
	for _, v := range list {
		v.unsafeRemoveParent()
	}
}
