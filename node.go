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

func (n *Node) setPrefix(indent int, prefix byte) error {
	if indent < 1 || indent > len(n.prefix)+1 {
		return errors.New("Invalid indent")
	} else if prefix > 3 {
		return errors.New("Invalid prefix")
	}

	n.unsafeSetPrefix(indent, prefix)

	return nil
}

// GetParent return the parent node of current node
func (n *Node) GetParent() (*Node, error) {
	if n.parent == nil {
		return nil, errors.New("No parent node found")
	}
	return n.parent, nil
}

func (n *Node) unsafeSetParent(node *Node) {
	n.prefix = append([]byte{}, node.prefix...)
	for i := range n.prefix {
		if n.prefix[i]&1 == 0 {
			n.prefix[i]++
		}
	}
	node.children = append(node.children, n)
	n.parent = node
	indent := n.GetIndent()
	n.unsafeSetPrefix(indent, 2)
	if len(node.children) > 1 {
		node.children[len(node.children)-2].unsafeSetPrefix(indent, 0)
	}
}

// SetParent set the parent node of current node
func (n *Node) SetParent(node *Node) error {
	if n.parent != nil {
		return errors.New("Already have a parent node")
	}

	n.unsafeSetParent(node)

	return nil
}

func (n *Node) unsafeRemoveParent() {
	if len(n.parent.children) > 1 && n.prefix[len(n.prefix)-1] == '2' {
		n.parent.children[len(n.parent.children)-2].unsafeSetPrefix(n.GetIndent(), 2)
	}
	for i, v := range n.parent.children {
		if v == n {
			n.parent.children = append(n.parent.children[:i], n.parent.children[i+1:]...)
			continue
		}
	}
	n.prefix = nil
	n.parent = nil
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

func (n *Node) unsafeAddChildren(list []*Node) {
	for _, v := range list {
		v.unsafeSetParent(n)
	}
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

func (n *Node) unsafeRemoveChildren(list []*Node) {
	for _, v := range list {
		v.unsafeRemoveParent()
	}
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

func (n *Node) unsafeAddSiblings(list []*Node) {
	if n.parent != nil {
		n.parent.unsafeAddChildren(list)
	} else if n.tree != nil {
		n.tree.unsafeAddRoots(list)
	}
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

func (n *Node) unsafeRemoveSiblings(list []*Node) {
	if n.parent != nil {
		n.parent.unsafeRemoveChildren(list)
	} else if n.tree != nil {
		n.tree.unsafeRemoveRoots(list)
	}
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

func (n *Node) unsafeSetTree(tree *Tree) {
	tree.roots = append(tree.roots, n)
	n.tree = tree
}

// SetTree set the tree of current node
func (n *Node) SetTree(tree *Tree) error {
	if n.tree != nil {
		return errors.New("Already have a tree")
	}

	n.unsafeSetTree(tree)

	return nil
}

func (n *Node) unsafeRemoveTree() {
	for i, v := range n.tree.roots {
		if v == n {
			n.tree.roots = append(n.tree.roots[:i], n.tree.roots[i+1:]...)
			continue
		}
	}
	n.tree = nil
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
