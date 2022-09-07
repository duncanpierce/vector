package simd

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"github.com/duncanpierce/vector/predicate"
)

func ternary[E any, WXYZ constraintsExt.Vector[E]](z *WXYZ, m *predicate.Bool, w, x, y *WXYZ, f func(a, b, c E) E) {
	predicate.RangeActive(len(*z), m, func(i, j int) {
		(*z)[i] = f((*w)[i], (*x)[i], (*y)[i])
	})
}

func binary[E any, XYZ constraintsExt.Vector[E]](z *XYZ, m *predicate.Bool, x, y *XYZ, f func(a, b E) E) {
	predicate.RangeActive(len(*z), m, func(i, j int) {
		(*z)[i] = f((*x)[i], (*y)[i])
	})
}

func binaryConvert[E, F any, Z constraintsExt.Vector[F], XY constraintsExt.Vector[E]](z *Z, m *predicate.Bool, x, y *XY, f func(a, b E) F) {
	predicate.RangeActive(len(*z), m, func(i, j int) {
		(*z)[i] = f((*x)[i], (*y)[i])
	})
}

func binaryBool[E any, XY constraintsExt.Vector[E]](z *predicate.Bool, m *predicate.Bool, x, y *XY, f func(a, b E) bool) {
	// TODO there is a risk here that a Bool is used for a wider vector, when reused for a narrower vector, it retains meaningless (?) upper bits
	// In particular, lane count will be unusable
	predicate.RangeActive(len(*x), m, func(i, j int) {
		predicate.Set(z, i, f((*x)[i], (*y)[i]))
	})
}

func unary[From, To any, Z constraintsExt.Vector[To], X constraintsExt.Vector[From]](z *Z, m *predicate.Bool, x *X, f func(a From) To) {
	predicate.RangeActive(len(*z), m, func(i, j int) {
		(*z)[i] = f((*x)[i])
	})
}

func unaryConvert[E, F any, Z constraintsExt.Vector[F], X constraintsExt.Vector[E]](z *Z, m *predicate.Bool, x *X, f func(a E) F) {
	predicate.RangeActive(len(*z), m, func(i, j int) {
		(*z)[i] = f((*x)[i])
	})
}

func unaryBool[E any, X constraintsExt.Vector[E]](z, m *predicate.Bool, x *X, f func(a E) bool) {
	predicate.RangeActive(len(*x), m, func(i, j int) {
		predicate.Set(z, i, f((*x)[i]))
	})
}
