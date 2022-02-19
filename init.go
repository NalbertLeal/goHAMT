package hamt

import (
	"fmt"
)

const (
	BITS  = 5
	SHIFT = 25 // BITS * (HAMT_HEIGHT - 1)
	MASK  = 0x1f
)

var (
	HAMTIsNilError        = fmt.Errorf("The HAMT is nil")
	DataNotFoundError     = fmt.Errorf("Data not found")
	IndexOutOfBoundsError = fmt.Errorf("Max allowed index is 1 billion")
)
