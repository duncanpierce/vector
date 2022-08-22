package scalable

import (
	"github.com/duncanpierce/vector/runtimeExt"
	"math/bits"
	"unsafe"
)

type (
	Bunch[Element any] interface {
		Load(source Reader[Element])
		Store(dest Writer[Element])
		Active() Predicate
	}

	array[Element any] interface {
		[2]Element | [4]Element | [8]Element | [16]Element | [32]Element | [64]Element
	}

	bunch[Element any, Array array[Element]] struct {
		vec    Array
		active Predicate
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
		b = &bunch[T, [2]T]{}
	case 4:
		b = &bunch[T, [4]T]{}
	case 8:
		b = &bunch[T, [8]T]{}
	case 16:
		b = &bunch[T, [16]T]{}
	case 32:
		b = &bunch[T, [32]T]{}
	default:
		b = &bunch[T, [64]T]{}
	}
	return
}

func (b *bunch[Element, Array]) Store(f Writer[Element]) {
	b.active.ForActive(func(index int) {
		f.Write(b.vec[index])
	})
}

func (b *bunch[Element, Array]) Load(source Reader[Element]) {
	for i := 0; i < len(b.vec); i++ {
		value, ok := source.Read()
		if !ok {
			return
		}
		b.vec[i] = value
		b.active = b.active.Set(i, true)
	}
}

func (b *bunch[Element, Array]) Active() Predicate {
	return b.active
}
