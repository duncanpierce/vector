package scalable

import "testing"

func TestVectorLengthSelection(t *testing.T) {
	b := NewBunch[int64]()
	_, ok := b.(bunch[int64, [8]int64])
	if !ok {
		t.Errorf("did not select expected vector length")
	}
}

func TestMisalignedVectorLengthSelection(t *testing.T) {
	b := NewBunch[[3]byte]()
	_, ok := b.(bunch[[3]byte, [16][3]byte])
	if !ok {
		t.Errorf("did not select expected vector length")
	}
}

func TestNarrowItemVectorLengthSelection(t *testing.T) {
	b := NewBunch[byte]()
	_, ok := b.(bunch[byte, [64]byte])
	if !ok {
		t.Errorf("did not select expected vector length")
	}
}
