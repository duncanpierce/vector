package vec

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"unsafe"
)

type (
	ScalableVector[E any, V constraintsExt.Vector[E]] struct {
		v V
	}

	Vector[E any] interface {
		Slice() []E
	}

	Lanes[E any] struct {
		// Arm SVE motivates Lanes[E] being different from Bool.
		// On AVX512, 1 bit corresponds to 1 element within vector register.
		// Arm SVE it may be wider because it each bit corresponds to 1 byte of vector register.
		mask [1]uint64
	}

	Bool struct {
		// mask has enough bits for each lane of the largest supported vector size
		mask uint64
	}
)

func (s *ScalableVector[E, V]) Slice() []E {
	// TODO unclear why I can't return (*s).v[:]
	return unsafe.Slice((*E)(unsafe.Pointer(&(*s).v)), len((*s).v))
}
