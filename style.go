package tree

// Style define the outlook of the tree
type Style struct {
	Node  string
	Line  string
	End   string
	Space string
}

func (s *Style) getPrefix(b byte) string {
	switch b {
	case 0:
		return s.Node
	case 1:
		return s.Line
	case 2:
		return s.End
	case 3:
		return s.Space
	default:
		return ""
	}
}
