package simd

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"github.com/duncanpierce/vector/predicate"
	"golang.org/x/exp/constraints"
)

/*
Merge generalises Copy, Zero, Blend, Permute, Rotate, Shuffle, Reverse.
z[i] = m[i] ? x[w[i]] : y[w[i]]
If m == nil, all predicate are active and y has no effect on the result.
If x == nil, x[i] takes the value of z[i].
If y == nil, y[i] takes the value of z[i].
If w == nil, w[i] takes the value of i.
It is difficult to pass nils from a type-parameterised function because there is no easy way to specify N, the length of the vector.
*/
func Merge[E any, W constraintsExt.Vector[int], XYZ constraintsExt.Vector[E]](z *XYZ, m *predicate.Bool, w *W, x, y *XYZ) {
	if x == nil {
		x = z
	}
	if y == nil {
		y = z
	}
	var temp XYZ
	predicate.RangeAll(len(*z), m, func(i int, b bool) {
		index := i
		if w != nil {
			index = (*w)[i]
		}
		if b {
			temp[i] = (*x)[index]
		} else {
			temp[i] = (*y)[index]
		}
	})
	for i := 0; i < len(*z); i++ {
		(*z)[i] = temp[i]
	}
}

func Copy[E any, XZ constraintsExt.Vector[E]](z *XZ, m *predicate.Bool, x *XZ) {
	unary(z, m, x, func(x E) E {
		return x
	})
}

/*
Blend returns a vector with each element drawn from a if the corresponding mask bit in m is set or from b if the bit is not set.
*/
func Blend[E any, XYZ constraintsExt.Vector[E]](z *XYZ, m *predicate.Bool, x, y *XYZ) {
	predicate.RangeAll(len(*z), m, func(i int, b bool) {
		if b {
			(*z)[i] = (*x)[i]
		} else {
			(*z)[i] = (*y)[i]
		}
	})
}

/*
Zero sets elements of a vector to their default (zero) value where the corresponding mask bit in m is not set.
*/
func Zero[E any, Z constraintsExt.Vector[E]](z *Z, m *predicate.Bool) {
	var zero E
	predicate.RangeInactive(len(*z), m, func(i, j int) {
		(*z)[i] = zero
	})
}

/*
Permute sets vector z to elements from x selected using indices in y. An element from x may appear more than once in the result, or may be absent.
*/
func Permute[E any, I constraints.Integer, XZ constraintsExt.Vector[E], Y constraintsExt.Vector[I]](z *XZ, m *predicate.Bool, x *XZ, y *Y) {
	predicate.RangeActive(len(*z), m, func(i, j int) {
		index := (*y)[i]
		(*z)[i] = (*x)[index]
	})
}

/*
Gather sets elements of z to the values referenced by pointers in x.
*/
func Gather[E any, Z constraintsExt.Vector[E], X constraintsExt.Vector[*E]](z *Z, m *predicate.Bool, x *X) {
	predicate.RangeActive(len(*z), m, func(i, j int) {
		(*z)[i] = *(*x)[i]
	})
}

/*
Address sets elements of vector z to the addresses of slice elements in x, indexed by y.
*/
func Address[E any, Z constraintsExt.Vector[*E], X constraintsExt.Vector[[]E], Y constraintsExt.Vector[int]](z *Z, m *predicate.Bool, x *X, y *Y) {
	predicate.RangeActive(len(*z), m, func(i, j int) {
		index := (*y)[i]
		(*z)[i] = &((*x)[i][index])
	})
}

/*
Scatter sets values pointed to by elements of z to corresponding values in x.
*/
func Scatter[E any, Z constraintsExt.Vector[*E], X constraintsExt.Vector[E]](z *Z, m *predicate.Bool, x *X) {
	predicate.RangeActive(len(*z), m, func(i, j int) {
		*(*z)[i] = (*x)[i]
	})
}

/*
Shuffle exchanges w consecutive elements from x with the neighbouring w consecutive elements and writes the result to z.
When w == 0, Shuffle is equivalent to Copy. When w == 1, each element is exchanged with its neighbour.
*/
func Shuffle[E any, XZ constraintsExt.Vector[E]](z *XZ, w int, x *XZ) {
	//	if w < 0 {
	//		panic("w cannot be less than 0")
	//	}
	//	temp := [64]E{}
	//	for _, xs := temp[:], (*x)[:]; len(xs) > 0; {
	//		//copy(ts, xs[:w])
	//	}
}

/*
Pairs splits vector x, placing alternate consecutive elements of x in z0 and z1. Panics if vectors z0 and z1 are not exactly half the length of x.
*/
func Pairs[E any, Z constraintsExt.Vector[E], X constraintsExt.Vector[E]](z0, z1 *Z, x *X) {
	if len(*z0)*2 != len(*x) {
		panic("results z0 and z1 must be half the length of x")
	}
	for i := 0; i < len(*z0); i++ {
		j := i * 2
		(*z0)[i] = (*x)[j]
		(*z1)[i] = (*x)[j+1]
	}
}

/*
Iota stores incrementing integer values in elements of z.
TODO should it store i, j or allow caller to decide?
TODO can only be implemented for complex numbers if we accumulate rather than multiplying (for which type construction complex128(int) does not work)
*/
func Iota[E constraintsExt.ConvertableNumber, Z constraintsExt.Vector[E]](z *Z, m *predicate.Bool, start, inc E) {
	predicate.RangeActive(len(*z), m, func(i, j int) {
		(*z)[i] = start + (E(i) * inc)
	})
}
