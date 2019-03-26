package tree

import (
	"bytes"
	"errors"
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
		buf.WriteString(str)
	}
	return nil
}

// Render generate the node structure in string with given style
func (n *Node) Render(style *Style) (string, error) {
	if style == nil {
		return "", errors.New("No style found")
	}

	buf := &bytes.Buffer{}

	if err := renderNode(buf, n, style); err != nil {
		return "", err
	} else if err := renderNodeList(buf, n.children, style); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// Render generate the tree structure in string with its style
func (t *Tree) Render() (string, error) {
	if t.Style == nil {
		return "", errors.New("No style found")
	}

	buf := &bytes.Buffer{}

	if err := renderNodeList(buf, t.roots, t.Style); err != nil {
		return "", err
	}

	return buf.String(), nil
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
