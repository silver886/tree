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

func (t *Tree) unsafeAddRoots(list []*Node) {
	for _, v := range list {
		v.unsafeSetTree(t)
	}
}

// AddRoots add the root node list of certain node
func (t *Tree) AddRoots(list []*Node) error {
	for _, v := range list {
		if v.tree != nil {
			return errors.New("Some nodes in the node list already have a tree")
		}
	}

	t.unsafeAddRoots(list)

	return nil
}

// RemoveRoots remove the root node list of certain node
func (t *Tree) RemoveRoots(list []*Node) error {
	for _, v := range list {
		if v.tree != t {
			return errors.New("Some nodes in the node list do not belong to this tree")
		}
	}

	t.unsafeRemoveRoots(list)

	return nil
}
