package vec8

import (
	"golang.org/x/exp/constraints"
	"math/bits"
)

func OnesCount[Element constraints.Unsigned](m Bool, a [Length]Element) (r [Length]int) {
	m.ForTrue(func(i int) {
		r[i] = bits.OnesCount64(uint64(a[i]))
	})
	return
}

func LeadingZeros[Element constraints.Unsigned](m Bool, a [Length]Element) (r [Length]int) {
	m.ForTrue(func(i int) {
		r[i] = bits.LeadingZeros64(uint64(a[i]))
	})
	return
}

func TrailingZeros[Element constraints.Unsigned](m Bool, a [Length]Element) (r [Length]int) {
	m.ForTrue(func(i int) {
		r[i] = bits.TrailingZeros64(uint64(a[i]))
	})
	return
}
