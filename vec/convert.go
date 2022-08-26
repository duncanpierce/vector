package vec

import "github.com/duncanpierce/vector/constraintsExt"

// Convert returns a vector with all elements converted to another numeric type. Complex numbers cannot be converted to other types.
func Convert[To, From constraintsExt.ConvertableNumber, Z constraintsExt.Vector[To], X constraintsExt.Vector[From]](z *Z, x *X) {
	lz := len(*z)
	if lz != len(*x) {
		panic("input and result vectors must be the same length")
	}
	for i := 0; i < lz; i++ {
		(*z)[i] = To((*x)[i])
	}
}
