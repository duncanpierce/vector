package vec

import (
	"github.com/duncanpierce/vector/runtimeExt"
	"math/bits"
	"unsafe"
)

// New returns a vector of the best length for the element type. The length is chosen to make best use of the CPU's vector instructions.
func New[E any]() (b Vector[E]) {
	return NewSized[E, E]()
}

// NewSized returns a vector sized to match the length of another vector with elements of type SizeFor.
// NewSized should be used when you require vectors with the same number of elements but different types.
func NewSized[E any, SizeFor any]() (b Vector[E]) {
	switch VectorLen[SizeFor]() {
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
