package hamt

func Get[Data any](h *HAMT, key uint32) (Data, error) {
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

	if nextH.data[index] == nil {
		return result, DataNotFoundError
	}
	result = nextH.data[index].(Data)
	return result, nil
}

func Append[Data any](old *HAMT, key uint32, data Data) (*HAMT, error) {
	if old == nil {
		return nil, HAMTIsNilError
	}
	if key > 1000000000 {
		return nil, IndexOutOfBoundsError
	}
	h := NewHAMT()

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

	nextH.data[index] = data

	h.size = old.size + 1
	return h, nil
}

func Update[Data any](old *HAMT, key uint32, data Data) (*HAMT, error) {
	if old == nil {
		return nil, HAMTIsNilError
	}
	if key > 1000000000 {
		return nil, IndexOutOfBoundsError
	}
	h := NewHAMT()

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

func Pop[Data any](old *HAMT, key uint32) (*HAMT, Data, error) {
	var result Data
	if old == nil {
		return nil, result, HAMTIsNilError
	}
	if key > 1000000000 {
		return nil, result, IndexOutOfBoundsError
	}

	h := NewHAMT()
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
