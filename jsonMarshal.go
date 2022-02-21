package hamt

import (
	"encoding/json"
)

func (h *HAMT[Data]) MarshalJSON() ([]byte, error) {
	arr := []Data{}
	c := h.initDFS()
	for {
		select {
		case x := <-c:
			if x == nil {
				return json.Marshal(arr)
			}
			arr = append(arr, x.(Data))
		}
	}
}

func (h *HAMT[Data]) initDFS() chan interface{} {
	c := make(chan interface{})
	go h.dfsNodes(c, h.root, 5)
	return c
}

func (h *HAMT[Data]) dfsNodes(c chan interface{}, n *node, height uint) {
	for i := 0; i < 32; i++ {
		if n.data[i] != nil {
			if height > 0 {
				h.dfsNodes(c, n.data[i].(*node), height-1)
			} else {
				c <- n.data[i]
			}
		}
	}
	if height == 5 {
		for i := 0; i < 32; i++ {
			if h.tail.data[i] != nil {
				c <- h.tail.data[i]
			}
		}
		c <- nil
	}
}
