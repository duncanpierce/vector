package vec2

import (
	"golang.org/x/exp/constraints"
	"math/bits"
	"vec/condition"
)

func ShiftRight[T, U constraints.Integer](m condition.Mask, a [N]T, b [N]U) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = a[i] >> b[i]
	})
	return
}

func ShiftLeft[T, U constraints.Integer](m condition.Mask, a [N]T, b [N]U) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = a[i] << b[i]
	})
	return
}

func RotateLeft[T constraints.Unsigned, U constraints.Integer](m condition.Mask, a [N]T, b [N]U) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = T(bits.RotateLeft(uint(a[i]), int(b[i])))
	})
	return
}
