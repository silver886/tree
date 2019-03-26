package tree

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
