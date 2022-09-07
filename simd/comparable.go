package simd

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"github.com/duncanpierce/vector/predicate"
)

func Equal[E comparable, XY constraintsExt.Vector[E]](z, m *predicate.Bool, x, y *XY) {
	binaryBool[E](z, m, x, y, func(a, b E) bool {
		return a == b
	})
}

func NotEqual[E comparable, XY constraintsExt.Vector[E]](z, m *predicate.Bool, x, y *XY) {
	binaryBool[E](z, m, x, y, func(a, b E) bool {
		return a != b
	})
}
