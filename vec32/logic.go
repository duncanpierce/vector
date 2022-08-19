package vec32

import (
	"golang.org/x/exp/constraints"
	"vector/mask"
)

func And[T constraints.Integer](m mask.Bits, a, b [N]T) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = a[i] & b[i]
	})
	return
}

func AndNot[T constraints.Integer](m mask.Bits, a, b [N]T) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = a[i] &^ b[i]
	})
	return
}

func Or[T constraints.Integer](m mask.Bits, a, b [N]T) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = a[i] | b[i]
	})
	return
}

func Xor[T constraints.Integer](m mask.Bits, a, b [N]T) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = a[i] ^ b[i]
	})
	return
}

func Not[T constraints.Integer](m mask.Bits, a [N]T) (r [N]T) {
	var all T
	all--
	m.ForTrue(N, func(i int) {
		r[i] = a[i] ^ all
	})
	return
}
