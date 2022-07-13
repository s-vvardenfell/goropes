package rope

type RopeNode struct {
	Left   *RopeNode // left child
	Right  *RopeNode //right child
	weight int       //left child size
	// height int       //for balancing
	Val string //sub-string if leaf node else empty
}

func NewRopeNode(weight int, val string) *RopeNode {
	return &RopeNode{
		weight: weight,
		Val:    val,
	}
}
