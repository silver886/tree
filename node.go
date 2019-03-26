package tree

import (
	"errors"
)

// Node is the structure of node
type Node struct {
	tree *Tree

	parent   *Node
	children []*Node

	prefix  []byte
	content string
}

// GetParent return the parent node of current node
func (n *Node) GetParent() (*Node, error) {
	if n.parent == nil {
		return nil, errors.New("No parent node found")
	}
	return n.parent, nil
}

// SetParent set the parent node of current node
func (n *Node) SetParent(node *Node) error {
	if n.parent != nil {
		return errors.New("Already have a parent node")
	}

	n.unsafeSetParent(node)

	return nil
}

// RemoveParent remove the parent node of current node
func (n *Node) RemoveParent() error {
	if n.parent == nil {
		return errors.New("No parent node to remove")
	}

	n.unsafeRemoveParent()

	return nil
}

// GetChildren return the child nodes of current node
func (n *Node) GetChildren() ([]*Node, error) {
	if len(n.children) == 0 {
		return nil, errors.New("No child nodes found")
	}
	return n.children, nil
}

// AddChildren add the child nodes of current node
func (n *Node) AddChildren(nodes []*Node) error {
	for _, v := range nodes {
		if v.parent != nil {
			return errors.New("Some nodes already have a parent node")
		}
	}

	n.unsafeAddChildren(nodes)

	return nil
}

// RemoveChildren remove the child nodes of current node
func (n *Node) RemoveChildren(nodes []*Node) error {
	for _, v := range nodes {
		if v.parent != n {
			return errors.New("Some nodes do not belong to this parent node")
		}
	}

	n.unsafeRemoveChildren(nodes)

	return nil
}

// GetSiblings return the sibling nodes of current node
func (n *Node) GetSiblings() ([]*Node, error) {
	if n.parent == nil {
		if n.tree == nil || len(n.tree.roots) == 0 {
			return nil, errors.New("No sibling nodes found")
		}
		return n.tree.roots, nil
	}

	if len(n.parent.children) == 0 {
		return nil, errors.New("No sibling nodes found")
	}
	return n.parent.children, nil
}

// AddSiblings add the sibling nodes of current node
func (n *Node) AddSiblings(nodes []*Node) error {
	if n.parent == nil && n.tree == nil {
		return errors.New("No available sibling nodes")
	}

	for _, v := range nodes {
		if v.parent != nil {
			return errors.New("Some nodes in the nodes already have a parent node")
		}
	}

	n.unsafeAddSiblings(nodes)

	return nil
}

// RemoveSiblings remove the sibling nodes of current node
func (n *Node) RemoveSiblings(nodes []*Node) error {
	if n.parent == nil && n.tree == nil {
		return errors.New("No available sibling nodes")
	}

	for _, v := range nodes {
		if v.parent != n {
			return errors.New("Some nodes in the nodes do not belong to the parent node of this node")
		}
	}

	n.unsafeRemoveSiblings(nodes)

	return nil
}

// GetTree return the tree of current node
func (n *Node) GetTree() (*Tree, error) {
	if n.tree == nil {
		return nil, errors.New("No tree found")
	}
	return n.tree, nil
}

// SetTree set the tree of current node
func (n *Node) SetTree(tree *Tree) error {
	if n.tree != nil {
		return errors.New("Already have a tree")
	}

	n.unsafeSetTree(tree)

	return nil
}

// RemoveTree remove the tree of current node
func (n *Node) RemoveTree() error {
	if n.tree == nil {
		return errors.New("No tree to remove")
	}

	n.unsafeRemoveTree()

	return nil
}

// GetIndent return the indentation of current node
func (n *Node) GetIndent() int {
	for i, tempNode := 0, n; ; i++ {
		if n, err := tempNode.GetParent(); err == nil {
			tempNode = n
		} else {
			return i
		}
	}
}

func (n *Node) setPrefix(indent int, prefix byte) error {
	if indent < 1 || indent > len(n.prefix)+1 {
		return errors.New("Invalid indent")
	} else if prefix > 3 {
		return errors.New("Invalid prefix")
	}

	n.unsafeSetPrefix(indent, prefix)

	return nil
}
