package simd

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"github.com/duncanpierce/vector/predicate"
)

func Broadcast[E any, Z constraintsExt.Vector[E]](z *Z, m *predicate.Bool, x E) {
	predicate.RangeActive(len(*z), m, func(i, j int) {
		(*z)[i] = x
	})
}
