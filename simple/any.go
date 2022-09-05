package simple

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"github.com/duncanpierce/vector/simple/lanes"
	"golang.org/x/exp/constraints"
)

func Copy[E any, XZ constraintsExt.Vector[E]](z *XZ, m *lanes.Bool, x *XZ) {
	unary(z, m, x, func(x E) E {
		return x
	})
}

/*
Blend returns a vector with each element drawn from a if the corresponding mask bit in m is set or from b if the bit is not set.
*/
func Blend[E any, XYZ constraintsExt.Vector[E]](z *XYZ, m *lanes.Bool, x, y *XYZ) {
	lanes.RangeAll[E](z, m, func(i, j int, b bool) {
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
func Zero[E any, Z constraintsExt.Vector[E]](z *Z, m *lanes.Bool) {
	var zero E
	lanes.RangeInactive[E](z, m, func(i, j int) {
		(*z)[i] = zero
	})
}

/*
Permute sets vector z to elements from x selected using indices in y. An element from x may appear more than once in the result, or may be absent.
*/
func Permute[E any, I constraints.Integer, XZ constraintsExt.Vector[E], Y constraintsExt.Vector[I]](z *XZ, m *lanes.Bool, x *XZ, y *Y) {
	lanes.RangeActive[E](z, m, func(i, j int) {
		index := (*y)[i]
		(*z)[i] = (*x)[index]
	})
}

/*
Gather sets elements of z to the values referenced by pointers in x.
*/
func Gather[E any, Z constraintsExt.Vector[E], X constraintsExt.Vector[*E]](z *Z, m *lanes.Bool, x *X) {
	lanes.RangeActive[E](z, m, func(i, j int) {
		(*z)[i] = *(*x)[i]
	})
}

/*
Address sets elements of vector z to the addresses of slice elements in x, indexed by y.
*/
func Address[E any, Z constraintsExt.Vector[*E], X constraintsExt.Vector[[]E], Y constraintsExt.Vector[int]](z *Z, m *lanes.Bool, x *X, y *Y) {
	lanes.RangeActive[*E](z, m, func(i, j int) {
		index := (*y)[i]
		(*z)[i] = &((*x)[i][index])
	})
}

/*
Scatter sets values pointed to by elements of z to corresponding values in x.
*/
func Scatter[E any, Z constraintsExt.Vector[*E], X constraintsExt.Vector[E]](z *Z, m *lanes.Bool, x *X) {
	lanes.RangeActive[*E](z, m, func(i, j int) {
		*(*z)[i] = (*x)[i]
	})
}

/*
Shuffle exchanges w consecutive elements from x with their neighbouring elements and writes the result to z.
When w == 0, Shuffle is equivalent to Copy. When w == 1, each element is exchanged with its neighbour.
*/
//func Shuffle[E any, XZ constraintsExt.Vector[E]](z *XZ, m *lanes.Bool, w int, x *XZ) {
//	if w < 0 {
//		panic("w cannot be less than 0")
//	}
//	temp := [64]E{}
//	for _, xs := temp[:], (*x)[:]; len(xs) > 0; {
//		//copy(ts, xs[:w])
//	}
//}
