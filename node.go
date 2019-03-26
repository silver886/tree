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

// GetChildren return the child node list of current node
func (n *Node) GetChildren() ([]*Node, error) {
	if len(n.children) == 0 {
		return nil, errors.New("No child nodes found")
	}
	return n.children, nil
}

// AddChildren add the child node list of current node
func (n *Node) AddChildren(list []*Node) error {
	for _, v := range list {
		if v.parent != nil {
			return errors.New("Some nodes in the node list already have a parent node")
		}
	}

	n.unsafeAddChildren(list)

	return nil
}

// RemoveChildren remove the child node list of current node
func (n *Node) RemoveChildren(list []*Node) error {
	for _, v := range list {
		if v.parent != n {
			return errors.New("Some nodes in the node list do not belong to this parent node")
		}
	}

	n.unsafeRemoveChildren(list)

	return nil
}

// GetSiblings return the sibling node list of current node
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

// AddSiblings add the sibling node list of current node
func (n *Node) AddSiblings(list []*Node) error {
	if n.parent == nil && n.tree == nil {
		return errors.New("No available sibling node list")
	}

	for _, v := range list {
		if v.parent != nil {
			return errors.New("Some nodes in the node list already have a parent node")
		}
	}

	n.unsafeAddSiblings(list)

	return nil
}

// RemoveSiblings remove the sibling node list of current node
func (n *Node) RemoveSiblings(list []*Node) error {
	if n.parent == nil && n.tree == nil {
		return errors.New("No available sibling node list")
	}

	for _, v := range list {
		if v.parent != n {
			return errors.New("Some nodes in the node list do not belong to the parent node of this node")
		}
	}

	n.unsafeRemoveSiblings(list)

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

// GetIndent find the indentation of current node
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
