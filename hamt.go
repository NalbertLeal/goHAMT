package hamt

import "fmt"

type HAMT[Data any] struct {
	root         *node
	size         int
	biggestIndex int
	tail         *node
	tailStart    int
	tailEnd      int
}

func NewHAMT[Data any]() *HAMT[Data] {
	return &HAMT[Data]{
		root:         newNode(),
		size:         0,
		biggestIndex: -1,
		tail:         newNode(),
		tailStart:    0,
		tailEnd:      31,
	}
}

func (old *HAMT[Data]) InsertAt(key int, data Data) (*HAMT[Data], error) {
	if old == nil {
		return nil, HAMTIsNilError
	}
	if key > 1000000000 {
		return nil, IndexOutOfBoundsError
	}

	if key >= old.tailStart && key <= old.tailEnd {
		h := NewHAMT[Data]()
		h.root = old.root
		h.size = old.size + 1
		if old.biggestIndex < key {
			h.biggestIndex = key
		} else {
			h.biggestIndex = old.biggestIndex
		}
		h.tailStart = old.tailStart
		h.tailEnd = old.tailEnd
		h.tail = copyNode(old.tail)

		index := key & MASK
		h.tail.data[index] = data
		return h, nil
	}

	if key > old.tailEnd {
		h := old.insertTailIntoTree(key)
		h.size = h.size + 1

		index := key & MASK
		h.tail.data[index] = data
		return h, nil
	}

	h := NewHAMT[Data]()
	h.root = copyNode(old.root)
	nextH := h.root
	nextOld := old.root
	index := 0
	mustCreateNode := false
	for level := SHIFT; level > 0; level -= BITS {
		index = (int(key) >> level) & MASK
		if mustCreateNode || nextOld.data[index] == nil {
			nextH.data[index] = newNode()
			mustCreateNode = true
		} else {
			nextH.data[index] = copyNode(nextOld.data[index].(*node))
			nextOld = nextOld.data[index].(*node)
		}
		nextH = nextH.data[index].(*node)
	}

	index = int(key) & MASK
	nextH.data[index] = data

	h.size = old.size + 1
	return h, nil
}

func (old *HAMT[Data]) Append(data Data) (*HAMT[Data], error) {
	if old == nil {
		return nil, HAMTIsNilError
	}
	if old.biggestIndex > 1000000000 {
		return nil, IndexOutOfBoundsError
	}

	index := (old.biggestIndex + 1) & MASK
	if old.biggestIndex > 0 && index == 0 {
		h := old.insertTailIntoTree(old.biggestIndex + 1)
		h.size = h.size + 1
		h.tail.data[index] = data
		return h, nil
	}
	h := NewHAMT[Data]()
	h.root = old.root
	h.size = old.size + 1
	h.biggestIndex = old.biggestIndex + 1
	h.tailStart = old.tailStart
	h.tailEnd = old.tailEnd
	h.tail = copyNode(old.tail)
	h.tail.data[index] = data
	return h, nil
}

func (old *HAMT[Data]) Update(key int, data Data) (*HAMT[Data], error) {
	if old == nil {
		return nil, HAMTIsNilError
	}
	if key > 1000000000 {
		return nil, IndexOutOfBoundsError
	}
	h := NewHAMT[Data]()

	h.root = copyNode(old.root)
	nextH := h.root
	nextOld := old.root
	index := 0
	for level := SHIFT; level > 0; level -= BITS {
		index = (int(key) >> level) & MASK
		if nextOld.data[index] == nil {
			return nil, DataNotFoundError
		} else {
			nextH.data[index] = copyNode(nextOld.data[index].(*node))
			nextOld = nextOld.data[index].(*node)
		}
		nextH = nextH.data[index].(*node)
	}

	nextH.data[index] = data

	h.size = old.size + 1
	return h, nil
}

func (old *HAMT[Data]) Pop(key int) (*HAMT[Data], Data, error) {
	var result Data
	if old == nil {
		return nil, result, HAMTIsNilError
	}
	if key > 1000000000 {
		return nil, result, IndexOutOfBoundsError
	}

	h := NewHAMT[Data]()
	h.root = copyNode(old.root)
	nextH := h.root
	nextOld := old.root
	index := 0
	for level := SHIFT; level > 0; level -= BITS {
		index = (int(key) >> level) & MASK
		if nextOld.data[index] == nil {
			return nil, result, DataNotFoundError
		} else {
			nextH.data[index] = copyNode(nextOld.data[index].(*node))
			nextOld = nextOld.data[index].(*node)
		}
		nextH = nextH.data[index].(*node)
	}

	result = nextH.data[index].(Data)
	nextH.data[index] = nil

	h.size = old.size - 1
	return h, result, nil
}

func (h *HAMT[Data]) Get(key int) (Data, error) {
	var result Data
	if h == nil {
		return result, HAMTIsNilError
	}

	if key >= h.tailStart && key <= h.tailEnd {
		index := key & MASK
		if h.tail.data[index] == nil {
			fmt.Println(h.tail)
			return result, DataNotFoundError
		}
		result = h.tail.data[index].(Data)
		return result, nil
	}

	nextH := h.root
	index := 0
	for level := SHIFT; level > 0; level -= BITS {
		index = (key >> level) & MASK
		if nextH.data[index] == nil {
			return result, DataNotFoundError
		}
		nextH = nextH.data[index].(*node)
	}

	index = int(key) & MASK
	if nextH.data[index] == nil {
		return result, DataNotFoundError
	}
	result = nextH.data[index].(Data)
	return result, nil
}

func (old *HAMT[Data]) insertTailIntoTree(biggestIndex int) *HAMT[Data] {
	h := NewHAMT[Data]()
	h.root = copyNode(old.root)
	h.size = old.size
	h.biggestIndex = biggestIndex
	h.tail = newNode()
	h.tailStart = biggestIndex - (biggestIndex % 32)
	h.tailEnd = h.tailStart + 31

	// now insert the old.tail into the new tree called "h"
	nextH := h.root
	nextOld := old.root
	var index int = 0
	mustCreateNode := false
	for level := SHIFT; level > 5; level -= BITS {
		index = (old.tailStart >> level) & MASK
		if mustCreateNode || nextOld.data[index] == nil {
			nextH.data[index] = newNode()
			mustCreateNode = true
		} else {
			nextH.data[index] = copyNode(nextOld.data[index].(*node))
			nextOld = nextOld.data[index].(*node)
		}
		nextH = nextH.data[index].(*node)
	}
	index = old.tailStart & MASK
	nextH.data[index] = old.tail

	return h
}
