package tree

import "bytes"

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

func (n *Node) unsafeGetPrefix(style *Style) string {
	buf := &bytes.Buffer{}
	for _, v := range n.prefix {
		buf.WriteString(style.getPrefix(v))
	}
	return buf.String()
}
