package tree

func (t *Tree) unsafeAddRoots(nodes []*Node) {
	for _, v := range nodes {
		v.unsafeSetTree(t)
	}
}

func (t *Tree) unsafeRemoveRoots(nodes []*Node) {
	for _, v := range nodes {
		v.unsafeRemoveTree()
	}
}
