package tree

func (n *Node) unsafeAddChildren(nodes []*Node) {
	for _, v := range nodes {
		v.unsafeSetParent(n)
	}
}

func (n *Node) unsafeRemoveChildren(nodes []*Node) {
	for _, v := range nodes {
		v.unsafeRemoveParent()
	}
}
