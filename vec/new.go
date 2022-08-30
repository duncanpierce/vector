package vec

import (
	"github.com/duncanpierce/vector/runtimeExt"
	"math/bits"
	"unsafe"
)

// Broadcast returns a value that repeats for all lanes of any vector it is combined with.
func Broadcast[E any](x E) *broadcast[E] {
	return &broadcast[E]{x}
}

// NewVector returns a vector of the best length for the element type. The length is chosen to make best use of the CPU's vector instructions.
func NewVector[E any]() ScalableVector[E] {
	return NewVectorFor[E, E]()
}

// NewVectorFor returns a ScalableVector sized to match the length of another vector with elements of type SizeFor.
// NewVectorFor should be used when you require vectors with the same number of elements but different types.
func NewVectorFor[E any, SizeFor any]() ScalableVector[E] {
	switch ScalableVectorLen[SizeFor]() {
	case 1:
		return &Vec1[E]{}
	case 2:
		return &Vec2[E]{}
	case 4:
		return &Vec4[E]{}
	case 8:
		return &Vec8[E]{}
	case 16:
		return &Vec16[E]{}
	case 32:
		return &Vec32[E]{}
	case 64:
		return &Vec64[E]{}
	default:
		panic("no suitable size found")
	}
}

// ScalableVectorLen returns the vector length that NewVector will use to create a vector of elements of type E.
func ScalableVectorLen[E any]() int {
	var el E
	size := uint64(unsafe.Sizeof(el))
	elementSize := 1 << (64 - bits.LeadingZeros64(size) - 1)
	if bits.OnesCount64(size) != 1 {
		// Go to next size up if size isn't a power of 2
		elementSize *= 2
	}
	return runtimeExt.VectorLenBytes() / elementSize
}

// NewBoolFor returns a ScalableBool sized to the number of elements in a ScalableVector of type SizeFor.
func NewBoolFor[SizeFor any]() ScalableBool {
	switch ScalableVectorLen[SizeFor]() {
	case 1:
		return &Bool1{}
	case 2:
		return &Bool2{}
	case 4:
		return &Bool4{}
	case 8:
		return &Bool8{}
	case 16:
		return &Bool16{}
	case 32:
		return &Bool32{}
	case 64:
		return &Bool64{}
	default:
		panic("no suitable size found")
	}
}
