package vec

import (
	"golang.org/x/exp/constraints"
	"math"
)

func Round[E constraints.Float, XZ Vector[E]](z, x XZ) {
	unary[E, E](z, x, func(x E) E {
		return E(math.Round(float64(x)))
	})
}

func RoundToEven[E constraints.Float, XZ Vector[E]](z, x XZ) {
	unary[E, E](z, x, func(x E) E {
		return E(math.RoundToEven(float64(x)))
	})
}
