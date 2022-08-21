package vec8

import (
	"golang.org/x/exp/constraints"
	"math/bits"
)

func ShiftRight[T, U constraints.Integer](m Bool, a [Length]T, b [Length]U) (r [Length]T) {
	m.ForTrue(func(i int) {
		r[i] = a[i] >> b[i]
	})
	return
}

func ShiftLeft[T, U constraints.Integer](m Bool, a [Length]T, b [Length]U) (r [Length]T) {
	m.ForTrue(func(i int) {
		r[i] = a[i] << b[i]
	})
	return
}

func RotateLeft[T constraints.Unsigned, U constraints.Integer](m Bool, a [Length]T, b [Length]U) (r [Length]T) {
	m.ForTrue(func(i int) {
		r[i] = T(bits.RotateLeft(uint(a[i]), int(b[i])))
	})
	return
}
