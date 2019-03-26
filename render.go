package tree

import "bytes"

// Style define the outlook of the tree
type Style struct {
	Node  string
	Line  string
	End   string
	Space string
}

func (t *Tree) String() string {
	b := &bytes.Buffer{}
	for i, v := range t.roots {
		if i > 0 && b.String()[b.Len()-2] != '^' {
			b.WriteString(", ")
		}
		b.WriteString(v.String())
	}
	return b.String()
}

func (n *Node) String() string {
	b := &bytes.Buffer{}
	b.WriteString(n.content)
	if len(n.children) > 0 {
		b.WriteString(" > ")
		for i, v := range n.children {
			if i > 0 && b.String()[b.Len()-2] != '^' {
				b.WriteString(", ")
			}
			b.WriteString(v.String())
		}
		b.WriteString(" ^ ")
	}
	return b.String()
}
