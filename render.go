package tree

import "bytes"

// Style define the outlook of the tree
type Style struct {
	Line  string
	Node  string
	End   string
	Space string
}

func (t *Tree) String() string {
	b := &bytes.Buffer{}
	for _, v := range t.roots {
		b.WriteString(v.String())
		b.WriteString(", ")
	}
	return b.String()
}

func (n *Node) String() string {
	b := &bytes.Buffer{}
	b.WriteString(n.content)
	if len(n.children) > 0 {
		b.WriteString(" > ")
		for _, v := range n.children {
			b.WriteString(v.String())
			b.WriteString(", ")
		}
		b.WriteString(" ^ ")
	}
	return b.String()
}
