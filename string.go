package tree

import (
	"bytes"
	"io"
)

func stringNodeList(nodeList []*Node) io.Reader {
	buf := &bytes.Buffer{}
	for i, v := range nodeList {
		if i > 0 && buf.String()[buf.Len()-2] != '^' {
			buf.WriteString(", ")
		}
		buf.WriteString(v.String())
	}
	return buf
}

func (n *Node) String() string {
	buf := &bytes.Buffer{}
	buf.WriteString(n.Content)
	if len(n.children) > 0 {
		buf.WriteString(" > ")
		buf.ReadFrom(stringNodeList(n.children))
		buf.WriteString(" ^ ")
	}
	return buf.String()
}

func (t *Tree) String() string {
	buf := &bytes.Buffer{}
	buf.ReadFrom(stringNodeList(t.roots))
	return buf.String()
}
