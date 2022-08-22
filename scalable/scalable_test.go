package scalable

import (
	"reflect"
	"testing"
)

func TestVectorLengthSelection(t *testing.T) {
	b := NewBunch[int64]()
	_, ok := b.(*bunch[int64, [8]int64])
	if !ok {
		t.Errorf("did not select expected vector length")
	}
}

func TestMisalignedVectorLengthSelection(t *testing.T) {
	b := NewBunch[[3]byte]()
	_, ok := b.(*bunch[[3]byte, [16][3]byte])
	if !ok {
		t.Errorf("did not select expected vector length")
	}
}

func TestNarrowItemVectorLengthSelection(t *testing.T) {
	b := NewBunch[byte]()
	_, ok := b.(*bunch[byte, [64]byte])
	if !ok {
		t.Errorf("did not select expected vector length")
	}
}

func TestNewBunchIsEmpty(t *testing.T) {
	b := NewBunch[int64]()
	b.Store(FailWriter[int64]{t})
}

func TestBunchCanConsumeSlice(t *testing.T) {
	b := NewBunch[byte]()
	original := "hello world from scalable vector land"
	v := []byte(original)
	slice := Slice(v)
	b.Load(slice)

	if b.Active().Count() != len(original) {
		t.Errorf("wrong count")
	}

	r := Extract(b)
	if !reflect.DeepEqual(string(r), original) {
		t.Errorf("wrong result - got %v", r)
	}

	sliceIo := slice.(*sliceIo[byte])
	if len(sliceIo.Elements) != 0 {
		t.Errorf("all elements should have been consumed from slice - still have %v", sliceIo.Elements)
	}
}

func TestBunchConsumesPartOfLargeSlice(t *testing.T) {
	b := NewBunch[int64]()
	original := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	slice := Slice(original)
	b.Load(slice)

	if b.Active().Count() != 8 {
		t.Errorf("wrong count - got %v", b.Active().Count())
	}

	r := Extract(b)
	if !reflect.DeepEqual(r, original[:8]) {
		t.Errorf("wrong result - got %v", r)
	}

	sliceIo := slice.(*sliceIo[int64])
	if len(sliceIo.Elements) != 8 {
		t.Errorf("elements should remain in slice - only have %v", sliceIo.Elements)
	}
}

func TestMultiply(t *testing.T) {
	b := NewBunch[int64]()
	b.Load(Slice([]int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}))

	r := NewBunch[int64]()
	r.Mul(b, b)

	if !reflect.DeepEqual(r, []int64{1, 4, 9, 16, 25, 36, 47, 64}) {
		t.Errorf("wrong result - got %v", r)
	}
}
