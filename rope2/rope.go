package rope2

import (
	"errors"
	"strings"
	"unicode/utf8"
)

type GoRope struct {
	root     *Node
	rawData  []rune
	size     int
	leafSize int
}

func NewGoRopeFromString(in string, leafSize int) (*GoRope, error) {
	if in == "" {
		return nil, errors.New("argument 'in' must not be an empty string")
	}
	if leafSize <= 0 {
		return nil, errors.New("leafSize must be positive integer")
	}

	rawData := make([]rune, 0, utf8.RuneCountInString(in))
	for _, r := range in {
		rawData = append(rawData, r)
	}

	r := &GoRope{
		rawData:  rawData,
		leafSize: leafSize,
	}

	r.root = r.createRope(0, len(rawData)-1)
	return r, nil
}

func (r *GoRope) createRope(left, right int) *Node {
	if left+r.leafSize > right {
		r.size += right - left + 1
		return NewNode(right-left+1, r.rawData[left:right+1])
	}

	mid := (left + right) / 2
	root := NewNode(mid-left+1, nil)
	root.Left = r.createRope(left, mid)
	root.Right = r.createRope(mid+1, right)
	return root
}

func (r *GoRope) String() string {
	traversal := make([]string, 0)
	inOrderRecursive(r.root, &traversal)
	sb := strings.Builder{}
	for i := range traversal {
		sb.WriteString(traversal[i])
	}
	return sb.String()
}

func inOrderRecursive(n *Node, tr *[]string) {
	if n != nil {
		if n.Data != nil {
			*tr = append(*tr, string(n.Data))
		}
		inOrderRecursive(n.Left, tr)
		inOrderRecursive(n.Right, tr)
	}
}

// Traverse the three and count nodes weight
// Should give the same value as GoRope.Size() field
func (r *GoRope) TotalWeight(root *Node) int {
	if root == nil {
		root = r.root
	}

	head := root
	count := 0

	for head != nil {
		count += head.Weight
		if head.Data != nil {
			break
		}
		head = head.Right
	}
	return count
}

// util for Index()
func index(n *Node, idx int) rune {
	if n.Weight <= idx && n.Right != nil {
		return index(n.Right, idx-n.Weight)
	}

	if n.Left != nil {
		return index(n.Left, idx)
	}

	return n.Data[idx]
}

// Returns a character at the specified index
func (r *GoRope) Index(idx int) (string, error) {
	if idx > r.size-1 {
		return "", errors.New("index out of range")
	}
	return string(index(r.root, idx)), nil
}

// Returns a copy of GoRope underlying data
func (r *GoRope) RawData() []rune {
	temp := make([]rune, len(r.rawData))
	_ = copy(temp, r.rawData)
	return temp
}

// Returns amoint of characters
func (r *GoRope) Size() int {
	return r.size
}

// Returns leaf size
func (r *GoRope) LeafSize() int {
	return r.leafSize
}

// func (r *GoRope) Root() *Node {
// 	return r.root
// }
