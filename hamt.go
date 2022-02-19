package hamt

type HAMT struct {
	root *node
	size uint32
}

func NewHAMT() *HAMT {
	return &HAMT{
		root: newNode(),
		size: 0,
	}
}
