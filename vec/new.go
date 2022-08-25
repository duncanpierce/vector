package vec

import (
	"github.com/duncanpierce/vector/runtimeExt"
	"math/bits"
	"unsafe"
)

func New[E any]() (b Vector[E]) {
	switch VectorLen[E]() {
	case 1:
		b = &ScalableVector[E, [1]E]{}
	case 2:
		b = &ScalableVector[E, [2]E]{}
	case 4:
		b = &ScalableVector[E, [4]E]{}
	case 8:
		b = &ScalableVector[E, [8]E]{}
	case 16:
		b = &ScalableVector[E, [16]E]{}
	case 32:
		b = &ScalableVector[E, [32]E]{}
	case 64:
		b = &ScalableVector[E, [64]E]{}
	default:
		panic("no suitable vector size found")
	}
	return
}

func VectorLen[E any]() int {
	var el E
	size := uint64(unsafe.Sizeof(el))
	elementSize := 1 << (64 - bits.LeadingZeros64(size) - 1)
	if bits.OnesCount64(size) != 1 {
		// Go to next size up if size isn't a power of 2
		elementSize *= 2
	}
	return runtimeExt.VectorLenBytes() / elementSize
}
