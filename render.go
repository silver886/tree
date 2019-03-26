package tree

import (
	"bytes"
	"errors"
	"io"
)

// Style define the outlook of the tree
type Style struct {
	Node  string
	Line  string
	End   string
	Space string
}

func renderNode(buf *bytes.Buffer, node *Node, style *Style) error {
	for _, v := range node.prefix {
		switch v {
		case 0:
			buf.WriteString(style.Node)
		case 1:
			buf.WriteString(style.Line)
		case 2:
			buf.WriteString(style.End)
		case 3:
			buf.WriteString(style.Space)
		default:
			return errors.New("Invalid prefix")
		}
	}

	buf.WriteString(node.content)
	buf.WriteRune('\n')

	return nil
}

func renderNodeList(buf *bytes.Buffer, list []*Node, style *Style) error {
	for _, v := range list {
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

	if err := renderNode(buf, n, style); err != nil {
		return nil, err
	} else if err := renderNodeList(buf, n.children, style); err != nil {
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

func stringNodeList(list []*Node) io.Reader {
	buf := &bytes.Buffer{}
	for i, v := range list {
		if i > 0 && buf.String()[buf.Len()-2] != '^' {
			buf.WriteString(", ")
		}
		buf.WriteString(v.String())
	}
	return buf
}

func (n *Node) String() string {
	buf := &bytes.Buffer{}
	buf.WriteString(n.content)
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
