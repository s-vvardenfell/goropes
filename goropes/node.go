package goropes

type Node struct {
	val   rune
	len   int64
	left  *Node
	right *Node
}
