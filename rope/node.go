package rope

type Node struct {
	Left   *Node // left child
	Right  *Node // right child
	Weight int   // left child size
	// Height int   // for balancing
	Data []rune
}

func NewNode(weight int, data []rune) *Node {
	return &Node{
		Weight: weight,
		Data:   data,
	}
}
