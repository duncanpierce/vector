package vec64

import (
	"golang.org/x/exp/constraints"
	"vector/constraintsExt"
	"vector/mask"
)

func Split[T constraintsExt.Number](a [N]T) (r, s [N / 2]T) {
	copy(r[:], a[:N/2])
	copy(s[:], a[N/2:])
	return
}

func Join[T constraintsExt.Number](a, b [N / 2]T) (r [N]T) {
	copy(r[:N/2], a[:])
	copy(r[N/2:], b[:])
	return
}

func Convert[S, T constraintsExt.Number](m mask.Bits, a [N]S) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = T(a[i])
	})
	return
}

func Blend[T constraintsExt.Number](m mask.Bits, a, b [N]T) (r [N]T) {
	m.For(N, func(i int, c bool) {
		if c {
			r[i] = b[i]
		} else {
			r[i] = a[i]
		}
	})
	return
}

func Broadcast[T constraintsExt.Number](m mask.Bits, a T) (r [N]T) {
	m.ForTrue(N, func(i int) {
		r[i] = a
	})
	return
}

func Permute[T constraints.Integer, U constraintsExt.Number](m mask.Bits, a [N]T, b [N]U) (r [N]U) {
	m.ForTrue(N, func(i int) {
		r[a[i]] = b[i]
	})
	return
}
