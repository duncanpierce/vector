package vec2

import (
	"golang.org/x/exp/constraints"
	"vec/condition"
)

func And[T constraints.Integer](m condition.Mask, a, b [N]T) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = a[i] & b[i]
	})
	return
}

func AndNot[T constraints.Integer](m condition.Mask, a, b [N]T) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = a[i] &^ b[i]
	})
	return
}

func Or[T constraints.Integer](m condition.Mask, a, b [N]T) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = a[i] | b[i]
	})
	return
}

func Xor[T constraints.Integer](m condition.Mask, a, b [N]T) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = a[i] ^ b[i]
	})
	return
}

func Not[T constraints.Integer](m condition.Mask, a [N]T) (r [N]T) {
	var all T
	all--
	m.ForTrue(N, func(i int) {
		r[i] = a[i] ^ all
	})
	return
}
