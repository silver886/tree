package tree

import (
	"bytes"
	"errors"
	"io"
)

func renderNode(buf *bytes.Buffer, node *Node, indent int, style *Style) {
	buf.WriteString(style.getPrefix(node.prefix[indent:]))
	buf.WriteString(node.Content)
	buf.WriteByte('\n')
}

func renderNodeList(nodeList []*Node, indent int, style *Style) io.Reader {
	buf := &bytes.Buffer{}
	for _, v := range nodeList {
		renderNode(buf, v, indent, style)
	}
	return buf
}

// Render generate the node structure in string with given style
func (n *Node) Render(style *Style) (io.Reader, error) {
	if style == nil {
		return nil, errors.New("No style found")
	}

	return renderNodeList(n.GetNodeList(), len(n.prefix), style), nil
}

// Render generate the tree structure in string with its style
func (t *Tree) Render() (io.Reader, error) {
	if t.Style == nil {
		return nil, errors.New("No style found")
	}

	nodeList, err := t.GetNodeList()
	if err != nil {
		return nil, err
	}

	return renderNodeList(nodeList, 0, t.Style), nil
}
