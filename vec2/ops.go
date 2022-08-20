package vec2

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math/bits"
	"vector/constraintsExt"
	"vector/mask"
)

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

func Zero[T constraintsExt.Number](m mask.Bits, a [N]T) (r [N]T) {
	m.For(N, func(i int, c bool) {
		if !c {
			r[i] = a[i]
		}
	})
	return
}

// README-DESIGN.md explains absence of mask parameter
func Broadcast[T constraintsExt.Number](a T) (r [N]T) {
	for i := range r {
		r[i] = a
	}
	return
}

func Permute[T constraints.Integer, U constraintsExt.Number](m mask.Bits, a [N]T, b [N]U) (r [N]U) {
	m.ForTrue(N, func(i int) {
		r[i] = b[a[i]]
	})
	return
}

/*
Interlace copies elements from two half-length vectors a and b into a single vector, taking n consecutive elements from a, then n consecutive elements from b.
This is repeated until all elements of a and b have been included in the result. For example, if n is 1,
the first few elements of the result will contain elements a[0], b[0], a[1], b[1], a[2], b[2], etc.
If n is 2 the first few elements of the result will contain a[0], a[1], b[0], b[1], a[2], a[3], b[2], b[3], etc.
Panics if n is greater than vector length or is not a power of 2.
*/
func Interlace[T constraintsExt.Number](n int, a, b [N / 2]T) (r [N]T) {
	if n > N {
		panic(fmt.Sprintf("n cannot be greater than vector length %v", N))
	}
	if bits.OnesCount(uint(n)) != 1 {
		panic("n must be a power of 2")
	}
	x, y, z := a[:], b[:], r[:]
	for len(z) > 0 {
		copy(z, x[:n])
		copy(z[n:], y[:n])
		x, y, z = x[n:], y[n:], z[n*2:]
	}
	return
}
