package hamt

type node struct {
	data [32]interface{}
}

func newNode() *node {
	return &node{}
}

func copyNode(old *node) *node {
	n := &node{}
	for i, _ := range old.data {
		n.data[i] = old.data[i]
	}
	return n
}
