package tree

import (
	"errors"
	"io"
)

// Parse create a tree from certain reader
func Parse(r io.Reader) *Tree {
	return nil
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

	var tempNodes []*Node

	for i, v := range lineList {
		n := &Node{
			Content: v.content,
		}
		if len(tempNodes) == 0 {
			if tree.AddRoots([]*Node{n}) != nil {
				return nil, errors.New("Cannot add a root")
			}
			tempNodes = append(tempNodes, n)
		} else {
			switch indent := v.indent - lineList[i-1].indent; indent {
			case 1:
				if tempNodes[len(tempNodes)-1].AddChildren([]*Node{n}) != nil {
					return nil, errors.New("Cannot add a child")
				}
				tempNodes = append(tempNodes, n)
			default:
				if indent > 0 {
					return nil, errors.New("Abnormal indentation")
				}
				tempNodes = tempNodes[:len(tempNodes)+indent]
				fallthrough
			case 0:
				if tempNodes[len(tempNodes)-1].AddSiblings([]*Node{n}) != nil {
					return nil, errors.New("Cannot add a sibling")
				}
				tempNodes[len(tempNodes)-1] = n
			}
		}
	}

	return tree, nil
}
