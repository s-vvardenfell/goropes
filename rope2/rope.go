package rope2

import (
	"errors"
	"strings"
	"unicode/utf8"
)

type GoRope struct {
	Root     *Node
	RawData  []rune
	Size     int
	LeafSize int
}

func NewGoRopeFromString(inStr string, leafSize int) (*GoRope, error) {
	if inStr == "" {
		return nil, errors.New("argument 'inStr' must not be an empty string")
	}
	if leafSize <= 0 {
		return nil, errors.New("leafSize must be positive integer")
	}

	rawData := make([]rune, 0, utf8.RuneCountInString(inStr))
	for _, r := range inStr {
		rawData = append(rawData, r)
	}

	r := &GoRope{
		RawData:  rawData,
		LeafSize: leafSize,
	}

	r.Root = r.createRope(0, len(rawData)-1)
	return r, nil
}

func (r *GoRope) createRope(left, right int) *Node {
	if left+r.LeafSize > right {
		r.Size += right - left + 1
		return NewNode(right-left+1, r.RawData[left:right+1])
	}

	mid := (left + right) / 2
	root := NewNode(mid-left+1, nil)
	root.Left = r.createRope(left, mid)
	root.Right = r.createRope(mid+1, right)
	return root
}

func (r *GoRope) String() string {
	traversal := make([]string, 0)
	inOrderRecursive(r.Root, &traversal)
	sb := strings.Builder{}
	for i := range traversal {
		sb.WriteString(traversal[i])
	}
	return sb.String()
}

func inOrderRecursive(n *Node, tr *[]string) {
	if n != nil {
		inOrderRecursive(n.Left, tr)
		*tr = append(*tr, string(n.Data))
		inOrderRecursive(n.Right, tr)
	}
}
