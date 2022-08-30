package vec

import (
	"golang.org/x/exp/constraints"
	"math/bits"
)

func OnesCount[E constraints.Unsigned, Z Vector[int], X Vector[E]](z Z, x X, m *Mask) {
	unary[int, E](m, z, x, func(x E) int {
		return bits.OnesCount64(uint64(x))
	})
}

func LeadingZeros[E constraints.Unsigned, Z Vector[int], X Vector[E]](z Z, x X, m *Mask) {
	unary[int, E](m, z, x, func(x E) int {
		return bits.LeadingZeros64(uint64(x))
	})
}

func TrailingZeros[E constraints.Unsigned, Z Vector[int], X Vector[E]](z Z, x X, m *Mask) {
	unary[int, E](m, z, x, func(x E) int {
		return bits.TrailingZeros64(uint64(x))
	})
}
