package vec64

import (
	"github.com/duncanpierce/vector/mask"
	"golang.org/x/exp/constraints"
	"math/bits"
)

func ShiftRight[T, U constraints.Integer](m mask.Bits, a [N]T, b [N]U) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = a[i] >> b[i]
	})
	return
}

func ShiftLeft[T, U constraints.Integer](m mask.Bits, a [N]T, b [N]U) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = a[i] << b[i]
	})
	return
}

func RotateLeft[T constraints.Unsigned, U constraints.Integer](m mask.Bits, a [N]T, b [N]U) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = T(bits.RotateLeft(uint(a[i]), int(b[i])))
	})
	return
}
