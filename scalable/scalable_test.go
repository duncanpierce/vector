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
	b.ForRange(func(v int64) {
		t.Fail()
	})
}

func TestBunchCanConsumeSlice(t *testing.T) {
	b := NewBunch[byte]()
	original := "hello world from scalable vector land"
	v := []byte(original)
	b.ConsumeSlice(&v)

	if b.Predicate().Count() != len(original) {
		t.Errorf("wrong count")
	}

	r := ReturnSlice(b)
	if !reflect.DeepEqual(string(r), original) {
		t.Errorf("wrong result - got %v", r)
	}
}

func TestBunchConsumesPartOfLargeSlice(t *testing.T) {
	b := NewBunch[int64]()
	original := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	v := make([]int64, len(original))
	copy(v, original)
	b.ConsumeSlice(&v)

	if b.Predicate().Count() != 8 {
		t.Errorf("wrong count - got %v", b.Predicate().Count())
	}

	r := ReturnSlice(b)
	if !reflect.DeepEqual(r, original[:8]) {
		t.Errorf("wrong result - got %v", r)
	}
}
