package rope

import (
	"errors"
	"io"
	"strings"
)

type Rope struct {
	Root     *RopeNode
	raw      string //should store []rune to work with non-ascii
	size     int
	leafsize int
}

func (r *Rope) createRope(lo, hi int) *RopeNode {
	if lo+r.leafsize > hi {
		r.size += hi - lo - 1
		return NewRopeNode(hi-lo+1, r.raw[lo:hi+1])
	}

	mid := (lo + hi) / 2
	root := NewRopeNode(mid-lo+1, "")
	root.Left = r.createRope(lo, mid)
	root.Right = r.createRope(mid+1, hi)
	return root
}

func NewRopeFromString(raw string, leafsize int) (*Rope, error) {
	if raw == "" {
		return nil, errors.New("argument 'raw' must not be an empty string")
	}
	if leafsize <= 0 {
		return nil, errors.New("leafsize must be positive integer")
	}

	var r Rope
	r.raw = raw
	r.leafsize = leafsize
	r.Root = r.createRope(0, len(raw)-1)
	return &r, nil
}

func NewRopeFromReader(r io.Reader, leafsize int) (*Rope, error) {
	buf, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return NewRopeFromString(string(buf), leafsize)
}

// func NewRopeFromSlice(val []string, leafsize int) (*Rope, error) {
// 	sb := strings.Builder{}
// 	for i := range val {
// 		sb.WriteString(val[i])
// 	}
// 	return NewRopeFromString(sb.String(), leafsize)
// }

func inOrderRecursive(n *RopeNode, tr *[]string) {
	if n != nil {
		inOrderRecursive(n.Left, tr)
		*tr = append(*tr, n.Val)
		inOrderRecursive(n.Right, tr)
	}
}

func (r *Rope) String() string {
	traversal := make([]string, 0)
	inOrderRecursive(r.Root, &traversal)
	sb := strings.Builder{}
	for i := range traversal {
		sb.WriteString(traversal[i])
	}
	return sb.String()
}

//TODO
// func (r *Rope) update() {

// }

// //TODO
// func (r *Rope) SetLeafSize(val int) error {
// 	if val == r.leafsize {
// 		return fmt.Errorf("rope already has leafsize %d", val)
// 	}
// 	if val < 0 {
// 		return errors.New("leafsize must be positive integer")
// 	}
// 	if val < 4 {
// 		logrus.Warning(
// 			"Smaller leafsizes will cause performance issues. Use big integers like 8, 16, 32, 64")
// 	}
// 	r.leafsize = val
// 	r.Root = r.refresh(r.Root)
// 	return nil
// }

// //TODO
// func (r *Rope) refresh(n *RopeNode) *RopeNode {
// 	return n
// }

func (r *Rope) LeafSize() int {
	return r.leafsize
}

func (r *Rope) RopeSize() int {
	return r.size
}
