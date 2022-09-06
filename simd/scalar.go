package simd

import (
	"github.com/duncanpierce/vector/constraintsExt"
	"github.com/duncanpierce/vector/simd/lanes"
)

func Broadcast[E any, Z constraintsExt.Vector[E]](z *Z, m *lanes.Bool, x E) {
	lanes.RangeActive[E](z, m, func(i, j int) {
		(*z)[i] = x
	})
}
