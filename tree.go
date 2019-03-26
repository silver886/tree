package tree

import (
	"errors"
)

// Tree is the structure of tree
type Tree struct {
	roots []*Node
	Style *Style
}

// GetRoots return the root node list of certain node
func (t *Tree) GetRoots() ([]*Node, error) {
	if len(t.roots) == 0 {
		return nil, errors.New("No root nodes found")
	}
	return t.roots, nil
}

func (t *Tree) unsafeAddRoots(nodes []*Node) {
	for _, v := range nodes {
		v.unsafeSetTree(t)
	}
}

// AddRoots add the root node list of certain node
func (t *Tree) AddRoots(nodes []*Node) error {
	for _, v := range nodes {
		if v.tree != nil {
			return errors.New("Some nodes in the node list already have a tree")
		}
	}

	t.unsafeAddRoots(nodes)

	return nil
}

// RemoveRoots remove the root node list of certain node
func (t *Tree) RemoveRoots(nodes []*Node) error {
	for _, v := range nodes {
		if v.tree != t {
			return errors.New("Some nodes in the node list do not belong to this tree")
		}
	}

	t.unsafeRemoveRoots(nodes)

	return nil
}
