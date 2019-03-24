package tree

import (
	"bufio"
	"bytes"
	"errors"
	"io"
)

type line struct {
	indent  int
	content string
}

func generateLineList(r io.Reader) ([]*line, error) {
	var lineList []*line
	buf := bufio.NewReader(r)
	lineBuf := &bytes.Buffer{}
	for {
		if cnt, isPrefix, err := buf.ReadLine(); err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		} else if isPrefix {
			if _, err := lineBuf.Write(cnt); err != nil {
				return nil, err
			}
		} else if lineBuf.Len() > 0 {
			lineList = appendLineList(lineList, lineBuf.String())
			lineBuf.Reset()
		} else {
			lineList = appendLineList(lineList, string(cnt))
		}
	}

	lineList, err := generalizeIndent(lineList)
	if err != nil {
		return nil, err
	}

	return lineList, nil
}

func appendLineList(list []*line, str string) []*line {
	if indent, content := parseLine(str); len(content) != 0 {
		list = append(list, &line{
			indent:  indent,
			content: content,
		})
	}
	return list
}

func parseLine(str string) (int, string) {
	for i, v := range str {
		if v != ' ' && v != '\t' && v != 0xA0 {
			return i, str[i:]
		}
	}
	return 0, ""
}

func generalizeIndent(list []*line) ([]*line, error) {
	var indents []int
	for _, v := range list {
		indents = append(indents, v.indent)
	}

	indent, err := calculateIndent(indents)
	if err != nil {
		return nil, err
	}

	for i := range list {
		list[i].indent = list[i].indent / indent
	}

	return list, nil
}

func calculateIndent(indents []int) (int, error) {
	indents = uniqueIndents(indents)

	min := indents[0]
	for _, v := range indents {
		if v < min {
			min = v
		}
	}

	for _, v := range indents {
		if v%min != 0 {
			return 0, errors.New("Indentation is not the same")
		}
	}

	return min, nil
}

func uniqueIndents(indents []int) []int {
	keys := make(map[int]bool)
	uniqueIndents := []int{}
	for _, v := range indents {
		if v == 0 {
			continue
		} else if _, ok := keys[v]; !ok {
			keys[v] = true
			uniqueIndents = append(uniqueIndents, v)
		}
	}
	return uniqueIndents
}
