package vec

import "github.com/duncanpierce/vector/constraintsExt"

// Convert returns a vector with all elements converted to another numeric type. Complex numbers cannot be converted to other types.
func Convert[From, To constraintsExt.ConvertableNumber, FromV constraintsExt.Vector[From], ToV constraintsExt.Vector[To]](z *ToV, x *FromV) {
	lz := len(*z)
	if lz != len(*x) {
		panic("input and result vectors must be the same length")
	}
	for i := 0; i < lz; i++ {
		(*z)[i] = To((*x)[i])
	}
}
