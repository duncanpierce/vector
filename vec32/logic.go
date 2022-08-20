package vec32

import (
	"golang.org/x/exp/constraints"
)

func And[T constraints.Integer](m Bool, a, b [N]T) (r [N]T) {
	m.ForTrue(func(i int) {
		r[i] = a[i] & b[i]
	})
	return
}

func AndNot[T constraints.Integer](m Bool, a, b [N]T) (r [N]T) {
	m.ForTrue(func(i int) {
		r[i] = a[i] &^ b[i]
	})
	return
}

func Or[T constraints.Integer](m Bool, a, b [N]T) (r [N]T) {
	m.ForTrue(func(i int) {
		r[i] = a[i] | b[i]
	})
	return
}

func Xor[T constraints.Integer](m Bool, a, b [N]T) (r [N]T) {
	m.ForTrue(func(i int) {
		r[i] = a[i] ^ b[i]
	})
	return
}

func Not[T constraints.Integer](m Bool, a [N]T) (r [N]T) {
	var all T
	all--
	m.ForTrue(func(i int) {
		r[i] = a[i] ^ all
	})
	return
}

func AndNotZero[T constraints.Integer](m Bool, a, b [N]T) (r Bool) {
	m.ForTrue(func(i int) {
		r.Set(i, a[i]&b[i] != 0)
	})
	return
}
