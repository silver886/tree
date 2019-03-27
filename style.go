package tree

import "bytes"

// Style define the outlook of the tree
type Style struct {
	Node  string
	Line  string
	End   string
	Space string
}

func (s *Style) getPrefix(indents []byte) string {
	buf := &bytes.Buffer{}
	for _, v := range indents {
		switch v {
		case 0:
			buf.WriteString(s.Node)
		case 1:
			buf.WriteString(s.Line)
		case 2:
			buf.WriteString(s.End)
		case 3:
			buf.WriteString(s.Space)
		default:
			buf.Reset()
			break
		}
	}
	return buf.String()
}
