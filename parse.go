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

	var nodeTempList []*Node

	for i, v := range lineList {
		n := &Node{
			content: v.content,
		}
		if len(nodeTempList) == 0 {
			if tree.AddRoots([]*Node{n}) != nil {
				return nil, errors.New("Cannot add a root")
			}
			nodeTempList = append(nodeTempList, n)
		} else {
			switch indent := v.indent - lineList[i-1].indent; indent {
			case 1:
				if nodeTempList[len(nodeTempList)-1].AddChildren([]*Node{n}) != nil {
					return nil, errors.New("Cannot add a child")
				}
				nodeTempList = append(nodeTempList, n)
			default:
				if indent > 0 {
					return nil, errors.New("Abnormal indentation")
				}
				nodeTempList = nodeTempList[:len(nodeTempList)+indent]
				fallthrough
			case 0:
				if nodeTempList[len(nodeTempList)-1].AddSiblings([]*Node{n}) != nil {
					return nil, errors.New("Cannot add a sibling")
				}
				nodeTempList[len(nodeTempList)-1] = n
			}
		}
	}

	return tree, nil
}
