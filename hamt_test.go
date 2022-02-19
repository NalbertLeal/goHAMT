package hamt

import (
	"testing"
)

func TestAppend(t *testing.T) {
	h := NewHAMT()
	_, err := Append(h, 5, 89)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}

func TestGet(t *testing.T) {
	h := NewHAMT()
	h2, err := Append(h, 5, 89)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	d, err := Get[int](h2, 5)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	if d != 89 {
		t.Errorf("d != 89")
	}
}
