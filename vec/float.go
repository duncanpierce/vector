package vec

import (
	"golang.org/x/exp/constraints"
	"math"
)

func Round[E constraints.Float, XZ Vector[E]](z, x XZ, m *Mask) {
	unary[E, E](m, z, x, func(x E) E {
		return E(math.Round(float64(x)))
	})
}

func RoundToEven[E constraints.Float, XZ Vector[E]](z, x XZ, m *Mask) {
	unary[E, E](m, z, x, func(x E) E {
		return E(math.RoundToEven(float64(x)))
	})
}

func IsInf[E constraints.Float, Z Bool, X Vector[E]](z Z, x X, sign int) {
	unaryBool[E](z, x, func(x E) bool {
		return math.IsInf(float64(x), sign)
	})
}

func IsNaN[E constraints.Float, Z Bool, X Vector[E]](z Z, x X) {
	unaryBool[E](z, x, func(x E) bool {
		return math.IsNaN(float64(x))
	})
}
