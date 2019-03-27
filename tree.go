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

// GetNodeList return the node list from current node
func (t *Tree) GetNodeList() ([]*Node, error) {
	roots, err := t.GetRoots()
	if err != nil {
		return nil, err
	}

	var nodeList []*Node
	for _, v := range roots {
		nodeList = append(nodeList, v.GetNodeList()...)
	}

	return nodeList, nil
}

// GetPrefixList return the prefix list
func (t *Tree) GetPrefixList() ([]string, error) {
	if t.Style == nil {
		return nil, errors.New("No style found")
	}

	nodeList, err := t.GetNodeList()
	if err != nil {
		return nil, err
	}

	var prefixList []string
	for _, v := range nodeList {
		prefixList = append(prefixList, t.Style.getPrefix(v.prefix))
	}

	return prefixList, nil
}
