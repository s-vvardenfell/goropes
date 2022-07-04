package rope

import (
	"unicode/utf8"
)

type Node struct {
	val   string
	len   int
	Left  *Node
	Right *Node
}

func NewNode(val string) *Node {
	return &Node{val, utf8.RuneCount([]byte(val)), nil, nil}
}
