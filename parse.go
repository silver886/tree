package tree

import (
	"errors"
	"io"
)

// ParseIndentList create a tree from certain reader with preset style
func ParseIndentList(indentList []int) (*Tree, error) {
	if indentList == nil {
		return nil, errors.New("Empty indent list")
	}
	if indentList[0] != 0 {
		return nil, errors.New("Tree must start with root")
	}

	tree := &Tree{}
	for _, v := range indentList {
		if err := tree.AddNode(v, &Node{}); err != nil {
			return nil, err
		}
	}

	return tree, nil
}

// ParseIndent create a tree from certain reader with indent
func ParseIndent(r io.Reader) (*Tree, error) {
	lineList, err := generateLineList(r)
	if err != nil {
		return nil, err
	}

	if lineList[0].indent != 0 {
		return nil, errors.New("Tree must start with root")
	}

	tree := &Tree{}
	for _, v := range lineList {
		if err := tree.AddNode(v.indent, &Node{
			Content: v.content,
		}); err != nil {
			return nil, err
		}
	}

	return tree, nil
}
