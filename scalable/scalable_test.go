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
	b.Load(&Slice[byte]{v})

	if b.Active().Count() != len(original) {
		t.Errorf("wrong count")
	}

	r := Extract(b)
	if !reflect.DeepEqual(string(r), original) {
		t.Errorf("wrong result - got %v", r)
	}
}

func TestBunchConsumesPartOfLargeSlice(t *testing.T) {
	b := NewBunch[int64]()
	original := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	b.Load(&Slice[int64]{original})

	if b.Active().Count() != 8 {
		t.Errorf("wrong count - got %v", b.Active().Count())
	}

	r := Extract(b)
	if !reflect.DeepEqual(r, original[:8]) {
		t.Errorf("wrong result - got %v", r)
	}
}

//func TestMultiply(t *testing.T) {
//	b := NewBunch[int64]()
//	data := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
//	b.ConsumeSlice(&data)
//
//	r := NewBunch[int64]()
//	r.Mul(b, Broadcast(10))
//
//	if !reflect.DeepEqual(r, []int64{1, 2, 3, 4, 5, 6, 7, 8}) {
//		t.Errorf("wrong result - got %v", r)
//	}
//}
