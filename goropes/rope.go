package goropes

import "io"

type Rope struct {
	node *Node
}

func NewRopeFromString(in string) *Rope {
	return &Rope{}
}

func NewRopeFromReader(in io.Reader) *Rope {
	return &Rope{}
}

func NewRopeFromBytes(in []byte) *Rope {
	return &Rope{}
}
