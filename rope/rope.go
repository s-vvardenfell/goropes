package rope

import (
	"io"
	"log"
	"strings"
)

type Rope struct {
	Root *Node
}

func NewRopeFromString(in string) *Rope {
	var r Rope
	s := strings.Fields(in)

	for i := range s {
		r.Insert(s[i])
	}
	return &r
}

func NewRopeFromReader(in io.Reader) *Rope {
	buf, err := io.ReadAll(in)
	if err != nil {
		log.Fatal(err)
	}
	return NewRopeFromString(string(buf))
}

// func (r *Rope) String() string {
// 	// var currentNode *Node
// 	// for currentNode != nil {

// 	// }
// 	// return fmt.Sprintf("val: %v, left: %v, right: %v\n", r.Node.val, r.Node.left, r.Node.right)
// }

func (r *Rope) Insert(val string) {
	r.Root = insertNode(r.Root, val)
}

func insertNode(root *Node, val string) *Node {
	if root == nil {
		return NewNode(val)
	}
	if val < root.val {
		root.Left = insertNode(root.Left, val)
	} else {
		root.Right = insertNode(root.Right, val)
	}
	return root
}
