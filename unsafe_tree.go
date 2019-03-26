package tree

func (t *Tree) unsafeRemoveRoots(list []*Node) {
	for _, v := range list {
		v.unsafeRemoveTree()
	}
}
