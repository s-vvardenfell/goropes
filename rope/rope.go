package rope

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type stringTuple struct {
	left  string
	right string
}

type Rope struct {
	Root     *RopeNode
	raw      string //should store runes to work with non-ascii
	size     int
	leafsize int
}

func (r *Rope) createRope(lo, hi int) *RopeNode {
	if lo+r.leafsize > hi {
		r.size += hi - lo + 1
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
	// r.Root = update(r.Root) // may need balancing after that
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

func update(root *RopeNode) *RopeNode {
	if root == nil || root.Val != "" { // comparison with "" won't work with runes
		return root
	}

	leftMost, rightMost := root.Right, root.Left

	for leftMost == nil && leftMost.Val == "" {
		leftMost = leftMost.Left
	}

	for rightMost == nil && rightMost.Val == "" {
		rightMost = rightMost.Right
	}

	if rightMost != nil {
		rightMost.Right = leftMost
	}

	if leftMost != nil {
		leftMost.Left = rightMost
	}
	return root
}

// Utility for Find() func.
// Returns a rope with single character at Root
func (r *Rope) find(root *RopeNode, idx int) (*Rope, error) {
	if idx < root.weight {
		if root != nil && root.Val == "" {
			return r.find(root.Left, idx)
		}
		return NewRopeFromString(string(root.Val[idx]), r.leafsize)
	}
	idx -= root.weight
	return r.find(root.Right, idx)
}

// Returns value of rope at index 'i'
// Returned value is also a Rope object
func (r *Rope) Find(idx int) (*Rope, error) {
	length := r.Size(nil)
	if 0 <= idx && idx < length {
		return r.find(r.Root, idx)
	} else if -1*length <= idx && idx < 0 {
		idx = length + idx
	} else {
		return nil, errors.New("index out of range")
	}
	return r.find(r.Root, idx)
}

func (r *Rope) concat(leftRoot *RopeNode, rightRoot *RopeNode, needUpdate bool, root *RopeNode) *RopeNode {
	if leftRoot != nil && rightRoot != nil {
		countLeft := r.Size(leftRoot)

		if rightRoot.Val != "" {
			countRight := r.Size(rightRoot)
			total := countLeft + countRight

			if total <= r.leafsize {
				leftThread := leftRoot.Left
				rightRoot.Val = leftRoot.Val + rightRoot.Val
				rightRoot.weight = len(rightRoot.Val)
				rightRoot.Left = leftThread

				if leftThread != nil {
					leftThread.Right = rightRoot
				}

				r.Delete(leftRoot)
				return rightRoot
			}
		}

		if root == nil {
			root = NewRopeNode(countLeft, "")
		} else {
			root.weight = countLeft
			root.Val = ""
		}

		root.Left, root.Right = root.Right, root.Left

		// if needUpdate {
		// 	root = r.update(root)
		// 	return r.balance(root)
		// }
	}

	if leftRoot != nil {
		return leftRoot
	}

	if rightRoot != nil {
		return rightRoot
	}
	return nil
}

func (r *Rope) Concat(other *Rope, inplace bool) *Rope {
	other.leafsize = r.leafsize
	newRoot := r.concat(r.Root, other.Root, true, nil)

	if inplace {
		r.Root = newRoot
		r.size = other.size
		return r
	}

	newRoot = r.copy(newRoot)
	newRope := Rope{leafsize: r.leafsize}

	newRope.Root = newRoot
	newRope.size = r.size + other.size
	return &newRope
}

func (r *Rope) copy(n *RopeNode) *RopeNode {
	if n != nil {
		newNode := NewRopeNode(n.weight, n.Val)
		newNode.height = n.height

		if n.Val != "" {
			return newNode
		}

		newNode.Left = r.copy(n.Left)
		newNode.Right = r.copy(n.Right)
		return newNode //update(newNode) //TODO
	}
	return n
}

func (r *Rope) Delete(n *RopeNode) int {
	return r.leafsize
}

func (r *Rope) LeafSize() int {
	return r.leafsize
}

func (r *Rope) RopeSize() int {
	return r.size
}

// For inner purposes
func (r *Rope) Size(root *RopeNode) int {
	if root == nil {
		root = r.Root
	}
	head := root
	count := 0

	for head != nil {
		count += head.weight
		if head.Val != "" {
			break
		}
		head = head.Right
	}
	return count
}

// Returns visual of Rope tree
func (r *Rope) Display() {
	lines, _, _, _ := r.displayAux(r.Root)
	for _, line := range lines {
		fmt.Println(line)
	}
}

// Returns list of strings, width, height and horizontal coordinate of the root
func (r *Rope) displayAux(root *RopeNode) ([]string, int, int, int) {
	// no child case
	if root != nil && root.Val != "" {
		line := visualizer(root)
		width := len(line)
		height := 1
		middle := width / 2
		return []string{line}, width, height, middle
	}

	left, n, p, x := r.displayAux(root.Left)
	rigth, m, q, y := r.displayAux(root.Right)
	s := visualizer(root)
	u := len(s)

	firstLine := strings.Repeat(" ", x+1) + strings.Repeat("_", n-x-1) + s + strings.Repeat("_", y) + strings.Repeat(" ", m-y)
	secondLine := strings.Repeat(" ", x) + "/" + strings.Repeat(" ", n-x-1+u+y) + "\\" + strings.Repeat(" ", m-y-1)

	if p < q {
		for i := 0; i < q-p; i++ {
			left = append(left, strings.Repeat(" ", n))
		}
	} else if q < p {
		for i := 0; i < p-q; i++ {
			rigth = append(rigth, strings.Repeat(" ", m))
		}
	}

	zippedLines := zip(left, rigth)
	lines := []string{firstLine, secondLine}

	for i := range zippedLines {
		lines = append(lines, zippedLines[i].left+strings.Repeat(" ", u)+zippedLines[i].right)
	}

	return lines, n + m + u, max(p, q) + 2, n + u/2
}

func visualizer(root *RopeNode) string {
	if root.Val != "" {
		return root.Val
	}
	return strconv.Itoa(root.weight)
}

func zip(a []string, b []string) []stringTuple /*, error)*/ {
	// if len(a) != len(b) {
	// 	return nil, fmt.Errorf("zip: arguments length must be same ")
	// }

	var t []stringTuple
	for index, value := range a {
		t = append(t, stringTuple{value, b[index]})
	}
	return t
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
