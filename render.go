package tree

import (
	"bytes"
	"errors"
	"io"
)

func renderNode(buf *bytes.Buffer, node *Node, style *Style) {
	for _, v := range node.prefix {
		buf.WriteString(style.getPrefix(v))
	}

	buf.WriteString(node.content)
	buf.WriteRune('\n')
}

func renderNodeList(buf *bytes.Buffer, nodeList []*Node, style *Style) error {
	for _, v := range nodeList {
		str, err := v.Render(style)
		if err != nil {
			return err
		}
		buf.ReadFrom(str)
	}
	return nil
}

// Render generate the node structure in string with given style
func (n *Node) Render(style *Style) (io.Reader, error) {
	if style == nil {
		return nil, errors.New("No style found")
	}

	buf := &bytes.Buffer{}
	renderNode(buf, n, style)

	if err := renderNodeList(buf, n.children, style); err != nil {
		return nil, err
	}

	return buf, nil
}

// Render generate the tree structure in string with its style
func (t *Tree) Render() (io.Reader, error) {
	if t.Style == nil {
		return nil, errors.New("No style found")
	}

	buf := &bytes.Buffer{}

	if err := renderNodeList(buf, t.roots, t.Style); err != nil {
		return nil, err
	}

	return buf, nil
}
