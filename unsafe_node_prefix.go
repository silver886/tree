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
