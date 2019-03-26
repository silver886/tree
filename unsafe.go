package tree

func (n *Node) unsafeSetPrefix(indent int, prefix byte) {
	for _, v := range n.children {
		v.unsafeSetPrefix(indent, prefix|1)
	}
	if indent--; indent == len(n.prefix) {
		n.prefix = append(n.prefix, prefix)
	} else {
		n.prefix[indent] = prefix
	}
}

func (n *Node) unsafeSetParent(node *Node) {
	n.prefix = append([]byte{}, node.prefix...)
	for i := range n.prefix {
		if n.prefix[i]&1 == 0 {
			n.prefix[i]++
		}
	}
	node.children = append(node.children, n)
	n.parent = node
	indent := n.GetIndent()
	n.unsafeSetPrefix(indent, 2)
	if len(node.children) > 1 {
		node.children[len(node.children)-2].unsafeSetPrefix(indent, 0)
	}
}

func (n *Node) unsafeRemoveParent() {
	if len(n.parent.children) > 1 && n.prefix[len(n.prefix)-1] == '2' {
		n.parent.children[len(n.parent.children)-2].unsafeSetPrefix(n.GetIndent(), 2)
	}
	for i, v := range n.parent.children {
		if v == n {
			n.parent.children = append(n.parent.children[:i], n.parent.children[i+1:]...)
			continue
		}
	}
	n.prefix = nil
	n.parent = nil
}

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

func (t *Tree) unsafeRemoveRoots(list []*Node) {
	for _, v := range list {
		v.unsafeRemoveTree()
	}
}
