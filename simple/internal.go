package simple

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"github.com/duncanpierce/vector/simple/lanes"
)

func binary[E any, XYZ constraintsExt.Vector[E]](z *XYZ, m *lanes.Bool, x, y *XYZ, f func(a, b E) E) {
	lanes.Range[E](z, m, func(i, j int) {
		(*z)[i] = f((*x)[i], (*y)[i])
	})
}

func binaryBool[E any, XY constraintsExt.Vector[E]](z *lanes.Bool, m *lanes.Bool, x, y *XY, f func(a, b E) bool) {
	// TODO there is a risk here that a Bool is used for a wider vector, when reused for a narrower vector, it retains meaningless (?) upper bits
	// In particular, lane count will be unusable
	lanes.Range[E](x, m, func(i, j int) {
		lanes.Set(z, i, f((*x)[i], (*y)[i]))
	})
}

func unary[E any, XYZ constraintsExt.Vector[E]](z *XYZ, m *lanes.Bool, x *XYZ, f func(a E) E) {
	lanes.Range[E](z, m, func(i, j int) {
		(*z)[i] = f((*x)[i])
	})
}
