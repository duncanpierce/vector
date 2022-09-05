package simple

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"github.com/duncanpierce/vector/simple/lanes"
	"golang.org/x/exp/constraints"
	"math/bits"
)

func OnesCount[E constraints.Unsigned, Z constraintsExt.Vector[int], X constraintsExt.Vector[E]](z *Z, m *lanes.Bool, x *X) {
	unary(z, m, x, func(x E) int {
		return bits.OnesCount64(uint64(x))
	})
}

func LeadingZeros[E constraints.Unsigned, Z constraintsExt.Vector[int], X constraintsExt.Vector[E]](z *Z, m *lanes.Bool, x *X) {
	unary(z, m, x, func(x E) int {
		return bits.LeadingZeros64(uint64(x))
	})
}

func TrailingZeros[E constraints.Unsigned, Z constraintsExt.Vector[int], X constraintsExt.Vector[E]](z *Z, m *lanes.Bool, x *X) {
	unary(z, m, x, func(x E) int {
		return bits.TrailingZeros64(uint64(x))
	})
}
