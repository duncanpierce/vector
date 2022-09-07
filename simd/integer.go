package simd

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"github.com/duncanpierce/vector/predicate"
	"golang.org/x/exp/constraints"
)

func And[E constraints.Integer, XYZ constraintsExt.Vector[E]](z *XYZ, m *predicate.Bool, x, y *XYZ) {
	binary(z, m, x, y, func(x, y E) E {
		return x & y
	})
}

func AndNot[E constraints.Integer, XYZ constraintsExt.Vector[E]](z *XYZ, m *predicate.Bool, x, y *XYZ) {
	binary(z, m, x, y, func(x, y E) E {
		return x &^ y
	})
}

func Or[E constraints.Integer, XYZ constraintsExt.Vector[E]](z *XYZ, m *predicate.Bool, x, y *XYZ) {
	binary(z, m, x, y, func(x, y E) E {
		return x | y
	})
}

func Xor[E constraints.Integer, XYZ constraintsExt.Vector[E]](z *XYZ, m *predicate.Bool, x, y *XYZ) {
	binary(z, m, x, y, func(x, y E) E {
		return x ^ y
	})
}

func Not[E constraints.Integer, XYZ constraintsExt.Vector[E]](z *XYZ, m *predicate.Bool, x, y *XYZ) {
	var allBits E
	allBits--
	unary(z, m, x, func(x E) E {
		return x ^ allBits
	})
}
