package simple

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"github.com/duncanpierce/vector/simple/lanes"
)

func Copy[E any, XZ constraintsExt.Vector[E]](z *XZ, m *lanes.Bool, x *XZ) {
	unary(z, m, x, func(x E) E {
		return x
	})
}

/*
Blend returns a vector with each element drawn from a if the corresponding mask bit in m is set or from b if the bit is not set.
*/
func Blend[E any, XYZ constraintsExt.Vector[E]](z *XYZ, m *lanes.Bool, x, y *XYZ) {
	lanes.RangeAll[E](z, m, func(i, j int, b bool) {
		if b {
			(*z)[i] = (*x)[i]
		} else {
			(*z)[i] = (*y)[i]
		}
	})
}

/*
Zero sets elements of a vector to their default (zero) value where the corresponding mask bit in m is not set.
*/
func Zero[E any, Z constraintsExt.Vector[E]](z *Z, m *lanes.Bool) {
	var zero E
	lanes.RangeInactive[E](z, m, func(i, j int) {
		(*z)[i] = zero
	})
}
