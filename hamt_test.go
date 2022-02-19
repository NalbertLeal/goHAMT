package hamt

import (
	"testing"
)

func TestAppend(t *testing.T) {
	h := NewHAMT[int]()
	_, err := h.Append(5, 89)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}

func TestGet(t *testing.T) {
	h := NewHAMT[int]()
	h2, err := h.Append(5, 89)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	d, err := h2.Get(5)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if d != 89 {
		t.Errorf("d != 89")
	}
}
