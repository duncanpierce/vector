package vec

import (
	"github.com/duncanpierce/vector/constraintsExt"
)

type (
	//ScalableVectorType[E constraintsExt.Number] interface {
	//	constraints.FixedVector[E]
	//}

	Vector1[E constraintsExt.Number]  [1]E
	Vector2[E constraintsExt.Number]  [2]E
	Vector4[E constraintsExt.Number]  [4]E
	Vector8[E constraintsExt.Number]  [8]E
	Vector16[E constraintsExt.Number] [16]E
	Vector32[E constraintsExt.Number] [32]E
	Vector64[E constraintsExt.Number] [64]E

	ScalableVector[E any, V constraintsExt.Vector[E]] struct {
		v V
	}

	Vector[E constraintsExt.Number] interface {
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
