package tree

func (t *Tree) unsafeRemoveRoots(nodes []*Node) {
	for _, v := range nodes {
		v.unsafeRemoveTree()
	}
}
