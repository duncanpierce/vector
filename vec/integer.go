package vec

import (
	"golang.org/x/exp/constraints"
)

func And[E constraints.Integer](z, x, y Vector[E], m *Mask) {
	binary(m, z, x, y, func(x, y E) E {
		return x & y
	})
}

func AndNot[E constraints.Integer](z, x, y Vector[E], m *Mask) {
	binary(m, z, x, y, func(x, y E) E {
		return x &^ y
	})
}

func Or[E constraints.Integer](z, x, y Vector[E], m *Mask) {
	binary(m, z, x, y, func(x, y E) E {
		return x | y
	})
}

func Xor[E constraints.Integer](z, x, y Vector[E], m *Mask) {
	binary(m, z, x, y, func(x, y E) E {
		return x ^ y
	})
}

func Not[E constraints.Integer](z, x Vector[E], m *Mask) {
	var allBits E
	allBits--
	unary(m, z, x, func(x E) E {
		return x ^ allBits
	})
}
