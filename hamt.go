package hamt

type HAMT[Data any] struct {
	root *node
	size uint32
}

func NewHAMT[Data any]() *HAMT[Data] {
	return &HAMT[Data]{
		root: newNode(),
		size: 0,
	}
}

func (old *HAMT[Data]) Append(key uint32, data Data) (*HAMT[Data], error) {
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

func (old *HAMT[Data]) Update(key uint32, data Data) (*HAMT[Data], error) {
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

func (old *HAMT[Data]) Pop(key uint32) (*HAMT[Data], Data, error) {
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

func (h *HAMT[Data]) Get(key uint32) (Data, error) {
	var result Data
	if h == nil {
		return result, HAMTIsNilError
	}

	nextH := h.root
	index := 0
	for level := SHIFT; level > 0; level -= BITS {
		index = (int(key) >> level) & MASK
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
