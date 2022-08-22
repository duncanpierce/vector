package scalable

import (
	"github.com/duncanpierce/vector/runtimeExt"
	"math/bits"
	"unsafe"
)

type (
	// Bunch holds mask of active elements
	Bunch[Element any] interface {
		GrabSlice(s *[]Element)
		GrabChan(c <-chan Element)
		ForRange(f func(value Element))
	}

	array[Element any] interface {
		[2]Element | [4]Element | [8]Element | [16]Element | [32]Element | [64]Element
	}

	bunch[Element any, Array array[Element]] struct {
		vec  Array
		pred Predicate
	}
)

func NewBunch[T any]() (b Bunch[T]) {
	var el T
	size := uint64(unsafe.Sizeof(el))
	elementSize := 1 << (64 - bits.LeadingZeros64(size) - 1)
	if bits.OnesCount64(size) != 1 {
		// Go to next size up if size isn't a power of 2
		elementSize *= 2
	}
	vectorLength := runtimeExt.VectorLenBytes() / elementSize
	switch vectorLength {
	case 2:
		b = bunch[T, [2]T]{}
	case 4:
		b = bunch[T, [4]T]{}
	case 8:
		b = bunch[T, [8]T]{}
	case 16:
		b = bunch[T, [16]T]{}
	case 32:
		b = bunch[T, [32]T]{}
	default:
		b = bunch[T, [64]T]{}
	}
	return
}

func (b bunch[Element, Array]) ForRange(f func(value Element)) {
	b.pred.ForActive(func(index int) {
		f(b.vec[index])
	})
}

func (b bunch[Element, Array]) GrabSlice(s *[]Element) {
	// can't do took := copy(b.vec[:], *s) AFAICT because we can't slice a generic of type `array` so we have to do it in long-form
	take := min(len(b.vec), len(*s))
	for i := 0; i < take; i++ {
		b.vec[i] = (*s)[i]
	}
	*s = (*s)[take:]
}

func (b bunch[Element, Array]) GrabChan(c <-chan Element) {
	for i := 0; i < len(b.vec); i++ {
		b.vec[i] = <-c
	}
}
