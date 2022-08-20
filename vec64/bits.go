package vec64

import (
	"golang.org/x/exp/constraints"
	"math/bits"
)

func OnesCount[T constraints.Unsigned](m Bool, a [N]T) (r [N]int) {
	m.ForTrue(func(i int) {
		r[i] = bits.OnesCount64(uint64(a[i]))
	})
	return
}

func LeadingZeros[T constraints.Unsigned](m Bool, a [N]T) (r [N]int) {
	m.ForTrue(func(i int) {
		r[i] = bits.LeadingZeros64(uint64(a[i]))
	})
	return
}

func TrailingZeros[T constraints.Unsigned](m Bool, a [N]T) (r [N]int) {
	m.ForTrue(func(i int) {
		r[i] = bits.TrailingZeros64(uint64(a[i]))
	})
	return
}
