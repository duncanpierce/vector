package vec2

import (
	"golang.org/x/exp/constraints"
)

func And[Element constraints.Integer](m Bool, a, b [Length]Element) (r [Length]Element) {
	m.ForTrue(func(i int) {
		r[i] = a[i] & b[i]
	})
	return
}

func AndNot[Element constraints.Integer](m Bool, a, b [Length]Element) (r [Length]Element) {
	m.ForTrue(func(i int) {
		r[i] = a[i] &^ b[i]
	})
	return
}

func Or[Element constraints.Integer](m Bool, a, b [Length]Element) (r [Length]Element) {
	m.ForTrue(func(i int) {
		r[i] = a[i] | b[i]
	})
	return
}

func Xor[Element constraints.Integer](m Bool, a, b [Length]Element) (r [Length]Element) {
	m.ForTrue(func(i int) {
		r[i] = a[i] ^ b[i]
	})
	return
}

func Not[Element constraints.Integer](m Bool, a [Length]Element) (r [Length]Element) {
	var all Element
	all--
	m.ForTrue(func(i int) {
		r[i] = a[i] ^ all
	})
	return
}

func AndNotZero[Element constraints.Integer](m Bool, a, b [Length]Element) (r Bool) {
	m.ForTrue(func(i int) {
		r.Set(i, a[i]&b[i] != 0)
	})
	return
}
