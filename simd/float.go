package simd

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"github.com/duncanpierce/vector/simd/lanes"
	"golang.org/x/exp/constraints"
	"math"
)

func Round[E constraints.Float, XZ constraintsExt.Vector[E]](z *XZ, m *lanes.Bool, x *XZ) {
	unary[E](z, m, x, func(x E) E {
		return E(math.Round(float64(x)))
	})
}
func RoundToEven[E constraints.Float, XZ constraintsExt.Vector[E]](z *XZ, m *lanes.Bool, x *XZ) {
	unary[E](z, m, x, func(x E) E {
		return E(math.RoundToEven(float64(x)))
	})
}

func IsInf[E constraints.Float, X constraintsExt.Vector[E]](z, m *lanes.Bool, x *X, sign int) {
	unaryBool[E](z, m, x, func(x E) bool {
		return math.IsInf(float64(x), sign)
	})
}

func IsNaN[E constraints.Float, X constraintsExt.Vector[E]](z, m *lanes.Bool, x *X) {
	unaryBool[E](z, m, x, func(x E) bool {
		return math.IsNaN(float64(x))
	})
}
