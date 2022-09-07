package simd

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"github.com/duncanpierce/vector/predicate"
	"golang.org/x/exp/constraints"
)

func Less[E constraints.Ordered, XY constraintsExt.Vector[E]](z, m *predicate.Bool, x, y *XY) {
	binaryBool(z, m, x, y, func(a, b E) bool {
		return a < b
	})
}

func LessOrEqual[E constraints.Ordered, XY constraintsExt.Vector[E]](z, m *predicate.Bool, x, y *XY) {
	binaryBool(z, m, x, y, func(a, b E) bool {
		return a <= b
	})
}

func Greater[E constraints.Ordered, XY constraintsExt.Vector[E]](z, m *predicate.Bool, x, y *XY) {
	binaryBool(z, m, x, y, func(a, b E) bool {
		return a > b
	})
}

func GreaterOrEqual[E constraints.Ordered, XY constraintsExt.Vector[E]](z, m *predicate.Bool, x, y *XY) {
	binaryBool(z, m, x, y, func(a, b E) bool {
		return a >= b
	})
}

func Min[E constraints.Ordered, XYZ constraintsExt.Vector[E]](z *XYZ, m *predicate.Bool, x, y *XYZ) {
	binary[E](z, m, x, y, func(a, b E) E {
		if a < b {
			return a
		}
		return b
	})
}

func Max[E constraints.Ordered, XYZ constraintsExt.Vector[E]](z *XYZ, m *predicate.Bool, x, y *XYZ) {
	binary[E](z, m, x, y, func(a, b E) E {
		if a > b {
			return a
		}
		return b
	})
}

func Abs[E constraintsExt.OrderedNumber, XZ constraintsExt.Vector[E]](z *XZ, m *predicate.Bool, x *XZ) {
	unary(z, m, x, func(a E) E {
		if a < 0 {
			a = -a
		}
		return a
	})
}
