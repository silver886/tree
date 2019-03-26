package tree

import "errors"

// Style define the outlook of the tree
type Style struct {
	Node  string
	Line  string
	End   string
	Space string
}

func (s *Style) getPrefix(b byte) (string, error) {
	switch b {
	case 0:
		return s.Node, nil
	case 1:
		return s.Line, nil
	case 2:
		return s.End, nil
	case 3:
		return s.Space, nil
	default:
		return "", errors.New("Invalid prefix")
	}
}
