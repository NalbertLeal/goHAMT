package hamt

import (
	"testing"
)

func TestInsertAt(t *testing.T) {
	h := NewHAMT[int]()
	h, err := h.InsertAt(5, 89)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if h.tail.data[5].(int) != 89 {
		t.Errorf("h.tail.data[5].(int) != 89")
		return
	}
	h, err = h.InsertAt(25, 924)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if h.tail.data[25].(int) != 924 {
		t.Errorf("h.tail.data[25].(int) != 924")
		return
	}

	// insertAt must create a new tail
	h, err = h.InsertAt(32773, 56)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if h.tail.data[5].(int) != 56 {
		t.Errorf("h.tail.data[32773].(int) != 56")
		return
	}
	d, _ := h.Get(5)
	if d != 89 {
		t.Errorf("d != 89: d = %d", d)
		return
	}
}

func TestAppend(t *testing.T) {
	h := NewHAMT[int]()
	h, err := h.Append(89)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	h, err = h.Append(25)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	d, _ := h.Get(0)
	d2, _ := h.Get(1)
	if d != 89 || d2 != 25 {
		t.Errorf("d != 89 || d2 != 25")
		return
	}

	// now test append in new tail
	for i := 2; i < 34; i++ {
		h, _ = h.Append(i)
	}
	d3, _ := h.Get(32)
	d4, _ := h.Get(33)
	if d3 != 32 || d4 != 33 {
		t.Errorf("d != 32 || d2 != 33")
		return
	}
}

func TestGet(t *testing.T) {
	h := NewHAMT[int]()
	h2, err := h.Append(89)
	h2, err = h2.Append(104)
	h2, err = h2.Append(7)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	d, err := h2.Get(0)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if d != 89 {
		t.Errorf("d != 89")
	}

	d, err = h2.Get(1)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if d != 104 {
		t.Errorf("d != 104")
	}

	d, err = h2.Get(2)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if d != 7 {
		t.Errorf("d != 7")
	}
}

func TestToJson(t *testing.T) {
	// create and insert
	h := NewHAMT[int]()
	h, _ = h.Append(1)
	h, _ = h.InsertAt(68, 5)
	h, _ = h.Append(3)
	h, _ = h.Append(9)

	// generate json
	j, err := h.MarshalJSON()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// verify result
	expected := []byte("[1,5,3,9]")
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
