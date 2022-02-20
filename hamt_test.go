package hamt

import (
	"testing"
)

func TestAppend(t *testing.T) {
	h := NewHAMT[int]()
	h, err := h.Append(5, 89)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	h, err = h.Append(89, 25)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	d, _ := h.Get(5)
	d2, _ := h.Get(89)
	if d != 89 && d2 != 25 {
		t.Errorf("d != 89 && d2 != 25")
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

func TestToJson(t *testing.T) {
	// create and insert
	h := NewHAMT[int]()
	h, _ = h.Append(89, 1)
	h, _ = h.Append(5, 5)
	h, _ = h.Append(4, 3)
	h, _ = h.Append(920, 9)

	// generate json
	j, err := h.MarshalJSON()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// verify result
	expected := []byte("[3,5,1,9]")
	if len(j) != len(expected) {
		t.Errorf("len(j) != len(expected)")
		return
	}
	for i, v := range expected {
		if v != j[i] {
			t.Errorf("v != j[i]:\n\tv = %d\n\tj[i] = %d", v, j[i])
			return
		}
	}
}
