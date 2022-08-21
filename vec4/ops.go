package vec4

import (
	"fmt"
	"github.com/duncanpierce/vector/constraintsExt"
	"golang.org/x/exp/constraints"
	"math/bits"
)

/*
Convert returns a vector with all elements converted to the type T.
*/
func Convert[ElementA, ElementB constraintsExt.Number](m Bool, a [Length]ElementA) (r [Length]ElementB) {
	m.ForTrue(func(i int) {
		r[i] = ElementB(a[i])
	})
	return
}

/*
Blend returns a vector with each element drawn from b if the corresponding mask bit in m is set or a if the bit is not set.
Unlike most other vector functions, Blend does not set elements in the result to zero if the corresponding mask bit is not set.
*/
func Blend[Element constraintsExt.Number](m Bool, a, b [Length]Element) (r [Length]Element) {
	m.For(func(i int, c bool) {
		if c {
			r[i] = b[i]
		} else {
			r[i] = a[i]
		}
	})
	return
}

/*
Zero returns a vector of elements drawn from a, except where the corresponding mask bit in m is not set, in which case, the element is set to zero.
*/
func Zero[Element constraintsExt.Number](m Bool, a [Length]Element) (r [Length]Element) {
	m.For(func(i int, c bool) {
		if !c {
			r[i] = a[i]
		}
	})
	return
}

/*
Broadcast returns a vector with a single value a copied into every element. This can be used to combine a vector with a scalar; for example: multiply every element of a vector by 10
`b := vec2.Mul(mask.All(), a, vec2.Broadcast(10))`

(Note: README-DESIGN.md explains absence of mask parameter)
*/
func Broadcast[Element constraintsExt.Number](a Element) (r [Length]Element) {
	for i := range r {
		r[i] = a
	}
	return
}

/*
Permute returns a masked vector containing elements from b chosen from indices in a. An element from b may appear more than once in the result, or may be absent.
*/
func Permute[Element constraints.Integer, U constraintsExt.Number](m Bool, a [Length]Element, b [Length]U) (r [Length]U) {
	m.ForTrue(func(i int) {
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
func Interlace[Element constraintsExt.Number](n int, a, b [Length / 2]Element) (r [Length]Element) {
	validInterlace(n)
	for x, y, z := a[:], b[:], r[:]; len(z) > 0; x, y, z = x[n:], y[n:], z[n*2:] {
		copy(z, x[:n])
		copy(z[n:], y[:n])
	}
	return
}

/*
Deinterlace splits a vector into 2 half-length vectors r and s, placing n consecutive elements in r, then n consecutive elements in s.
This is repeated until all the elements of a are present in r and s.
This is the reverse of Interlace.
*/
func Deinterlace[Element constraintsExt.Number](n int, a [Length]Element) (r, s [Length / 2]Element) {
	validInterlace(n)
	for x, y, z := a[:], r[:], s[:]; len(x) > 0; x, y, z = x[n*2:], y[n:], z[n:] {
		copy(y, x[:n])
		copy(z, x[n:n*2])
	}
	return
}

func validInterlace(n int) {
	if n > Length {
		panic(fmt.Sprintf("n cannot be greater than vector length %v", Length))
	}
	if bits.OnesCount(uint(n)) != 1 {
		panic("n must be a power of 2")
	}
}
