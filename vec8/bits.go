package vec8

import (
	"golang.org/x/exp/constraints"
	"math/bits"
	"vector/mask"
)

func OnesCount[T constraints.Unsigned](m mask.Bits, a [N]T) (r [N]int) {
	m.ForTrue(N, func(i int) {
		r[i] = bits.OnesCount64(uint64(a[i]))
	})
	return
}

func LeadingZeros[T constraints.Unsigned](m mask.Bits, a [N]T) (r [N]int) {
	m.ForTrue(N, func(i int) {
		r[i] = bits.LeadingZeros64(uint64(a[i]))
	})
	return
}

func TrailingZeros[T constraints.Unsigned](m mask.Bits, a [N]T) (r [N]int) {
	m.ForTrue(N, func(i int) {
		r[i] = bits.TrailingZeros64(uint64(a[i]))
	})
	return
}
